package validate

import (
	"fmt"
	"reflect"
)

func SelectOneOf[T comparable](name string, values ...T) (T, error) {
	var zero T
	var result T
	for _, value := range values {
		if !reflect.ValueOf(value).IsNil() {
			if result != zero && !reflect.ValueOf(result).IsNil() {
				return zero, fmt.Errorf("only one of %s must be set", name)
			}
			result = value
		}
	}
	return result, nil
}
