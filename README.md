# gopoc  
## Go Proof of Concept application (development)

The purpose of this repository is to explore patterns for the following requirements:

* Create installable application that runs from command line, i.e. create a main() function
* Leverage global const and var declarations
* Use an init() function to initialize program
* Setup/include basic logging that:
  1. Writes to a log file
  2. Also writes to the console
* Use the "flag" package for command line arguments
* Launch a command that calls another shell script
  1. Send the output to the Logger
* Extract functionality into funcs that main() calls

## Install and run

```bash
go get github.com/aquacode/gopoc
go install
gopoc
```
