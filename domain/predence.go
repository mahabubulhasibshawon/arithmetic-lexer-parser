package domain

const (
	_ int = iota
	PrecedenceLowest
	PrecedenceSum     // + or -
	PrecedenceProduct // * or /
)

var precedences = map[TokenType]int{
	TokenAdd: PrecedenceSum,
	TokenSub: PrecedenceSum,
	TokenMul: PrecedenceProduct,
	TokenDiv: PrecedenceProduct,
}
