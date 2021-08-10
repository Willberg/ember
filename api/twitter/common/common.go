package common

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type Env struct {
	Ak  string `yaml:"apiKey"`
	Ask string `yaml:"apiSecretKey"`
	Bt  string `yaml:"bearerToken"`
	At  string `yaml:"accessToken"`
	Ats string `yaml:"accessTokenSecret"`
}

var e = Env{}

func Parse(p string) *Env {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &e)
	if err != nil {
		log.Fatal(err)
	}

	return &e
}

func GetBearer(p string) string {
	return "Bearer " + e.Bt
}
