package main

import "fmt"

func test_all() {
	overall_success := true

	overall_success = overall_success && test_plus_mult()
	overall_success = overall_success && test_and()
	overall_success = overall_success && test_or()
	overall_success = overall_success && test_negation()
	overall_success = overall_success && test_equal()
	overall_success = overall_success && test_less()

	overall_success = overall_success && test_decl()
	overall_success = overall_success && test_assign()

	if overall_success {
		fmt.Printf("\n=====> [*] Overall overall test SUCCESS")
	} else {
		fmt.Printf("\n=====> [*] Overall overall test FAIL")
	}
}

func test_plus_mult() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	ast := plus(mult(number(1), number(2)), number(0))
	overall_success = overall_success && test_expressions(ast, Val{ValueInt, 2, false}, TyInt)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}

func test_and() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	ast := and(boolean(false), number(0))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, false}, TyIllTyped)
	ast = and(boolean(true), number(0))
	overall_success = overall_success && test_expressions(ast, Val{Undefined, 0, false}, TyIllTyped)
	ast = and(boolean(true), boolean(true))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = and(boolean(true), boolean(false))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, false}, TyBool)
	ast = and(neg(equal(boolean(false), boolean(true))), boolean(true))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}

func test_or() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	ast := or(boolean(false), number(0))
	overall_success = overall_success && test_expressions(ast, Val{Undefined, 0, false}, TyIllTyped)
	ast = or(boolean(true), number(0))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyIllTyped)
	ast = or(boolean(true), boolean(true))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = or(boolean(true), boolean(false))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = or(equal(boolean(false), boolean(true)), boolean(false))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, false}, TyBool)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}

func test_negation() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	ast := neg(boolean(true))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, false}, TyBool)
	ast = neg(boolean(false))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = neg(number(3))
	overall_success = overall_success && test_expressions(ast, Val{Undefined, 0, false}, TyIllTyped)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}

func test_equal() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	ast := equal(boolean(true), boolean(false))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, false}, TyBool)
	ast = equal(boolean(true), boolean(true))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = equal(boolean(false), boolean(false))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = equal(number(3), boolean(false))
	overall_success = overall_success && test_expressions(ast, Val{Undefined, 0, false}, TyIllTyped)
	ast = equal(number(3), number(3))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = equal(number(5), number(3))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, false}, TyBool)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}

func test_less() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	ast := less(number(3), number(5))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, true}, TyBool)
	ast = less(number(5), number(3))
	overall_success = overall_success && test_expressions(ast, Val{ValueBool, 0, false}, TyBool)
	ast = less(boolean(true), number(3))
	overall_success = overall_success && test_expressions(ast, Val{Undefined, 0, false}, TyIllTyped)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}

func test_decl() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	assign_stmt := decl("x", boolean(false))
	overall_success = overall_success && test_stmt(assign_stmt, ValState{"x": Val{ValueBool, 0, false}}, TyState{"x": TyBool}, false)
	assign_stmt = decl("x", number(3))
	overall_success = overall_success && test_stmt(assign_stmt, ValState{"x": Val{ValueInt, 3, false}}, TyState{"x": TyInt}, false)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}

func test_assign() bool {
	fmt.Printf("\n##### ---------- New test ---------- #####")
	overall_success := true
	assign_stmt := seq(decl("x", boolean(false)), assign("x", boolean(true)))
	overall_success = overall_success && test_stmt(assign_stmt, ValState{"x": Val{ValueBool, 0, true}}, TyState{"x": TyBool}, false)
	assign_stmt = seq(decl("x", boolean(false)), assign("x", number(3)))
	overall_success = overall_success && test_stmt(assign_stmt, ValState{"x": Val{ValueBool, 0, false}}, TyState{"x": TyBool}, true)
	assign_stmt = assign("x", boolean(true))
	overall_success = overall_success && test_stmt(assign_stmt, ValState{}, TyState{}, true)
	if overall_success {
		fmt.Printf("\n===> [*] Overall test SUCCESS")
	} else {
		fmt.Printf("\n===> [*] Overall test FAIL")
	}
	return overall_success
}
