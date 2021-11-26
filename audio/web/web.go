package main

import (
	"ember/audio/common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/tts", handleTts)
	http.HandleFunc("/play", handlePlay)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleTts(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("s")
	spd, err := strconv.Atoi(r.URL.Query().Get("spd"))
	if err != nil {
		fmt.Fprint(w, err)
	}
	body := common.GetMp3(s, spd)
	b, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Fprint(w, err)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=tts.mp3")
	w.Header().Set("Content-Type", "audio/mpeg")
	fmt.Fprint(w, string(b))
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	f := common.ReadLocalMp3("/home/john/Music/bensound-sunny.mp3")
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprint(w, err)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=tts.mp3")
	w.Header().Set("Content-Type", "audio/mpeg")
	fmt.Fprint(w, string(b))
}
