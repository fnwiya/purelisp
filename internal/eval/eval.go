package eval

import (
	"fmt"

	"github.com/fnwiya/purelisp/internal/functions"
	"github.com/fnwiya/purelisp/internal/types"
)

// Environment は変数の環境を表す構造体
type Environment struct {
	parent *Environment
	vars   map[string]types.Value
}

// NewEnvironment は新しい環境を作成する関数
func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent: parent,
		vars:   make(map[string]types.Value),
	}
}

// SetVar は環境に変数を設定する関数
func (e *Environment) SetVar(name string, value types.Value) {
	e.vars[name] = value
}

// GetVar は環境から変数を取得する関数
func (e *Environment) GetVar(name string) (types.Value, bool) {
	if val, ok := e.vars[name]; ok {
		return val, true
	}
	if e.parent != nil {
		return e.parent.GetVar(name)
	}
	return nil, false
}

// Eval は式を評価する関数
func Eval(expr types.Value, env *Environment) (types.Value, error) {
	if expr == nil {
		return types.Nil, nil
	}

	switch {
	case expr.IsAtom():
		atom := expr.(*types.Atom)
		if val, ok := env.GetVar(atom.Value); ok {
			return val, nil
		}
		return expr, nil
	default:
		list := expr.(*types.List)
		if list.Car == nil {
			return types.Nil, nil
		}

		// 特殊形式の処理
		if atom, ok := list.Car.(*types.Atom); ok {
			switch atom.Value {
			case "quote":
				if list.Cdr == nil || list.Cdr.IsAtom() {
					return nil, fmt.Errorf("invalid quote form")
				}
				return list.Cdr.(*types.List).Car, nil
			case "lambda":
				if list.Cdr == nil || list.Cdr.IsAtom() {
					return nil, fmt.Errorf("invalid lambda form")
				}
				return &Lambda{
					Params: list.Cdr.(*types.List).Car,
					Body:   list.Cdr.(*types.List).Cdr.(*types.List).Car,
					Env:    env,
				}, nil
			case "define":
				if list.Cdr == nil || list.Cdr.IsAtom() {
					return nil, fmt.Errorf("invalid define form")
				}
				name := list.Cdr.(*types.List).Car.(*types.Atom).Value
				val, err := Eval(list.Cdr.(*types.List).Cdr.(*types.List).Car, env)
				if err != nil {
					return nil, err
				}
				env.SetVar(name, val)
				return val, nil
			case "cond":
				return evalCond(list.Cdr, env)
			}
		}

		// 関数適用
		fn, err := Eval(list.Car, env)
		if err != nil {
			return nil, err
		}

		// 引数の評価
		var args types.Value = types.Nil
		var current types.Value = types.Nil
		for p := list.Cdr; p != nil && !p.IsAtom(); p = p.(*types.List).Cdr {
			arg, err := Eval(p.(*types.List).Car, env)
			if err != nil {
				return nil, err
			}
			newList := &types.List{Car: arg, Cdr: types.Nil}
			if args == types.Nil {
				args = newList
				current = newList
			} else {
				current.(*types.List).Cdr = newList
				current = newList
			}
		}

		switch f := fn.(type) {
		case *Lambda:
			return f.Apply(args, env)
		case *functions.NativeFunction:
			return f.Fn(args), nil
		default:
			return nil, fmt.Errorf("not a function: %v", fn)
		}
	}
}

// Lambda はラムダ式を表す構造体
type Lambda struct {
	Params types.Value
	Body   types.Value
	Env    *Environment
}

func (l *Lambda) IsAtom() bool {
	return false
}

func (l *Lambda) String() string {
	return fmt.Sprintf("(lambda %v %v)", l.Params, l.Body)
}

// Apply はラムダ式を適用する関数
func (l *Lambda) Apply(args types.Value, env *Environment) (types.Value, error) {
	newEnv := NewEnvironment(l.Env)
	params := l.Params
	argList := args

	for !params.IsAtom() && !argList.IsAtom() {
		param := params.(*types.List).Car.(*types.Atom).Value
		arg := argList.(*types.List).Car
		newEnv.SetVar(param, arg)
		params = params.(*types.List).Cdr
		argList = argList.(*types.List).Cdr
	}

	return Eval(l.Body, newEnv)
}

// evalCond はcond式を評価する関数
func evalCond(expr types.Value, env *Environment) (types.Value, error) {
	if expr == nil || expr.IsAtom() {
		return types.Nil, nil
	}

	clause := expr.(*types.List).Car
	if clause.IsAtom() {
		return nil, fmt.Errorf("invalid cond clause: %v", clause)
	}

	condition := clause.(*types.List).Car
	result := clause.(*types.List).Cdr.(*types.List).Car

	condVal, err := Eval(condition, env)
	if err != nil {
		return nil, err
	}

	if !condVal.IsAtom() || condVal.(*types.Atom).Value != "NIL" {
		return Eval(result, env)
	}

	return evalCond(expr.(*types.List).Cdr, env)
}
