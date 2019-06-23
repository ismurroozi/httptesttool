package libs

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Printf("ERROR|TestAbort|err message:%v", err)
		os.Exit(1)
	}
}
