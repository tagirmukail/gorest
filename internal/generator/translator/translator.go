package translator

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

type FieldType int

const (
	UnknownField FieldType = iota
	ArrayField
	BooleanField
	ComponentField
	CustomField
	DateField
	DateTimeField
	FileField
	FloatField
	FreeFormObject
	IntegerField
	StringField
	StructField
	UnixTimeField
)

type TypeDef struct {
	Name   string
	GoType string
	Fields []Field
	Fields2 []FieldI

	Level int
	Place string
}

func (d TypeDef) HasNoStringFields() bool {
	for _, f := range d.Fields {
		if f.Type != StringField {
			return true
		}
	}
	return false
}

func (d TypeDef) HasNoFileFields() bool {
	for _, f := range d.Fields {
		if f.Type != FileField {
			return true
		}
	}
	return false
}

func (d TypeDef) ContextErrorRequired() bool {
	for _, f := range d.Fields2 {
		if f.ContextErrorRequired() {
			return true
		}
	}
	return false
}

type FieldI interface {
	Name2() string
	StrVarName() string
	SecondsVarName() string
	IsRequired() bool
	CheckDefault() bool

	ContextErrorRequired() bool
	ImportsRequired() []string
	BuildCode() (string, error)
}

type ComponentFieldImpl struct {
	Field
}

func (c *ComponentFieldImpl) BuildCode() (string, error) {
	return "Test me ComponentFiledImpl", nil
}

type BooleanFieldImpl struct {
	Field
}

func (c *BooleanFieldImpl) BuildCode() (string, error) {
	var booleanFieldTemplate = template.Must(template.New("booleanField").Parse(
		`
         {{ .StrVarName }}, existed := c.Request.URL.Query().Get(key)
switch strings.ToLower({{ .StrVarName }}) {
    case "1", "true", "t":
        result.{{ .Name }} = true
    {{- if .IsRequired }}
    case "0", "false", "f":
        result.{{ .Name }} = false
    case "":
        errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "is absent", nil))
    default:
        errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as boolean", nil))
    {{- else }}
    case "0", "false", "f", "":
        result.{{ .Name }} = false
    default:
        errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as boolean", nil))
    {{- end }}
	}`))



	return "Test me BooleanFieldImpl", nil
}

func (c *BooleanFieldImpl) ContextErrorRequired() bool {
	return true
}

//func (c *ComponentFieldImpl) ImportsRequired() []string {
//	return []string{}
//}

// Field represents struct field
type Field struct {
	Name      string    // UserID
	GoType    string    // int64
	Parameter string    // user_id
	Required  bool
	Type      FieldType // IntegerField
	Schema    openapi3.SchemaType
}

func (f Field) Name2() string {
	return f.Name
}

func (f Field) ContextErrorRequired() bool {
	return false
}

func (f Field) ImportsRequired() []string {
	return []string{}
}

func (f Field) StrVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Str"
}

func (f Field) SecondsVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Sec"
}

func (f Field) IsStruct() bool {
	return f.Type == StructField
}

func (f Field) IsComponent() bool {
	return f.Type == ComponentField
}

func (f Field) IsCustom() bool {
	return f.Type == CustomField
}

func (f Field) IsArray() bool {
	return f.Type == ArrayField
}

func (f Field) IsBoolean() bool {
	return f.Type == BooleanField
}

func (f Field) IsString() bool {
	return f.Type == StringField
}

func (f Field) IsInteger() bool {
	return f.Type == IntegerField
}

func (f Field) IsFloat() bool {
	return f.Type == FloatField
}

func (f Field) IsDate() bool {
	return f.Type == DateField
}

func (f Field) IsDateTime() bool {
	return f.Type == DateTimeField
}

func (f Field) IsUnixTime() bool {
	return f.Type == UnixTimeField
}

func (f Field) IsFile() bool {
	return f.Type == FileField
}

func (f Field) IsRequired() bool {
	return f.Required
}

func (f Field) CheckDefault() bool {
	return f.Schema.Default != nil
}

func ProcessRootSchema(schema openapi3.SchemaType) ([]TypeDef, error) {
	queue := list.New()
	queue.PushBack(schema)

	result := make([]TypeDef, 0)
	for {
		el := queue.Back()
		if el == nil {
			break
		}
		queue.Remove(el)

		if sch, ok := el.Value.(openapi3.SchemaType); ok {
			def, err := ProcessObjSchema(sch, queue)
			if err != nil {
				return nil, err
			}
			result = append(result, def)
		} else {
			return nil, fmt.Errorf("unprocessible entity: %v", el.Value)
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result, nil
}

// Will build internal type definition by reading an open api Schema object. Queue will have all schemas, on which this
// scheme depends
func ProcessObjSchema(schema openapi3.SchemaType, queue *list.List) (def TypeDef, err error) {
	if schema.Type != openapi3.ObjectType {
		// TODO(a.telyshev): More complex processing
		// err = fmt.Errorf("schema must be `object`, got: `%s`", schema.Type)
		def.Name = MakeIdentifier(schema.Name)

		var goType string

		switch schema.Type {
		case openapi3.BooleanType:
			goType = "bool"
		case openapi3.IntegerType:
			goType = "int64"
		case openapi3.NumberType:
			goType = "float64"
		case openapi3.StringType:
			goType = "string"
		case openapi3.ArrayType:
			goType = "[]" + GetNameFromRef(schema.Items.Ref)
		default:
			err = fmt.Errorf("unsupported type: `%s` of schema %v", schema.Type, schema)
			return
		}

		def.GoType = goType
		return
	}

	def.Name = MakeIdentifier(schema.Name)
	def.GoType = "struct"

	requiredMap := make(map[string]bool)
	for _, propName := range schema.Required {
		requiredMap[propName] = true
	}

	for propName, propSchema := range schema.Properties {
		propID := MakeIdentifier(propName)
		propSchema.Name = propID

		var field Field
		field, err = determineType(def.Name, *propSchema, propName, queue)
		if err != nil {
			return
		}

		isRequired, ok := requiredMap[propName]
		if ok && isRequired {
			field.Required = true
		}

		var field2 FieldI
		field2, err = determineType2(def.Name, *propSchema, propName, isRequired, queue)
		if err != nil {
			return
		}
		if field2 != nil {
			def.Fields2 = append(def.Fields2, field2)
		}

		def.Fields = append(def.Fields, field)
	}
	sort.Slice(def.Fields, func(i, j int) bool {
		return def.Fields[i].Name < def.Fields[j].Name
	})
	sort.Slice(def.Fields2, func(i, j int) bool {
		return def.Fields2[i].Name2() < def.Fields2[j].Name2()
	})
	return
}


func determineType2(parentName string, schema openapi3.SchemaType, parameter string, required bool, queue *list.List) (FieldI, error) {
	if schema.Ref != "" {
		return &ComponentFieldImpl{
			Field {
					Type:      ComponentField,
					Name:      schema.Name,
					Parameter: parameter,
					GoType:    GetNameFromRef(schema.Ref),
					Required:  required,
			},
		}, nil
	}

	switch schema.Type {
	case openapi3.BooleanType:
		return &BooleanFieldImpl{
			Field{
				Type:      BooleanField,
				Name:      schema.Name,
				Parameter: parameter,
				GoType:    "bool",
				Schema:    schema,
			},
		}, nil
	}


	return nil, nil
}

func determineType(parentName string, schema openapi3.SchemaType, parameter string, queue *list.List) (Field, error) {
	if schema.Ref != "" {
		return Field{
			Type:      ComponentField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    GetNameFromRef(schema.Ref),
		}, nil
	}

	switch schema.Type {
	case openapi3.ArrayType:
		childName := parentName + MakeTitledIdentifier(schema.Name)
		t, err := determineType(childName, *schema.Items, parameter, queue)
		if err != nil {
			return Field{}, err
		}
		return Field{
			Type:      ArrayField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "[]" + t.GoType,
			Schema:    schema,
		}, nil

	case openapi3.BooleanType:
		return Field{
			Type:      BooleanField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "bool",
			Schema:    schema,
		}, nil

	case openapi3.IntegerType:
		switch schema.Format {
		case "":
			schema.BitSize = 0

		case openapi3.Integer32bit:
			schema.BitSize = 32

		case openapi3.Integer64bit:
			schema.BitSize = 64

		default:
			return Field{
				Type:      CustomField,
				Name:      schema.Name,
				Parameter: parameter,
				GoType:    MakeTitledIdentifier(string(schema.Format)),
				Schema:    schema,
			}, nil
		}
		return Field{
			Type:      IntegerField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "int64",
			Schema:    schema,
		}, nil

	case openapi3.NumberType:
		switch schema.Format {
		case "":
			schema.BitSize = 0

		case openapi3.NumberFloat:
			schema.BitSize = 32

		case openapi3.NumberDouble:
			schema.BitSize = 64

		default:
			return Field{
				Type:      CustomField,
				Name:      schema.Name,
				Parameter: parameter,
				GoType:    MakeTitledIdentifier(string(schema.Format)),
				Schema:    schema,
			}, nil
		}
		return Field{
			Type:      FloatField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "int64",
			Schema:    schema,
		}, nil

	case openapi3.ObjectType:
		noProperties := len(schema.ObjectSchema.Properties) == 0
		noAdditionalProperties := schema.AdditionalProperties == nil || len(schema.AdditionalProperties.Properties) == 0

		if noProperties && noAdditionalProperties {
			return Field{
				Type:      FreeFormObject,
				Parameter: parameter,
				Name:      schema.Name,
				GoType:    "json.RawMessage",
				Schema:    schema,
			}, nil
		}

		name := schema.Name
		goType := parentName + MakeTitledIdentifier(schema.Name)

		schema.Name = goType

		queue.PushBack(schema)
		return Field{
			Type:   StructField,
			Name:   name,
			GoType: goType,
			Schema: schema,
		}, nil

	case openapi3.StringType:
		fieldType := StringField
		goType := "string"

		switch schema.Format {
		case openapi3.Binary:
			fieldType = FileField
			goType = "*multipart.FileHeader"

		case openapi3.Date:
			fieldType = DateField
			goType = "time.Time" //nolint:goconst

		case openapi3.DateTime:
			fieldType = DateTimeField
			goType = "time.Time"

		case openapi3.UnixTime:
			fieldType = UnixTimeField
			goType = "time.Time"
		}

		return Field{
			Type:      fieldType,
			GoType:    goType,
			Name:      schema.Name,
			Parameter: parameter,
			Schema:    schema,
		}, nil

	default:
		return Field{}, fmt.Errorf("%s.%s: unknown data type: %v", parentName, parameter, schema.Type)
	}
}

func MakeIdentifier(s string) string {
	result := strcase.ToCamel(strings.ReplaceAll(s, " ", "_"))

	for _, suff := range [...]string{
		"Api",
		"Edo",
		"Db",
		"Http",
		"Id",
		"Inn",
		"Json",
		"Sql",
		"Uid",
		"Url",
	} {
		if strings.HasSuffix(result, suff) {
			result = result[:len(result)-len(suff)] + strings.ToUpper(suff)
			break
		}
	}
	return result
}

func MakeTitledIdentifier(s string) string {
	return strings.Title(MakeIdentifier(s))
}

func GetNameFromRef(s string) string {
	return s[len("#/components/schemas/"):]
}
