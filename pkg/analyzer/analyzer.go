package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "sqlmustcontext",
	Doc:      "check that all sql calls use their context variants",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var (
	disallowedMethods = []string{
		"Exec",
		"Ping",
		"Prepare",
		"Query",
		"QueryRow",
	}

	allowedTypes = []string{
		"*database/sql.DB",
		"*database/sql.Tx",
	}
)

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	callFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(callFilter, func(node ast.Node) {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return
		}

		// Check if the function is a method call on a *sql.DB or *sql.Tx
		if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := selExpr.X.(*ast.Ident); ok {
				obj := pass.TypesInfo.ObjectOf(ident)
				if obj != nil {
					typ := obj.Type()
					if SliceContains(allowedTypes, typ.String()) {
						// Check if the method name is one of the non-context methods
						methodName := selExpr.Sel.Name
						if SliceContains(disallowedMethods, methodName) {
							pass.Reportf(callExpr.Pos(), "use %sContext instead of %s", methodName, methodName)
							return
						}
					}
				}
			}
		}
	})

	return nil, nil
}

func SliceContains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}
	return false
}
