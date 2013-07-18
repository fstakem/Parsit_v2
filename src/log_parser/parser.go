package log_parser

import (

)

type Parser interface {
	CreateLog(name string, lines []string)		*StructuredLog
	CreateLogLine(tokens []string)				*StructuredLogLine
}

