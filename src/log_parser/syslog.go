package log_parser

import (

)

type Syslog struct {
	Name			string
	LogLines		[]SyslogLine
}

func NewSyslog() (syslog *Syslog) {
	syslog.Name = "Generic Log"

	return syslog
}

