package gorequests

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	cookiejar "github.com/chyroc/persistent-cookiejar"
)

type Request struct {
	// internal
	cachedurl     string
	persistentJar *cookiejar.Jar
	lock          sync.RWMutex
	err           error
	logger        Logger

	// request
	context      context.Context     // request context
	isIgnoreSSL  bool                // request  ignore ssl verify
	header       http.Header         // request header
	querys       map[string][]string // request query
	isNoRedirect bool                // request ignore redirect
	timeout      time.Duration       // request timeout
	url          string              // request url
	method       string              // request method
	rawBody      []byte              // []byte of body
	body         io.Reader           // request body
	fullUrl      string

	// resp
	resp      *http.Response
	bytes     []byte
	isRead    bool
	isRequest bool
}

func New(method, url string) *Request {
	r := &Request{
		url:     url,
		method:  method,
		header:  map[string][]string{},
		querys:  make(map[string][]string),
		context: context.TODO(),
		logger:  NewStdoutLogger(),
	}
	return r
}

func (r *Request) SetError(err error) *Request {
	r.err = err
	return r
}
