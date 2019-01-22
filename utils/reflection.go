package utils

import (
	"errors"
	"reflect"
)

// copied from github.com/oleiade/reflections/blob/master/reflections.go

func Fields(obj interface{}) ([]string, error) {
	if !HasValidType(obj, []reflect.Kind{reflect.Struct, reflect.Ptr}) {
		return nil, errors.New("Cannot use GetField on a non-struct interface")
	}

	objValue := ReflectValue(obj)
	objType := objValue.Type()
	fieldsCount := objType.NumField()

	var allFields []string
	for i := 0; i < fieldsCount; i++ {
		field := objType.Field(i)
		if IsExportableField(field) {
			allFields = append(allFields, field.Name)
		}
	}

	return allFields, nil
}
func ReflectValue(obj interface{}) reflect.Value {
	var val reflect.Value

	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		val = reflect.ValueOf(obj).Elem()
	} else {
		val = reflect.ValueOf(obj)
	}

	return val
}

func IsExportableField(field reflect.StructField) bool {
	// PkgPath is empty for exported fields.
	return field.PkgPath == ""
}

func HasValidType(obj interface{}, types []reflect.Kind) bool {
	for _, t := range types {
		if reflect.TypeOf(obj).Kind() == t {
			return true
		}
	}

	return false
}
