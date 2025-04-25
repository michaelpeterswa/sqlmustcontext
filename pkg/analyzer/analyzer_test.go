package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/michaelpeterswa/sqlmustcontext/pkg/analyzer"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, analyzer.Analyzer, "test")
}

func TestSliceContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    any
		item     any
		expected bool
	}{
		{
			name:     "string item exists in string slice",
			slice:    []string{"db", "tx", "conn"},
			item:     "db",
			expected: true,
		},
		{
			name:     "string item does not exist in string slice",
			slice:    []string{"db", "tx", "conn"},
			item:     "query",
			expected: false,
		},
		{
			name:     "int item exists in int slice",
			slice:    []int{1, 2, 3},
			item:     2,
			expected: true,
		},
		{
			name:     "int item does not exist in int slice",
			slice:    []int{1, 2, 3},
			item:     4,
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			item:     "db",
			expected: false,
		},
		{
			name:     "item is empty string",
			slice:    []string{"db", "tx", "conn"},
			item:     "",
			expected: false,
		},
		{
			name:     "slice contains empty string",
			slice:    []string{"", "tx", "conn"},
			item:     "",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var result bool
			switch slice := tc.slice.(type) {
			case []string:
				result = analyzer.SliceContains(slice, tc.item.(string))
			case []int:
				result = analyzer.SliceContains(slice, tc.item.(int))
			default:
				assert.Fail(t, "untested type")
			}
			assert.Equal(t, tc.expected, result)
		})
	}
}
