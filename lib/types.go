package lib


type TokenList map[string]float64

type Doc struct {
	Content string    `json:"content"`
	Heading string    `json:"heading"`
	Id      string    `json:"id"`
	Path    string    `json:"path"`
	Tokens  TokenList `json:"tokens"`
}
