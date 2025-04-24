package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseStatus string

const (
	StatusSuccess ResponseStatus = "success"
	StatusError   ResponseStatus = "error"
)

type Response struct {
	Status    ResponseStatus `json:"status,omitempty"`
	Message   string         `json:"message,omitempty"`
	Data      interface{}    `json:"data,omitempty"`
	Errors    interface{}    `json:"errors,omitempty"`
	Meta      *Meta          `json:"meta,omitempty"`
	Timestamp string         `json:"timestamp"`
}

type Meta struct {
	CurrentPage  int   `json:"current_page,omitempty"`
	From         int   `json:"from,omitempty"`
	LastPage     int   `json:"last_page,omitempty"`
	PerPage      int   `json:"per_page,omitempty"`
	To           int   `json:"to,omitempty"`
	Total        int64 `json:"total,omitempty"`
	TotalPages   int   `json:"total_pages,omitempty"`
	TotalRecords int64 `json:"total_records,omitempty"`
}

type ResponseBuilder struct {
	response Response
	err      error
}

func NewResponse() *ResponseBuilder {
	return &ResponseBuilder{
		response: Response{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}
}

func (rb *ResponseBuilder) WithStatus(status ResponseStatus) *ResponseBuilder {
	rb.response.Status = status
	return rb
}

func (rb *ResponseBuilder) WithMessage(message string) *ResponseBuilder {
	rb.response.Message = message
	return rb
}

func (rb *ResponseBuilder) WithData(data interface{}) *ResponseBuilder {
	rb.response.Data = data
	return rb
}

func (rb *ResponseBuilder) WithErrors(errors interface{}) *ResponseBuilder {
	rb.response.Errors = errors
	return rb
}

func (rb *ResponseBuilder) WithMeta(meta *Meta) *ResponseBuilder {
	rb.response.Meta = meta
	return rb
}

func (rb *ResponseBuilder) Send(w http.ResponseWriter, status int) error {
	if rb.err != nil {
		return fmt.Errorf("failed to send response: %w", rb.err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(rb.response); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}

	return nil
}

func Success(w http.ResponseWriter, data interface{}, message string) error {
	return NewResponse().
		WithStatus(StatusSuccess).
		WithMessage(message).
		WithData(data).
		Send(w, http.StatusOK)
}

func Created(w http.ResponseWriter, data interface{}, message string) error {
	return NewResponse().
		WithStatus(StatusSuccess).
		WithMessage(message).
		WithData(data).
		Send(w, http.StatusCreated)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func Error(w http.ResponseWriter, status int, message string, errors interface{}) error {
	return NewResponse().
		WithStatus(StatusError).
		WithMessage(message).
		WithErrors(errors).
		Send(w, status)
}

func BadRequest(w http.ResponseWriter, message string, errors interface{}) error {
	return Error(w, http.StatusBadRequest, message, errors)
}

func Unauthorized(w http.ResponseWriter, message string) error {
	return Error(w, http.StatusUnauthorized, message, nil)
}

func Forbidden(w http.ResponseWriter, message string) error {
	return Error(w, http.StatusForbidden, message, nil)
}

func NotFound(w http.ResponseWriter, message string) error {
	return Error(w, http.StatusNotFound, message, nil)
}

func ValidationError(w http.ResponseWriter, message string, errors interface{}) error {
	return Error(w, http.StatusUnprocessableEntity, message, errors)
}

func InternalServerError(w http.ResponseWriter, message string) error {
	return Error(w, http.StatusInternalServerError, message, nil)
}

func ServiceUnavailable(w http.ResponseWriter, message string) error {
	return Error(w, http.StatusServiceUnavailable, message, nil)
}

func PaginatedResponse(w http.ResponseWriter, data interface{}, meta *Meta, message string) error {
	return NewResponse().
		WithStatus(StatusSuccess).
		WithMessage(message).
		WithData(data).
		WithMeta(meta).
		Send(w, http.StatusOK)
}
