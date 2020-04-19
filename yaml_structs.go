package main

import (
	"fmt"
)

type KeyValuePairs map[string]string

type Request struct {
	Method  []string      `yaml:"method"`
	Url     string        `yaml:"url"`
	Query   KeyValuePairs `yaml:"query"`
	Headers KeyValuePairs `yaml:"headers"`
}

func (req *Request) String() string {
	return fmt.Sprintf("%v %s", req.Method, req.Url)
}

type Response struct {
	Headers KeyValuePairs `yaml:"headers"`
	Latency int           `yaml:"latency"`
	Body    string        `yaml:"body"`
	Status  int           `yaml:"status"`
}

func (res *Response) String() string {
	return fmt.Sprintf("%d", res.Status)
}

type Stub struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

func (s *Stub) String() string {
	return fmt.Sprintf("%s -> %s", s.Request.String(), s.Response.String())
}
