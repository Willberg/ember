package main

import (
	"ember/pattern/singleton"
	"fmt"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	var n sync.WaitGroup
	for i := 0; i < 10; i++ {
		n.Add(1)
		go func(n *sync.WaitGroup) {
			defer n.Done()
			hungryInstance := singleton.GetHungrySingletonInstance()
			lazyInstance := singleton.GetLazySingletonInstance()
			fmt.Printf("%v %v\n", &hungryInstance.Name, &lazyInstance.Name)
		}(&n)
	}
	n.Wait()
}
