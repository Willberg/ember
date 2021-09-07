package singleton

import (
	"fmt"
	"sync"
)

var lazyInstance *singleton
var hungryInstance *singleton = &singleton{Name: "hungry"}
var mu sync.Mutex

type singleton struct {
	Name string
}

// 饿汉式,导入包的时候就创建对象
func GetHungrySingletonInstance() *singleton {
	fmt.Println(&hungryInstance)
	return hungryInstance
}

// 懒汉式,懒加载,初次调用时创建对象
func GetLazySingletonInstance() *singleton {
	if lazyInstance == nil {
		mu.Lock()
		if lazyInstance == nil {
			lazyInstance = &singleton{Name: "lazy"}
		}
		mu.Unlock()
	}
	fmt.Println(&lazyInstance)
	return lazyInstance
}

var once sync.Once
var onceInstance *singleton

// sync.once保证只执行一次,内部自行判断,具有双锁检查效果
func GetLazySingletonInstanceByOnce() *singleton {
	once.Do(func() {
		onceInstance = &singleton{Name: "Lazy"}
	})
	fmt.Println(&onceInstance)
	return onceInstance
}
