package validate

import (
	"fmt"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	eng "github.com/go-playground/locales/en"
	uni "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

var (
	Validate  *validator.Validate
	Translate uni.Translator
)

func init() {
	Validate = validator.New()

	en := eng.New()
	Translator := uni.New(en, en)

	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	Translate, _ = Translator.GetTranslator("en")
	_ = enTrans.RegisterDefaultTranslations(Validate, Translate)

}

func StructValidation(s any) (errs response.ErrorLogs) {
	err := Validate.Struct(s)
	if err == nil {
		return errs
	}
	vErrs := err.(validator.ValidationErrors)
	for _, e := range vErrs {
		translatedErr := fmt.Errorf(e.Translate(Translate))
		errs = append(errs, response.ErrorLog{
			RootCause:  translatedErr.Error(),
			Query:      fmt.Sprintf("%s", e.Value()),
			StatusCode: "400",
		})
	}
	return errs
}
