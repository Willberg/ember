package singleton

import (
	"fmt"
	"sync"
)

var lazyInstance *singleton
var hungryInstance *singleton
var mu sync.Mutex

type singleton struct {
	Name string
}

func GetHungrySingletonInstance() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if lazyInstance == nil {
		lazyInstance = &singleton{Name: "hungry"}
	}

	fmt.Println(&lazyInstance)
	return lazyInstance
}

func GetLazySingletonInstance() *singleton {
	if hungryInstance == nil {
		mu.Lock()
		if hungryInstance == nil {
			hungryInstance = &singleton{Name: "lazy"}
		}
		mu.Unlock()
	}
	fmt.Println(&hungryInstance)
	return hungryInstance
}
