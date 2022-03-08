package logging

import "log"

func Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Printf("[WARN] "+format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args...)
}

func Tracef(format string, args ...interface{}) {
	log.Printf("[TRACE] "+format, args...)
}
