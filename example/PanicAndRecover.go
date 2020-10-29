package example

import (
	"errors"
	"fmt"
)

/**
处理异常，panic直接抛出异常之名错误，recover捕获异常
*/
type PanicAndRecover struct {
}

func (receiver *PanicAndRecover) ReceivePanic() {
	fmt.Println("\nPanicAndRecover.ReceivePanic():")
	defer receiver.CoverPanic()
	panic("this is panic error")
}

func (receiver *PanicAndRecover) RecoverPanic() {
	fmt.Println("\nPanicAndRecover.RecoverPanic():")
	defer func() {
		msg := recover()
		fmt.Println("recover msg:", msg)
	}()
	panic("this is panic error")
}

func (receiver *PanicAndRecover) CoverPanic() {
	fmt.Println("\nPanicAndRecover.CoverPanic():")
	msg := recover()
	fmt.Println("panic recover msg:", msg)
}

func (receiver *PanicAndRecover) CoverPanicGood() {
	fmt.Println("\nPanicAndRecover.CoverPanicGood():")
	msg := recover()
	switch msg.(type) {
	case string:
		fmt.Println("string recover msg:", msg)
	case error:
		fmt.Println("error recover msg:", msg)
	default:
		fmt.Println("unknown recover msg:", msg)
	}
}

func (receiver *PanicAndRecover) RecoverPanicTest() {
	fmt.Println("\nPanicAndRecover.RcoverPanicTest():")
	defer receiver.CoverPanicGood()
	//panic("this is panic error")

	defer receiver.CoverPanicGood()
	panic(errors.New("this is errors.New()"))
}
