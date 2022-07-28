package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string // 错误的字段 TagListRequest.State
	Message string // 错误的具体信息
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v) // 将请求体的参数解析到 v 结构体上
	if err != nil {
		v := c.Value("trans") // 获取翻译器（中间件挂载的字段）
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors) // 如果解析参数失败，获取失败内容
		if !ok {
			return false, errs
		}

		fmt.Printf(">>> trans err: %v", verrs)

		for key, value := range verrs.Translate(trans) { // 遍历翻译之后的错误信息
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
