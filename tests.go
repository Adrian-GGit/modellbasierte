package main

func test_plus_mult() {
	ast := plus(mult(number(1), number(2)), number(0))
	run(ast)
}

func test_and() {
	ast := and(boolean(false), number(0))
	run(ast)
	ast2 := and(boolean(true), number(0))
	run(ast2)
	ast3 := and(boolean(true), boolean(true))
	run(ast3)
	ast4 := and(boolean(true), boolean(false))
	run(ast4)
	ast5 := and(neg(equal(boolean(false), boolean(true))), boolean(true))
	run(ast5)
}

func test_or() {
	ast := or(boolean(false), number(0))
	run(ast)
	ast2 := or(boolean(true), number(0))
	run(ast2)
	ast3 := or(boolean(true), boolean(true))
	run(ast3)
	ast4 := or(boolean(true), boolean(false))
	run(ast4)
	ast5 := or(equal(boolean(false), boolean(true)), boolean(true))
	run(ast5)
}

func test_negation() {
	ast := neg(boolean(true))
	run(ast)
	ast2 := neg(boolean(false))
	run(ast2)
	ast3 := neg(number(3))
	run(ast3)
}

func test_equal() {
	ast := equal(boolean(true), boolean(false))
	run(ast)
	ast2 := equal(boolean(true), boolean(true))
	run(ast2)
	ast3 := equal(boolean(false), boolean(false))
	run(ast3)
	ast4 := equal(number(3), boolean(false))
	run(ast4)
	ast5 := equal(number(3), number(3))
	run(ast5)
	ast6 := equal(number(5), number(3))
	run(ast6)
}

func test_less() {
	ast := less(number(3), number(5))
	run(ast)
	ast2 := less(number(5), number(3))
	run(ast2)
	ast3 := less(boolean(true), number(3))
	run(ast3)
}

func test_assign() {
	// assign_stmt := assign("x", boolean(false))
	assign_stmt := Assign{"x", boolean(false)} // assign should be only possible after declare (Decl)
	runStmt(assign_stmt)
}

// func test_if_then_else() {
// 	ifthenelse_stmt = ifthenelse(equal(boolean(true), boolean(true)), assign("x", boolean(true)), assign("x", boolean(false)))
// }
