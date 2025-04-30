package main

import (
	"fmt"
	"os"

	"github.com/fnwiya/purelisp/internal/types"
	purelisp "github.com/fnwiya/purelisp/pkg"
)

func main() {
	env := purelisp.NewEnvironment(nil)

	// 基本的な関数を環境に登録
	funcs := purelisp.NewNativeFunctions()
	for name, fn := range funcs {
		env.SetVar(name, fn)
	}

	// サンプルコードの実行
	expr := &types.List{
		Car: &types.Atom{Value: "cons"},
		Cdr: &types.List{
			Car: &types.Atom{Value: "a"},
			Cdr: &types.List{
				Car: &types.Atom{Value: "b"},
				Cdr: types.Nil,
			},
		},
	}

	result, err := purelisp.Eval(expr, env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", result)
}
