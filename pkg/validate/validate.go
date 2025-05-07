package validate

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	localzh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	translationszh "github.com/go-playground/validator/v10/translations/zh"
)

// example
//type Args struct {
//	Name string `validate:"required"`
//	Age  int    `validate:"required"`
//}
//
//func Hello(c *gin.Context) {
//	var args Args
//	// 参数验证
//	err := c.ShouldBind(&args)
//	validErrs, _ := err.(validate.ErrValidators)
//	// 用法 1
//	// 检测是否有验证错误
//	if validErrs.HasErrors() {
//		//validErrs[0].Field 字段名
//		//validErrs[0].Message 字段错误信息
//
//		// 输出所有错误
//		c.JSON(http.StatusOK, controller.NewErrParamsResponse(validErrs.Errors()))
//		return
//
//		// 输出第一个错误
//		c.JSON(http.StatusOK, controller.NewErrParamsResponse(validErrs.Error()))
//		return
//	}
//
//	// 用法 2
//	// 参数验证的错误处理
//	if errors.As(err, &validate.ErrValidators{}) {
//		c.JSON(http.StatusOK, controller.NewErrParamsResponse(err.Error()))
//		return
//	}
//}

type ErrValidators []ErrValidator

func (e ErrValidators) Errors() string {
	buff := bytes.NewBufferString("")

	for _, err := range e {
		buff.WriteString(err.Field)
		buff.WriteString(":")
		buff.WriteString(err.Message)
		buff.WriteString(";")
	}

	return strings.TrimSpace(buff.String())
}

func (e ErrValidators) Error() string {
	if e.HasErrors() {
		return fmt.Sprintf(e[0].Message)
	}
	return ""
}

func (e ErrValidators) HasErrors() bool {
	return len(e) > 0
}

type ErrValidator struct {
	Field   string
	Message string
}

func (e *ErrValidator) String() string {
	if e == nil {
		return ""
	}
	return e.Message
}

var ginValidator *defaultValidator
var validate *validator.Validate
var uni *ut.UniversalTranslator
var trans ut.Translator
var _ binding.StructValidator = &defaultValidator{}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    *ut.Translator
}

func init() {
	validate = validator.New()
	localZH := localzh.New()
	uni = ut.New(localZH, localZH)
	trans, _ = uni.GetTranslator("zh")
	_ = translationszh.RegisterDefaultTranslations(validate, trans)
	customValidator()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		var name string
		// name 为汉字描述的字段
		name = strings.SplitN(fld.Tag.Get("name"), ",", 2)[0]
		if name == "" {
			name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		}
		if name == "-" {
			return ""
		}
		return name
	})
	ginValidator = &defaultValidator{
		validate: validate,
	}
}

// GinValidator 初始化验证器
func GinValidator() *defaultValidator {
	return ginValidator
}

// Default 默认验证器
func Default() *validator.Validate {
	return validate
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			errs := err.(validator.ValidationErrors)
			errValidators := make(ErrValidators, len(errs), len(errs))
			for i, e := range errs {
				errValidators[i] = ErrValidator{
					Field:   e.Field(),
					Message: e.Translate(*v.trans),
				}
			}
			return errValidators
		}
	}
	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validate
		v.trans = &trans
		// v.validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(fe.Tag(), fe.Field())
	return t
}

// customValidator 统一注册自定义验证器
func customValidator() {
}
