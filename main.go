package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	//"golang.org/x/tools"
	"io/ioutil"
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
func ForAddGosched(forStmt *ast.ForStmt){
	for _, stmt := range forStmt.Body.List{
		AddGosched(&stmt)
	}
	tGosched := ast.Ident {Name: "Gosched"}
	tRuntime := ast.Ident {Name: "runtime"}
	tFun := ast.SelectorExpr{X: &tRuntime, Sel: &tGosched}
	shedCall := &ast.ExprStmt{&ast.CallExpr { Fun : &tFun}}
	forStmt.Body.List = append(forStmt.Body.List, shedCall)
}

func AddGosched(stat  *ast.Stmt){
	switch stat := (*stat).(type) {
	case ast.ForStmt:
		ForStmtAddGosched(&stat)
	case ast.RangeStmt:
		RangeStmtAddGosched(&stat)
	}

	for _, stat := range stat .Body.List{
		switch stat.(type) {
		case *ast.ForStmt:
			ForStmtAddGosched(stat)
		case *ast.RangeStmt:
			RangeStmtAddGosched(stat)
		}
	}
}


func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
import ("runtime")
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
	for i := 1;; 1=1 {
	}
	tem()
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	//ast.Print(fset, f)
	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			for _, stmt := range decl.Body.List{
				switch stmt := stmt.(type) {
				case *ast.ForStmt:
					AddGosched(stmt)
				}
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
}