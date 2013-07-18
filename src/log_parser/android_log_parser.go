package log_parser

import (
	"errors"
	"log"
	"time"
	"strconv"
	"strings"
)

type AndroidLogParser struct {
	
}

func NewAndroidLogParser() (log_parser *AndroidLogParser) {
	return log_parser
}

func (this *AndroidLogParser) CreateLog(name string, lines []string) (android_log *AndroidLog, err error) {
	log.Println("Parsing an android log with %d lines.", len(lines))
	
	var log_lines []AndroidLogLine
	var log_tokenizer = NewAndroidLogTokenizer()
	
	for i:= 0; i < len(lines); i++ {
		tokens := log_tokenizer.CreateTokens(&lines[i])
		log_line, err := this.CreateLogLine(tokens)
		
		if err != nil {
			log_lines = append(log_lines, *log_line)
		} else {
			log.Println(err.Error())
		}
	}
	
	if len(log_lines) == 0 {
		return android_log, errors.New("No android log lines found.")
	}
	
	android_log = NewAndroidLog()
	android_log.Name = name
	android_log.LogLines = log_lines
	
	log.Println("Found %d log lines.", len(log_lines))
	
	return android_log, nil
}

func (this *AndroidLogParser) CreateLogLine(tokens []string) (log_line *AndroidLogLine, err error) {
	if len(tokens) < 6 {
		return log_line, errors.New("The android log line does not contain the correct number of tokens.")
	}
	
	timestamp, err := this.parseTimeStamp(tokens[0], tokens[1])
	if err == nil {
		return log_line, err
	}
	
	pid, err	 := this.parsePid(tokens[2])
	if err == nil {
		return log_line, err
	}
		
	tid, err	 := this.parseTid(tokens[3])
	if err == nil {
		return log_line, err
	}
	
	level, err := this.parseLogLevel(tokens[4])
	if err == nil {
		return log_line, err
	}
	
	source := this.parseSource(tokens[5])
	msg_tokens := tokens[6:]
	msg := strings.Join(msg_tokens, " ")
	
	log_line = NewAndroidLogLine(timestamp, pid, tid, level, source, msg)
	
	return log_line, nil
}

func (this *AndroidLogParser) parseTimeStamp(token_date string, token_time string) (timestamp time.Time, err error) {
	time_token := token_date + " " + token_time
	timestamp, err = time.Parse("02-23 10:13:43.573", time_token)
	if err != nil {
		err_str := "Error parsing the timestamp: " + time_token
		return timestamp, errors.New(err_str)
	}
	
	return timestamp, nil
}

func (this *AndroidLogParser) parsePid(token string) (pid int, err error) {
	pid64, err := strconv.ParseInt(token, 10, 32)
	if err != nil {
		err_str := "Error parsing the pid: " + token
		return pid, errors.New(err_str)
	}
	
	pid = int(pid64)
	
	return pid, nil
}

func (this *AndroidLogParser) parseTid(token string) (tid int, err error) {
	tid64, err := strconv.ParseInt(token, 10, 32)
	if err != nil {
		err_str := "Error parsing the tid: " + token
		return tid, errors.New(err_str)
	}
	
	tid = int(tid64)
	
	return tid, nil
}

func (this *AndroidLogParser) parseLogLevel(token string) (log_level AndroidLogLevel, err error) {
    if token == "V" {
    		return Verbose, nil
    } else if token == "I" {
    		return Info, nil
    } else if token == "D" {
    		return Debug, nil
    } else if token == "W" {
    		return Warn, nil
    } else if token == "E" {
    		return Error, nil
    } else if token == "A" {
    		return Assert, nil
    }
    
    err_str := "Error parsing the level: " + token
	return log_level, errors.New(err_str)
}

func (this *AndroidLogParser) parseSource(token string) (source string) {
	source = token[:-1]
	return source
}






