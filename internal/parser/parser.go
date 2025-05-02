package parser

import (
	"fmt"
	"strings"

	"github.com/fnwiya/purelisp/internal/types"
)

// Parse は文字列からLispの式をパースする
func Parse(input string) (types.Value, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, fmt.Errorf("empty input")
	}

	// トークン化
	tokens := tokenize(input)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("no tokens found")
	}

	// パース
	value, remaining, err := parseTokens(tokens)
	if err != nil {
		return nil, err
	}
	if len(remaining) > 0 {
		return nil, fmt.Errorf("unexpected tokens after expression")
	}
	return value, nil
}

func tokenize(input string) []string {
	// スペースで分割し、括弧を個別のトークンとして扱う
	input = strings.ReplaceAll(input, "(", " ( ")
	input = strings.ReplaceAll(input, ")", " ) ")
	return strings.Fields(input)
}

func parseTokens(tokens []string) (types.Value, []string, error) {
	if len(tokens) == 0 {
		return nil, tokens, fmt.Errorf("unexpected end of input")
	}

	token := tokens[0]
	remaining := tokens[1:]

	if token == "(" {
		// リストの開始
		var elements []types.Value
		for len(remaining) > 0 && remaining[0] != ")" {
			value, newRemaining, err := parseTokens(remaining)
			if err != nil {
				return nil, remaining, err
			}
			elements = append(elements, value)
			remaining = newRemaining
		}
		if len(remaining) == 0 {
			return nil, remaining, fmt.Errorf("unclosed parenthesis")
		}
		// ")" を消費
		remaining = remaining[1:]

		// リストを構築
		result := types.Nil
		for i := len(elements) - 1; i >= 0; i-- {
			result = &types.List{
				Car: elements[i],
				Cdr: result,
			}
		}
		return result, remaining, nil
	} else if token == ")" {
		return nil, remaining, fmt.Errorf("unexpected )")
	} else {
		// アトム
		return &types.Atom{Value: token}, remaining, nil
	}
}
