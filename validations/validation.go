package validations

import (
	"blog-app/entities/web"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func InitValidations() {
	validate = validator.New()
}

func Validation(body interface{}) (bool, []*web.ValidationErrorResponse) {
	var errors []*web.ValidationErrorResponse
	if err := validate.Struct(body); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errors = append(errors, &web.ValidationErrorResponse{
				Field:   e.Field(),
				Message: e.Translate(trans),
			})
		}

		return false, errors
	}
	return true, nil
}
