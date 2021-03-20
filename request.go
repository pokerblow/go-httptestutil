package httptestutil

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestRouter struct {
	router       *gin.Engine
	basePath     string
	extraHeaders map[string]string
}

func (r *TestRouter) BasePath(basePath string) *TestRouter {
	r.basePath = basePath
	return r
}

// Adds extra headers to the requests
func (r *TestRouter) Headers(headers map[string]string) *TestRouter {
	r.extraHeaders = headers
	return r
}

func NewRouter(router *gin.Engine) *TestRouter {
	if router == nil {
		log.Fatal("router parameter can't be nil")
	}
	return &TestRouter{router: router}
}

func (r *TestRouter) Request(t *testing.T, method, path string, bodyJson *string, headers map[string]string) *httptest.ResponseRecorder {
	fullPath := fmt.Sprintf("%s%s", r.basePath, path)
	req, err := http.NewRequest(method, fullPath, body(bodyJson))
	if err != nil {
		t.Fatal(err)
	}

	for headerName, v := range r.extraHeaders {
		req.Header.Set(headerName, v)
	}

	for headerName, v := range headers {
		req.Header.Set(headerName, v)
	}

	rr := httptest.NewRecorder()
	r.router.ServeHTTP(rr, req)
	return rr
}

func (r *TestRouter) GET(t *testing.T, path string, headers map[string]string) *httptest.ResponseRecorder {
	return r.Request(t, http.MethodGet, path, nil, headers)
}

func (r *TestRouter) POST(t *testing.T, path string, bodyJson string, headers map[string]string) *httptest.ResponseRecorder {
	return r.Request(t, http.MethodPost, path, &bodyJson, headers)
}

func (r *TestRouter) PUT(t *testing.T, path string, bodyJson string, headers map[string]string) *httptest.ResponseRecorder {
	return r.Request(t, http.MethodPut, path, &bodyJson, headers)
}

func (r *TestRouter) PATCH(t *testing.T, path string, bodyJson string, headers map[string]string) *httptest.ResponseRecorder {
	return r.Request(t, http.MethodPatch, path, &bodyJson, headers)
}

func (r *TestRouter) DELETE(t *testing.T, path string, headers map[string]string) *httptest.ResponseRecorder {
	return r.Request(t, http.MethodPut, path, nil, headers)
}

func body(json *string) io.Reader {
	if json == nil {
		return nil
	}
	return bytes.NewBuffer([]byte(*json))
}
