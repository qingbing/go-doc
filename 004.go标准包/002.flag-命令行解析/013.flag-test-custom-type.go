package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

/*
flag 接受自定义的类型，类型必须实现接口
type flag.Value interface {
	String() string
	Set(string) error
}
*/
type myDuration []time.Duration

func (d *myDuration) String() string {
	return fmt.Sprint(*d)
}
func (d *myDuration) Set(val string) error {
	if len(*d) > 0 {
		return errors.New("已经设置了值")
	}
	for _, dt := range strings.Split(val, ",") {
		duration, err := time.ParseDuration(dt)
		if nil != err {
			return err
		}
		*d = append(*d, duration)
	}
	return nil
}

var (
	myD myDuration
)

func init() {
	flag.Var(&myD, "d", "时间切片，用','分割")
}

//  -d 1s,20ms,30ns => [1s 20ms 30ns]
func main() {
	flag.Parse()
	fmt.Println(myD)
}
