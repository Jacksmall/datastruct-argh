package main

import (
	"fmt"
	"runtime"
)

type Context struct{}

type Handler interface {
	// 自身业务
	Do(c *Context) (err error)
	// 设置下一个处理对象
	SetNext(h Handler) Handler
	// 执行
	Run(c *Context) (err error)
}

type Next struct {
	nextHandler Handler
}

func (n *Next) SetNext(h Handler) Handler {
	n.nextHandler = h
	return h
}

func (n *Next) Run(c *Context) (err error) {
	if n.nextHandler != nil {
		if err = n.nextHandler.Do(c); err != nil {
			return
		}
		return n.nextHandler.Run(c)
	}
	return
}

type NullHandler struct {
	Next
}

func (h *NullHandler) Do(c *Context) (err error) {
	return
}

type ArgumentHandler struct {
	Next
}

func (h *ArgumentHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "参数校验handler")
	return
}

type AddressInfoHandler struct {
	Next
}

func (h *AddressInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取地址信息handler")
	return
}

type CartInfoHandler struct {
	Next
}

func (h *CartInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取购物车信息handler")
	return
}

type StockInfoHandler struct {
	Next
}

func (h *StockInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取商品库存信息handler")
	return
}

type PromotionInfoHandler struct {
	Next
}

func (h *PromotionInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取优惠信息handler")
	return
}

type ShipInfoHandler struct {
	Next
}

func (h *ShipInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取运费信息handler")
	return
}

type PromotionUseInfoHandler struct {
	Next
}

func (h *PromotionUseInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取优惠使用信息handler")
	return
}

type StockSubtractHandler struct {
	Next
}

func (h *StockSubtractHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "扣库存。。。")
	return
}

type CartDelHandler struct {
	Next
}

func (h *CartDelHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "清理购物车")
	return
}

type DBTableOrderHandler struct {
	Next
}

func (h *DBTableOrderHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "写订单表...")
	return
}

type DBTableOrderSkusHandler struct {
	Next
}

func (h *DBTableOrderSkusHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "写订单sku表...")
	return
}

type DBTableOrderPromotionHandler struct {
	Next
}

func (h *DBTableOrderPromotionHandler) Do(c *Context) (er error) {
	fmt.Println(runFuncName(), "写订单优惠信息表...")
	return
}

// ...

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	nullHandler := &NullHandler{}
	nullHandler.SetNext(&ArgumentHandler{}).
		SetNext(&AddressInfoHandler{}).
		SetNext(&CartInfoHandler{}).
		SetNext(&StockInfoHandler{}).
		SetNext(&PromotionInfoHandler{}).
		SetNext(&PromotionUseInfoHandler{}).
		SetNext(&StockSubtractHandler{}).
		SetNext(&CartDelHandler{}).
		SetNext(&DBTableOrderHandler{}).
		SetNext(&DBTableOrderSkusHandler{}).
		SetNext(&DBTableOrderPromotionHandler{})

	if err := nullHandler.Run(&Context{}); err != nil {
		fmt.Println("Error | Fatal:" + err.Error())
		return
	}
	fmt.Println("Success")
}
