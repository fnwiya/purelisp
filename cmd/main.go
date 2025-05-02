package main

import (
	"fmt"
	"os"

	purelisp "github.com/fnwiya/purelisp/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: purelisp <expression>")
		os.Exit(1)
	}

	env := purelisp.NewEnvironment()

	// 基本的な関数を環境に登録
	funcs := purelisp.NewNativeFunctions()
	for name, fn := range funcs {
		env.SetVar(name, fn)
	}

	// コマンドライン引数から式をパース
	expr, err := purelisp.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		os.Exit(1)
	}

	result, err := purelisp.Eval(expr, env)
	if err != nil {
		fmt.Printf("Eval error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
