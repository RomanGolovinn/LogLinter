package loglinter

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/RomanGolovinn/loglinter/internal/rules"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loglinter",
	Doc:      "checks log messages for formatting and sensitive data",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	activeRules := rules.GetAllRules()

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		checkNode(n, pass, activeRules)
	})

	return nil, nil
}

func checkNode(n ast.Node, pass *analysis.Pass, activeRules []rules.Rule) {
	call := n.(*ast.CallExpr)

	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return
	}

	obj := pass.TypesInfo.Uses[ident]
	if obj == nil {
		return
	}

	pkgName, ok := obj.(*types.PkgName)
	if !ok {
		return
	}

	importPath := pkgName.Imported().Path()
	isLogPackage := importPath == "log" || importPath == "log/slog" || importPath == "go.uber.org/zap"
	if !isLogPackage {
		return
	}

	methodName := sel.Sel.Name
	isLogMethod := methodName == "Info" || methodName == "Error" || methodName == "Warn" || methodName == "Debug" || methodName == "Fatal"
	if !isLogMethod {
		return
	}

	if len(call.Args) == 0 {
		return
	}

	msg, ok := extractString(call.Args[0])
	if !ok {
		return
	}

	for _, rule := range activeRules {
		if errMessage := rule(msg); errMessage != "" {
			pass.Reportf(call.Args[0].Pos(), "%s", errMessage)
		}
	}
}

func extractString(expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			return strings.Trim(e.Value, `"`), true
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			leftStr, _ := extractString(e.X)
			rightStr, _ := extractString(e.Y)
			return leftStr + " " + rightStr, true
		}
	case *ast.Ident:
		return e.Name, true
	}
	return "", false
}
