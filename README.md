# Usage
go run .

# Documentation
## ast.go
Helper functions to build expressions and statements.
Example:
```
func ifthenelse(con Exp, stmtIf Stmt, stmtElse Stmt) Stmt {
	return IfThenElse{con, stmtIf, stmtElse}
}
```
Build IfThenElse struct with given parameters. IfThenElse receives a condition expression that ought to be checked and two statements - for if and else branch respectively.

## eval.go
Collection of functions which execute a certain expression or statement.
Example:
```
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
		fmt.Printf("\nif-then-else eval fail")
	}
}
```
The parameter is an IfThenElse statement, consisting of a condition expression and two statements. The condition has to be evaluated first to receive the information about the flag of the expression. Based on this flag either the thenStmt or the elseStmt is evaluated.

## infer.go
Collection of functions which infer a certain expression or statement.
Example:
```
func (ifthenelse IfThenElse) check(t TyState) bool {
	ty := ifthenelse.cond.infer(t)
	if ty != TyBool {
		return false
	}
	if !ifthenelse.thenStmt.check(t) {
		return false
	}
	if !ifthenelse.elseStmt.check(t) {
		return false
	}
	return true
}
```
A function which takes an IfThenElse statement as its only argument and returns a boolean indicating whether or not the the passed statement is of the expected type. First the statement's condition expression is infered. If this fails the function immediately returns false. Otherwise the thenStmt and elseStmt are infered respectively. In case all checks succeed the function returns true.

## pretty.go
This file consists of functions to building a human readable string for an expression or statement. 
Example:
```
func (ifthenelse IfThenElse) pretty() string {
	return "if " + ifthenelse.cond.pretty() + " {\n\t" + ifthenelse.thenStmt.pretty() + "\n} else {\n\t" + ifthenelse.elseStmt.pretty() + "\n}"
}
```
This function takes an IfThenElse statement as an argument and returns a string using the familiar formatting, known by any programmer.
