## Usage
go run .

## Documentation
# ast.go
Helper functions to build expressions and statements.
Example:
```
func ifthenelse(con Exp, stmtIf Stmt, stmtElse Stmt) Stmt {
	return IfThenElse{con, stmtIf, stmtElse}
}
```
Build IfThenElse struct with given parameters. IfThenElse receives a condition expression that ought to be checked and two statements - for if and else branch.

# eval.go
