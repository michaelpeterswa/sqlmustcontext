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
