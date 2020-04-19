package main

import (
	"fmt"

	"github.com/fatih/color"
)

type KeyValuePairs map[string]string

type Request struct {
	Methods    []string      `yaml:"methods"`
	Path       string        `yaml:"path"`
	PathPrefix string        `yaml:"path-prefix"`
	Query      KeyValuePairs `yaml:"query"`
	Headers    KeyValuePairs `yaml:"headers"`
}

func (req *Request) String() string {
	var route string
	if req.Path != "" {
		route = req.Path
	} else {
		route = req.PathPrefix
	}
	return fmt.Sprintf("%s %s", color.MagentaString(fmt.Sprintf("%v", req.Methods)), color.GreenString(route))
}

type Response struct {
	Headers KeyValuePairs `yaml:"headers"`
	Latency int           `yaml:"latency"`
	Body    string        `yaml:"body"`
	File    string        `yaml:"file"`
	Status  int           `yaml:"status"`
}

func (res *Response) String() string {
	return color.YellowString(fmt.Sprintf("%d", res.Status))
}

type Stub struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

func (s *Stub) String() string {
	return fmt.Sprintf("%s %s %s", s.Request.String(), color.CyanString("->"), s.Response.String())
}
