package main

import (
	"runtime"

	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"

	"os"
	"os/exec"
)

const (
	gopocLogfile   string = "gopoc-%v.log"
	gopocLogPrefix string = "gopoc: "
	gopocScript    string = "./gopoctester.sh"
	gopocBatch     string = "gopoctester.bat"
)

var (
	sleepDurationSec int

	logFile     *os.File
	logger      *log.Logger
	multiWriter io.Writer
)

func init() {

	flag.IntVar(&sleepDurationSec, "sleepDurationSec", 5, "Set the number of seconds to sleep during cmd execution")

	var logFileError error
	logFile, logFileError = os.Create("gopoc.log")
	if logFileError != nil {
		panic(logFileError)
	}

	multiWriter = io.MultiWriter(logFile, os.Stdout)

	logger = log.New(multiWriter, gopocLogPrefix, log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetOutput(multiWriter)

}

func main() {

	flag.Parse()

	defer logFile.Close() // close logFile after main() exits

	logger.Printf("Go POC")

	strformat := "help-%s"
	yesser := fmt.Sprintf(strformat, "me!")
	logger.Printf(yesser)

	if runtime.GOOS == "windows" {
		runCommand(gopocBatch, strconv.Itoa(sleepDurationSec))
	} else {
		runCommand(gopocScript, strconv.Itoa(sleepDurationSec))
	}

	logger.Printf("Done!")

}

func runCommand(name string, args ...string) {

	logger.Printf("Executing Cmd:")
	logger.Printf("\tname: %s", name)
	logger.Printf("\targs: %s", args)

	cmd := exec.Command(name, args...)
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

}
