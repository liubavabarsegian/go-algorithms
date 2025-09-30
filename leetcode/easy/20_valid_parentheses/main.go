package main

func isValid(s string) bool {
	// { -> записать в стек
	// ( -> записать в стек
	// } -> сравить и закрыть или записать
	// ) -> сравнить и закрыть или записать

	symbols := make([]byte, 0, len(s))

	for _, v := range []byte(s) {
		if len(symbols) > 0 && (v == '}' || v == ')' || v == ']') && canClose(symbols[len(symbols)-1], v) {
			symbols = symbols[0 : len(symbols)-1]
		} else {
			symbols = append(symbols, v)
		}
	}

	return len(symbols) == 0
}

func canClose(a byte, b byte) bool {
	return a == '{' && b == '}' || a == '(' && b == ')' || a == '[' && b == ']'
}
