package main

import (
	"cron-parser/parser"
	"cron-parser/util"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go \"*/15 0 1,15 * 1-5 /usr/bin/find\"")
		os.Exit(1)
	}
	cronExpression := os.Args[1]
	output, err := parser.ParseCron(cronExpression)
	if err != nil {
		log.Fatalf("error : %v", err)
	}
	util.PrintCronOutput(output)
}
