package funcreturn

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "funcreturn",
	Doc:  "Checks whether there is a newline between functions",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		endOfFunc := 0
		beginOfFunc := 0
		prevFunc := false
		ast.Inspect(file, func(n ast.Node) bool {
			switch fn := n.(type) {
			case *ast.FuncDecl:
				currentFuncPos := pass.Fset.Position(fn.Pos()).Line
				if prevFunc && beginOfFunc != endOfFunc && currentFuncPos-endOfFunc == 1 {
					pass.Reportf(fn.Pos(), "There is no newline before function")
				}
				beginOfFunc = pass.Fset.Position(fn.Pos()).Line
				endOfFunc = pass.Fset.Position(fn.End()).Line
				prevFunc = true
				return false
			default:
				prevFunc = false
			}

			return true
		})
	}
	return nil, nil
}
