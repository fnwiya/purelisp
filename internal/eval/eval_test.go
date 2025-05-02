package eval

import (
	"testing"

	"github.com/fnwiya/purelisp/internal/functions"
	"github.com/fnwiya/purelisp/internal/types"
)

func TestEval(t *testing.T) {
	env := NewEnvironment(nil)
	funcs := functions.NewNativeFunctions()
	for name, fn := range funcs {
		env.SetVar(name, fn)
	}

	tests := []struct {
		name    string
		expr    types.Value
		want    types.Value
		wantErr bool
	}{
		{
			name: "quote",
			expr: &types.List{
				Car: &types.Atom{Value: "quote"},
				Cdr: &types.List{
					Car: &types.Atom{Value: "a"},
					Cdr: types.Nil,
				},
			},
			want: &types.Atom{Value: "a"},
		},
		{
			name: "cons",
			expr: &types.List{
				Car: &types.Atom{Value: "cons"},
				Cdr: &types.List{
					Car: &types.List{
						Car: &types.Atom{Value: "quote"},
						Cdr: &types.List{
							Car: &types.Atom{Value: "a"},
							Cdr: types.Nil,
						},
					},
					Cdr: &types.List{
						Car: &types.List{
							Car: &types.Atom{Value: "quote"},
							Cdr: &types.List{
								Car: &types.Atom{Value: "b"},
								Cdr: types.Nil,
							},
						},
						Cdr: types.Nil,
					},
				},
			},
			want: &types.List{
				Car: &types.Atom{Value: "a"},
				Cdr: &types.Atom{Value: "b"},
			},
		},
		{
			name: "car",
			expr: &types.List{
				Car: &types.Atom{Value: "car"},
				Cdr: &types.List{
					Car: &types.List{
						Car: &types.Atom{Value: "quote"},
						Cdr: &types.List{
							Car: &types.List{
								Car: &types.Atom{Value: "a"},
								Cdr: &types.Atom{Value: "b"},
							},
							Cdr: types.Nil,
						},
					},
					Cdr: types.Nil,
				},
			},
			want: &types.Atom{Value: "a"},
		},
		{
			name: "cdr",
			expr: &types.List{
				Car: &types.Atom{Value: "cdr"},
				Cdr: &types.List{
					Car: &types.List{
						Car: &types.Atom{Value: "quote"},
						Cdr: &types.List{
							Car: &types.List{
								Car: &types.Atom{Value: "a"},
								Cdr: &types.Atom{Value: "b"},
							},
							Cdr: types.Nil,
						},
					},
					Cdr: types.Nil,
				},
			},
			want: &types.Atom{Value: "b"},
		},
		{
			name: "atom - true",
			expr: &types.List{
				Car: &types.Atom{Value: "atom"},
				Cdr: &types.List{
					Car: &types.List{
						Car: &types.Atom{Value: "quote"},
						Cdr: &types.List{
							Car: &types.Atom{Value: "a"},
							Cdr: types.Nil,
						},
					},
					Cdr: types.Nil,
				},
			},
			want: types.T,
		},
		{
			name: "atom - false",
			expr: &types.List{
				Car: &types.Atom{Value: "atom"},
				Cdr: &types.List{
					Car: &types.List{
						Car: &types.Atom{Value: "quote"},
						Cdr: &types.List{
							Car: &types.List{
								Car: &types.Atom{Value: "a"},
								Cdr: &types.Atom{Value: "b"},
							},
							Cdr: types.Nil,
						},
					},
					Cdr: types.Nil,
				},
			},
			want: types.Nil,
		},
		{
			name: "eq - true",
			expr: &types.List{
				Car: &types.Atom{Value: "eq"},
				Cdr: &types.List{
					Car: &types.List{
						Car: &types.Atom{Value: "quote"},
						Cdr: &types.List{
							Car: &types.Atom{Value: "a"},
							Cdr: types.Nil,
						},
					},
					Cdr: &types.List{
						Car: &types.List{
							Car: &types.Atom{Value: "quote"},
							Cdr: &types.List{
								Car: &types.Atom{Value: "a"},
								Cdr: types.Nil,
							},
						},
						Cdr: types.Nil,
					},
				},
			},
			want: types.T,
		},
		{
			name: "eq - false",
			expr: &types.List{
				Car: &types.Atom{Value: "eq"},
				Cdr: &types.List{
					Car: &types.List{
						Car: &types.Atom{Value: "quote"},
						Cdr: &types.List{
							Car: &types.Atom{Value: "a"},
							Cdr: types.Nil,
						},
					},
					Cdr: &types.List{
						Car: &types.List{
							Car: &types.Atom{Value: "quote"},
							Cdr: &types.List{
								Car: &types.Atom{Value: "b"},
								Cdr: types.Nil,
							},
						},
						Cdr: types.Nil,
					},
				},
			},
			want: types.Nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eval(tt.expr, env)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !equalValue(got, tt.want) {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
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
