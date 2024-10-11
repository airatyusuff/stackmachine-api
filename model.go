package main

type Command struct {
	Text string `json:"text"`
}

type Result struct {
	Status int `json:"status"`
	Data   int `json:"data"`
}
type Error struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"error_msg"`
}
