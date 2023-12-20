package lexer

type TokenCategory int32
const (
	HEADING TokenCategory = 0
	BOLD    TokenCategory = 1
	ITALIC  TokenCategory = 2
	SPACE   TokenCategory = 3
	TEXT    TokenCategory = 4
)

// A Token is an intermediary data type used when parsing the markup text.
// Each character in the text is converted to a Token corresponding it's category.
type Token struct {
	Category TokenCategory
	Value rune
	Level int32
}

func Lex(runes []rune) ([]Token, error) {
	fileLength := len(runes)
	tokens := make([]Token, 0, fileLength)
	for i := 0; i < fileLength; i++ {
		switch currentRune := runes[i]; currentRune {
		case '#':
			level := 1
			nextRune := runes[i+1]
			for nextRune == '#' {
				level += 1
				i += 1
				nextRune = runes[i+1]
			}
			tokens = append(tokens, Token{HEADING, currentRune, int32(level)})
		case '*':
			tokens = append(tokens, Token{BOLD, currentRune, 0})
		case '/':
			tokens = append(tokens, Token{ITALIC, currentRune, 0})
		case ' ':
			tokens = append(tokens, Token{SPACE, currentRune, 0})
		default:
			tokens = append(tokens, Token{TEXT, currentRune, 0})
		}
	}

	return tokens, nil
}