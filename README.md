# Usage
go run .

# Documentation
The [IMP Project](https://sulzmann.github.io/ModelBasedSW/imp.html#(1)) as implemented for the course "Modellbasierte Softwareentwicklung". The task was to implement an evualator and a typechecker for the project. Additionally tests were required to show that the implementation works as intended.

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

## tests.go
Numerous tests to check various expressions and statements.
Example:
```
func test_ifthenelse() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	assign_stmt := seq(decl("x", boolean(true)), ifthenelse(equal(variable("x"), boolean(true)), assign("x", boolean(true)), assign("x", boolean(false))))
	overall_success = overall_success && test_stmt(assign_stmt, ValState{"x": Val{ValueBool, 0, true}}, TyState{"x": TyBool}, false)
	....
}
```
A snippet from the tests for IfThenElse statements. Since we are executing multiple tests in this function we track the overall success in a variable. Then the AST for an IfThenElse statement is built and afterwards passed to the global test function for statements alongside the exptected maps for value state and type state. The final argument of the test function is a boolean which indicates wheter or not the type check is allowed to fail. The test function returns boolean true if the test succeeds and boolean false if not. This return value is then ANDed with the value of the current overall success variable.
