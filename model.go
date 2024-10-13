package main

type Command struct {
	Text string `json:"text"`
}

type Result struct {
	Status   int    `json:"status"`
	Data     int    `json:"data,omitempty"`
	ErrorMsg string `json:"error_msg,omitempty"`
}
