package v1

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

const (
	host = "http://localhost:9001"
)

func TestPingPong(t *testing.T) {
	Test(t,
		Description("ping pong test"),
		Get(host+"/ping"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Equal("pong"),
	)
}
