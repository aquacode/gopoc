package main 

import (

	"log"
	"io"
	"bufio"

	"os"
	"os/exec"

)

const (

	gopocLogPrefix string = "gopoc: "
	gopocScript string = "./gopoctester.sh"

)

var (

	logFile *os.File 
	logger *log.Logger
	multiWriter io.Writer

)

func init() {

	var logFileError error
	logFile, logFileError = os.Create("gopoc.log")
	if logFileError != nil {
		panic(logFileError)
	}

	multiWriter = io.MultiWriter(logFile, os.Stdout)

	logger = log.New(multiWriter, gopocLogPrefix, log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetOutput(multiWriter)

}

func main() {

	defer logFile.Close()

	logger.Printf("Hello Go POC")

	cmd := exec.Command(gopocScript)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		logger.Printf(err.Error())
		panic(err)
	}
	
	multiReader := io.MultiReader(stdout, stderr)
	in := bufio.NewScanner(multiReader)
	for in.Scan() {
		logger.Printf(in.Text())
	}
	if err := in.Err(); err != nil {
		logger.Printf(err.Error())
	}

	cmd.Wait()

	logger.Printf("Done!")

}
