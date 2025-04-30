package types

// Value はLISPの値を表すインターフェース
type Value interface {
	IsAtom() bool
	String() string
}

// Atom はアトムを表す構造体
type Atom struct {
	Value string
}

func (a *Atom) IsAtom() bool {
	return true
}

func (a *Atom) String() string {
	return a.Value
}

// List はリストを表す構造体
type List struct {
	Car Value
	Cdr Value
}

func (l *List) IsAtom() bool {
	return false
}

func (l *List) String() string {
	if l.Car == nil && l.Cdr == nil {
		return "()"
	}
	return "(" + l.Car.String() + " . " + l.Cdr.String() + ")"
}

// Nil は空リストを表す定数
var Nil = &List{}

// T は真値を表す定数
var T = &Atom{Value: "T"}
