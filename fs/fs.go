package fs

import (
	"encoding/json"
	"os"
)

func ReadJson(path string, inf interface{}) (interface{}, bool) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return err, false
	}
	err = json.Unmarshal(bs, inf)
	if err != nil {
		return err, false
	}
	return inf, true
}
