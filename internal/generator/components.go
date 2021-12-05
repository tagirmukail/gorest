package generator

import (
	"io"
	"sort"
	"strings"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func (g *Generator) fetchComponentName(content openapi3.ContentType) string {
	if content.Schema == nil {
		return ""
	}

	if content.Schema.Ref == "" {
		return ""
	}

	return strings.Replace(content.Schema.Ref, "#/components/schemas/", "", 1)
}

func (g *Generator) defineComponentContentTypes(pathMap openapi3.PathMap) map[string][]string {
	var componentToContentTypes = make(map[string]map[string]struct{})

	for _, r := range pathMap {
		var (
			postContentMap  map[openapi3.MimeType]openapi3.ContentType
			putContentMap   map[openapi3.MimeType]openapi3.ContentType
			patchContentMap map[openapi3.MimeType]openapi3.ContentType
		)

		if r.Post != nil && r.Post.RequestBody != nil {
			postContentMap = r.Post.RequestBody.Content
		}

		if r.Put != nil && r.Put.RequestBody != nil {
			putContentMap = r.Put.RequestBody.Content

		}

		if r.Patch != nil && r.Patch.RequestBody != nil {
			patchContentMap = r.Patch.RequestBody.Content
		}

		for contentType, content := range postContentMap {
			componentName := g.fetchComponentName(content)
			if componentName == "" {
				continue
			}

			_, ok := componentToContentTypes[componentName]
			if !ok {
				componentToContentTypes[componentName] = make(map[string]struct{})
			}

			componentToContentTypes[componentName][contentType] = struct{}{}
		}

		for contentType, content := range putContentMap {
			componentName := g.fetchComponentName(content)
			if componentName == "" {
				continue
			}

			_, ok := componentToContentTypes[componentName]
			if !ok {
				componentToContentTypes[componentName] = make(map[string]struct{})
			}

			componentToContentTypes[componentName][contentType] = struct{}{}
		}

		for contentType, content := range patchContentMap {
			componentName := g.fetchComponentName(content)
			if componentName == "" {
				continue
			}

			_, ok := componentToContentTypes[componentName]
			if !ok {
				componentToContentTypes[componentName] = make(map[string]struct{})
			}

			componentToContentTypes[componentName][contentType] = struct{}{}
		}
	}

	var result = make(map[string][]string)
	for component, contentTypes := range componentToContentTypes {
		for contentType := range contentTypes {
			result[component] = append(result[component], contentType)
		}
	}

	return result
}

func (g *Generator) makeComponents(wr io.Writer, sp openapi3.Spec) error {
	sortedComponents := make([]string, 0, len(sp.Components.Schemas))

	componentToContentTypes := g.defineComponentContentTypes(sp.Paths)

	for k := range sp.Components.Schemas {
		sortedComponents = append(sortedComponents, k)
	}

	sort.Strings(sortedComponents)

	for _, name := range sortedComponents {
		schema := sp.Components.Schemas[name]
		schema.Name = name

		defs, err := translator.ProcessRootSchema(*schema)
		if err != nil {
			return err
		}

		contentTypes := componentToContentTypes[name]

		for _, d := range defs {
			if err := g.makeStruct(wr, d, true, contentTypes); err != nil {
				return err
			}

			if err := g.makeValidateFunc(wr, d); err != nil {
				return err
			}
		}
	}

	return nil
}
