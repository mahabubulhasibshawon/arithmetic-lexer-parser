package domain

type TokenType int

const (
	TokenEof TokenType = iota
	TokenNumber
	TokenAdd
	TokenSub
	TokenMul
	TokenDiv
	TokenLParen
	TokenRParen
)

type Token struct {
	Type  TokenType
	Value string
}
