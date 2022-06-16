package main

import (
	"fmt"
	"runtime"
)

type Context struct {
	// ...
}

type Handler interface {
	Do(c *Context) error
}

type NullHandler struct {
	Tasks []Handler
}

func (h *NullHandler) Do(c *Context) (err error) {
	return
}

type CurrencyGBTHandler struct{}

func (h *CurrencyGBTHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "积分处理...")
	return
}

type CurrencyCNYHandler struct{}

func (h *CurrencyCNYHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "现金处理...")
	return
}

type GoodsHandler struct{}

func (h *GoodsHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "商品处理...")
	return
}

type ChannelHandler struct{}

func (h *ChannelHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "渠道处理...")
	return
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {

	nullHandler := &NullHandler{}
	nullHandler.Tasks = []Handler{
		&CurrencyGBTHandler{},
		&CurrencyCNYHandler{},
		&GoodsHandler{},
		&ChannelHandler{},
	}

	c := &Context{}

	for _, h := range nullHandler.Tasks {
		if err := h.Do(c); err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	}

	fmt.Println("Success")
}
