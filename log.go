package main

import "log"

func logInfo(msg string) {
	log.Printf("[INFO] %s", msg)
}

func logWarn(msg string) {
	log.Printf("[WARN] %s", msg)
}

func logFatal(msg string) {
	log.Fatalf("[FATAL] %s", msg)
}

func logDebug(msg string) {
	log.Printf("[DEBUG] %s", msg)
}
