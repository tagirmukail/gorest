// Code generated by gorest; DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const handlerNameKey = "handler"

type ParamPlace int

const (
	UndefinedPlace ParamPlace = iota
	InBody
	InCookie
	InHeader
	InPath
	InQuery
)

type ContentType int

const (
	UndefinedContentType ContentType = iota
	AppJSON
	AppXML
	AppFormUrlencoded
	TextPlain
)

type FieldError struct {
	In      ParamPlace
	Field   string
	Message string
	Reason  error
}

func NewFieldError(in ParamPlace, f string, msg string, err error) FieldError {
	return FieldError{
		In:      in,
		Field:   f,
		Message: msg,
		Reason:  err,
	}
}

type PaymentGatewayAPI interface {
	// GET /v1/example/:year/:user
	Example(in ExampleRequest, c *gin.Context)

	// GET /v1/payment
	GetPayment(in GetPaymentRequest, c *gin.Context)
	// POST /v1/payment
	ProvidePayment(in ProvidePaymentRequest, c *gin.Context)

	// Service methods
	ProcessMakeRequestErrors(errors []FieldError)
	ProcessValidateErrors(errors []FieldError)
}

type PaymentGatewayAPIServer struct {
	Srv PaymentGatewayAPI
}

// _PaymentGatewayAPI_GetPayment_Handler

type GetPaymentRequest struct {
	Query GetPaymentRequestQuery
}

func (t GetPaymentRequest) Validate() []FieldError {
	return nil
}

type GetPaymentRequestQuery struct {
	ID string
}

func (t GetPaymentRequestQuery) Validate() []FieldError {
	return nil
}

func MakeGetPaymentRequest(c *gin.Context) (result GetPaymentRequest, errors []FieldError) {
	result.Query, errors = MakeGetPaymentRequestQuery(c)
	if errors != nil {
		return
	}
	return
}

func MakeGetPaymentRequestQuery(c *gin.Context) (result GetPaymentRequestQuery, errors []FieldError) {
	result.ID, _ = c.GetQuery("id")
	return
}

func (server PaymentGatewayAPIServer) _PaymentGatewayAPI_GetPayment_Handler(c *gin.Context) {
	c.Set(handlerNameKey, "GetPayment")

	req, errors := MakeGetPaymentRequest(c)
	if len(errors) > 0 {
		server.Srv.ProcessMakeRequestErrors(errors)
		return
	}

	errors = req.Validate()
	if len(errors) > 0 {
		server.Srv.ProcessValidateErrors(errors)
		return
	}

	server.Srv.GetPayment(req, c)
}

// _PaymentGatewayAPI_ProvidePayment_Handler

type ProvidePaymentRequest struct {
	Body ProvidePaymentRequestBody
}

func (t ProvidePaymentRequest) Validate() []FieldError {
	return nil
}

type ProvidePaymentRequestBody struct {
	JSON Payment
	Type ContentType
}

func (t ProvidePaymentRequestBody) Validate() []FieldError {
	return nil
}

func MakeProvidePaymentRequest(c *gin.Context) (result ProvidePaymentRequest, errors []FieldError) {
	result.Body, errors = MakeProvidePaymentRequestBody(c)
	if errors != nil {
		return
	}
	return
}

func MakeProvidePaymentRequestBody(c *gin.Context) (result ProvidePaymentRequestBody, errors []FieldError) {
	switch c.Request.Header.Get("Content-Type") {
	case "application/json":
		result.Type = AppJSON
		if err := json.NewDecoder(c.Request.Body).Decode(result.JSON); err != nil {
			errors = append(errors, NewFieldError(InBody, "requestBody", "can't decode body from JSON", err))
		}
	}
	return
}

func (server PaymentGatewayAPIServer) _PaymentGatewayAPI_ProvidePayment_Handler(c *gin.Context) {
	c.Set(handlerNameKey, "ProvidePayment")

	req, errors := MakeProvidePaymentRequest(c)
	if len(errors) > 0 {
		server.Srv.ProcessMakeRequestErrors(errors)
		return
	}

	errors = req.Validate()
	if len(errors) > 0 {
		server.Srv.ProcessValidateErrors(errors)
		return
	}

	server.Srv.ProvidePayment(req, c)
}

// _PaymentGatewayAPI_Example_Handler

type ExampleRequest struct {
	Path  ExampleRequestPath
	Query ExampleRequestQuery
}

func (t ExampleRequest) Validate() []FieldError {
	return nil
}

type ExampleRequestPath struct {
	User ExampleRequestPathUser
	Year int64
}

func (t ExampleRequestPath) Validate() []FieldError {
	return nil
}

type ExampleRequestPathUser struct {
	FirstName string
	Role      string
}

func (t ExampleRequestPathUser) Validate() []FieldError {
	return nil
}

type ExampleRequestQuery struct {
	From     time.Time
	FromDate time.Time
	Sum      Decimal
	To       time.Time
}

func (t ExampleRequestQuery) Validate() []FieldError {
	return nil
}

func MakeExampleRequest(c *gin.Context) (result ExampleRequest, errors []FieldError) {
	result.Path, errors = MakeExampleRequestPath(c)
	if errors != nil {
		return
	}

	result.Query, errors = MakeExampleRequestQuery(c)
	if errors != nil {
		return
	}
	return
}

func MakeExampleRequestQuery(c *gin.Context) (result ExampleRequestQuery, errors []FieldError) {
	var err error

	fromStr, _ := c.GetQuery("from")
	result.From, err = time.Parse(time.RFC3339, fromStr)
	if err != nil {
		errors = append(errors, NewFieldError(InQuery, "from", "can't parse as RFC3339 time", err))
	}

	fromDateStr, _ := c.GetQuery("fromDate")
	result.FromDate, err = time.Parse(time.RFC3339, fromDateStr)
	if err != nil {
		errors = append(errors, NewFieldError(InQuery, "fromDate", "can't parse as RFC3339 time", err))
	}

	sumStr, _ := c.GetQuery("sum")
	result.Sum = Decimal{}
	if err = result.Sum.SetFromString(sumStr); err != nil {
		errors = append(errors, NewFieldError(InQuery, "sum", fmt.Sprintf("can't create from string '%s'", sumStr), err))
	}

	toStr, _ := c.GetQuery("to")
	toSec, err := strconv.ParseInt(toStr, 10, 64)
	if err != nil {
		errors = append(errors, NewFieldError(InQuery, "to", "can't parse as 64 bit integer", err))
	} else {
		result.To = time.Unix(toSec, 0)
	}
	return
}

func MakeExampleRequestPath(c *gin.Context) (result ExampleRequestPath, errors []FieldError) {
	var err error

	yearStr, _ := c.Params.Get("year")
	result.Year, err = strconv.ParseInt(yearStr, 10, 64)
	if err != nil {
		errors = append(errors, NewFieldError(InPath, "year", "can't parse as 64 bit integer", err))
	}
	return
}

func (server PaymentGatewayAPIServer) _PaymentGatewayAPI_Example_Handler(c *gin.Context) {
	c.Set(handlerNameKey, "Example")

	req, errors := MakeExampleRequest(c)
	if len(errors) > 0 {
		server.Srv.ProcessMakeRequestErrors(errors)
		return
	}

	errors = req.Validate()
	if len(errors) > 0 {
		server.Srv.ProcessValidateErrors(errors)
		return
	}

	server.Srv.Example(req, c)
}

// Router
func RegisterRoutes(r *gin.Engine, api PaymentGatewayAPI) {
	e := &PaymentGatewayAPIServer{api}

	r.Handle("GET", "/v1/example/:year/:user", e._PaymentGatewayAPI_Example_Handler)

	r.Handle("GET", "/v1/payment", e._PaymentGatewayAPI_GetPayment_Handler)
	r.Handle("POST", "/v1/payment", e._PaymentGatewayAPI_ProvidePayment_Handler)
}

type Payment struct {
	MerchantID string          `json:"merchant_id"`
	Meta       json.RawMessage `json:"meta"`
	PaymentID  ID              `json:"payment_id"`
	Sum        Decimal         `json:"sum"`
}

type ID string

type Payments []Payment

// Custom types

type FromStringSetter interface {
	SetFromString(string) error
}

var _ json.Marshaler = (*Decimal)(nil)
var _ json.Unmarshaler = (*Decimal)(nil)
var _ FromStringSetter = (*Decimal)(nil)
