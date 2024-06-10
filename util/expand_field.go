package util

import (
	"strconv"
	"strings"
)

// ExpandField parse the field and get desire time output
func ExpandField(field string, min, max int) []string {
	// Handle wildcard '*' which represents every value in the range
	if field == "*" {
		return generateRange(min, max, 1)
	}

	values := []string{}
	parts := strings.Split(field, ",")

	for _, part := range parts {
		if strings.Contains(part, "/") && strings.Contains(part, "-") {
			//Handle range with steps (e.g., 1-12/2)
			subParts := strings.Split(part, "/")
			rangePart := subParts[0]
			step, _ := strconv.Atoi(subParts[1])
			rangeParts := strings.Split(rangePart, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			values = append(values, generateRange(start, end, step)...)
		} else if strings.Contains(part, "/") {
			// Handle step values (e.g., */15)
			subParts := strings.Split(part, "/")
			step, _ := strconv.Atoi(subParts[1])
			if subParts[0] == "*" {
				values = append(values, generateRange(min, max, step)...)
			}
		} else if strings.Contains(part, "-") {
			// Handle ranges (e.g., 1-5)
			subParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(subParts[0])
			end, _ := strconv.Atoi(subParts[1])
			values = append(values, generateRange(start, end, 1)...)
		} else {
			// Handle single values (e.g., 1)
			values = append(values, part)
		}
	}

	return values
}

// generateRange generates a list of string values from start to end with a given step
func generateRange(start, end, step int) []string {
	var values []string
	for i := start; i <= end; i += step {
		values = append(values, strconv.Itoa(i))
	}
	return values
}
