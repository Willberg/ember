package main

import "ember/audio/common"

func main() {
	common.PlayNetMp3("你好骚啊", 7)
	common.PlayLocalMp3("/home/john/Music/bensound-sunny.mp3")
}
