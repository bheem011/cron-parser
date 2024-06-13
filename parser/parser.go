package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type CronOutput struct {
	Minute     []string
	Hour       []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    []string
}

type FieldParser interface {
	Parse(field string) ([]string, error)
}

type MinuteParser struct{}
type HourParser struct{}
type DayOfMonthParser struct{}
type MonthParser struct{}
type DayOfWeekParser struct{}
type CommandParser struct{}

func (p *MinuteParser) Parse(field string) ([]string, error) {
	return parseField(field, 0, 59)
}

func (p *HourParser) Parse(field string) ([]string, error) {
	return parseField(field, 0, 23)
}

func (p *DayOfMonthParser) Parse(field string) ([]string, error) {
	return parseField(field, 1, 31)
}

func (p *MonthParser) Parse(field string) ([]string, error) {
	return parseField(field, 1, 12)
}

func (p *DayOfWeekParser) Parse(field string) ([]string, error) {
	return parseField(field, 0, 6)
}

func (p *CommandParser) Parse(field string) ([]string, error) {
	return []string{field}, nil
}

// ParseCron parses the cron string to give cron output
func ParseCron(cronExpression string) (CronOutput, error) {

	fields := strings.Fields(cronExpression)
	if len(fields) != 6 {
		return CronOutput{}, fmt.Errorf("invalid cron expression")
	}

	parsers := []FieldParser{
		&MinuteParser{},
		&HourParser{},
		&DayOfMonthParser{},
		&MonthParser{},
		&DayOfWeekParser{},
		&CommandParser{},
	}

	var output CronOutput

	for i, field := range fields {
		parsedValues, err := parsers[i].Parse(field)
		if err != nil {
			return CronOutput{}, err
		}
		switch i {
		case 0:
			output.Minute = parsedValues
		case 1:
			output.Hour = parsedValues
		case 2:
			output.DayOfMonth = parsedValues
		case 3:
			output.Month = parsedValues
		case 4:
			output.DayOfWeek = parsedValues
		case 5:
			output.Command = parsedValues
		}
	}
	return output, nil

}

// parseField parses the individual cron field
func parseField(field string, min, max int) ([]string, error) {
	if field == "*" {
		return generateRange(min, max, 1), nil
	}
	values := []string{}
	parts := strings.Split(field, ",")
	for _, part := range parts {
		if strings.Contains(part, "/") && strings.Contains(part, "-") {
			//Handle range with steps (e.g., 1-12/2)
			subParts := strings.Split(part, "/")
			rangePart := subParts[0]
			step, err := strconv.Atoi(subParts[1])
			if err != nil {
				return values, err
			}
			rangeParts := strings.Split(rangePart, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			err = rangeValidator(field, min, max, start, end)
			if err != nil {
				return values, err
			}
			values = append(values, generateRange(start, end, step)...)
		} else if strings.Contains(part, "/") {
			// Handle step values (e.g., */15)
			subParts := strings.Split(part, "/")
			step, err := strconv.Atoi(subParts[1])
			if err != nil {
				return values, err
			}
			if subParts[0] == "*" {
				values = append(values, generateRange(min, max, step)...)
			} else {
				start, err := strconv.Atoi(subParts[0])
				if err != nil {
					return values, err
				}
				err = rangeValidator(field, min, max, start, start)
				if err != nil {
					return values, err
				}
				values = append(values, generateRange(start, max, step)...)
			}
		} else if strings.Contains(part, "-") {
			// Handle ranges (e.g., 1-5)
			subParts := strings.Split(part, "-")
			start, err := strconv.Atoi(subParts[0])
			if err != nil {
				return values, err
			}
			end, err := strconv.Atoi(subParts[1])
			if err != nil {
				return values, err
			}
			err = rangeValidator(field, min, max, start, end)
			if err != nil {
				return values, err
			}
			values = append(values, generateRange(start, end, 1)...)
		} else {
			// Handle single values (e.g., 1)
			start, err := strconv.Atoi(part)
			if err != nil {
				return values, err
			}
			err = rangeValidator(field, min, max, start, start)
			if err != nil {
				return values, err
			}
			values = append(values, part)
		}
	}

	return values, nil
}

// generateRange generates a list of string values from start to end with a given step
func generateRange(start, end, step int) []string {
	var values []string
	for i := start; i <= end; i += step {
		values = append(values, strconv.Itoa(i))
	}
	return values
}

// rangeValidator validate the field range
func rangeValidator(field string, min, max, start, end int) error {
	var err error
	if start < min || end > max {
		err = fmt.Errorf("%s is not in the range, valid range is %d-%d", field, min, max)
	}
	return err
}
