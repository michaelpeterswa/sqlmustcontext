package main

import (
	"github.com/michaelpeterswa/sqlmustcontext/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
