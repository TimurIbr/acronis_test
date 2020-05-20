package AddGoshed

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
)

func ForStmtAddGosched(forStmt *ast.ForStmt){
	tGosched := ast.Ident {Name: "Gosched"}
	tRuntime := ast.Ident {Name: "runtime"}
	tFun := ast.SelectorExpr{X: &tRuntime, Sel: &tGosched}
	shedCall := &ast.ExprStmt{&ast.CallExpr { Fun : &tFun}}
	forStmt.Body.List = append(forStmt.Body.List, shedCall)
}
func RangeStmtAddGosched(rangeStmt *ast.RangeStmt){
	tGosched := ast.Ident {Name: "Gosched"}
	tRuntime := ast.Ident {Name: "runtime"}
	tFun := ast.SelectorExpr{X: &tRuntime, Sel: &tGosched}
	shedCall := &ast.ExprStmt{&ast.CallExpr { Fun : &tFun}}
	rangeStmt.Body.List = append(rangeStmt.Body.List, shedCall)
}

func AddGosched(stat  *ast.Stmt){
	tStat := *stat
	switch st := tStat.(type) {
	case *ast.ForStmt:
		ForStmtAddGosched(st)
	case *ast.RangeStmt:
		RangeStmtAddGosched(st)
	}
}
type  Walker1 struct {
	n int
}

func (w *Walker1)Visit(n ast.Node) ast.Visitor{
	//ast.Print(nil, n)
	switch nT := n.(type) {
	case *ast.ForStmt:
		ForStmtAddGosched(nT)
	case *ast.RangeStmt:
		RangeStmtAddGosched(nT)
	}
	return w
}

func AddImportRuntime(fset *token.FileSet, f *ast.File){
	//f.Imports = append(f.Imports , &ast.ImportSpec{Path:&ast.BasicLit{Kind:token.STRING, Value:"\"runtime\""}})
	hasRuntime := false
	for _, impr := range f.Imports{
		if impr.Path.Value == "\"runtime\""{
			hasRuntime = true
		}
	}
	if !hasRuntime{
		importRuntime := &ast.GenDecl{
			TokPos: f.Package,
			Tok:    token.IMPORT,
			Specs:  []ast.Spec{&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: "\"runtime\""}}},
		}
		f.Decls = append([]ast.Decl{importRuntime}, f.Decls...)
		ast.SortImports(fset,f)
	}
}

func AddGoschedToFile(fileName string, src string) string{
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, src, 0)
	if err != nil {
		fmt.Print(fmt.Errorf("unable to parse file %v: %v", fileName, err))
		panic(err)
	}
	AddImportRuntime(fset, f)
	w := Walker1{}
	ast.Walk(&w, f)
	var programm bytes.Buffer
	err = printer.Fprint(&programm, fset, f)
	if err != nil{
		panic(err)
	}
	return programm.String()
}

