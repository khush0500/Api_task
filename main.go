package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var Videos = make(map[string]*Video)

func allVideos(w http.ResponseWriter, r *http.Request) {
	fmt.Println(chi.URLParam(r, "id"))
	out, err := json.MarshalIndent(Videos, "", "     ")

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	video, ok := Videos[id]
	if ok {
		video.Views = video.Views + 1
		jsonval := jsonResponse{
			Status: "OK",
			Msg:    "View Inc",
			Video:  *video,
		}

		out, _ := json.MarshalIndent(jsonval, "", "     ")

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		jsonval := jsonResponse{
			Status: "FAIL",
			Msg:    "No Video with given Id",
		}
		out, _ := json.MarshalIndent(jsonval, "", "     ")

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

	}

}

func addVideo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var video Video
	json.Unmarshal(reqBody, &video)
	rand := RandStringRunes(10)
	_, ok := Videos[rand]
	for ok {
		rand = RandStringRunes(15)
		_, ok = Videos[rand]
	}
	video.Id = RandStringRunes(10)
	fmt.Println(video)
	Videos[video.Id] = &video

	out, err := json.MarshalIndent(video, "", "     ")

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	handleRequests()
}
