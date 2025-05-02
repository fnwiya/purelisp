package functions

import (
	"testing"

	"github.com/fnwiya/purelisp/internal/types"
)

func TestNativeFunctions(t *testing.T) {
	funcs := NewNativeFunctions()

	tests := []struct {
		name     string
		funcName string
		args     types.Value
		want     types.Value
	}{
		{
			name:     "atom - アトムの場合",
			funcName: "atom",
			args: &types.List{
				Car: &types.Atom{Value: "a"},
				Cdr: types.Nil,
			},
			want: types.T,
		},
		{
			name:     "atom - リストの場合",
			funcName: "atom",
			args: &types.List{
				Car: &types.List{
					Car: &types.Atom{Value: "a"},
					Cdr: &types.Atom{Value: "b"},
				},
				Cdr: types.Nil,
			},
			want: types.Nil,
		},
		{
			name:     "eq - 同じアトム",
			funcName: "eq",
			args: &types.List{
				Car: &types.Atom{Value: "a"},
				Cdr: &types.List{
					Car: &types.Atom{Value: "a"},
					Cdr: types.Nil,
				},
			},
			want: types.T,
		},
		{
			name:     "eq - 異なるアトム",
			funcName: "eq",
			args: &types.List{
				Car: &types.Atom{Value: "a"},
				Cdr: &types.List{
					Car: &types.Atom{Value: "b"},
					Cdr: types.Nil,
				},
			},
			want: types.Nil,
		},
		{
			name:     "car",
			funcName: "car",
			args: &types.List{
				Car: &types.List{
					Car: &types.Atom{Value: "a"},
					Cdr: &types.Atom{Value: "b"},
				},
				Cdr: types.Nil,
			},
			want: &types.Atom{Value: "a"},
		},
		{
			name:     "cdr",
			funcName: "cdr",
			args: &types.List{
				Car: &types.List{
					Car: &types.Atom{Value: "a"},
					Cdr: &types.Atom{Value: "b"},
				},
				Cdr: types.Nil,
			},
			want: &types.Atom{Value: "b"},
		},
		{
			name:     "cons",
			funcName: "cons",
			args: &types.List{
				Car: &types.Atom{Value: "a"},
				Cdr: &types.List{
					Car: &types.Atom{Value: "b"},
					Cdr: types.Nil,
				},
			},
			want: &types.List{
				Car: &types.Atom{Value: "a"},
				Cdr: &types.Atom{Value: "b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, ok := funcs[tt.funcName]
			if !ok {
				t.Fatalf("function %s not found", tt.funcName)
			}
			nativeFn := fn.(*NativeFunction)
			got := nativeFn.Fn(tt.args)
			if !equalValue(got, tt.want) {
				t.Errorf("%s() = %v, want %v", tt.funcName, got, tt.want)
			}
		})
	}
}

// equalValue は2つのValue型の値が等しいかどうかを判定する
func equalValue(v1, v2 types.Value) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}
	if v1.IsAtom() != v2.IsAtom() {
		return false
	}
	if v1.IsAtom() {
		return v1.(*types.Atom).Value == v2.(*types.Atom).Value
	}
	list1 := v1.(*types.List)
	list2 := v2.(*types.List)
	return equalValue(list1.Car, list2.Car) && equalValue(list1.Cdr, list2.Cdr)
}
