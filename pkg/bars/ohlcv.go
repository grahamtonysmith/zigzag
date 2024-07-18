package bars

import "fmt"

type OHLCV struct {
	o float32
	h float32
	l float32
	c float32
	v int
}

func (bar *OHLCV) String() string {
	return fmt.Sprintf("open=%f high=%f low=%f close=%f volume=%d", bar.o, bar.h, bar.l, bar.c, bar.v)
}
