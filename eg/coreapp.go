package main

import (
	"log"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	qtrt.SetDebugFFICall(true)

	argv := []string{"./coreapp", "-v", "-x"}
	app := qtcore.NewQCoreApplication(len(argv), argv, 0)
	log.Println(app)
	go func() {
		time.Sleep(5 * time.Second)
		app.Exit(0)
	}()
	app.Exec()
}
