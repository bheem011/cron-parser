package util

import (
	"cron-parser/parser"
	"fmt"
	"strings"
)

func PrintCronOutput(cron parser.CronOutput) {
	fmt.Printf("%-14s%s\n", "minute:", strings.Join(cron.Minute, " "))
	fmt.Printf("%-14s%s\n", "hour:", strings.Join(cron.Hour, " "))
	fmt.Printf("%-14s%s\n", "day of Month:", strings.Join(cron.DayOfMonth, " "))
	fmt.Printf("%-14s%s\n", "month:", strings.Join(cron.Month, " "))
	fmt.Printf("%-14s%s\n", "day of Week:", strings.Join(cron.DayOfWeek, " "))
	fmt.Printf("%-14s%s\n", "command:", strings.Join(cron.Command, " "))
}
