package helper

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func ValidateStruct(s any) map[string][]string {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	errors := make(map[string][]string)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		tag := fieldType.Tag.Get("validation")

		if jsonTag == "" {
			jsonTag = fieldType.Name
		}

		if field.Kind() == reflect.Struct {
			// Recursively validate nested struct
			nestedErrors := ValidateStruct(field.Addr().Interface())
			for nestedField, nestedErrs := range nestedErrors {
				fieldName := fmt.Sprintf("%s.%s", jsonTag, nestedField)
				errors[fieldName] = nestedErrs
			}
		} else if field.Kind() == reflect.Array || field.Kind() == reflect.Slice {
			var errArr []string
			for i := 0; i < field.Len(); i++ {
				elem := field.Index(i)
				if elem.Kind() == reflect.Struct {
					// Recursively validate nested struct
					nestedErrors := ValidateStruct(elem.Addr().Interface())
					for nestedField, nestedErrs := range nestedErrors {
						for _, errStr := range nestedErrs {
							fieldName := fmt.Sprintf("row: %d - %s - %s", (i + 1), nestedField, errStr)
							errArr = append(errArr, fieldName)
						}
					}
				}

			}

			if len(errArr) > 0 {
				errors[jsonTag] = errArr
			}
		}

		if tag != "" {
			tagValues := strings.Split(tag, ";")

			// First priority: Check "required" validation
			if contains(tagValues, "required") && !isRequired(field.Interface()) {
				errors[jsonTag] = append(errors[jsonTag], fmt.Sprintf("%s is required", jsonTag))
				continue
			}
		}
	}

	return errors
}

func contains(slice []string, item string) bool {
	return slices.Contains(slice, item)
}

func isRequired(value any) bool {
	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		return val.String() != ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return val.Float() != 0.0
	case reflect.Bool:
		return val.Bool()
	case reflect.Slice, reflect.Array:
		return val.Len() != 0
	case reflect.Ptr, reflect.Interface:
		return !val.IsNil()
	default:
		return !reflect.DeepEqual(value, reflect.Zero(val.Type()).Interface())
	}
}
