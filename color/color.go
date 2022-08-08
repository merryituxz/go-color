package color

import (
	"fmt"

	"github.com/merryituxz/go-color/consts"
)

type colorPrinter struct {
	currentColor string
}

var cp colorPrinter

func setColor() {
	fmt.Print(cp.currentColor)
}

func resetColor() {
	fmt.Print(consts.DEFAULT)
}

func Black() *colorPrinter {
	cp.currentColor = consts.BLACK
	return &cp
}

func (cp *colorPrinter) Black() *colorPrinter {
	cp.currentColor = consts.BLACK
	return cp
}

func Red() *colorPrinter {
	cp.currentColor = consts.RED
	return &cp
}

func (cp *colorPrinter) Red() *colorPrinter {
	cp.currentColor = consts.RED
	return cp
}

func Green() *colorPrinter {
	cp.currentColor = consts.GREEN
	return &cp
}

func (cp *colorPrinter) Green() *colorPrinter {
	cp.currentColor = consts.GREEN
	return cp
}

func Yellow() *colorPrinter {
	cp.currentColor = consts.YELLOW
	return &cp
}

func (cp *colorPrinter) Yellow() *colorPrinter {
	cp.currentColor = consts.YELLOW
	return cp
}

func Blue() *colorPrinter {
	cp.currentColor = consts.BLUE
	return &cp
}

func (cp *colorPrinter) Blue() *colorPrinter {
	cp.currentColor = consts.BLUE
	return cp
}

func Magenta() *colorPrinter {
	cp.currentColor = consts.MAGENTA
	return &cp
}

func (cp *colorPrinter) Magenta() *colorPrinter {
	cp.currentColor = consts.MAGENTA
	return cp
}

func Cyan() *colorPrinter {
	cp.currentColor = consts.CYAN
	return &cp
}

func (cp *colorPrinter) Cyan() *colorPrinter {
	cp.currentColor = consts.CYAN
	return cp
}

func White() *colorPrinter {
	cp.currentColor = consts.WHITE
	return &cp
}

func (cp *colorPrinter) White() *colorPrinter {
	cp.currentColor = consts.WHITE
	return cp
}

func (cp *colorPrinter) Println(a ...interface{}) *colorPrinter {
	setColor()
	defer resetColor()
	fmt.Println(a...)
	return cp
}

func (cp *colorPrinter) Printf(format string, a ...interface{}) *colorPrinter {
	setColor()
	defer resetColor()
	fmt.Printf(format, a...)
	return cp
}

func (cp *colorPrinter) Print(a ...interface{}) *colorPrinter {
	fmt.Print(a...)
	return cp
}

func Println(a ...interface{}) *colorPrinter {
	fmt.Println(a...)
	return &cp
}

func Printf(format string, a ...interface{}) *colorPrinter {
	fmt.Printf(format, a...)
	return &cp
}

func Print(a ...interface{}) *colorPrinter {
	fmt.Print(a...)
	return &cp
}
