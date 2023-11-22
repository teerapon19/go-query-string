package query

import "testing"

func TestConvertToSnakeCase(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
	}{
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
		{"SCREAMING_SNAKE_CASE", "screaming_snake_case"},
		{"snake_case", "snake_case"},
		{"MixedCase123", "mixed_case123"},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := convertToSnakeCase(tc.name)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

func BenchmarkConvertToSnakeCase(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			convertToSnakeCase("camelCase")
		}
	})
}
