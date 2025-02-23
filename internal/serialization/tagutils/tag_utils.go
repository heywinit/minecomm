package tagutils

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrInvalidLen = errors.New("minecomm: invalid len tag")

// CheckDependency
// returns false either if the dependency has not been found or the dependency was not met
// returns true if the dependency has been found AND it's true
func CheckDependency(inter reflect.Value, field reflect.StructField) bool {
	depend := field.Tag.Get("depends_on")
	if depend == "" {
		return true
	}

	val := inter.FieldByName(depend)
	return val.Bool()
}

func GetLength(inter reflect.Value, field reflect.StructField) (int, error) {
	lengthTag := field.Tag.Get("len")
	if lengthTag == "" {
		return -1, nil
	} else if length, err := strconv.Atoi(lengthTag); err == nil {
		return length, nil
	} else if lenField := inter.FieldByName(lengthTag); lenField.IsValid() {
		return int(lenField.Int()), nil
	} else {
		return 0, ErrInvalidLen
	}
}
