package lexer

type Finder func(data []byte) ([]byte, []byte, []byte, error)
