package parser

import (
	"testing"

	"github.com/fnwiya/purelisp/internal/types"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    types.Value
		wantErr bool
	}{
		{
			name:  "単一のアトム",
			input: "a",
			want:  &types.Atom{Value: "a"},
		},
		{
			name:  "quote式",
			input: "(quote a)",
			want: &types.List{
				Car: &types.Atom{Value: "quote"},
				Cdr: &types.List{
					Car: &types.Atom{Value: "a"},
					Cdr: types.Nil,
				},
			},
		},
		{
			name:  "cons式",
			input: "(cons (quote a) (quote b))",
			want: &types.List{
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
		},
		{
			name:    "空の入力",
			input:   "",
			wantErr: true,
		},
		{
			name:    "閉じ括弧がない",
			input:   "(quote a",
			wantErr: true,
		},
		{
			name:    "開き括弧がない",
			input:   "quote a)",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !equalValue(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
