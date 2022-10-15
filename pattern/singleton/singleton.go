package singleton

import (
	"fmt"
	"sync"
)

var lazyInstance *Singleton
var hungryInstance *Singleton = &Singleton{Name: "hungry"}
var mu sync.Mutex

type Singleton struct {
	Name string
}

// 饿汉式,导入包的时候就创建对象
func GetHungrySingletonInstance() *Singleton {
	fmt.Println(&hungryInstance)
	return hungryInstance
}

// 懒汉式,懒加载,初次调用时创建对象
func GetLazySingletonInstance() *Singleton {
	if lazyInstance == nil {
		mu.Lock()
		if lazyInstance == nil {
			lazyInstance = &Singleton{Name: "lazy"}
		}
		mu.Unlock()
	}
	fmt.Println(&lazyInstance)
	return lazyInstance
}

var once sync.Once
var onceInstance *Singleton

// sync.once保证只执行一次,内部自行判断,具有双锁检查效果
func GetLazySingletonInstanceByOnce() *Singleton {
	once.Do(func() {
		onceInstance = &Singleton{Name: "Lazy"}
	})
	fmt.Println(&onceInstance)
	return onceInstance
}
