package main

import (
    "log"
    "os/exec"
	"time"
	"strings"
)

func main(){
	
	list_domain := []string{"https://blackhatintelligence.net","https://google.com","https://id.quora.com"}

	// domain := "https://google.com"
	output := strings.Join(list_domain, `,`)

	cmd := exec.Command(
		"/opt/google/chrome/google-chrome",
		output,
	)

	if err := cmd.Start(); err != nil {
		log.Fatalln("can't open browser", err)
	}
	
	time.Sleep(5 * time.Second)
	
	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("close the opened browser failed", err)
	}

}