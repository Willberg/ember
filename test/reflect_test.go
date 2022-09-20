package test

import (
	"fmt"
	"reflect"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

// 被反射调用的函数可以是对外不可见的
func changeString(name string) (int, string) {
	return 1, name + "1"
}

func TestFnReflect(t *testing.T) {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(Add)
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// 反射调用函数
	retList := funcValue.Call(paramList)
	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int())

	f := reflect.ValueOf(changeString)
	params := []reflect.Value{reflect.ValueOf("test")}
	ret := f.Call(params)
	fmt.Println(ret[0].Int(), ret[1].String())
}

type Animal struct {
	Name string
}

// 被反射调用的方法必须是可见的， 因此Change 必须是可见的，不能是change
func (a *Animal) Change(val int, name string) (int, string) {
	if val == 1 {
		a.Name = name
	}
	return val, a.Name
}

func (a *Animal) PrintName() {
	fmt.Println(a.Name)
}

func TestCreate(t *testing.T) {
	typeAnimal := reflect.TypeOf(Animal{})
	animal := reflect.New(typeAnimal)
	animal.Elem().FieldByName("Name").SetString("Cat")      // 利用反射设置结构体字段值
	fmt.Println(animal.Elem().FieldByName("Name").String()) // Hank
}

func TestMethodReflect(t *testing.T) {
	animal := Animal{"cat"}
	val := reflect.ValueOf(&animal)
	method := val.MethodByName("PrintName")
	method.Call([]reflect.Value{})

	f := val.MethodByName("Change")
	params := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf("dog"),
	}
	ret := f.Call(params)
	fmt.Println(ret[0].Int(), ret[1].String(), animal.Name)
}
