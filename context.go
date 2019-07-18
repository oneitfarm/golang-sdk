package golang_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type CallConfig struct {
	Method string
	Path   string
	Header *http.Header
	Body   io.Reader
}

type CallOption func(config *CallConfig)

type ApiContext struct {
	requestContext RequestContext
}

func (ctx *ApiContext) init() {

}

func (ctx *ApiContext) ReturnOk(data interface{}) {
	ctx.Return(&Result{
		State: 1,
		Data:  data,
		Msg:   "ok",
	})
}

func (ctx *ApiContext) ReturnError(state int, msg string) {
	ctx.Return(&Result{
		State: state,
		Msg:   msg,
	})
}

func (ctx *ApiContext) Return(result *Result) {
	var err error
	data, err := json.Marshal(result)
	if err != nil {
		panic(fmt.Errorf("json marshal error:%v", err))
	}
	ctx.requestContext.GetWriter().Header().Set("Content-Type", "application/json")
	_, err = ctx.requestContext.GetWriter().Write(data)
	if err != nil {
		panic(fmt.Errorf("write response error:%v", err))
	}
}

func (ctx *ApiContext) Get(service string, api string) *Result {
	return ctx.CallWithOptions(service, WithApi(api), WithMethod("get"))
}

func (ctx *ApiContext) PostForm(service string, api string, form map[string]string) *Result {
	return ctx.CallWithOptions(service, WithApi(api), WithMethod("post"), WithForm(form))
}

func (ctx *ApiContext) PostJson(service string, api string, jsonData map[string]string) *Result {
	return ctx.CallWithOptions(service, WithApi(api), WithMethod("post"), WithJson(jsonData))
}

func (ctx *ApiContext) CallWithOptions(service string, options ...CallOption) *Result {
	config := &CallConfig{}
	for _, opt := range options {
		opt(config)
	}
	return ctx.Call(service, config)
}

func (ctx *ApiContext) Call(service string, config *CallConfig) *Result {
	serviceUrl := GetSdk().GetServiceUrl(service)

	if serviceUrl == "" {
		return ErrorResult(ErrServiceNotFound, fmt.Sprintf("service %s not found", service))
	}

	request, err := http.NewRequest(config.Method, fmt.Sprintf("%s/%s", serviceUrl, config.Path), config.Body)
	if err != nil {
		return ErrorResult(ErrUnknownError, err.Error())
	}

	for k, v := range *config.Header {
		request.Header[k] = v
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return ErrorResult(ErrUnknownError, "make request error")
	}

	statusCode := response.StatusCode

	if statusCode < 200 || statusCode >= 400 {
		switch statusCode {
		case http.StatusUnauthorized:
			return ErrorResult(ErrResponse401, "Unauthorized")
		case http.StatusForbidden:
			return ErrorResult(ErrResponse403, "No permission")
		case http.StatusNotFound:
			return ErrorResult(ErrResponse404, "api not exist")
		default:
			return ErrorResult(ErrResponseOther, fmt.Sprintf("http request err:%s", response.Status))
		}
	}

	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ErrorResult(ErrResponseOther, fmt.Sprintf("read response error:%v", err))
	}

	result := &Result{}
	err = json.Unmarshal(bodyData, &response)
	if err != nil {
		return ErrorResult(ErrResponseContentTypeError, fmt.Sprintf("invalid json format"))
	}

	return result
}

type RequestContext interface {
	GetRequest() *http.Request
	GetWriter() http.ResponseWriter
}

func WithApi(api string) CallOption {
	return func(config *CallConfig) {
		config.Path = api
	}
}

func WithMethod(method string) CallOption {
	return func(config *CallConfig) {
		config.Method = method
	}
}

func WithJson(object interface{}) CallOption {
	return func(config *CallConfig) {
		config.Header.Set("Content-Type", "application/json")

		data, err := json.Marshal(object)
		if err != nil {
			return
		}
		config.Body = bytes.NewReader(data)
	}
}

func WithForm(form map[string]string) CallOption {
	return func(config *CallConfig) {
		config.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		data := url.Values{}
		for k, v := range form {
			data[k] = []string{v}
		}

		config.Body = strings.NewReader(data.Encode())
	}
}

func WithUploadFile(fileName string, fileData io.Reader) CallOption {
	return func(config *CallConfig) {
		// TODO: handle upload file
	}
}
