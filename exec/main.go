package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	curl := exec.Command("curl", "-H", "Referer: https://www.baidu.com/", "https://pics5.baidu.com/feed/d833c895d143ad4b7e3f4a17173b35a5a50f0693.jpeg?token=b2261fc4f81d73587f4a39feebe29bcb")
	out, err := curl.Output()
	if err == nil {
		err = ioutil.WriteFile("/home/john/Downloads/1.jpg", out, 0660)
		if err != nil {
			fmt.Printf("%v", err)
		}
		return
	}
	fmt.Printf("%v", err)
}
