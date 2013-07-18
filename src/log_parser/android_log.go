package log_parser

import (

)

type AndroidLogLevel int

const (
	Verbose AndroidLogLevel = iota+1
    Info  
    Debug
    Warn
    Error
    Assert
)

type AndroidLog struct {
	Name				string
	LogLines			[]AndroidLogLine
}

func NewAndroidLog() (log *AndroidLog) {
	log.Name = "Android Log"

	return log
}


