package main

type Type int

const (
	TyIllTyped Type = 0
	TyInt      Type = 1
	TyBool     Type = 2
)

// Value State is a mapping from variable names to values
type ValState map[string]Val

// Type State is a mapping from variable names to types
type TyState map[string]Type

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
	printExp Exp
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
