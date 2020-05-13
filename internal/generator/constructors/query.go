package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/constructors/fields"
	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeQueryParamsConstructor receive a query params struct definition and generate corresponding constructor
func MakeQueryParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return queryParamsConstructorTemplate.Execute(wr, def)
}

var queryParamsConstructorTemplate = template.Must(template.New("queryParamsConstructor").Funcs(fields.BaseConstructor).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
    {{- else if .ContextErrorRequired }}
	var err error
	{{ end }}

    {{- range .Fields2 }}
		{{ .BuildCode }}
    {{ end -}}

	{{- range $, $field := .Fields }}
	{{- with $field }}
        {{- if or .CheckDefault .IsRequired}}
        {{ .StrVarName }}, ok := c.GetQuery("{{ .Parameter }}")
        if !ok {
           {{- if .CheckDefault }}
           {{ .StrVarName }} = "{{ .Schema.Default }}"
           {{- else }}
           errors = append(errors, NewFieldError(InPath, "{{ .Parameter }}", "is absent", nil))
           {{- end }}
        }
        {{- else }}
        {{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
        {{- end }}

		{{- BaseValueFieldConstructor . "InQuery" }}

	{{- end -}}
	{{ end -}}
	return
}
`))
