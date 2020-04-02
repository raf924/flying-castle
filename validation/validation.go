package validation

import (
	"net/url"
	"os"
	"reflect"
	"regexp"
)

type validationType string

const (
	Path         validationType = "path"
	RealPath     validationType = "real_path"
	Url          validationType = "url"
	Mail         validationType = "mail"
	Alphanumeric validationType = "alphanum"
	Alpha        validationType = "alpha"
	Password     validationType = "password"
)

var validationMap map[validationType]func(value string) bool

func init() {
	validationMap = make(map[validationType]func(value string) bool)
	validationMap[Path] = func(value string) bool {
		return true
	}
	validationMap[RealPath] = func(value string) bool {
		_, err := os.Stat(value)
		return err == nil
	}
	validationMap[Alphanumeric] = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	validationMap[Alpha] = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	validationMap[Url] = func(value string) bool {
		_, err := url.Parse(value)
		return err == nil
	}
	validationMap[Mail] = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString
	validationMap[Password] = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{8,}$").MatchString
}

func ValidateField(valueField reflect.Value, typeField reflect.StructField, value interface{}) (reflect.Value, bool) {
	tag := typeField.Tag
	isRequired, hasRequired := tag.Lookup("required")
	realValue := reflect.Indirect(reflect.ValueOf(value))
	if hasRequired && isRequired == "true" && (!realValue.IsValid() || realValue.String() == "") {
		return reflect.ValueOf(nil), false
	}
	if valueField.Kind() == reflect.String {
		if !ValidateTag(realValue.String(), tag) {
			return reflect.ValueOf(nil), false
		}
	}
	return realValue, true
}

func ValidateTag(value string, tag reflect.StructTag) bool {
	if valueType, ok := tag.Lookup("type"); ok {
		return ValidateValue(value, validationType(valueType))
	}
	return true
}

func ValidateValue(value string, valueType validationType) bool {
	if validation, ok := validationMap[valueType]; ok {
		return validation(value)
	} else {
		return false
	}
}
