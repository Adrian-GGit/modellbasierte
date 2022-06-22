package main

import "fmt"

func (stmt Seq) eval(s ValState) {
	stmt[0].eval(s)
	stmt[1].eval(s)
}

func (decl Decl) eval(s ValState) {
	v := decl.rhs.eval(s)
	x := (string)(decl.lhs)
	s[x] = v
}

func (assign Assign) eval(s ValState) {
	v := assign.rhs.eval(s)
	x := (string)(assign.lhs)
	s[x] = v
}

func (ifthenelse IfThenElse) eval(s ValState) {
	v := ifthenelse.cond.eval(s)
	if v.flag == ValueBool {
		switch {
		case v.valB:
			ifthenelse.thenStmt.eval(s)
		case !v.valB:
			ifthenelse.elseStmt.eval(s)
		}
	} else {
		fmt.Printf("if-then-else eval fail")
	}
}

func (while While) eval(s ValState) {
	v := while.cond.eval(s)
	if v.flag == ValueBool {
		if v.valB {
			while.whileStmt.eval(s)
			while.eval(s)
		}
	} else {
		fmt.Printf("while eval fail")
	}
}

func (print Print) eval(s ValState) {
	// TODO: soll hier iwas geprinted weren?
	print.printStmt.eval(s)
}

// Evaluator

func (x Bool) eval(s ValState) Val {
	return mkBool((bool)(x))
}

func (x Num) eval(s ValState) Val {
	return mkInt((int)(x))
}

func (e Mult) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI * n2.valI)
	}
	return mkUndefined()
}

func (e Plus) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI + n2.valI)
	}
	return mkUndefined()
}

func (e And) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == false:
		return mkBool(false)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB && b2.valB)
	}
	return mkUndefined()
}

func (e Or) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == true:
		return mkBool(true)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB || b2.valB)
	}
	return mkUndefined()
}

func (e Neg) eval(s ValState) Val {
	b1 := e[0].eval(s)
	if b1.flag == ValueBool {
		return mkBool(!b1.valB)
	}
	return mkUndefined()
}

func (e Equal) eval(s ValState) Val {
	nb1 := e[0].eval(s)
	nb2 := e[1].eval(s)
	switch {
	case nb1.flag == ValueInt && nb2.flag == ValueInt:
		return mkBool(nb1.valI == nb2.valI)
	case nb1.flag == ValueBool && nb2.flag == ValueBool:
		return mkBool(nb1.valB == nb2.valB)
	}
	return mkUndefined()
}

func (e Less) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkBool(n1.valI < n2.valI)
	}
	return mkUndefined()
}
