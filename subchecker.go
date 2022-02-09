package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
	"log"
    "os/exec"
	"time"
)

func printFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    
	defer f.Close()
    
	scanner := bufio.NewScanner(f)
	
	var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
		io.WriteString(os.Stdout, scanner.Text())
        io.WriteString(os.Stdout, "\n")
    }

	cmd := exec.Command(
		"/opt/firefox/firefox",
		lines...,
	)

	if err := cmd.Start(); err != nil {
		log.Fatalln("can't open browser", err)
	}
	
	time.Sleep(5 * time.Second)
	
	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("close the opened browser failed", err)
	}

    return nil
}

func main() {
    filename := ""
    arguments := os.Args
    if len(arguments) == 1 {
        io.Copy(os.Stdout, os.Stdin)
        return
    }

    for i := 1; i < len(arguments); i++ {
        filename = arguments[i]
        err := printFile(filename)
        if err != nil {
            fmt.Println(err)
        }
    }
}