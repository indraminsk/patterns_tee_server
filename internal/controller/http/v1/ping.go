package v1

import (
	"net/http"
	"patterns/tee/server/internal/entity"
)

type PingHandler struct {
	logger entity.ILogger
}

func newPingHandler(logger entity.ILogger) *PingHandler {
	return &PingHandler{
		logger: logger,
	}
}

func (r *PingHandler) ping(writer http.ResponseWriter, request *http.Request) {
	r.logger.Info("received ping request")

	writer.WriteHeader(http.StatusOK)

	_, err := writer.Write([]byte("pong"))
	if err != nil {
		r.logger.Error("ping request: something went wrong", err)
		return
	}
}
