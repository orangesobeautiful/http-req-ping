package main

import (
	"crypto/tls"
	"encoding/json"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

func main() {
	const readTimeoutDuration = 10 * time.Second

	s := http.Server{
		Addr:        ":8080",
		Handler:     &receviceAllHandler{},
		ReadTimeout: readTimeoutDuration,
	}

	slog.Info("start server", slog.String("addr", s.Addr))
	err := s.ListenAndServe()
	if err != nil {
		slog.Error("ListenAndServe: ", slog.Any("err", err))
	}
}

type receviceAllHandler struct{}

type RequestInfo struct {
	Method        string
	URLPath       string
	Proto         string
	ProtoMajor    int
	ProtoMinor    int
	Header        http.Header
	ContentLength int64
	Host          string
	Form          url.Values
	PostForm      url.Values
	MultipartForm *multipart.Form
	Trailer       http.Header
	RemoteAddr    string
	RequestURI    string
	TLS           *tls.ConnectionState
}

func newRequestInfo(r *http.Request) *RequestInfo {
	return &RequestInfo{
		Method:        r.Method,
		URLPath:       r.URL.String(),
		Proto:         r.Proto,
		ProtoMajor:    r.ProtoMajor,
		ProtoMinor:    r.ProtoMinor,
		Header:        r.Header,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		Form:          r.Form,
		PostForm:      r.PostForm,
		MultipartForm: r.MultipartForm,
		Trailer:       r.Trailer,
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		TLS:           r.TLS,
	}
}

func (*receviceAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("receive request", slog.String("Method", r.Method), slog.String("URL", r.URL.String()))
	reqInfo := newRequestInfo(r)

	res, err := json.MarshalIndent(reqInfo, "", "  ")
	if err != nil {
		slog.Error("json.MarshalIndent: ", slog.Any("err", err))
		return
	}

	_, err = w.Write(res)
	if err != nil {
		slog.Error("w.Write: ", slog.Any("err", err))
		return
	}
}
