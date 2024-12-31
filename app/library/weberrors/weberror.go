package weberrors

import (
	"go.uber.org/zap"
)

/**
 * 参考記事
 * https://qiita.com/iwakiri0104/items/24d70205aaa2415f0491
 * https://qiita.com/yagi_eng/items/2957ef04cebcdeaae1d6
 */

type WebError struct {
	Code       int
	Msg        string
	StackTrace string
}

// Error error interfaceを実装
func (we *WebError) Error() string {
	return we.Msg
}

// New コンストラクタ
func New(code int, msg string) *WebError {
	stack := zap.Stack("").String
	return &WebError{
		Code:       code,
		Msg:        msg,
		StackTrace: stack,
	}
}
