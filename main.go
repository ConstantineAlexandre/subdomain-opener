package main

import (
    "log"
    "os/exec"
	"time"
)

func main(){
	
	list_domain := []string{"https://blackhatintelligence.net","https://blackhatintelligence.net"}

	cmd := exec.Command(
		"/opt/google/chrome/google-chrome",
		list_domain...,
	)

	if err := cmd.Start(); err != nil {
		log.Fatalln("can't open browser", err)
	}
	
	time.Sleep(5 * time.Second)
	
	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("close the opened browser failed", err)
	}

}