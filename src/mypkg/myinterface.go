package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnVale() string
}

type MyType int

func (m MyType) MethodWithoutParameters() { //必須定義
	fmt.Println("MethodWithoutParameters called!")
}

func (m MyType) MethodWithParameters(f float64) { //必須定義
	fmt.Println("MethodWithParameters called with", f)
}

func (m MyType) MethodWithReturnVale() string { //必須定義
	return "MethodWithReturnVale called!"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called!")
}
