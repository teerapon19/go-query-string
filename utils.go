package query

import (
	"regexp"
	"strings"
)

const (
	tag               = "query"
	seperator         = "&"
	equal             = "="
	tagIgnore         = "-"
	tagNameFollowType = "name:type"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func convertToSnakeCase(name string) (snake string) {
	snake = matchFirstCap.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
