package log_parser

import (

)

type Tokenizer interface {
	CreateTokens(line *string)	[]string
}
