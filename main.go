package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)
func ForStmtAddGosched(forStmt *ast.ForStmt){
	for _, stmt := range forStmt.Body.List{
		AddGosched(&stmt)
	}
	tGosched := ast.Ident {Name: "Gosched"}
	tRuntime := ast.Ident {Name: "runtime"}
	tFun := ast.SelectorExpr{X: &tRuntime, Sel: &tGosched}
	shedCall := &ast.ExprStmt{&ast.CallExpr { Fun : &tFun}}
	forStmt.Body.List = append(forStmt.Body.List, shedCall)
}
func RangeStmtAddGosched(rangeStmt *ast.RangeStmt){
	for _, stmt := range rangeStmt.Body.List{
		AddGosched(&stmt)
	}
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

func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
import ("runtime"
	"go/ast"
	"go/parser"
	"go/token")
func tem(){
	for i := 1; i < 10; i+=1 {
		println("Hello, innerfuncWorld!")
		for j := 1; i < 10; i+=1 {
			println("Hello,SUPER innerfuncWorld!")
		}
	}
}
func main() {
	for i := 1; i < 10; i+=1 {
		println("Hello, World!") // reaully hello
		runtime.Gosched() // hellololo
	}
	a := func(i int) int{
		for i := 1;; 1=1 {
			i += 1
			tem()
		}
		return i + 1
	}
	for i := 1;; 1=1 {
		i += 1
		a(i)
	}

}

`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	f.Imports = append(f.Imports, &ast.ImportSpec{Path:&ast.BasicLit{Kind:token.STRING, Value:"runtime"}})
	ast.SortImports(fset,f)
	//ast.Print(fset, f.Imports)
	w := Walker1{}
	ast.Walk(&w, f)
	//ast.Print(fset, f)
	old_code := `	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			for _, stmt := range d.Body.List{
				stmt.Pos()
				//AddGosched(&stmt)
			}
		}
	}

	var programm bytes.Buffer
	printer.Fprint(&programm, fset, f)
	//fmt.Print(programm.String())
	err = ioutil.WriteFile("../for_test.txt", []byte(programm.String()), 0777)
	if err != nil {
		// Если произошла ошибка выводим ее в консоль
		fmt.Println(err)
	}
`
	old_code = old_code
}