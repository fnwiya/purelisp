package functions

import (
	"github.com/fnwiya/purelisp/internal/types"
)

// NativeFunction は組み込み関数を表す構造体
type NativeFunction struct {
	name string
	Fn   func(args types.Value) types.Value
}

func (n *NativeFunction) IsAtom() bool {
	return false
}

func (n *NativeFunction) String() string {
	return "#<native-function:" + n.name + ">"
}

// NewNativeFunctions は基本的な組み込み関数を作成する
func NewNativeFunctions() map[string]types.Value {
	funcs := make(map[string]types.Value)

	funcs["atom"] = &NativeFunction{
		name: "atom",
		Fn: func(args types.Value) types.Value {
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
		},
	}

	funcs["eq"] = &NativeFunction{
		name: "eq",
		Fn: func(args types.Value) types.Value {
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
		},
	}

	funcs["car"] = &NativeFunction{
		name: "car",
		Fn: func(args types.Value) types.Value {
			if args == nil || args.IsAtom() {
				return types.Nil
			}
			list := args.(*types.List)
			if list.Car == nil {
				return types.Nil
			}
			if list.Car.IsAtom() {
				return types.Nil
			}
			return list.Car.(*types.List).Car
		},
	}

	funcs["cdr"] = &NativeFunction{
		name: "cdr",
		Fn: func(args types.Value) types.Value {
			if args == nil || args.IsAtom() {
				return types.Nil
			}
			list := args.(*types.List)
			if list.Car == nil {
				return types.Nil
			}
			if list.Car.IsAtom() {
				return types.Nil
			}
			return list.Car.(*types.List).Cdr
		},
	}

	funcs["cons"] = &NativeFunction{
		name: "cons",
		Fn: func(args types.Value) types.Value {
			if args == nil || args.IsAtom() {
				return types.Nil
			}
			list := args.(*types.List)
			if list.Car == nil || list.Cdr == nil || list.Cdr.IsAtom() {
				return types.Nil
			}
			return &types.List{
				Car: list.Car,
				Cdr: list.Cdr.(*types.List).Car,
			}
		},
	}

	return funcs
}
