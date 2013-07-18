package log_analyzer

import (
	"reflect"
	"log_parser"
)

type Rule struct {
	Field		string
	Expression	string
}

func(this *Rule) evaluate(log_line *log_parser.StructuredLogLine) (equ bool) {
	
	reflect.Value.FieldByName(this.Field)
	
	return false
}

