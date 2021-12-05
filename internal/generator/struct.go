package generator

import (
	"io"
	"mime"
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var primitiveTypeTemplate = template.Must(template.New("primitiveType").Parse(`
type {{ .Name }} {{ .GoType }}
`))

var structTemplate = template.Must(template.New("struct").Parse(`
type {{ .Name }} struct {
	{{- range $, $field := .Fields }}
	{{ $field.Name }} {{ $field.GoType -}}
	{{ end }}
}
`))

var structWithJSONTagsTemplate = template.Must(template.New("structWithJSON").Parse(`
{{ $mimeTypes := .MimeTypes }}
type {{ .Def.Name }} struct {
	{{- range $, $field := .Def.Fields -}}
	{{ $field.Name }} {{ $field.GoType -}} ` + "`{{- range $, $mType := $mimeTypes -}} {{ $mType }}:\"{{ $field.Parameter }}\" {{ end -}}`" + `
	{{ end -}}
}
`))

// TODO(a.telyshev): Test me
func (g *Generator) makeStruct(wr io.Writer, def translator.TypeDef, withJSONTags bool, contentTypes []string) error {
	if def.GoType != "struct" {
		return primitiveTypeTemplate.Execute(wr, def)
	}

	for _, f := range def.Fields {
		if f.IsCustom() {
			g.customFields[f.GoType] = f
		}
	}
	if withJSONTags {
		mimeTypes := g.contentTypesToMimeTypes(contentTypes)

		return structWithJSONTagsTemplate.Execute(wr, map[string]interface{}{
			"Def":       def,
			"MimeTypes": mimeTypes,
		})
	}
	return structTemplate.Execute(wr, def)
}

func (g *Generator) contentTypesToMimeTypes(ct []string) []string {
	var result = make([]string, 0)

	jsonExist := false
	for _, contentType := range ct {
		extensions, _ := mime.ExtensionsByType(contentType)
		for _, ext := range extensions {
			result = append(result, strings.Replace(ext, ".", "", 1))
		}

		if contentType == "application/json" {
			jsonExist = true
		}
	}

	if !jsonExist {
		result = append(result, "json")
	}

	return result
}
