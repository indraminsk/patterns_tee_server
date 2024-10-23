package v1

import (
	"net/http"
	"patterns/tee/server/internal/entity"
)

type Handler struct {
	ServMux *http.ServeMux
}

func NewHandler(logger entity.ILogger) Handler {
	handler := Handler{
		ServMux: http.NewServeMux(),
	}

	pingHandler := newPingHandler(logger)
	handler.ServMux.HandleFunc("GET /ping", pingHandler.ping)

	return handler
}
