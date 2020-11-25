package customvalidator

import (
	"unicode"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	"github.com/seijihg/api_template_mongodb/models"

	ut "github.com/go-playground/universal-translator"
)

// CheckUserValid is user validation function.
func CheckUserValid(userStruct models.User) validator.ValidationErrorsTranslations {

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()

	// Manual validation for password.
	_ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {

		var (
			hasMinLen  = false
			hasUpper   = false
			hasLower   = false
			hasNumber  = false
			hasSpecial = false
		)

		// Check length
		if len(fl.Field().String()) >= 6 {
			hasMinLen = true
		}

		// Check Upper, Lower, numbers and symbols are included.
		for _, value := range fl.Field().String() {
			switch {
			case unicode.IsUpper(value):
				hasUpper = true
			case unicode.IsLower(value):
				hasLower = true
			case unicode.IsNumber(value):
				hasNumber = true
			case unicode.IsPunct(value) || unicode.IsSymbol(value):
				hasSpecial = true
			}
		}
		return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
	})

	// Custom errors message
	// Required
	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	// Email.
	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
	// Password.
	_ = validate.RegisterTranslation("password", trans, func(ut ut.Translator) error {
		return ut.Add("password", "{0} is not strong enough", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())
		return t
	})

	// Validate struct and return nil or errors
	err := validate.Struct(userStruct)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return errs.Translate(trans)
	}

	return nil
}
