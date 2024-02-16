package error

import "fmt"

type errorCode int

const (
	BadRequest  errorCode = 400
	ServerError errorCode = 500
)

/*
error类型在Go语言中可以是任何自定义的结构体，只要这个结构体实现了error接口。
在Go中，error是一个内置的接口类型，定义如下：
type error interface {
    Error() string
}
*/

type Err struct {
	ErrorCode errorCode
	Err       error
}

// 实现error接口。MyError类型现在可以作为error使用。
func (e *Err) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ErrorMessage: %v", e.ErrorCode, e.Err)
}
