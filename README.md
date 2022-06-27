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
Build IfThenElse struct with given parameters. IfThenElse receives a condition expression that ought to be checked and two statements - for if and else branch.

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
The parameter is an IfThenElse statement, consisting of a condition expression and two statements. The condition has to be evaluated aswell to receive the information about the flag of the expression. If the expression is type of boolean then the IfThenElse Statement can be evaluated. In case the expression is true the thenStmt is evaluated, otherwise the elseStmt.

## infer.go
