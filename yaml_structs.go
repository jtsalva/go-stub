package main

import (
	"fmt"
)

type KeyValuePairs map[string]string

type Request struct {
	Methods []string      `yaml:"methods"`
	Path    string        `yaml:"path"`
	Query   KeyValuePairs `yaml:"query"`
	Headers KeyValuePairs `yaml:"headers"`
}

func (req *Request) String() string {
	return fmt.Sprintf("%v %s", req.Methods, req.Path)
}

type Response struct {
	Headers KeyValuePairs `yaml:"headers"`
	Latency int           `yaml:"latency"`
	Body    string        `yaml:"body"`
	File    string        `yaml:"file"`
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
