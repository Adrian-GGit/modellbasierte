package main

// Type inferencer/checker

// Statements

func (stmt Seq) check(t TyState) bool {
	if !stmt[0].check(t) {
		return false
	}
	return stmt[1].check(t)
}

func (decl Decl) check(t TyState) bool {
	ty := decl.rhs.infer(t)
	if ty == TyIllTyped {
		return false
	}

	x := (string)(decl.lhs)
	t[x] = ty
	return true
}

func (assign Assign) check(t TyState) bool {
	x := (string)(assign.lhs)
	y := assign.rhs.infer(t)
	if t[x] != TyIllTyped && y != TyIllTyped {
		return t[x] == assign.rhs.infer(t)
	} else {
		return false
	}
}

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

func (while While) check(t TyState) bool {
	ty := while.cond.infer(t)
	if ty != TyBool {
		return false
	}
	if !while.whileStmt.check(t) {
		return false
	}
	return true
}

func (print Print) check(t TyState) bool {
	if print.printExp.infer(t) != TyIllTyped {
		return true
	} else {
		return false
	}
}

// Expressions

func (x Var) infer(t TyState) Type {
	y := (string)(x)
	ty, ok := t[y]
	if ok {
		return ty
	} else {
		return TyIllTyped
	}
}

func (x Bool) infer(t TyState) Type {
	return TyBool
}

func (x Num) infer(t TyState) Type {
	return TyInt
}

func (e Mult) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyInt
	}
	return TyIllTyped
}

func (e Plus) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyInt
	}
	return TyIllTyped
}

func (e And) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

func (e Or) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

func (e Neg) infer(t TyState) Type {
	t1 := e[0].infer(t)
	if t1 == TyBool {
		return TyBool
	}
	return TyIllTyped
}

func (e Equal) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool || t1 == TyInt && t2 == TyInt {
		return TyBool
	}
	return TyIllTyped
}

func (e Less) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyBool
	}
	return TyIllTyped
}
