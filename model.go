package main

type Command struct {
	Text string `json:"text"`
}

type Result struct {
	Status   int    `json:"status"`
	Data     int    `json:"data"`
	ErrorMsg string `json:"error_msg,omitempty"`
}
