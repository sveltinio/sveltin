package validator

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Init setup the validator used to validate the sveltin.json file.
func Init() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(tagNameAsJsonValue)
	_ = validate.RegisterValidation("dateiso", validateISODate)
	return validate
}

func tagNameAsJsonValue(fld reflect.StructField) string {
	return fld.Tag.Get("json")
}

func validateISODate(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return re.MatchString(str)
}
