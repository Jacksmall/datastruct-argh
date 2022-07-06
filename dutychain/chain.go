/*
 * 责任链代码示例，适用场景:下单操作
 */
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

/*
 * go run .
 * 输出:
 * main.(*CurrencyGBTHandler).Do 积分处理...
 * main.(*CurrencyCNYHandler).Do 现金处理...
 * main.(*GoodsHandler).Do 商品处理...
 * main.(*ChannelHandler).Do 渠道处理...
 * Success
 */

func main() {
	// NullHandler 结构体 有tasks slice 用来接收责任链中多个handler 处理者
	// 多个handler都是struct 都实现了 Handler 接口中的方法 所以可以作为Handler 被tasks接收
	nullHandler := &NullHandler{}
	nullHandler.Tasks = []Handler{
		&CurrencyGBTHandler{},
		&CurrencyCNYHandler{},
		&GoodsHandler{},
		&ChannelHandler{},
	}

	c := &Context{}

	for _, h := range nullHandler.Tasks {
		// 执行每个task 的 Do() 方法
		if err := h.Do(c); err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	}

	fmt.Println("Success")
}
