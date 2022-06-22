package main

import (
	"fmt"
	"reflect"
)

// Types

type Type int

const (
	TyIllTyped Type = 0
	TyInt      Type = 1
	TyBool     Type = 2
)

func showType(t Type) string {
	var s string
	switch {
	case t == TyInt:
		s = "Int"
	case t == TyBool:
		s = "Bool"
	case t == TyIllTyped:
		s = "Illtyped"
	}
	return s
}

// Value State is a mapping from variable names to values
type ValState map[string]Val

// Type State is a mapping from variable names to types
type TyState map[string]Type

// Interface

type Exp interface {
	pretty() string
	eval(s ValState) Val
	infer(t TyState) Type
}

type Stmt interface {
	pretty() string
	eval(s ValState)
	check(t TyState) bool
}

type Seq [2]Stmt
type Decl struct {
	lhs string
	rhs Exp
}
type Assign struct {
	lhs string
	rhs Exp
}
type IfThenElse struct {
	cond     Exp
	thenStmt Stmt
	elseStmt Stmt
}

type While struct {
	cond      Exp
	whileStmt Stmt
}

type Print struct {
	printStmt Stmt
}

type Bool bool
type Num int
type Mult [2]Exp
type Plus [2]Exp
type And [2]Exp
type Or [2]Exp
type Var string
type Neg [1]Exp
type Equal [2]Exp
type Less [2]Exp

// TODO: maybe implement grouping (...) as expression

// Examples

func test_expressions(e Exp, expected_val Val, expected_type Type) bool {
	val_states := make(ValState)
	type_states := make(TyState)
	fmt.Printf("\n---------- New EXP test case ----------")
	fmt.Printf("\n %s", e.pretty())
	e.eval(val_states)
	type_check := e.infer(type_states) == expected_type
	if type_check {
		fmt.Printf("\n[*] Typecheck SUCCESS")
	} else {
		fmt.Printf("\n[*] Typecheck FAIL")
	}
	eval_check := e.eval(val_states) == expected_val
	if eval_check {
		fmt.Printf("\n[*] Evalcheck SUCCESS")
	} else {
		fmt.Printf("\n[*] Evalcheck FAIL")
	}
	if type_check && eval_check {
		fmt.Printf("\n=> [*] Overall SUCCESS")
		return true
	} else {
		fmt.Printf("\n=> [!] Overall FAIL")
		return false
	}
}

func test_stmt(stmt Stmt, expected_vals ValState, expected_types TyState, allow_check_fail bool) bool {
	val_states := make(ValState)
	type_states := make(TyState)
	fmt.Printf("\n---------- New STMT test case ----------")
	fmt.Printf("\n %s", stmt.pretty())
	stmt.eval(val_states)
	type_check := stmt.check(type_states)
	if type_check {
		fmt.Printf("\n[*] Typecheck SUCCESS")
	} else {
		fmt.Printf("\n[*] Typecheck FAIL")
	}
	if allow_check_fail && !type_check {
		fmt.Printf("- but was allowed to FAIL")
		type_check = true
	} else if allow_check_fail && type_check {
		fmt.Printf("- but shouldve FAIL")
		type_check = false
	}
	compare_val_states := reflect.DeepEqual(val_states, expected_vals)
	compare_type_states := reflect.DeepEqual(type_states, expected_types)
	if compare_val_states {
		fmt.Printf("\n[*] Eval SUCCESS")
	} else {
		fmt.Printf("\n[!] Eval FAIL")
	}
	if compare_type_states {
		fmt.Printf("\n[*] Type SUCCESS")
	} else {
		fmt.Printf("\n[!] Type FAIL")
	}
	if type_check && compare_val_states && compare_type_states {
		fmt.Printf("\n=> [*] Overall SUCCESS")
		return true
	} else {
		fmt.Printf("\n=> [!] Overall FAIL")
		return false
	}
}
