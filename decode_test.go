package query_test

import (
	"fmt"
	"testing"

	"github.com/teerapon19/go-query-string"
)

func TestUnmarshal(t *testing.T) {
	t.Run("run test", func(t *testing.T) {

		type QueryParams struct {
			Test int `query:"test"`
		}

		var q QueryParams

		query.Unmarshal("test=1", &q)

		fmt.Printf("%+v", q)
	})
}
