package httptestutil

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	if router == nil {
		log.Fatal("router parameter can't be nil")
	}
	return &Router{router: router}
}

func (r *Router) Request(t *testing.T, method, path string, bodyJson *string, headers map[string]string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, body(bodyJson))
	if err != nil {
		t.Fatal(err)
	}

	for header, v := range headers {
		req.Header.Set(header, v)
	}

	rr := httptest.NewRecorder()
	r.router.ServeHTTP(rr, req)
	return rr
}

func body(json *string) io.Reader {
	if json == nil {
		return nil
	}
	return bytes.NewBuffer([]byte(*json))
}
