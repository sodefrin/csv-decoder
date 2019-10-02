package csvdecoder

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"time"

	"cloud.google.com/go/civil"
)

var (
	ErrInvalidInterface = errors.New("invalid interface")
	ErrParse            = errors.New("parse error")
)

func Decode(out interface{}, r io.Reader) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("%w; out must be pointer", ErrInvalidInterface)
	}

	v = v.Elem()
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("%w; out must be slice", ErrInvalidInterface)
	}

	csvreader := csv.NewReader(r)

	column, err := csvreader.Read()
	if err != nil {
		return err
	}

	for {
		values, err := csvreader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		elem := reflect.New(v.Type().Elem()).Elem()

		if err := decode(elem, column, values); err != nil {
			return err
		}
		v.Set(reflect.Append(v, elem))
	}

	return nil
}

func decode(v reflect.Value, column, values []string) error {
	if v.Kind() == reflect.Ptr {
		v.Set(reflect.New(v.Type().Elem()))
		return decode(v.Elem(), column, values)
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("%w; v must be slice but have %s", ErrInvalidInterface, v.Kind())
	}

	valueMap := map[string]string{}
	if len(column) != len(values) {
		return fmt.Errorf("%w; column must have %d elems", ErrParse, len(column))
	}

	for i, key := range column {
		valueMap[key] = values[i]
	}

	return decodeStruct(v, valueMap)
}

func decodeStruct(v reflect.Value, values map[string]string) error {
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		str := getStrFromMap(v.Type().Field(i), values)

		if str == "" {
			continue
		}

		if err := parseField(value, str); err != nil {
			return err
		}
	}

	return nil
}

func getStrFromMap(field reflect.StructField, values map[string]string) string {
	str, ok := values[field.Tag.Get("json")]
	if ok && str != "" {
		return str
	}

	str, ok = values[field.Name]
	if ok && str != "" {
		return str
	}
	return ""
}

func parseField(value reflect.Value, str string) error {
	if value.Kind() == reflect.Ptr {
		value.Set(reflect.New(value.Type().Elem()))
		return parseField(value.Elem(), str)
	}

	switch value.Interface().(type) {
	case int, int8, int16, int32, int64:
		ret, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(ret)
	case uint, uint8, uint16, uint32, uint64:
		ret, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		value.SetUint(ret)
	case float32, float64:
		ret, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		value.SetFloat(ret)
	case bool:
		ret, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		value.SetBool(ret)
	case string:
		value.SetString(str)
	case time.Time:
		ret, err := time.Parse(time.RFC3339, str)
		if err != nil {
			return err
		}
		value.Set(reflect.ValueOf(ret))
	case civil.Date:
		ret, err := civil.ParseDate(str)
		if err != nil {
			return err
		}
		value.Set(reflect.ValueOf(ret))
	}
	return nil
}
