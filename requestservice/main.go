package requestservice

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var RequesterService RequesterServiceinterface = &requesterService{}

type RequesterServiceinterface interface {
	Requester(method, path string, body io.Reader) (*http.Response, error)
	SetMethod(method string) error
	SetUrl(urls string) error
	SetPath(path string)
	SetBody(body io.Reader)
	GetMethod() string
	GetUrl() string
	GetPath() string
	GetBody() io.Reader
}
type requesterService struct {
	method string
	url    string
	path   string
	body   io.Reader
}

func NewRequesterService(url string) RequesterServiceinterface {
	var nr requesterService
	if err := nr.SetUrl(url); err != nil {
		log.Fatal("something is wrong with url", url)
	}
	return &nr

}
func (r *requesterService) Requester(method, path string, body io.Reader) (*http.Response, error) {
	if err := r.SetMethod(method); err != nil {
		return nil, err
	}
	r.SetPath(path)
	r.SetBody(body)
	req, err := http.NewRequest(r.GetMethod(), r.GetUrl()+r.GetPath(), r.GetBody())
	if err != nil {
		log.Printf("something went wrong with the %s of %s request", r.GetUrl(), r.GetPath())
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
func (r *requesterService) SetMethod(method string) error {
	// validate methods
	var error error
	// if method != "POST" || method != "GET" || method != "DELETE" || method != "PUT" {
	// 	error = fmt.Errorf("wrong method atached")
	// 	return error
	// }
	r.method = method
	return error
}
func (r *requesterService) SetUrl(urls string) error {
	// validate url
	var error error
	_, err := url.Parse(urls)
	if err != nil {
		error = fmt.Errorf("wrong method atached")
		return error
	}
	r.url = urls
	return error
}
func (r *requesterService) SetPath(path string) {
	// validate path
	r.path = path
}
func (r *requesterService) SetBody(body io.Reader) {
	// validate body
	r.body = body
}
func (r *requesterService) GetMethod() string {
	return r.method
}
func (r *requesterService) GetUrl() string {
	return r.url
}
func (r *requesterService) GetPath() string {
	return r.path
}
func (r *requesterService) GetBody() io.Reader {
	return r.body
}
