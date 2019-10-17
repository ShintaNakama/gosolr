package gosolr

type Result struct {
	ResponseHeader Header
	Response       Response
}

type Header struct {
	Status int
	Qtime  int
	Params map[string]interface{} `json:"params"`
}

type Response struct {
	NumFound int
	Start    int
	Docs     []map[string]interface{} `json:"docs"`
}
