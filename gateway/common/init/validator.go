package init

import (
	"fmt"
	"reflect"

	"gateway/common/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func InitTrans(locale string) (err error) {
	// 修改 Gin 框架中的 validator 引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取 json tag 的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("comment")
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
