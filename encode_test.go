package query_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/teerapon19/go-query-string"
)

func TestMarshal(t *testing.T) {

	t.Run("error unsupport type", func(t *testing.T) {

		type QueryParams struct {
			Interface interface{} `query:"interface"`
		}

		_, err := query.Marshal(QueryParams{
			Interface: QueryParams{},
		})

		expectError := fmt.Errorf("type does not support")
		Equal(t, expectError, err)
	})

	t.Run("with string", func(t *testing.T) {

		type QueryParams struct {
			Text string `query:"text"`
		}

		expect := "text=TEXT"
		actual, err := query.Marshal(QueryParams{
			Text: "TEXT",
		})
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, expect, actual)
	})

	t.Run("with bool", func(t *testing.T) {

		type QueryParams struct {
			IsTrue bool `query:"is_true"`
		}

		expect := "is_true=true"
		actual, err := query.Marshal(QueryParams{
			IsTrue: true,
		})
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, expect, actual)
	})

	t.Run("with int int8 int16 int32 int64 and multiple value", func(t *testing.T) {

		type QueryParams struct {
			Int   int   `query:"int"`
			Int8  int8  `query:"int8"`
			Int16 int16 `query:"int16"`
			Int32 int32 `query:"int32"`
			Int64 int64 `query:"int64"`
		}

		expect := "int=1&int8=2&int16=3&int32=4&int64=5"
		actual, err := query.Marshal(QueryParams{
			Int:   1,
			Int8:  2,
			Int16: 3,
			Int32: 4,
			Int64: 5,
		})
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, expect, actual)
	})

	t.Run("with float32 float64 and multiple value", func(t *testing.T) {

		type QueryParams struct {
			Float32 float32 `query:"float32"`
			Float64 float64 `query:"float64"`
		}

		expect := "float32=9.123456&float64=9.1234567890123"
		actual, err := query.Marshal(QueryParams{
			Float32: 9.123456,
			Float64: 9.1234567890123,
		})
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, expect, actual)
	})

	t.Run("without tag", func(t *testing.T) {

		type QueryParams struct {
			IsTrue bool
		}

		expect := "is_true=true"
		actual, err := query.Marshal(QueryParams{
			IsTrue: true,
		})
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, expect, actual)
	})

	t.Run("with tag case:normal", func(t *testing.T) {

		type QueryParams struct {
			IsTrue bool `query:"case:normal"`
		}

		expect := "IsTrue=true"
		actual, err := query.Marshal(QueryParams{
			IsTrue: true,
		})
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, expect, actual)
	})

	t.Run("with pointer field", func(t *testing.T) {

		type QueryParams struct {
			Page *int
		}

		expect := "page=1"
		page := 1
		actual, err := query.Marshal(QueryParams{
			Page: &page,
		})
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, expect, actual)
	})

}

func Equal(t *testing.T, expect interface{}, actual interface{}) {
	if !reflect.DeepEqual(expect, actual) {
		t.Fatalf("expect: %v, actual: %v", expect, actual)
	}
}
