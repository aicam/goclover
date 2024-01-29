package apicall

import (
	"encoding/json"
	"reflect"
)

func MarshalJSONFiltered(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	fields := make(map[string]interface{})
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if !isZero(field) {
			tag := rv.Type().Field(i).Tag.Get("json")
			key := tag
			if key == "" {
				key = rv.Type().Field(i).Name
			}
			fields[key] = field.Interface()
		}
	}

	return json.Marshal(fields)
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		// Recursively check all fields in nested struct
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false // At least one field is not empty
			}
		}
		return true
	}
	return false
}
