package logger

import (
	"fmt"
	"log"
)

func Info(v ...any) {

	log.Println("[INFO]", fmt.Sprint(v...))

}

func Error(v ...any) {

	log.Println("[ERROR]", fmt.Sprint(v...))

}
