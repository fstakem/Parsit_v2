package log_parser

import (
	"strings"
)

type SyslogTokenizer struct {
}

func NewSyslogTokenizer() (syslog_tokenizer *SyslogTokenizer) {
	return syslog_tokenizer
}

func (this *SyslogTokenizer) CreateTokens(line *string) (tokens []string) {
	return strings.Fields(*line)
}
