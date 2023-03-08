package logging

import "log"

func Errorf(format string, args ...any) {
	log.Printf("[ERROR] "+format, args...)
}

func Warnf(format string, args ...any) {
	log.Printf("[WARN] "+format, args...)
}

func Infof(format string, args ...any) {
	log.Printf("[INFO] "+format, args...)
}

func Debugf(format string, args ...any) {
	log.Printf("[DEBUG] "+format, args...)
}

func Tracef(format string, args ...any) {
	log.Printf("[TRACE] "+format, args...)
}
