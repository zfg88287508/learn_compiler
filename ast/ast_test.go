package ast

import "testing"

func TestNewNode(t *testing.T) {
	x := NewNode(Program,[]rune("calculator"))



}

func TestDumpAST(t *testing.T) {
	x := NewNode(Program,[]rune("calculator"))
	//打印AST
	DumpAST(x,"\t")
}
