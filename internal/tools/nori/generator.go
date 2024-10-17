package nori

import (
	"context"
	_ "embed"
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/goccy/nori/nori"
)

type File struct {
	Name     string
	Messages []*Message
	Enums    []*Enum
	Rule     *FileRule
}

type FileRule struct {
	Exports []*Export
}

type Export struct {
	Headers []string
	Funcs   []*FunctionDef
	Methods []*MethodDef
}

type FunctionDef struct {
	Name   string
	Alias  string
	Args   []*Type
	Return *Type
}

type MethodDef struct {
	*FunctionDef
	Receiver *Type
}

type TypeKind = nori.TypeKind

type Type struct {
	Kind            TypeKind
	Pointer         uint64
	Const           bool
	Addr            bool
	IsFuncBasePtr   bool
	IsRepeated      bool
	ArrayNum        uint64
	ArgArrayNum     uint64
	ArgStringLength uint64
	Ref             any
}

type Message struct {
	Name           string
	Fields         []*Field
	NestedMessages []*Message
	Parent         *Message
	Rule           *MessageRule
}

type MessageRule struct {
	Funcptr        *FunctionDef
	Anonymous      bool
	Alias          string
	HasConstructor bool
}

type Field struct {
	Name    string
	Type    *Type
	Rule    *FieldRule
	Message *Message
	Oneof   string
}

type FieldRule struct {
	Type  *Type
	Alias string
}

type Enum struct {
	Name   string
	Values []*EnumValue
	Rule   *EnumRule
}

type EnumRule struct {
	Alias string
}

type EnumValue struct {
	Name string
	Rule *EnumValueRule
}

type EnumValueRule struct {
	Alias string
}

func (m *Message) FullName() string {
	return strings.Join(append(m.ParentMessageNames(), m.Name), "_")
}

func (m *Message) ParentMessageNames() []string {
	if m.Parent == nil {
		return []string{}
	}
	return append(m.Parent.ParentMessageNames(), m.Parent.Name)
}

func (f *Field) FullName() string {
	return f.Message.FullName() + "_" + f.Name
}

func (t *Type) Is64Bit() bool {
	if t.Pointer != 0 {
		return false
	}
	return t.Kind == nori.TypeKind_UINT64 || t.Kind == nori.TypeKind_INT64 || t.Kind == nori.TypeKind_DOUBLE
}

func (t *Type) IsString() bool {
	if t.IsRepeated {
		return false
	}
	if t.Pointer > 1 {
		return false
	}
	if t.Kind == nori.TypeKind_STRING && t.Pointer == 0 {
		return true
	}
	if t.Kind == nori.TypeKind_CHARPTR && t.Pointer <= 1 {
		return true
	}
	return false
}

func (t *Type) IsStringKind() bool {
	if t.Kind == nori.TypeKind_STRING {
		return true
	}
	if t.Kind == nori.TypeKind_CHARPTR {
		return true
	}
	return false
}

func (t *Type) IsFloatKind() bool {
	return t.Kind == nori.TypeKind_FLOAT || t.Kind == nori.TypeKind_DOUBLE
}

func (t *Type) IsFunction() bool {
	if t == nil {
		return false
	}
	if t.Ref != nil {
		msg, ok := t.Ref.(*Message)
		if ok {
			if msg.Rule != nil && msg.Rule.Funcptr != nil {
				return true
			}
		}
	}
	return false
}

func Generate(ctx context.Context, req *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	files, err := newResolver().Resolve(req.GetProtoFile())
	if err != nil {
		return nil, err
	}
	lastFile := files[len(files)-1]
	if lastFile.Name == "nori/nori.proto" {
		return nil, nil
	}
	cFile, err := generateCFile(lastFile)
	if err != nil {
		return nil, err
	}
	goFile, err := generateGoFile(lastFile)
	if err != nil {
		return nil, err
	}
	return &pluginpb.CodeGeneratorResponse{
		File: []*pluginpb.CodeGeneratorResponse_File{
			{
				Name:    proto.String("bind.c"),
				Content: proto.String(string(cFile)),
			},
			{
				Name:    proto.String("bind.go"),
				Content: proto.String(string(goFile)),
			},
		},
	}, nil
}

type Resolver struct {
	pkgMap           map[string]struct{}
	messageMap       map[string]*Message
	fieldMap         map[string]*Field
	enumMap          map[string]*Enum
	enumValueMap     map[string]*EnumValue
	messageRuleMap   map[*Message]*nori.MessageRule
	fieldRuleMap     map[*Field]*nori.FieldRule
	enumRuleMap      map[*Enum]*nori.EnumRule
	enumValueRuleMap map[*EnumValue]*nori.EnumValueRule
}

func newResolver() *Resolver {
	return &Resolver{
		pkgMap:           make(map[string]struct{}),
		messageMap:       make(map[string]*Message),
		fieldMap:         make(map[string]*Field),
		enumMap:          make(map[string]*Enum),
		enumValueMap:     make(map[string]*EnumValue),
		messageRuleMap:   make(map[*Message]*nori.MessageRule),
		fieldRuleMap:     make(map[*Field]*nori.FieldRule),
		enumRuleMap:      make(map[*Enum]*nori.EnumRule),
		enumValueRuleMap: make(map[*EnumValue]*nori.EnumValueRule),
	}
}

func (r *Resolver) Resolve(defs []*descriptorpb.FileDescriptorProto) ([]*File, error) {
	if err := r.resolveReference(defs); err != nil {
		return nil, err
	}
	files, err := r.resolveFiles(defs)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *Resolver) resolveReference(defs []*descriptorpb.FileDescriptorProto) error {
	for _, def := range defs {
		if err := r.resolveFileReference(def); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) resolveFileReference(def *descriptorpb.FileDescriptorProto) error {
	pkgName := def.GetPackage()
	r.pkgMap[pkgName] = struct{}{}
	if err := r.resolveMessageReferences(pkgName, nil, def.GetMessageType()); err != nil {
		return err
	}
	if err := r.resolveEnumReferences(pkgName, nil, def.GetEnumType()); err != nil {
		return err
	}
	return nil
}

func (r *Resolver) resolveMessageReferences(pkgName string, parentMsgNames []string, defs []*descriptorpb.DescriptorProto) error {
	for _, def := range defs {
		if err := r.resolveMessageReference(pkgName, parentMsgNames, def); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) resolveMessageReference(pkgName string, parentMsgNames []string, def *descriptorpb.DescriptorProto) error {
	msgNames := append(parentMsgNames, def.GetName())
	fqdn := fmt.Sprintf("%s.%s", pkgName, strings.Join(msgNames, "."))
	if _, exists := r.messageMap[fqdn]; exists {
		return nil
	}
	r.messageMap[fqdn] = &Message{
		Name: def.GetName(),
	}
	if err := r.resolveMessageReferences(pkgName, msgNames, def.GetNestedType()); err != nil {
		return err
	}
	if err := r.resolveEnumReferences(pkgName, msgNames, def.GetEnumType()); err != nil {
		return err
	}
	if err := r.resolveFieldReferences(pkgName, msgNames, def.GetField()); err != nil {
		return err
	}
	return nil
}

func (r *Resolver) resolveFieldReferences(pkgName string, msgNames []string, defs []*descriptorpb.FieldDescriptorProto) error {
	for _, def := range defs {
		if err := r.resolveFieldReference(pkgName, msgNames, def); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) resolveFieldReference(pkgName string, msgNames []string, def *descriptorpb.FieldDescriptorProto) error {
	fqdn := fmt.Sprintf("%s.%s.%s", pkgName, strings.Join(msgNames, "."), def.GetName())
	if _, exists := r.fieldMap[fqdn]; exists {
		return nil
	}
	r.fieldMap[fqdn] = &Field{
		Name: def.GetName(),
	}
	return nil
}

func (r *Resolver) resolveEnumReferences(pkgName string, parentMsgNames []string, defs []*descriptorpb.EnumDescriptorProto) error {
	for _, def := range defs {
		if err := r.resolveEnumReference(pkgName, parentMsgNames, def); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) resolveEnumReference(pkgName string, parentMsgNames []string, def *descriptorpb.EnumDescriptorProto) error {
	enumName := def.GetName()
	fqdn := fmt.Sprintf("%s.%s", pkgName, strings.Join(append(parentMsgNames, enumName), "."))
	if _, exists := r.enumMap[fqdn]; exists {
		return nil
	}
	r.enumMap[fqdn] = &Enum{
		Name: enumName,
	}
	if err := r.resolveEnumValueReferences(pkgName, parentMsgNames, enumName, def.GetValue()); err != nil {
		return err
	}
	return nil
}

func (r *Resolver) resolveEnumValueReferences(pkgName string, parentMsgNames []string, enumName string, defs []*descriptorpb.EnumValueDescriptorProto) error {
	for _, def := range defs {
		if err := r.resolveEnumValueReference(pkgName, parentMsgNames, enumName, def); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) resolveEnumValueReference(pkgName string, parentMsgNames []string, enumName string, def *descriptorpb.EnumValueDescriptorProto) error {
	enumValueName := def.GetName()
	fqdn := fmt.Sprintf("%s.%s", pkgName, strings.Join(append(parentMsgNames, enumName, enumValueName), "."))
	if _, exists := r.enumValueMap[fqdn]; exists {
		return nil
	}
	r.enumValueMap[fqdn] = &EnumValue{
		Name: enumValueName,
	}
	return nil
}

func (r *Resolver) resolveFiles(defs []*descriptorpb.FileDescriptorProto) ([]*File, error) {
	ret := make([]*File, 0, len(defs))
	for _, def := range defs {
		file, err := r.resolveFile(def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, file)
	}
	return ret, nil
}

func (r *Resolver) resolveFile(def *descriptorpb.FileDescriptorProto) (*File, error) {
	ruleDef, err := getExtensionRule[*nori.FileRule](def.GetOptions(), nori.E_File)
	if err != nil {
		return nil, err
	}
	pkgName := def.GetPackage()
	msgs, err := r.resolveMessages(pkgName, def.GetMessageType())
	if err != nil {
		return nil, err
	}
	enums, err := r.resolveEnums(pkgName, def.GetEnumType())
	if err != nil {
		return nil, err
	}
	for _, msg := range r.messageMap {
		if err := r.resolveMessageRule(pkgName, msg, r.messageRuleMap[msg]); err != nil {
			return nil, err
		}
	}
	for _, field := range r.fieldMap {
		if err := r.resolveFieldRule(pkgName, field, r.fieldRuleMap[field]); err != nil {
			return nil, err
		}
	}
	for _, enum := range r.enumMap {
		if err := r.resolveEnumRule(enum, r.enumRuleMap[enum]); err != nil {
			return nil, err
		}
	}
	for _, value := range r.enumValueMap {
		if err := r.resolveEnumValueRule(value, r.enumValueRuleMap[value]); err != nil {
			return nil, err
		}
	}
	rule, err := r.resolveFileRule(pkgName, ruleDef)
	if err != nil {
		return nil, err
	}
	return &File{
		Name:     def.GetName(),
		Messages: msgs,
		Enums:    enums,
		Rule:     rule,
	}, nil
}

func (r *Resolver) resolveFileRule(pkgName string, def *nori.FileRule) (*FileRule, error) {
	exports, err := r.resolveExports(pkgName, def.GetExport())
	if err != nil {
		return nil, err
	}
	return &FileRule{
		Exports: exports,
	}, nil
}

func (r *Resolver) resolveExports(pkgName string, defs []*nori.Export) ([]*Export, error) {
	ret := make([]*Export, 0, len(defs))
	for _, def := range defs {
		export, err := r.resolveExport(pkgName, def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, export)
	}
	return ret, nil
}

func (r *Resolver) resolveExport(pkgName string, def *nori.Export) (*Export, error) {
	funcs, err := r.resolveFunctionDefs(pkgName, def.GetFunc())
	if err != nil {
		return nil, err
	}
	mtds, err := r.resolveMethodDefs(pkgName, def.GetMethod())
	if err != nil {
		return nil, err
	}
	return &Export{
		Headers: def.GetHeader(),
		Funcs:   funcs,
		Methods: mtds,
	}, nil
}

func (r *Resolver) resolveFunctionDefs(pkgName string, defs []*nori.FunctionDef) ([]*FunctionDef, error) {
	ret := make([]*FunctionDef, 0, len(defs))
	for _, def := range defs {
		fn, err := r.resolveFunctionDef(pkgName, def)
		if err != nil {
			return nil, err
		}
		if fn == nil {
			continue
		}
		ret = append(ret, fn)
	}
	return ret, nil
}

func (r *Resolver) resolveFunctionDef(pkgName string, def *nori.FunctionDef) (*FunctionDef, error) {
	if def == nil {
		return nil, nil
	}
	name := def.GetName()
	alias := name
	if v := def.GetAlias(); v != "" {
		alias = v
	}
	args, err := r.resolveTypes(pkgName, def.GetArgs())
	if err != nil {
		return nil, err
	}
	retType, err := r.resolveType(pkgName, def.GetReturn())
	if err != nil {
		return nil, err
	}
	return &FunctionDef{
		Name:   name,
		Alias:  alias,
		Args:   args,
		Return: retType,
	}, nil
}

func (r *Resolver) resolveMethodDefs(pkgName string, defs []*nori.MethodDef) ([]*MethodDef, error) {
	ret := make([]*MethodDef, 0, len(defs))
	for _, def := range defs {
		mtd, err := r.resolveMethodDef(pkgName, def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, mtd)
	}
	return ret, nil
}

func (r *Resolver) resolveMethodDef(pkgName string, def *nori.MethodDef) (*MethodDef, error) {
	name := def.GetName()
	alias := def.GetAlias()
	if alias == "" {
		alias = name
	}
	args, err := r.resolveTypes(pkgName, def.GetArgs())
	if err != nil {
		return nil, err
	}
	retType, err := r.resolveType(pkgName, def.GetReturn())
	if err != nil {
		return nil, err
	}
	recv := def.GetRecv()
	if !r.existsPackage(recv) {
		recv = fmt.Sprintf("%s.%s", pkgName, recv)
	}
	msg, exists := r.messageMap[recv]
	if !exists {
		return nil, fmt.Errorf("failed to find message from %s at resolving method receiver", recv)
	}
	return &MethodDef{
		Receiver: &Type{
			Kind:    nori.TypeKind_STRUCT,
			Ref:     msg,
			Pointer: 1,
		},
		FunctionDef: &FunctionDef{
			Name:   name,
			Alias:  alias,
			Args:   args,
			Return: retType,
		},
	}, nil
}

func (r *Resolver) resolveTypes(pkgName string, defs []*nori.Type) ([]*Type, error) {
	ret := make([]*Type, 0, len(defs))
	for _, def := range defs {
		typ, err := r.resolveType(pkgName, def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, typ)
	}
	return ret, nil
}

func (r *Resolver) resolveType(pkgName string, def *nori.Type) (*Type, error) {
	if def == nil {
		return nil, nil
	}
	typeKind := def.GetKind()
	refName := def.GetRef()
	if !r.existsPackage(refName) {
		refName = fmt.Sprintf("%s.%s", pkgName, refName)
	}

	var ref any
	switch typeKind {
	case nori.TypeKind_STRUCT:
		msg, exists := r.messageMap[refName]
		if !exists {
			return nil, fmt.Errorf("failed to find message from %s at resolving type", refName)
		}
		ref = msg
	case nori.TypeKind_FUNCPTR:
		msg, exists := r.messageMap[refName]
		if !exists {
			return nil, fmt.Errorf("failed to find message from %s at resolving type", refName)
		}
		if msg.Rule == nil || msg.Rule.Funcptr == nil {
			return nil, fmt.Errorf("%s message doesn't specify funcptr but used as a funcptr", refName)
		}
		ref = msg
	case nori.TypeKind_ENUM:
		enum, exists := r.enumMap[refName]
		if !exists {
			return nil, fmt.Errorf("failed to find enum from %s at resolving type", refName)
		}
		ref = enum
	}
	isRepeated := def.GetArray()
	if def.ArrayNum != nil {
		isRepeated = true
	}
	if def.ArrayNumArg != nil {
		isRepeated = true
	}
	return &Type{
		Kind:            typeKind,
		Ref:             ref,
		Const:           def.GetConst(),
		Addr:            def.GetAddr(),
		IsFuncBasePtr:   def.GetFuncbaseptr(),
		IsRepeated:      isRepeated,
		ArrayNum:        def.GetArrayNum(),
		ArgArrayNum:     def.GetArrayNumArg(),
		ArgStringLength: def.GetStringLengthArg(),
		Pointer:         def.GetPointer(),
	}, nil
}

func (r *Resolver) existsPackage(fqdn string) bool {
	name := strings.TrimPrefix(fqdn, ".")
	if !strings.Contains(name, ".") {
		return false
	}
	names := strings.Split(name, ".")
	for lastIdx := len(names) - 1; lastIdx > 0; lastIdx-- {
		pkgName := strings.Join(names[:lastIdx], ".")
		if _, exists := r.pkgMap[pkgName]; exists {
			return true
		}
	}
	return false
}

func (r *Resolver) resolveFieldType(pkgName string, kind descriptorpb.FieldDescriptorProto_Type, typeName string, isRepeated bool) (*Type, error) {
	typeName = strings.TrimPrefix(typeName, ".") // trim leading dot character.
	if !r.existsPackage(typeName) {
		typeName = fmt.Sprintf("%s.%s", pkgName, typeName)
	}
	var (
		ref      any
		typeKind TypeKind
	)
	switch kind {
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		typeKind = nori.TypeKind_DOUBLE
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		typeKind = nori.TypeKind_FLOAT
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		typeKind = nori.TypeKind_INT64
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		typeKind = nori.TypeKind_UINT64
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		typeKind = nori.TypeKind_INT32
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		typeKind = nori.TypeKind_UINT64
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		typeKind = nori.TypeKind_UINT32
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		typeKind = nori.TypeKind_BOOL
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		typeKind = nori.TypeKind_STRING
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		typeKind = nori.TypeKind_STRUCT
		msg, exists := r.messageMap[typeName]
		if !exists {
			return nil, fmt.Errorf("failed to find message from %s at resolving type", typeName)
		}
		ref = msg
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		typeKind = nori.TypeKind_STRING
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		typeKind = nori.TypeKind_UINT32
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		typeKind = nori.TypeKind_ENUM
		enum, exists := r.enumMap[typeName]
		if !exists {
			return nil, fmt.Errorf("failed to find enum from %s at resolving type", typeName)
		}
		ref = enum
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		typeKind = nori.TypeKind_INT32
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		typeKind = nori.TypeKind_INT64
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		typeKind = nori.TypeKind_INT32
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		typeKind = nori.TypeKind_INT64
	default:
		return nil, fmt.Errorf("found expected type kind %v", typeKind)
	}
	return &Type{
		Kind:       typeKind,
		Ref:        ref,
		IsRepeated: isRepeated,
	}, nil
}

func (r *Resolver) resolveMessages(pkgName string, defs []*descriptorpb.DescriptorProto) ([]*Message, error) {
	ret := make([]*Message, 0, len(defs))
	for _, def := range defs {
		msg, err := r.resolveMessage(pkgName, nil, def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, msg)
	}
	return ret, nil
}

func (r *Resolver) resolveMessageRule(pkgName string, msg *Message, def *nori.MessageRule) error {
	funcptr, err := r.resolveFunctionDef(pkgName, def.GetFuncptr())
	if err != nil {
		return err
	}
	if funcptr != nil {
		var funcBasePtrCount int
		for _, arg := range funcptr.Args {
			if arg.IsFuncBasePtr {
				funcBasePtrCount++
			}
		}
		if funcBasePtrCount != 1 {
			return fmt.Errorf("failed to resolve %s funcptr. funcbaseptr flag must be enabled for one of the arguments", msg.Name)
		}
	}
	hasConstructor := true
	if def != nil && def.Constructor != nil {
		hasConstructor = def.GetConstructor()
	}
	if funcptr != nil {
		hasConstructor = false
	}
	msg.Rule = &MessageRule{
		Funcptr:        funcptr,
		Anonymous:      def.GetAnonymous(),
		Alias:          def.GetAlias(),
		HasConstructor: hasConstructor,
	}
	return nil
}

func (r *Resolver) resolveMessage(pkgName string, parentMsgNames []string, def *descriptorpb.DescriptorProto) (*Message, error) {
	msgName := def.GetName()
	msgNames := append(parentMsgNames, msgName)
	fqdn := fmt.Sprintf("%s.%s", pkgName, strings.Join(msgNames, "."))
	msg, exists := r.messageMap[fqdn]
	if !exists {
		return nil, fmt.Errorf("failed to find message from %s", fqdn)
	}
	ruleDef, err := getExtensionRule[*nori.MessageRule](def.GetOptions(), nori.E_Message)
	if err != nil {
		return nil, err
	}
	r.messageRuleMap[msg] = ruleDef
	var oneofNames []string
	for _, oneofDef := range def.GetOneofDecl() {
		oneofNames = append(oneofNames, oneofDef.GetName())
	}
	fields, err := r.resolveFields(pkgName, msgNames, oneofNames, msg, def.GetField())
	if err != nil {
		return nil, err
	}
	msg.Fields = fields
	for _, nested := range def.GetNestedType() {
		nestedMsg, err := r.resolveMessage(pkgName, msgNames, nested)
		if err != nil {
			return nil, err
		}
		nestedMsg.Parent = msg
		msg.NestedMessages = append(msg.NestedMessages, nestedMsg)
	}
	return msg, nil
}

func (r *Resolver) resolveFields(pkgName string, msgNames, oneofNames []string, msg *Message, defs []*descriptorpb.FieldDescriptorProto) ([]*Field, error) {
	ret := make([]*Field, 0, len(defs))
	for _, def := range defs {
		field, err := r.resolveField(pkgName, msgNames, oneofNames, msg, def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, field)
	}
	return ret, nil
}

func (r *Resolver) resolveField(pkgName string, msgNames, oneofNames []string, msg *Message, def *descriptorpb.FieldDescriptorProto) (*Field, error) {
	ruleDef, err := getExtensionRule[*nori.FieldRule](def.GetOptions(), nori.E_Field)
	if err != nil {
		return nil, err
	}
	fieldName := def.GetName()
	fqdn := strings.Join(append(append([]string{pkgName}, msgNames...), fieldName), ".")
	field, exists := r.fieldMap[fqdn]
	if !exists {
		return nil, fmt.Errorf("failed to find field from %s", fqdn)
	}
	fieldType, err := r.resolveFieldType(pkgName, def.GetType(), def.GetTypeName(), def.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
	if err != nil {
		return nil, err
	}
	r.fieldRuleMap[field] = ruleDef
	field.Message = msg
	field.Type = fieldType

	if def.OneofIndex != nil {
		field.Oneof = oneofNames[def.GetOneofIndex()]
	}
	return field, nil
}

func (r *Resolver) resolveFieldRule(pkgName string, field *Field, def *nori.FieldRule) error {
	typ, err := r.resolveType(pkgName, def.GetType())
	if err != nil {
		return err
	}
	if typ != nil {
		if typ.Kind != 0 {
			field.Type.Kind = typ.Kind
		}
		if typ.Ref != nil {
			field.Type.Ref = typ.Ref
		}
		field.Type.Pointer = typ.Pointer
		field.Type.Const = typ.Const
		field.Type.Addr = typ.Addr
		field.Type.IsFuncBasePtr = typ.IsFuncBasePtr
		field.Type.ArrayNum = typ.ArrayNum
		field.Type.ArgArrayNum = typ.ArgArrayNum
		field.Type.IsRepeated = typ.IsRepeated
		if typ.IsFunction() {
			field.Type.Kind = nori.TypeKind_FUNCPTR
		}
	}
	if field.Type.IsFunction() {
		field.Type.Kind = nori.TypeKind_FUNCPTR
	}
	field.Rule = &FieldRule{
		Type:  typ,
		Alias: def.GetAlias(),
	}
	return nil
}

func (r *Resolver) resolveEnums(pkgName string, defs []*descriptorpb.EnumDescriptorProto) ([]*Enum, error) {
	ret := make([]*Enum, 0, len(defs))
	for _, def := range defs {
		enum, err := r.resolveEnum(pkgName, def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, enum)
	}
	return ret, nil
}

func (r *Resolver) resolveEnum(pkgName string, def *descriptorpb.EnumDescriptorProto) (*Enum, error) {
	enumName := def.GetName()
	fqdn := fmt.Sprintf("%s.%s", pkgName, enumName)
	enum, exists := r.enumMap[fqdn]
	if !exists {
		return nil, fmt.Errorf("failed to find enum from %s", fqdn)
	}
	ruleDef, err := getExtensionRule[*nori.EnumRule](def.GetOptions(), nori.E_Enum)
	if err != nil {
		return nil, err
	}
	r.enumRuleMap[enum] = ruleDef
	values, err := r.resolveEnumValues(pkgName, enumName, def.GetValue())
	if err != nil {
		return nil, err
	}
	enum.Values = values
	return enum, nil
}

func (r *Resolver) resolveEnumRule(enum *Enum, def *nori.EnumRule) error {
	alias := enum.Name
	if v := def.GetAlias(); v != "" {
		alias = v
	}
	enum.Rule = &EnumRule{
		Alias: alias,
	}
	return nil
}

func (r *Resolver) resolveEnumValues(pkgName, enumName string, defs []*descriptorpb.EnumValueDescriptorProto) ([]*EnumValue, error) {
	ret := make([]*EnumValue, 0, len(defs))
	for _, def := range defs {
		value, err := r.resolveEnumValue(pkgName, enumName, def)
		if err != nil {
			return nil, err
		}
		ret = append(ret, value)
	}
	return ret, nil
}

func (r *Resolver) resolveEnumValue(pkgName, enumName string, def *descriptorpb.EnumValueDescriptorProto) (*EnumValue, error) {
	valueName := def.GetName()
	fqdn := fmt.Sprintf("%s.%s.%s", pkgName, enumName, valueName)
	value, exists := r.enumValueMap[fqdn]
	if !exists {
		return nil, fmt.Errorf("failed to find enum value from %s", fqdn)
	}
	ruleDef, err := getExtensionRule[*nori.EnumValueRule](def.GetOptions(), nori.E_EnumValue)
	if err != nil {
		return nil, err
	}
	r.enumValueRuleMap[value] = ruleDef
	return value, nil
}

func (r *Resolver) resolveEnumValueRule(enumValue *EnumValue, def *nori.EnumValueRule) error {
	alias := enumValue.Name
	if v := def.GetAlias(); v != "" {
		alias = v
	}
	enumValue.Rule = &EnumValueRule{
		Alias: alias,
	}
	return nil
}

func getExtensionRule[T proto.Message](opts proto.Message, extType protoreflect.ExtensionType) (T, error) {
	var ret T

	typ := reflect.TypeOf(ret)
	if typ.Kind() != reflect.Ptr {
		return ret, fmt.Errorf("proto.Message value must be pointer type")
	}
	v := reflect.New(typ.Elem()).Interface().(proto.Message)

	if opts == nil {
		return ret, nil
	}
	if !proto.HasExtension(opts, extType) {
		return ret, nil
	}

	extFullName := extType.TypeDescriptor().Descriptor().FullName()

	if setRuleFromDynamicMessage(opts, extFullName, v) {
		return v.(T), nil
	}

	ext := proto.GetExtension(opts, extType)
	if ext == nil {
		return ret, fmt.Errorf("%s extension does not exist", extFullName)
	}
	rule, ok := ext.(T)
	if !ok {
		return ret, fmt.Errorf("%s extension cannot not be converted from %T", extFullName, ext)
	}
	return rule, nil
}

// setRuleFromDynamicMessage if each options are represented dynamicpb.Message type, convert and set it to rule instance.
// NOTE: compile proto files by compiler package, extension is replaced by dynamicpb.Message.
func setRuleFromDynamicMessage(opts proto.Message, extFullName protoreflect.FullName, rule proto.Message) bool {
	isSet := false
	opts.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if !fd.IsExtension() {
			return true
		}
		if fd.FullName() != extFullName {
			return true
		}
		ext := proto.GetExtension(opts, dynamicpb.NewExtensionType(fd))
		if ext == nil {
			return true
		}
		msg, ok := ext.(*dynamicpb.Message)
		if !ok {
			return true
		}
		bytes, err := proto.Marshal(msg)
		if err != nil {
			return true
		}
		if err := proto.Unmarshal(bytes, rule); err != nil {
			return true
		}

		isSet = true

		return true
	})
	return isSet
}
