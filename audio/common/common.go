package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func PlayNetMp3(s string, spd int) {
	f := GetMp3(s, spd)
	if f != nil {
		playMp3(f)
	}
}

func PlayLocalMp3(p string) {
	f := ReadLocalMp3(p)
	if f != nil {
		playMp3(f)
	}
}

func ReadLocalMp3(p string) *os.File {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func playMp3(f io.ReadCloser) {
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func GetMp3(s string, spd int) io.ReadCloser {
	//url := fmt.Sprintf("https://api.oick.cn/txt/apiz.php?text=%s&spd=%d", s, spd)
	url := fmt.Sprintf("https://fanyi.baidu.com/gettts?lan=zh&text=%s&spd=%d&source=web", s, spd)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp.Body
}
