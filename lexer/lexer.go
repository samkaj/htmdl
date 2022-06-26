package lexer

type contents int

const (
	HeadingOne contents = iota
	HeadingTwo
	HeadingThree
	HeadingFour
	HeadingFive
	HeadingSix
)

type token struct {
	tokenContents contents
}

func Lex(bytes *[]byte) []token {
	index := 0
	tokens := make([]token, 0)

	for index < len(*bytes) {
		c := (*bytes)[index]
		if c == byte('#') {
			headingLevel := 1
			for c == byte('#') {
				index++
				headingLevel++
				c = (*bytes)[index]
			}
			tokens = append(tokens, token{contents(headingLevel)})
		}
		index++
	}

	return tokens
}
