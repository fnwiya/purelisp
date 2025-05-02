package functions

import (
	"fmt"

	"github.com/fnwiya/purelisp/internal/types"
)

// NativeFunction は組み込み関数を表す構造体
type NativeFunction struct {
	Name string
	Fn   func(types.Value) types.Value
}

func (f *NativeFunction) IsAtom() bool {
	return false
}

func (f *NativeFunction) String() string {
	return fmt.Sprintf("#<function %s>", f.Name)
}

func Atom(args types.Value) types.Value {
	if args == nil || args.IsAtom() {
		return types.Nil
	}
	list := args.(*types.List)
	if list.Car == nil {
		return types.Nil
	}
	if list.Car.IsAtom() {
		return types.T
	}
	return types.Nil
}

func Eq(args types.Value) types.Value {
	if args == nil || args.IsAtom() {
		return types.Nil
	}
	list := args.(*types.List)
	if list.Car == nil || list.Cdr == nil || list.Cdr.IsAtom() {
		return types.Nil
	}
	v1 := list.Car
	v2 := list.Cdr.(*types.List).Car
	if v1.IsAtom() && v2.IsAtom() {
		if v1.(*types.Atom).Value == v2.(*types.Atom).Value {
			return types.T
		}
	}
	return types.Nil
}

func Car(args types.Value) types.Value {
	if args.IsAtom() {
		return types.Nil
	}
	list := args.(*types.List)
	if list.Car == nil || list.Car.IsAtom() {
		return types.Nil
	}
	return list.Car.(*types.List).Car
}

func Cdr(args types.Value) types.Value {
	if args.IsAtom() {
		return types.Nil
	}
	list := args.(*types.List)
	if list.Car == nil || list.Car.IsAtom() {
		return types.Nil
	}
	return list.Car.(*types.List).Cdr
}

func Cons(args types.Value) types.Value {
	if args.IsAtom() {
		return types.Nil
	}
	list := args.(*types.List)
	if list.Cdr == nil || list.Cdr.IsAtom() {
		return types.Nil
	}
	return &types.List{
		Car: list.Car,
		Cdr: list.Cdr.(*types.List).Car,
	}
}

func Quote(args types.Value) types.Value {
	if args.IsAtom() {
		return types.Nil
	}
	return args.(*types.List).Car
}

// NewNativeFunctions は基本的な組み込み関数を作成する
func NewNativeFunctions() map[string]types.Value {
	funcs := make(map[string]types.Value)
	funcs["atom"] = &NativeFunction{Name: "atom", Fn: Atom}
	funcs["eq"] = &NativeFunction{Name: "eq", Fn: Eq}
	funcs["car"] = &NativeFunction{Name: "car", Fn: Car}
	funcs["cdr"] = &NativeFunction{Name: "cdr", Fn: Cdr}
	funcs["cons"] = &NativeFunction{Name: "cons", Fn: Cons}
	funcs["quote"] = &NativeFunction{Name: "quote", Fn: Quote}
	return funcs
}
