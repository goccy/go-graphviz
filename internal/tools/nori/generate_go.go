package nori

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"github.com/goccy/nori/nori"
)

//go:embed templates/bind.go.tmpl
var bindGo string

type GoFile struct {
	*File
	WasmName string
}

type GoExportFunction struct {
	Receiver string
	GoName   string
	WasmName string
	Args     []*GoArg
	Return   *GoReturn
}

type GoArg struct {
	Index     int
	IsLastArg bool
	Src       string
	Dst       string
	Value     *GoValue
}

type GoReturn struct {
	Index int
	Src   string
	Dst   string
	Value *GoValue
}

type GoValue struct {
	typ *Type
}

func (v *GoValue) Elem() *GoValue {
	typ := *v.typ
	typ.Pointer--
	return &GoValue{typ: &typ}
}

func (v *GoValue) Is64Bit() bool {
	return v.typ.Is64Bit()
}

func (v *GoValue) IsFunction() bool {
	return v.typ.IsFunction()
}

func (v *GoValue) IsPtrValue() bool {
	return isGoPtrValue(v.typ)
}

func (v *GoValue) IsStringKind() bool {
	return v.typ.IsStringKind()
}

func (v *GoValue) IsFloatKind() bool {
	return v.typ.IsFloatKind()
}

func (v *GoValue) IsSlice() bool {
	return v.typ.IsRepeated
}

func (v *GoValue) FuncName() string {
	if !v.typ.IsFunction() {
		return ""
	}
	return v.typ.Ref.(*Message).FullName()
}

func (v *GoValue) GoType() string {
	return v.toTypeText(v.typ)
}

// GoInterfaceType consider callback function type when creates GoType text.
func (v *GoValue) GoInterfaceType() string {
	return v.toIfaceTypeText(v.typ)
}

func (v *GoValue) WasmTypeConverter() string {
	if isGoPtrValue(v.typ) {
		return "mod.toPtrWasmValue"
	}
	var typeName string
	switch v.typ.Kind {
	case nori.TypeKind_STRUCT:
		typeName = "mod.toObject"
	case nori.TypeKind_INT, nori.TypeKind_ENUM:
		typeName = "mod.toInt"
	case nori.TypeKind_INT32:
		typeName = "mod.toInt32"
	case nori.TypeKind_INT64:
		typeName = "mod.toInt64"
	case nori.TypeKind_UINT:
		typeName = "mod.toUint"
	case nori.TypeKind_UINT32:
		typeName = "mod.toUint32"
	case nori.TypeKind_UINT64:
		typeName = "mod.toUint64"
	case nori.TypeKind_VOIDPTR:
		typeName = "mod.toAny"
	case nori.TypeKind_CHARPTR, nori.TypeKind_STRING:
		typeName = "mod.toString"
	case nori.TypeKind_BOOL:
		typeName = "mod.toBool"
	case nori.TypeKind_FUNCPTR:
		typeName = "mod.toFunc"
	case nori.TypeKind_FLOAT:
		typeName = "mod.toFloat"
	case nori.TypeKind_DOUBLE:
		typeName = "mod.toDouble"
	default:
		typeName = "mod.toUint"
	}
	if v.typ.IsRepeated {
		typeName += "Array"
	}
	return typeName + "WasmValue"
}

func (v *GoValue) GoTypeConverter() string {
	var typeName string
	switch v.typ.Kind {
	case nori.TypeKind_STRUCT:
		typeName = fmt.Sprintf("new%s", toPublicGoVariable(v.typ.Ref.(*Message).Name))
	case nori.TypeKind_INT:
		typeName = "mod.toInt"
	case nori.TypeKind_INT32:
		typeName = "mod.toInt32"
	case nori.TypeKind_INT64:
		typeName = "mod.toInt64"
	case nori.TypeKind_UINT:
		typeName = "mod.toUint"
	case nori.TypeKind_UINT32:
		typeName = "mod.toUint32"
	case nori.TypeKind_UINT64:
		typeName = "mod.toUint64"
	case nori.TypeKind_ENUM:
		typeName = toPublicGoVariable(v.typ.Ref.(*Enum).Name)
	case nori.TypeKind_VOIDPTR:
		typeName = "mod.toAny"
	case nori.TypeKind_CHARPTR, nori.TypeKind_STRING:
		typeName = "mod.toString"
	case nori.TypeKind_BOOL:
		typeName = "mod.toBool"
	case nori.TypeKind_FUNCPTR:
		typeName = "mod.toAny"
	case nori.TypeKind_FLOAT:
		typeName = "mod.toFloat32"
	case nori.TypeKind_DOUBLE:
		typeName = "mod.toFloat64"
	default:
		typeName = "mod.toAny"
	}
	if v.typ.IsRepeated {
		typeName += "Slice"
	}
	return typeName
}

func (v *GoValue) toIfaceTypeText(typ *Type) string {
	ret := v.toTypeText(typ)
	if typ.Kind == nori.TypeKind_FUNCPTR {
		return fmt.Sprintf("*CallbackFunc[%s]", ret)
	}
	return ret
}

func (v *GoValue) toTypeText(typ *Type) string {
	var typeName string
	switch typ.Kind {
	case nori.TypeKind_STRUCT:
		typeName = "*" + toPublicGoVariable(typ.Ref.(*Message).Name)
	case nori.TypeKind_INT:
		typeName = "int"
	case nori.TypeKind_UINT:
		typeName = "uint"
	case nori.TypeKind_VOIDPTR:
		typeName = "any"
	case nori.TypeKind_CHARPTR:
		typeName = "string"
	case nori.TypeKind_STRING:
		typeName = "string"
	case nori.TypeKind_BOOL:
		typeName = "bool"
	case nori.TypeKind_UINT64:
		typeName = "uint64"
	case nori.TypeKind_INT64:
		typeName = "int64"
	case nori.TypeKind_ENUM:
		typeName = toPublicGoVariable(typ.Ref.(*Enum).Name)
	case nori.TypeKind_FUNCPTR:
		def := typ.Ref.(*Message).Rule.Funcptr
		args := make([]string, 0, len(def.Args))
		for _, arg := range def.Args {
			args = append(args, v.toTypeText(arg))
		}
		ret := "error"
		if def.Return != nil {
			ret = fmt.Sprintf("(%s, error)", v.toTypeText(def.Return))
		}
		typeName = fmt.Sprintf("func(context.Context, %s)%s", strings.Join(args, ","), ret)
	case nori.TypeKind_DOUBLE:
		typeName = "float64"
	case nori.TypeKind_INT32:
		typeName = "int32"
	case nori.TypeKind_UINT32:
		typeName = "uint32"
	case nori.TypeKind_FLOAT:
		typeName = "float32"
	}
	if typ.IsRepeated {
		return "[]" + typeName
	}
	if typ.Pointer == 1 {
		typeName = "*" + strings.TrimPrefix(typeName, "*")
	} else if typ.Pointer > 1 {
		return strings.Repeat("*", int(typ.Pointer-1)) + typeName
	}
	return typeName
}

func isGoPtrValue(typ *Type) bool {
	if !typ.IsRepeated && typ.Pointer > 0 {
		switch typ.Kind {
		case nori.TypeKind_STRUCT:
			if typ.Pointer == 2 {
				return true
			}
		case nori.TypeKind_CHARPTR:
			if typ.Pointer == 2 {
				return true
			}
		default:
			return true
		}
	}
	return false
}

type GoExportType struct {
	Name           string
	HasConstructor bool
	Fields         []*GoExportField
	EnumValues     []*GoExportEnumValue
}

type GoExportField struct {
	GoName   string
	WasmName string
	Src      string
	Dst      string
	Value    *GoValue
}

type GoExportEnumValue struct {
	GoName   string
	WasmName string
}

type GoCallbackFunction struct {
	Name   string
	Args   []*GoArg
	Return *GoReturn
}

func (f *GoFile) ExportCallbackFunctions() []*GoCallbackFunction {
	var ret []*GoCallbackFunction
	for _, msg := range f.Messages {
		ret = append(ret, f.toCallbackFunctions(msg)...)
	}
	return ret
}

func (f *GoFile) toCallbackFunctions(msg *Message) []*GoCallbackFunction {
	var ret []*GoCallbackFunction
	for _, m := range msg.NestedMessages {
		ret = append(ret, f.toCallbackFunctions(m)...)
	}
	if msg.Rule == nil {
		return ret
	}
	if msg.Rule.Funcptr == nil {
		return ret
	}
	funcptr := msg.Rule.Funcptr
	var args []*GoArg
	for idx, typ := range funcptr.Args {
		args = append(args, &GoArg{
			Index:     idx,
			IsLastArg: idx == len(funcptr.Args)-1,
			Value:     &GoValue{typ: typ},
			Src:       fmt.Sprintf("stack[%d]", idx),
			Dst:       "ret",
		})
	}
	var retValue *GoReturn
	if funcptr.Return != nil {
		retValue = &GoReturn{
			Value: &GoValue{typ: funcptr.Return},
		}
	}
	return append(ret, &GoCallbackFunction{
		Name:   msg.FullName(),
		Args:   args,
		Return: retValue,
	})
}

func (f *GoFile) ExportMessages() []*GoExportType {
	var ret []*GoExportType
	for _, msg := range f.Messages {
		ret = append(ret, f.toExportTypes(msg)...)
	}
	return ret
}

func (f *GoFile) toExportTypes(m *Message) []*GoExportType {
	var ret []*GoExportType
	for _, msg := range m.NestedMessages {
		ret = append(ret, f.toExportTypes(msg)...)
	}

	if m.Rule != nil && m.Rule.Funcptr != nil {
		return ret
	}

	exportType := &GoExportType{
		Name:           toPublicGoVariable(m.Name),
		HasConstructor: m.Rule.HasConstructor,
	}
	for _, field := range m.Fields {
		exportType.Fields = append(exportType.Fields, &GoExportField{
			Src:      "p",
			Dst:      "ret",
			Value:    &GoValue{typ: field.Type},
			WasmName: field.FullName(),
			GoName:   toPublicGoVariable(field.Name),
		})
	}
	return append(ret, exportType)
}

func (f *GoFile) ExportEnums() []*GoExportType {
	var ret []*GoExportType
	for _, enum := range f.Enums {
		enumValues := make([]*GoExportEnumValue, 0, len(enum.Values))
		for _, value := range enum.Values {
			if value.Rule == nil || value.Rule.Alias == "" {
				continue
			}
			enumValues = append(enumValues, &GoExportEnumValue{
				GoName:   value.Name,
				WasmName: value.Rule.Alias,
			})
		}
		ret = append(ret, &GoExportType{
			Name:       toPublicGoVariable(enum.Name),
			EnumValues: enumValues,
		})
	}
	return ret
}

func (f *GoFile) ExportFunctions() []*GoExportFunction {
	if f.Rule == nil {
		return nil
	}
	var ret []*GoExportFunction
	for _, export := range f.Rule.Exports {
		for _, fn := range export.Funcs {
			ret = append(ret, f.toExportFunction(fn, nil))
		}
	}
	return ret
}

func (f *GoFile) ExportMethods() []*GoExportFunction {
	if f.Rule == nil {
		return nil
	}
	var ret []*GoExportFunction
	for _, export := range f.Rule.Exports {
		for _, mtd := range export.Methods {
			ret = append(ret, f.toExportFunction(mtd.FunctionDef, mtd.Receiver))
		}
	}
	return ret
}

func (f *GoFile) toExportFunction(fn *FunctionDef, receiver *Type) *GoExportFunction {
	argsDef := fn.Args
	var args = make([]*GoArg, 0, len(argsDef))
	for idx, typ := range argsDef {
		args = append(args, &GoArg{
			Index:     idx,
			IsLastArg: idx == len(argsDef)-1,
			Src:       "p",
			Dst:       "_arg" + fmt.Sprint(idx),
			Value:     &GoValue{typ: typ},
		})
	}
	var retValue *GoReturn
	if fn.Return != nil {
		retValue = &GoReturn{
			Index: len(argsDef),
			Value: &GoValue{typ: fn.Return},
			Src:   "p",
			Dst:   "ret",
		}
	}
	var names []string
	if receiver != nil {
		names = append(names, receiver.Ref.(*Message).Name)
	}
	names = append(names, fn.Name)
	var receiverName string
	if receiver != nil {
		receiverName = toPublicGoVariable(receiver.Ref.(*Message).Name)
	}
	return &GoExportFunction{
		Receiver: receiverName,
		GoName:   toPublicGoVariable(fn.Name),
		WasmName: strings.Join(names, "_"),
		Args:     args,
		Return:   retValue,
	}
}

func generateGoFile(file *File) ([]byte, error) {
	parsed, err := template.New("").Funcs(map[string]any{
		"map": createMap,
	}).Parse(bindGo)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template bind.go.tmpl: %w", err)
	}
	var buf bytes.Buffer
	if err := parsed.Execute(&buf, &GoFile{File: file, WasmName: "graphviz"}); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to format %s: %w", buf.String(), err)
	}
	return src, nil
}

func toPublicGoVariable(s string) string {
	if len(s) == 0 {
		return ""
	}
	up := strings.ToUpper(string(s[0]))
	if len(s) == 1 {
		return up
	}
	return toGoVariable(up + s[1:])
}

func toPrivateGoVariable(s string) string {
	if len(s) == 0 {
		return ""
	}
	up := strings.ToLower(string(s[0]))
	if len(s) == 1 {
		return up
	}
	return toGoVariable(up + s[1:])
}

func toGoVariable(s string) string {
	ret := make([]rune, 0, len(s))
	var isUpper bool
	for _, c := range s {
		if c == '_' {
			isUpper = true
			continue
		}
		if isUpper {
			ret = append(ret, []rune(strings.ToUpper(string(c)))...)
			isUpper = false
		} else {
			ret = append(ret, c)
		}
	}
	return string(ret)
}

func createMap(pairs ...any) (map[string]any, error) {
	if len(pairs)%2 != 0 {
		return nil, errors.New("the number of arguments must be divisible by two")
	}

	m := make(map[string]any, len(pairs)/2)
	for i := 0; i < len(pairs); i += 2 {
		key, ok := pairs[i].(string)

		if !ok {
			return nil, fmt.Errorf("cannot use type %T as map key", pairs[i])
		}
		m[key] = pairs[i+1]
	}
	return m, nil
}
