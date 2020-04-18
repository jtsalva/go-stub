package main

import "fmt"

type Query map[string]string
type Headers map[string]string

type Request struct {
	Method []string `yaml:"method"`
	Url string `yaml:"url"`
	Query Query `yaml:"query"`
	Headers Headers `yaml:"headers"`
}

func (req *Request) QueryPairs() []string {
	var queryPairs []string
	for key, value := range req.Query {
		queryPairs = append(queryPairs, key, value)
	}
	return queryPairs
}

func (req *Request) HeaderPairs() []string {
	var headerPairs []string
	for key, value := range req.Headers {
		headerPairs = append(headerPairs, key, value)
	}
	return headerPairs
}

func (req *Request) String() string {
	return fmt.Sprintf("%v %s", req.Method, req.Url)
}

type Response struct {
	Headers Headers `yaml:"headers"`
	Latency int `yaml:"latency"`
	Body string `yaml:"body"`
	Status int `yaml:"status"`
}

func (res *Response) String() string {
	return fmt.Sprintf("%d", res.Status)
}

type Stub struct {
	Request Request `yaml:"request"`
	Response Response `yaml:"response"`
}

func (s *Stub) String() string {
	return fmt.Sprintf("%s -> %s", s.Request.String(), s.Response.String())
}