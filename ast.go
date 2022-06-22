package main

// Helper functions to build ASTs by hand

//Statements

// func seq(x, y Stmt) {
// 	return (Seq)([2]Stmt{x, y})
// }

// func assign(lhs string, rhs Exp) {
// 	return Assign{lhs, rhs}
// }

// func decl(lhs string, rhs Exp) {
// 	return Decl{lhs, rhs}
// }

func ifthenelse(con Exp, stmtIf Stmt, stmtElse Stmt) Stmt {
	return IfThenElse{con, stmtIf, stmtElse}
}

func while(con Exp, whileStmt Stmt) Stmt {
	return While{con, whileStmt}
}

func print(printStmt Stmt) Stmt {
	return Print{printStmt}
}

// Expressions

func number(x int) Exp {
	return Num(x)
}

func boolean(x bool) Exp {
	return Bool(x)
}

func plus(x, y Exp) Exp {
	return (Plus)([2]Exp{x, y})
}

func mult(x, y Exp) Exp {
	return (Mult)([2]Exp{x, y})
}

func and(x, y Exp) Exp {
	return (And)([2]Exp{x, y})
}

func or(x, y Exp) Exp {
	return (Or)([2]Exp{x, y})
}

func neg(x Exp) Exp {
	return (Neg)([1]Exp{x})
}

func equal(x, y Exp) Exp {
	return (Equal)([2]Exp{x, y})
}

func less(x, y Exp) Exp {
	return (Less)([2]Exp{x, y})
}
