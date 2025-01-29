package token

type Token struct {
	Type            string
	Tag             int
	BlockStartIndex int
	Text            []int
	Matches         []int
	Children        []Token
}
