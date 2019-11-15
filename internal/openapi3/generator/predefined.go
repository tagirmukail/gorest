package generator

const Predefined = `// Code generated by gorest; DO NOT EDIT.

package %s

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const handlerNameKey = "handler"

type ParamPlace int

const (
	Query ParamPlace = iota
	Path
	Cookie
	Header
)

type Response struct {
	Code    int              ` + "`" + `json:"code"` + "`" + `
	Message string           ` + "`" + `json:"message,omitempty"` + "`" + `
	Data    *json.RawMessage ` + "`" + `json:"data,omitempty"` + "`" + `
	Errors  []FieldError     ` + "`" + `json:"errors,omitempty"` + "`" + `
}

type FieldError struct {
	Field   string ` + "`" + `json:"field"` + "`" + `
	Message string ` + "`" + `json:"message"` + "`" + `
	Reason  string ` + "`" + `json:"reason"` + "`" + `
}

func ExtractParameter(c *gin.Context, paramName string, from ParamPlace) string {
	switch from {
	case Query:
		return c.Query(paramName)

	case Path:
		return c.Param(paramName)

	case Cookie:
		cookie, err := c.Request.Cookie(paramName)
		if err == http.ErrNoCookie {
			return ""
		}
		return cookie.Value

	case Header:
		return c.Request.Header.Get(paramName)

	default:
		panic(fmt.Sprintf("extract '%s': unknown param place: %v", paramName, from))
	}
}

func ExtractParameterWithDefault(c *gin.Context, paramName string, from ParamPlace, defaultValue string) interface{} {
	switch from {
	case Query:
		return c.DefaultQuery(paramName, defaultValue)

	case Path:
		pathParam, ok := c.Params.Get(paramName)
		if !ok {
			return defaultValue
		}
		return pathParam

	case Cookie:
		cookie, err := c.Request.Cookie(paramName)
		if err == http.ErrNoCookie  {
			return defaultValue
		}
		return cookie.Value

	case Header:
		header := c.Request.Header.Get(paramName)
		if header == "" {
			return defaultValue
		}
		return header

	default:
		panic(fmt.Sprintf("extract '%s' with default: unknown param place: %v", paramName, from))
	}
}`
