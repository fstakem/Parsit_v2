package log_parser

import (
	"time"
)

type SyslogLine struct {
	Timestamp		time.Time
}

func NewSyslogLine() (syslog_line *SyslogLine) {
	return nil
}

