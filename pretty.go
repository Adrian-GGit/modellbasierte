package main

import "strconv"

/////////////////////////
// Stmt instances

func (stmt Seq) pretty() string {
	return stmt[0].pretty() + "\n" + stmt[1].pretty()
}

func (decl Decl) pretty() string {
	return decl.lhs + " := " + decl.rhs.pretty()
}

func (assign Assign) pretty() string {
	return assign.lhs + " = " + assign.rhs.pretty()
}

func (ifthenelse IfThenElse) pretty() string {
	return "if " + ifthenelse.cond.pretty() + " {\n\t" + ifthenelse.thenStmt.pretty() + "\n} else {\n\t" + ifthenelse.elseStmt.pretty() + "\n}"
}

func (while While) pretty() string {
	return "while " + while.cond.pretty() + " {\n\t" + while.whileStmt.pretty() + "\n}"
}

func (print Print) pretty() string {
	return "print(" + print.printExp.pretty() + ")"
}

/////////////////////////
// Exp instances

func (x Var) pretty() string {
	return (string)(x)
}

func (x Bool) pretty() string {
	if x {
		return "true"
	} else {
		return "false"
	}

}

func (x Num) pretty() string {
	return strconv.Itoa(int(x))
}

func (e Mult) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += " * "
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Plus) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += " + "
	x += e[1].pretty()
	x += ")"

	return x
}

func (e And) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += " && "
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Or) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += " || "
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Neg) pretty() string {

	var x string
	x = "("
	x += "!"
	x += e[0].pretty()
	x += ")"

	return x
}

func (e Equal) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += " == "
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Less) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += " < "
	x += e[1].pretty()
	x += ")"

	return x
}
