package main

import (
	"encoding/base64"
	"flag"
	"github.com/bayucandra/shell-notify/internal"
	"github.com/gen2brain/beeep"
	"log"
	"os"
)

func main() {

	var err error

	iconPath := iconExtract()

	if os.Getenv("GO_ENV") == "development" || os.Getenv("GO_ENV") == "test" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	titlePtr := flag.String("title", "Shell notify", "Title of notification")
	messagePtr := flag.String("message", "-", "Message of notification")
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("error, you must specify message to notify")
	}

	err = beeep.Notify(*titlePtr, *messagePtr, iconPath)
	if err != nil {
		panic(err)
	}

}

func iconExtract() string {
	dec, err := base64.StdEncoding.DecodeString(internal.Icon)
	if err != nil {
		panic(err)
	}

	basePath := os.TempDir() + "biqdev.com"
	err = os.MkdirAll(basePath, 0755)

	if err != nil {
		panic(err)
	}

	extractPath := basePath + "/shell-notify.png"

	f, err := os.Create(extractPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	return extractPath
}
