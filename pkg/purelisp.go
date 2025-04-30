package purelisp

import (
	"github.com/fnwiya/purelisp/internal/eval"
	"github.com/fnwiya/purelisp/internal/functions"
	"github.com/fnwiya/purelisp/internal/types"
)

// Environment は変数の環境を表す構造体
type Environment = eval.Environment

// NewEnvironment は新しい環境を作成する関数
func NewEnvironment(parent *Environment) *Environment {
	return eval.NewEnvironment(parent)
}

// Eval は式を評価する関数
func Eval(expr types.Value, env *Environment) (types.Value, error) {
	return eval.Eval(expr, env)
}

// NewNativeFunctions は基本的な組み込み関数を作成する
func NewNativeFunctions() map[string]types.Value {
	return functions.NewNativeFunctions()
}
