package common

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type env struct {
	AK string `yaml:"apiKey"`
}

var e = env{}

func GetApiKey(p string) string {
	if e.AK != "" {
		return e.AK
	}

	data, err := os.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &e)
	if err != nil {
		log.Fatal(err)
	}
	return e.AK
}
