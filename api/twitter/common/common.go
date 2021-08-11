package common

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"sync"
)

type Env struct {
	Ak  string `yaml:"apiKey"`
	Ask string `yaml:"apiSecretKey"`
	Bt  string `yaml:"bearerToken"`
	At  string `yaml:"accessToken"`
	Ats string `yaml:"accessTokenSecret"`
}

type envS struct {
	env *Env
	mu  sync.Mutex
}

var es = envS{env: nil}

func Parse(p string) *Env {
	if es.env != nil {
		return es.env
	}

	es.mu.Lock()
	defer es.mu.Unlock()

	if es.env == nil {
		es.env = &Env{}
	}

	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &es.env)
	if err != nil {
		log.Fatal(err)
	}

	return es.env
}

func GetBearer(p string) string {
	es.env = Parse(p)
	return "Bearer " + es.env.Bt
}
