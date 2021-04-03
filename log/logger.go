package log

import "fmt"

func Error(err error) {
	fmt.Println(err.Error())
}
