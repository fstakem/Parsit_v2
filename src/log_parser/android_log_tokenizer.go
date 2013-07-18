package log_parser

import (
	"strings"
)

type AndroidLogTokenizer struct {
}

func NewAndroidLogTokenizer() (log_tokenizer *AndroidLogTokenizer) {
	return log_tokenizer
}

func (this *AndroidLogTokenizer) CreateTokens(line *string) (tokens []string) {
	return strings.Fields(*line)
}
