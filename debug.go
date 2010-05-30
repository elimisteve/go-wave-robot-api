package waveapi

import (
	l "log"
)

var Debug = false

func log(v ...interface{}) {
	if !Debug {
		return
	}
	l.Stderr(v)
}

func logf(f string, v ...interface{}) {
	if !Debug {
		return
	}
	l.Stderrf(f, v)
}
