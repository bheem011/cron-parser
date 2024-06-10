package main

import (
	"fmt"
	"os"
	"strings"

	"log"

	"cron-parser/util"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go \"*/15 0 1,15 * 1-5 /usr/bin/find\"")
		os.Exit(1)
	}

	cron := os.Args[1]

	if !util.ValidCornParser(cron) {
		log.Fatalf("Invalid cron string format.")
	}

	parts := strings.Fields(cron)
	if len(parts) != 6 {
		log.Fatalf("Invalid number of parts present expected 6 found %d\n", len(parts))
	}

	minuteField := parts[0]
	hourField := parts[1]
	dayOfMonthField := parts[2]
	monthField := parts[3]
	dayOfWeekField := parts[4]
	command := parts[5]

	minutes := util.ExpandField(minuteField, 0, 59)
	hours := util.ExpandField(hourField, 0, 23)
	daysOfMonth := util.ExpandField(dayOfMonthField, 1, 31)
	months := util.ExpandField(monthField, 1, 12)
	daysOfWeek := util.ExpandField(dayOfWeekField, 0, 6)

	PrintTimes(minutes, hours, daysOfMonth, months, daysOfWeek, command)
}

func PrintTimes(minutes, hours, daysOfMonth, months, daysOfWeek []string, command string) {
	fmt.Printf("%-14s%s\n", "minute", strings.Join(minutes, " "))
	fmt.Printf("%-14s%s\n", "hour", strings.Join(hours, " "))
	fmt.Printf("%-14s%s\n", "day of month", strings.Join(daysOfMonth, " "))
	fmt.Printf("%-14s%s\n", "month", strings.Join(months, " "))
	fmt.Printf("%-14s%s\n", "day of week", strings.Join(daysOfWeek, " "))
	fmt.Printf("%-14s%s\n", "command", command)
}
