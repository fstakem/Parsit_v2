package log_parser

import (
	"errors"
)

type SyslogParser struct {
	syslog_tokenizer 	*SyslogTokenizer
}

func NewSyslogParser() (syslog_parser *SyslogParser) {
	return syslog_parser
}

func (this *SyslogParser) CreateLog(name string, lines []string) (syslog *Syslog, err error) {
	var log_lines []SyslogLine
	this.syslog_tokenizer = NewSyslogTokenizer()
	
	for i:= 0; i < len(lines); i++ {
		tokens := this.syslog_tokenizer.CreateTokens(&lines[i])
		log_line, err := this.CreateLogLine(tokens)
		
		if err != nil {
			log_lines = append(log_lines, *log_line)
		}
	}
	
	if len(log_lines) == 0 {
		return syslog, errors.New("No syslog lines found.")
	}
	
	syslog = NewSyslog()
	syslog.Name = name
	syslog.LogLines = log_lines
	
	return syslog, nil
}

func (this *SyslogParser) CreateLogLine(tokens []string) (syslog_line *SyslogLine, err error) {
	if len(tokens) < 1 {
		return syslog_line, errors.New("The syslog line does not contain the correct number of tokens.")
	}
	
	// TODO -> parse tokens
	
	// TODO -> pass in tokens and create constructor
	syslog_line = NewSyslogLine()
	
	return syslog_line, nil
}

// TODO -> parse functions : timestamp etc




