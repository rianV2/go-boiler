package helpers

import (
	"encoding/base64"
	"reflect"
	"time"
)

func isZero(v reflect.Value, dv interface{}) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool() == false

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0

	case reflect.Ptr, reflect.Interface:
		return isZero(v.Elem(), dv)

	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i), dv) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.String, reflect.Map:
		return v.Len() == 0

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if !isZero(v.Field(i), dv) {
				return false
			}
		}
		return true

		// reflect.Chan, reflect.UnsafePointer, reflect.Func
	default:
		return v.IsNil()
	}
}

// IsZero reports whether v is zero struct
// Does not support cycle pointers for performance, so as json
func IsZero(v interface{}) bool {
	return isZero(reflect.ValueOf(v), v)
}

func LimitIntWithDefault(minInt int, maxInt int, intVal *int, defaultVal int) int {
	if intVal == nil {
		return defaultVal
	}
	return LimitInt(minInt, maxInt, *intVal)
}

func LimitInt(minInt int, maxInt int, intVal int) int {
	if intVal > maxInt {
		return maxInt
	}
	if intVal < minInt {
		return minInt
	}
	return intVal
}

func EncodeByteToBase64(byteData []byte) string {
	return base64.URLEncoding.EncodeToString(byteData)
}

func DecodeBase64ToByte(text string) []byte {
	result, _ := base64.URLEncoding.DecodeString(text)
	return result
}

func SecondOrNotNilString(val1, val2 *string) *string {

	if val1 == nil && val2 == nil {
		return nil
	}

	if val1 == nil && val2 != nil {
		return val2
	}

	if val1 != nil && val2 == nil {
		return val1
	}

	if *val1 == *val2 {
		return val1
	}

	return val2
}

func LatestTime(val1, val2 *time.Time) *time.Time {

	if val1 == nil && val2 == nil {
		return nil
	}

	if val1 == nil && val2 != nil {
		return val2
	}

	if val1 != nil && val2 == nil {
		return val1
	}

	if *val1 == *val2 {
		return val1
	}

	unixTime1 := ValTimeUnix(val1)
	unixTime2 := ValTimeUnix(val2)

	if unixTime1 > unixTime2 {
		return val1
	}

	return val2
}
