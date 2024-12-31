package weberrors

import (
	"go.uber.org/zap"
)

/**
 * 参考記事
 * https://zenn.dev/yagi_eng/articles/go-error-handling
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
