package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func StringPointer(str string) *string {
	return &str
}

func StringToTime(timeStr *string) *time.Time {
	if timeStr == nil {
		return nil
	}

	if strings.TrimSpace(*timeStr) == "" {
		return nil
	}

	formats := []string{time.RFC3339, "2006-01-02"}

	for _, v := range formats {
		tt, err := time.Parse(v, *timeStr)
		if err != nil {
			continue
		}
		return &tt
	}

	return nil
}

func StringToInt64(str *string) *int64 {
	if str == nil {
		return nil
	}
	intVal, err := strconv.ParseInt(*str, 10, 64)
	if err != nil {
		return nil
	}
	return &intVal
}

func StringToInt(str *string) *int {
	if str == nil {
		return nil
	}
	intVal, err := strconv.Atoi(*str)
	if err != nil {
		return nil
	}
	return &intVal
}

func NumberToInt(num *json.Number) *int {
	if nil == num {
		return nil
	}
	i, err := num.Int64()
	if err != nil {
		return nil
	}
	i2 := int(i)
	return &i2
}

func NumberToInt64(num *json.Number) *int64 {
	if nil == num {
		return nil
	}
	i, err := num.Int64()
	if err != nil {
		return nil
	}
	return &i
}

func NumberToFloat64(num *json.Number) *float64 {
	if nil == num {
		return nil
	}
	f, err := num.Float64()
	if err != nil {
		return nil
	}
	return &f
}

func Int64ToString(num *int64) *string {
	if num == nil {
		return nil
	}

	s := strconv.FormatInt(*num, 10)
	return &s
}

func IntToString(num *int) *string {
	if num == nil {
		return nil
	}

	s := strconv.Itoa(*num)
	return &s
}

func IntPointer(i int) *int {
	return &i
}

func Int64Pointer(i int64) *int64 {
	return &i
}

func BoolPointer(i bool) *bool {
	return &i
}

func Float64Pointer(i float64) *float64 {
	return &i
}

func TimePointer(time time.Time) *time.Time {
	return &time
}

func JsonNumberPointer(num json.Number) *json.Number {
	return &num
}

func InterfaceToString(interfaceData interface{}) *string {
	stringData := fmt.Sprintf("%v", interfaceData)
	return &stringData
}

func Val(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	refVal := reflect.ValueOf(value)
	if refVal.Kind() != reflect.Ptr {
		return value
	}

	if refVal.IsNil() {
		return nil
	}

	return refVal.Elem().Interface()
}

func ValStr(val *string) string {
	if val == nil {
		return ""
	}

	return *val
}

func ValTimeUnix(val *time.Time) int64 {
	if val == nil {
		return 0
	}

	return val.Unix()
}

func EqualValueStr(a *string, b *string) bool {
	return ValStr(a) == ValStr(b)
}

func ValInt64(val *int64) int64 {
	if val == nil {
		return 0
	}

	return *val
}
