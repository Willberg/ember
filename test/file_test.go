package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestFileMode(t *testing.T) {
	//dir := "/home/john/mine/workplace/c/linux-5.13.7/scripts/dtc/include-prefixes"
	dir := "/home/john/mine/workplace/c/linux-5.13.7"
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		switch mode := entry.Mode(); {
		case mode.IsRegular():
			fmt.Println("regular file: " + entry.Name())
		case mode.IsDir():
			fmt.Println("directory: " + entry.Name())
		case mode&os.ModeSymlink != 0:
			fmt.Println("symbolic link: " + entry.Name())
			fmt.Println(mode.Perm())
			abPath, err := filepath.EvalSymlinks(filepath.Join(dir, entry.Name()))
			if err == nil {
				fmt.Println(abPath)
			}
		case mode&os.ModeNamedPipe != 0:
			fmt.Println("named pipe")
		}
	}
}
