package http

import (
	"crypto/tls"
	"fmt"
	"kaya-backend/utils/constants"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/parnurzeal/gorequest"
)

// Env ..
type Env struct {
	DebugClient bool   `envconfig:"DEBUG_CLIENT" default:"true"`
	Timeout     string `envconfig:"TIMEOUT" default:"60s"`
	RetryBad    int    `envconfig:"RETRY_BAD" default:"1"`
}

var (
	httpEnv Env
)

func init() {
	if err := envconfig.Process("HTTP", &httpEnv); err != nil {
		fmt.Println("Failed to get HTTP env:", err)
	}
}

// HTTPGet func
func HTTPGet(url string, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	reqagent := request.Get(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPost func
func HTTPPost(url string, jsondata interface{}) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPostWithHeader func
func HTTPPostWithHeader(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	reqagent := request.Post(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPutWithHeader func
func HTTPPutWithHeader(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Put(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPDeleteWithHeader func
func HTTPDeleteWithHeader(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Delete(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// SendHttpRequest ..
func SendHttpRequest(method string, url string, header http.Header, body interface{}) ([]byte, error) {
	var data []byte
	var err error
	switch method {
	case constants.HttpMethodGet:
		data, err = HTTPGet(url, header)
	case constants.HttpMethodPost:
		data, err = HTTPPostWithHeader(url, body, header)
	case constants.HttpMethodPut:
		data, err = HTTPPutWithHeader(url, body, header)
	case constants.HttpMethodDelete:
		data, err = HTTPDeleteWithHeader(url, body, header)
	}
	return data, err
}
