package pure

import (
	"fmt"
	"reflect"
)

func GetAbstractName[T string | any](haystack T) string {

	var pointer any = &haystack

	switch pointer.(type) {
	case *string, string:
		return fmt.Sprint(haystack)
	}

	reflector := reflect.TypeOf(haystack)

	if reflector.Kind() == reflect.Struct {
		return reflector.String()
	}

	return reflector.Elem().String()
}
