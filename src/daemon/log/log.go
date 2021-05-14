package mlog

import (
	"log"
	"os"
)

var ILogger = log.New(os.Stdout, "[INFO]", log.LstdFlags|log.Lshortfile)
var WLogger = log.New(os.Stdout, "[WARNING]", log.LstdFlags|log.Lshortfile)
var ELogger = log.New(os.Stdout, "[ERROR]", log.LstdFlags|log.Lshortfile)
