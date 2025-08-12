package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var ErrNotStruct = errors.New("wrong argument given, should be a struct")
var ErrInvalidValidatorSyntax = errors.New("invalid validator syntax")
var ErrValidateForUnexportedFields = errors.New("validation for unexported field is not allowed")

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for _, ve := range v {
		sb.WriteString(fmt.Sprintf("Field '%s': %v\n", ve.Field, ve.Err))
	}
	return sb.String()
}

// Тип функции "правило валидации". Возвращает true, если
// переданное значение удовлетворяет правилу валидации. Иначе false.
type ruleFunc func(vv reflect.Value) bool

// Правило максимума
func ruleFuncMax(n int) ruleFunc {
	return func(vv reflect.Value) bool {
		switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vv.Int() <= int64(n)
		case reflect.String:
			return len(vv.String()) <= n
		}
		return false
	}
}

// Правило минимума
func ruleFuncMin(n int) ruleFunc {
	return func(vv reflect.Value) bool {
		switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vv.Int() >= int64(n)
		case reflect.String:
			return len(vv.String()) >= n
		}
		return false
	}
}

// Правило для длины
func ruleFuncLen(n int) ruleFunc {
	return func(vv reflect.Value) bool {
		if vv.Kind() == reflect.String {
			return len(vv.String()) == n
		}
		return false
	}
}

// Правило для включения
func ruleFuncIn(values []string) ruleFunc {
	return func(vv reflect.Value) bool {
		switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// Создаём множество значений
			intSet := make(map[int64]struct{}, len(values))
			for _, s := range values {
				sInt, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return false
				}
				intSet[sInt] = struct{}{}
			}
			// Делаем проверку принадлежности
			if _, ok := intSet[vv.Int()]; !ok {
				return false
			}
			return true
		case reflect.String:
			// Создаём множество значений
			strSet := make(map[string]struct{}, len(values))
			for _, s := range values {
				strSet[s] = struct{}{}
			}
			// Делаем проверку принадлежности
			if _, ok := strSet[vv.String()]; !ok {
				return false
			}
			return true
		}
		return false
	}
}

// Возвращает функцию согласно переданному тэгу. Возможные значения:
// len:n — длина строки должна быть равна n
// in:a,b,c — значение должно быть в списке
// min:n — значение/длина >= n
// max:n — значение/длина <= n
func parseTag(tag string) (ruleFunc, error) {

	if tag == "" {
		return nil, errors.New("empty tag")
	}

	tagParts := strings.SplitN(tag, ":", 2)
	if len(tagParts) != 2 {
		return nil, errors.New("invalid tag format")
	}

	tagKey := tagParts[0]
	tagValue := tagParts[1]

	if tagValue == "" {
		return nil, errors.New("invalid tag format")
	}

	switch tagKey {
	case "max":
		n, err := strconv.Atoi(tagValue)
		if err != nil {
			return nil, ErrInvalidValidatorSyntax
		}
		return ruleFuncMax(n), nil
	case "min":
		n, err := strconv.Atoi(tagValue)
		if err != nil {
			return nil, ErrInvalidValidatorSyntax
		}
		return ruleFuncMin(n), nil
	case "len":
		n, err := strconv.Atoi(tagValue)
		if err != nil {
			return nil, ErrInvalidValidatorSyntax
		}
		return ruleFuncLen(n), nil
	case "in":
		tagValueParts := strings.SplitN(tagValue, ",", -1)
		return ruleFuncIn(tagValueParts), nil
	}

	return nil, ErrInvalidValidatorSyntax
}

func validateField(fv reflect.Value, tag string) error {

	rf, err := parseTag(tag)
	if err != nil {
		return err
	}
	if rf == nil {
		return ErrInvalidValidatorSyntax
	}
	if !rf(fv) {
		return errors.Errorf(`validation error. expected '%s' , got '%v'`, tag, fv)
	}
	return nil
}

func Validate(v any) error {

	vValue := reflect.ValueOf(v)
	if vValue.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	var allErrs ValidationErrors

	vType := vValue.Type()
	for i := 0; i < vValue.NumField(); i++ {

		fTag, ok := vType.Field(i).Tag.Lookup("validate")

		// Если в тэге не задан "validate", то поле структуры пропускается
		if !ok {
			continue
		}

		// Проверяем, что валидируемое поле экспортируемо
		if !vType.Field(i).IsExported() {
			allErrs = append(allErrs, ValidationError{
				Field: vType.Field(i).Name,
				Err:   ErrValidateForUnexportedFields,
			})
			continue
		}

		fValue := vValue.Field(i)

		// Валидируем значение поля структуры согласно тэгу
		err := validateField(fValue, fTag)
		if err != nil {
			allErrs = append(allErrs, ValidationError{
				Field: vType.Field(i).Name,
				Err:   err,
			})
		}
	}

	if len(allErrs) > 0 {
		return allErrs
	}

	return nil
}
