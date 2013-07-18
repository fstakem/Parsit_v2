package log_parser

import (
	"time"
)

type AndroidLogLine struct {
	Timestamp		time.Time
	Pid				int
	Tid				int
	Level			AndroidLogLevel
	Source 			string
	Message			string
}

func NewAndroidLogLine(timestamp time.Time, pid int, tid int, 
					   level AndroidLogLevel, source string,
					   msg string) (log_line *AndroidLogLine) {
					   
	log_line.Timestamp = timestamp
	log_line.Pid = pid
	log_line.Tid = tid
	log_line.Level = level
	log_line.Source = source
	log_line.Message = msg
				   
	return log_line
}


