package main

type Query map[string]string
type Headers map[string]string

type Request struct {
	Method []string `yaml:"method"`
	Url string `yaml:"url"`
	Query Query `yaml:"query"`
	Headers Headers `yaml:"headers"`
}

type Response struct {
	Headers Headers `yaml:"headers"`
	Latency int `yaml:"latency"`
	Body string `yaml:"body"`
	Status int `yaml:"status"`
}

type Stub struct {
	Request Request `yaml:"request"`
	Response Response `yaml:"response"`
}
