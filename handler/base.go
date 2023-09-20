package handler

import (
	"encoding/json"
	"net/http"
	"context"

	"go_grpc/lib"
	"go_grpc"
	"go_grpc/model/proto/healthz"
	// "google.golang.org/protobuf/types/known/emptypb"
)

type handlerService struct {
	Backend *backend.Backend
}

func NewHandler(backend *backend.Backend) handlerService {
	return handlerService{
		Backend: backend,
	}
}

func (h *handlerService) Healthz(ctx context.Context, req *healthz.HealthCheckRequest) (*healthz.HealthCheckResponse, error) {
	// reference https://github.com/go-training/grpc-health-check/blob/master/main.go
	return &healthz.HealthCheckResponse{Status: healthz.HealthCheckResponse_SERVING}, nil
}

type ResponseBody struct {
	Data    interface{}  `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
	Meta    ResponseMeta `json:"meta"`
}

type ResponseMeta struct {
	HTTPStatus int   `json:"http_status"`
	Total      *uint `json:"total,omitempty"`
	Offset     *uint `json:"offset,omitempty"`
	Limit      *uint `json:"limit,omitempty"`
	Page       *uint `json:"page,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Field   string `json:"field,omitempty"`
}

type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

func writeError(w http.ResponseWriter, err error) {
	var resp interface{}
	code := http.StatusInternalServerError

	switch errOrig := err.(type) {
	case lib.CustomError:
		resp = ErrorBody{
			Errors: []ErrorInfo{
				{
					Message: errOrig.Message,
					Code:    errOrig.Code,
					Field:   errOrig.Field,
				},
			},
			Meta: ResponseMeta{
				HTTPStatus: errOrig.HTTPCode,
			},
		}

		code = errOrig.HTTPCode
	default:
		resp = ResponseBody{
			Message: "Internal Server Error",
			Meta: ResponseMeta{
				HTTPStatus: code,
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func writeSuccess(w http.ResponseWriter, data interface{}, message string, meta ResponseMeta) {
	resp := ResponseBody{
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.HTTPStatus)
	json.NewEncoder(w).Encode(resp)
}

func writeResponse(w http.ResponseWriter, resp interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
