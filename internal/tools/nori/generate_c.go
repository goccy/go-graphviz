package nori

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/goccy/nori/nori"
)

func generateCFile(file *File) ([]byte, error) {
	parsed, err := template.New("").Funcs(map[string]any{
		"map": createMap,
	}).Parse(bindC)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template bind.c.tmpl: %w", err)
	}
	var buf bytes.Buffer
	if err := parsed.Execute(&buf, &CFile{file}); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.Bytes(), nil
}

//go:embed templates/bind.c.tmpl
var bindC string

type CFile struct {
	*File
}

type CExportFunction struct {
	Name     string
	Args     []*CArg
	Return   *CReturn
	Function string
	FuncArgs []*CArg
}

type CCallbackFunction struct {
	Name   string
	Args   []*CArg
	Return string
}

type CArg struct {
	Index     int
	IsLastArg bool
	Value     *CValue
	Dst       string
	Src       string
}

type CReturn struct {
	Index int
	Value *CValue
}

type CValue struct {
	Src           string
	Dst           string
	IsReturnValue bool
	typ           *Type
}

func (v *CValue) IsFunction() bool {
	if v.typ.Kind == nori.TypeKind_FUNCPTR {
		return true
	}
	if v.typ.Ref == nil {
		return false
	}
	msg, ok := v.typ.Ref.(*Message)
	if !ok {
		return false
	}
	return msg.Rule != nil && msg.Rule.Funcptr != nil
}

func (v *CValue) FuncName() string {
	if !v.IsFunction() {
		return ""
	}
	return v.typ.Ref.(*Message).FullName()
}

func (v *CValue) IsString() bool {
	return v.typ.IsString()
}

func (v *CValue) IsFloat() bool {
	return v.typ.IsFloatKind()
}

func (v *CValue) IsSlice() bool {
	return v.typ.IsRepeated
}

func (v *CValue) IsStruct() bool {
	return v.typ.Pointer == 0 && v.typ.Kind == nori.TypeKind_STRUCT
}

func (v *CValue) Type() *Type {
	return v.typ
}

func (v *CValue) IsPtrValue() bool {
	return !v.typ.IsRepeated && v.typ.Pointer > 0
}

func (v *CValue) IsStringKind() bool {
	return v.typ.IsStringKind()
}

func (v *CValue) FixedArrayNum() uint64 {
	return v.typ.ArrayNum
}

func (v *CValue) ArrayNumArgIndex() int {
	if v.typ.ArgArrayNum == 0 {
		return 0
	}
	return int(v.typ.ArgArrayNum) - 1
}

func (v *CValue) Elem() *CValue {
	cloned := *v.typ
	cloned.IsRepeated = false
	if cloned.Pointer > 1 {
		cloned.Pointer--
	}
	return &CValue{
		Src: fmt.Sprintf("%s[i]", v.Src),
		Dst: "v",
		typ: &cloned,
	}
}

func (v *CValue) ArgStringLength() int {
	return int(v.typ.ArgStringLength)
}

func (v *CValue) IntPtr() string {
	if v.IsPtrValue() || v.IsStruct() || v.IsStringKind() || v.IsFloat() {
		return ""
	}
	return "(intptr_t)"
}

func (v *CValue) GoType() string {
	if v.typ.IsRepeated {
		return "GoSlice *"
	}
	ptr := strings.Repeat("*", int(v.typ.Pointer))
	switch v.typ.Kind {
	case nori.TypeKind_STRUCT:
		typeName := v.typ.Ref.(*Message).Rule.Alias
		if v.typ.Pointer == 0 {
			return typeName + " *"
		}
		return typeName + " " + ptr
	case nori.TypeKind_INT:
		return "int" + ptr
	case nori.TypeKind_UINT:
		return "unsigned int" + ptr
	case nori.TypeKind_VOIDPTR:
		if v.typ.Pointer == 0 {
			return "void *"
		}
		return "GoString " + ptr
	case nori.TypeKind_CHARPTR, nori.TypeKind_STRING:
		if v.typ.Pointer == 0 {
			return "GoString *"
		}
		return "GoString " + ptr
	case nori.TypeKind_BOOL:
		return "bool" + ptr
	case nori.TypeKind_UINT64:
		return "unsigned long long int" + ptr
	case nori.TypeKind_INT64:
		return "long long int" + ptr
	case nori.TypeKind_ENUM:
		return "int" + ptr
	case nori.TypeKind_FUNCPTR:
		return "void *"
	case nori.TypeKind_DOUBLE:
		return "GoString *" + ptr
	case nori.TypeKind_INT32:
		return "long int" + ptr
	case nori.TypeKind_UINT32:
		return "unsigned long int" + ptr
	case nori.TypeKind_FLOAT:
		return "GoString *" + ptr
	}
	return ""
}

func (v *CValue) WasmType() (ret string) {
	defer func() {
		if v.IsReturnValue {
			ret += "*"
		}
	}()
	if v.typ.IsRepeated {
		return "GoSlice *"
	}
	if v.typ.Pointer > 0 {
		return "void *"
	}
	var typeName string
	switch v.typ.Kind {
	case nori.TypeKind_STRUCT:
		typeName = "void *"
	case nori.TypeKind_INT:
		typeName = "int"
	case nori.TypeKind_UINT:
		typeName = "unsigned int"
	case nori.TypeKind_VOIDPTR:
		typeName = "void *"
	case nori.TypeKind_CHARPTR:
		typeName = "void *"
	case nori.TypeKind_STRING:
		typeName = "void *"
	case nori.TypeKind_BOOL:
		typeName = "bool"
	case nori.TypeKind_UINT64:
		typeName = "unsigned long long int"
	case nori.TypeKind_INT64:
		typeName = "long long int"
	case nori.TypeKind_ENUM:
		typeName = "int"
	case nori.TypeKind_FUNCPTR:
		typeName = "void *"
	case nori.TypeKind_DOUBLE:
		typeName = "double"
	case nori.TypeKind_INT32:
		typeName = "long int"
	case nori.TypeKind_UINT32:
		typeName = "unsigned long int"
	case nori.TypeKind_FLOAT:
		typeName = "float"
	}
	return typeName
}

func (v *CValue) CType() string {
	var (
		typeName string
		ptrNum   int
	)
	switch v.typ.Kind {
	case nori.TypeKind_STRUCT:
		msgRule := v.typ.Ref.(*Message).Rule
		typeName = msgRule.Alias
	case nori.TypeKind_INT:
		typeName = "int"
	case nori.TypeKind_UINT:
		typeName = "unsigned int"
	case nori.TypeKind_VOIDPTR:
		typeName = "void *"
	case nori.TypeKind_CHARPTR:
		typeName = "char"
		ptrNum = 1
	case nori.TypeKind_STRING:
		typeName = "char"
		ptrNum = 1
	case nori.TypeKind_BOOL:
		typeName = "bool"
	case nori.TypeKind_UINT64:
		typeName = "unsigned long long int"
	case nori.TypeKind_INT64:
		typeName = "long long int"
	case nori.TypeKind_ENUM:
		enumRule := v.typ.Ref.(*Enum).Rule
		typeName = enumRule.Alias
	case nori.TypeKind_FUNCPTR:
		typeName = "void *"
	case nori.TypeKind_DOUBLE:
		typeName = "double"
	case nori.TypeKind_INT32:
		typeName = "long int"
	case nori.TypeKind_UINT32:
		typeName = "unsigned long int"
	case nori.TypeKind_FLOAT:
		typeName = "float"
	}
	var attrs []string
	if v.typ.Const {
		attrs = append(attrs, "const")
	}
	if v.typ.Pointer != 0 || v.typ.IsRepeated {
		typeName = strings.TrimSuffix(typeName, "*")
	}
	attrs = append(attrs, typeName)
	if v.typ.Pointer != 0 {
		ptrNum = int(v.typ.Pointer)
	}
	if v.typ.IsRepeated {
		ptrNum++
	}
	if ptrNum != 0 {
		attrs = append(attrs, strings.Repeat("*", ptrNum))
	}
	if v.typ.Addr {
		attrs = append(attrs, "&")
	}
	return strings.Join(attrs, " ")
}

func (v *CValue) Converter() string {
	var (
		prefixes []string
		suffixes []string
	)
	if v.typ.Const {
		prefixes = append(prefixes, "const")
	}
	var ptrNum int
	if v.typ.Pointer == 0 && (v.typ.Kind == nori.TypeKind_CHARPTR || v.typ.Kind == nori.TypeKind_VOIDPTR) {
		ptrNum = 1
	} else {
		ptrNum = int(v.typ.Pointer)
	}
	if v.typ.IsRepeated {
		ptrNum++
	}
	if ptrNum != 0 {
		suffixes = append(suffixes, strings.Repeat("*", ptrNum))
	} else if v.typ.Kind == nori.TypeKind_STRUCT {
		ptrNum++
		suffixes = append(suffixes, strings.Repeat("*", ptrNum))
	}
	if v.typ.Addr {
		suffixes = append(suffixes, "&")
	}

	var typeName string
	switch v.typ.Kind {
	case nori.TypeKind_INT:
		return v.toCastValue("int", prefixes, suffixes)
	case nori.TypeKind_INT32:
		return v.toCastValue("long int", prefixes, suffixes)
	case nori.TypeKind_INT64:
		return v.toCastValue("long long int", prefixes, suffixes)
	case nori.TypeKind_UINT:
		return v.toCastValue("unsigned int", prefixes, suffixes)
	case nori.TypeKind_UINT32:
		return v.toCastValue("unsigned long int", prefixes, suffixes)
	case nori.TypeKind_UINT64:
		return v.toCastValue("unsigned long long int", prefixes, suffixes)
	case nori.TypeKind_FLOAT:
		return v.toCastValue("float", prefixes, suffixes)
	case nori.TypeKind_DOUBLE:
		return v.toCastValue("double", prefixes, suffixes)
	case nori.TypeKind_VOIDPTR:
		if v.typ.Pointer == 0 {
			return fmt.Sprintf(
				"(%s)",
				strings.Join(append(prefixes, "void *"), " "),
			)
		}
		return v.toCastValue("void", prefixes, suffixes)
	case nori.TypeKind_CHARPTR, nori.TypeKind_STRING:
		if ptrNum == 0 {
			return fmt.Sprintf(
				"(%s)",
				strings.Join(append(prefixes, "char *"), " "),
			)
		}
		return v.toCastValue("char", prefixes, suffixes)
	case nori.TypeKind_BOOL:
		return v.toCastValue("bool", prefixes, suffixes)
	case nori.TypeKind_ENUM:
		enumRule := v.typ.Ref.(*Enum).Rule
		typeName = enumRule.Alias
		return v.toCastValue(typeName, prefixes, suffixes)
	case nori.TypeKind_STRUCT:
		typeName := v.typ.Ref.(*Message).Rule.Alias
		conv := v.toCastValue(typeName, prefixes, suffixes)
		if v.typ.Pointer == 0 && !v.typ.IsRepeated {
			// need dereference.
			return "*" + conv
		}
		return conv
	case nori.TypeKind_FUNCPTR:
		return "(void *)"
	}
	return ""
}

func (v *CValue) toCastValue(typeName string, prefixes, suffixes []string) string {
	return "(" + strings.Join(append(append(prefixes, typeName), suffixes...), " ") + ")"
}

func (f *CFile) Headers() []string {
	if f.Rule == nil {
		return nil
	}
	headerMap := make(map[string]struct{})
	for _, export := range f.Rule.Exports {
		for _, header := range export.Headers {
			headerMap[header] = struct{}{}
		}
	}
	ret := make([]string, 0, len(headerMap))
	for header := range headerMap {
		ret = append(ret, header)
	}
	sort.Strings(ret)
	return ret
}

type CExportType struct {
	Name           string
	Type           string
	HasConstructor bool
	Fields         []*CExportField
}

type CExportField struct {
	Name         string
	ReceiverType string
	Type         string
	Value        *CValue
	ArrayNum     uint64
}

type CExportEnum struct {
	Name   string
	Values []*CExportEnumValue
}

type CExportEnumValue struct {
	Name string
}

func (f *CFile) ExportMessages() []*CExportType {
	var ret []*CExportType
	for _, msg := range f.Messages {
		ret = append(ret, f.toExportTypes(msg)...)
	}
	return ret
}

func (f *CFile) ExportEnums() []*CExportEnum {
	ret := make([]*CExportEnum, 0, len(f.Enums))
	for _, enum := range f.Enums {
		ret = append(ret, f.toExportEnum(enum))
	}
	return ret
}

func (f *CFile) toExportEnum(enum *Enum) *CExportEnum {
	values := make([]*CExportEnumValue, 0, len(enum.Values))
	for _, value := range enum.Values {
		if value.Rule == nil || value.Rule.Alias == "" {
			continue
		}
		values = append(values, &CExportEnumValue{
			Name: value.Rule.Alias,
		})
	}
	return &CExportEnum{
		Name:   enum.Rule.Alias,
		Values: values,
	}
}

func (f *CFile) toExportTypes(msg *Message) []*CExportType {
	var ret []*CExportType
	for _, m := range msg.NestedMessages {
		ret = append(ret, f.toExportTypes(m)...)
	}
	if msg.Rule.Alias == "" {
		return ret
	}
	exportType := &CExportType{
		Name:           msg.FullName(),
		Type:           msg.Rule.Alias,
		HasConstructor: msg.Rule.HasConstructor,
	}
	for _, field := range msg.Fields {
		accessor := f.fieldAccessor(field)
		typeText := (&CValue{typ: field.Type}).CType()
		if field.Type.Pointer == 0 && field.Type.Kind == nori.TypeKind_STRUCT {
			typeText += "*"
		}
		exportType.Fields = append(exportType.Fields, &CExportField{
			Name:         field.FullName(),
			ReceiverType: msg.Rule.Alias,
			Value: &CValue{
				Src: fmt.Sprintf("recv->%s", accessor),
				Dst: "v",
				typ: field.Type,
			},
			Type:     typeText,
			ArrayNum: field.Type.ArrayNum,
		})
	}
	ret = append(ret, exportType)
	return ret
}

func (f *CFile) fieldAccessor(field *Field) string {
	if field.Rule != nil && field.Rule.Alias != "" {
		if field.Oneof != "" {
			return field.Oneof + "." + field.Rule.Alias
		}
		return field.Rule.Alias
	}
	if field.Oneof != "" {
		return field.Oneof + "." + field.Name
	}
	return field.Name
}

func (f *CFile) ExportCallbackFunctions() []*CCallbackFunction {
	var ret []*CCallbackFunction
	for _, msg := range f.Messages {
		ret = append(ret, f.toCallbackFunctions(msg)...)
	}
	return ret
}

func (f *CFile) toCallbackFunctions(msg *Message) []*CCallbackFunction {
	var ret []*CCallbackFunction
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
	var args []*CArg
	for idx, typ := range funcptr.Args {
		args = append(args, &CArg{
			Index:     idx,
			IsLastArg: idx == len(funcptr.Args)-1,
			Value: &CValue{
				Src: fmt.Sprintf("_arg%d", idx),
				Dst: fmt.Sprintf("arg%d", idx),
				typ: typ,
			},
		})
	}
	returnType := "void"
	if funcptr.Return != nil {
		returnType = (&CValue{typ: funcptr.Return, IsReturnValue: true}).CType()
	}
	return append(ret, &CCallbackFunction{
		Name:   msg.FullName(),
		Args:   args,
		Return: returnType,
	})
}

func (f *CFile) ExportFunctions() []*CExportFunction {
	if f.Rule == nil {
		return nil
	}
	var ret []*CExportFunction
	for _, export := range f.Rule.Exports {
		for _, fn := range export.Funcs {
			ret = append(ret, f.toExportFunction(fn, nil))
		}
		for _, mtd := range export.Methods {
			ret = append(ret, f.toExportFunction(mtd.FunctionDef, mtd.Receiver))
		}
	}
	return ret
}

func (f *CFile) toExportFunction(fn *FunctionDef, receiver *Type) *CExportFunction {
	argsDef := fn.Args
	if receiver != nil {
		argsDef = append([]*Type{receiver}, argsDef...)
	}
	var (
		args     = make([]*CArg, 0, len(argsDef))
		funcArgs = make([]*CArg, 0, len(argsDef))
	)
	for idx, typ := range argsDef {
		args = append(args, &CArg{
			Index:     idx,
			IsLastArg: fn.Return == nil && idx == len(argsDef)-1,
			Value: &CValue{
				typ: typ,
			},
		})
		funcArgs = append(funcArgs, &CArg{
			Index:     idx,
			IsLastArg: idx == len(argsDef)-1,
			Value: &CValue{
				typ: typ,
			},
			Dst: fmt.Sprintf("arg%d", idx),
			Src: fmt.Sprintf("_arg%d", idx),
		})
	}
	var retValue *CReturn
	if fn.Return != nil {
		args = append(args, &CArg{
			Index:     len(argsDef),
			IsLastArg: true,
			Value: &CValue{
				IsReturnValue: true,
				typ:           fn.Return,
			},
		})
		retValue = &CReturn{
			Index: len(argsDef),
			Value: &CValue{
				Src: "ret",
				Dst: "v",
				typ: fn.Return,
			},
		}
	}
	var names []string
	if receiver != nil {
		names = append(names, receiver.Ref.(*Message).FullName())
	}
	names = append(names, fn.Name)
	return &CExportFunction{
		Name:     strings.Join(names, "_"),
		Args:     args,
		Function: fn.Alias,
		FuncArgs: funcArgs,
		Return:   retValue,
	}
}
