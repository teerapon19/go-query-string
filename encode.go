package query

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	TAG       = "query"
	SEPERATOR = "&"
	EQUAL     = "="
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func Marshal(v any) (string, error) {
	e := newEncoder(v)
	return e.do()
}

type encode struct {
	qb  bytes.Buffer
	obj reflect.Value
}

func newEncoder(obj any) *encode {
	vObj := reflect.ValueOf(obj)

	if reflect.TypeOf(obj).Kind() == reflect.Pointer {
		vObj = vObj.Elem()
	}

	return &encode{
		qb:  bytes.Buffer{},
		obj: vObj,
	}
}

type encodeError struct{ error }

func (e *encode) error(err error) {
	panic(encodeError{err})
}

func (e *encode) do() (query string, err error) {
	defer func() {
		if r := recover(); r != nil {
			if je, ok := r.(encodeError); ok {
				err = je.error
			} else {
				panic(r)
			}
		}
	}()

	for i := 0; i < e.obj.NumField(); i++ {
		e.pair(e.obj.Type().Field(i), e.obj.Field(i))
		if i != e.obj.NumField()-1 {
			e.qb.WriteString(SEPERATOR)
		}
	}

	query = e.qb.String()
	return
}

func (e *encode) pair(fs reflect.StructField, v reflect.Value) {
	bb := bytes.Buffer{}
	key := e.key(fs)
	bb.WriteString(key)
	bb.WriteString(EQUAL)
	bb.WriteString(e.valueToString(v))
	e.qb.Write(bb.Bytes())
}

func (e *encode) key(v reflect.StructField) string {
	name := v.Tag.Get(TAG)

	if name == "case:normal" {
		return v.Name
	}

	return e.convertToSnakeCase(v.Name)
}

func (e *encode) convertToSnakeCase(name string) string {
	snake := matchFirstCap.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (e *encode) valueToString(v reflect.Value) (s string) {

	// deref pointer
	if v.Kind() == reflect.Pointer {
		s = e.valueToString(v.Elem())
		return
	}

	switch v.Kind() {
	case reflect.String:
		s = v.String()
	case reflect.Bool:
		s = strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s = strconv.FormatInt(v.Int(), 10)
	case reflect.Float32:
		s = strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		s = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	default:
		e.error(fmt.Errorf("type does not support"))
	}

	return
}
