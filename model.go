package main

type Video struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Views int    `json:"views"`
}

type jsonResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Video  Video
}
