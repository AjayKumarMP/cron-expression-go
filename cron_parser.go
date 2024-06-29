package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Cron represents a parsed cron expression.
type Cron struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

// ParseCronExpression parses a cron expression string into a CronExpression struct.
func ParseCron(expr string) (*Cron, error) {
	parts := strings.Fields(expr)
	if len(parts) != 6 {
		return nil, fmt.Errorf("invalid cron expression: %s", expr)
	}

	return &Cron{
		Minute:     parts[0],
		Hour:       parts[1],
		DayOfMonth: parts[2],
		Month:      parts[3],
		DayOfWeek:  parts[4],
		Command:    parts[5],
	}, nil
}

// expandField expands a cron field into a list of values.
func expandField(field string, min int, max int) ([]int, error) {
	var values []int
	if field == "*" {
		for i := min; i <= max; i++ {
			values = append(values, i)
		}
	} else if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		base := parts[0]
		step, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		start := min
		if base != "*" {
			start, err = strconv.Atoi(base)
			if err != nil {
				return nil, err
			}
		}
		for i := start; i <= max; i += step {
			values = append(values, i)
		}
	} else if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		for _, part := range parts {
			value, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
	} else if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		for i := start; i <= end; i++ {
			values = append(values, i)
		}
	} else {
		value, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

// ExpandCronExpression expands all fields in a CronExpression into their individual values.
func ExpandCron(cronExpr *Cron) (map[string][]int, error) {
	expanded := make(map[string][]int)

	minuteValues, err := expandField(cronExpr.Minute, 0, 59)
	if err != nil {
		return nil, err
	}
	expanded["minute"] = minuteValues

	hourValues, err := expandField(cronExpr.Hour, 0, 23)
	if err != nil {
		return nil, err
	}
	expanded["hour"] = hourValues

	dayOfMonthValues, err := expandField(cronExpr.DayOfMonth, 1, 31)
	if err != nil {
		return nil, err
	}
	expanded["day of month"] = dayOfMonthValues

	monthValues, err := expandField(cronExpr.Month, 1, 12)
	if err != nil {
		return nil, err
	}
	expanded["month"] = monthValues

	dayOfWeekValues, err := expandField(cronExpr.DayOfWeek, 0, 6)
	if err != nil {
		return nil, err
	}
	expanded["day of week"] = dayOfWeekValues

	return expanded, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: your-program \"<cron expression>\"")
		return
	}

	expr := os.Args[1]
	cronExpr, err := ParseCron(expr)
	if err != nil {
		fmt.Println("Error parsing cron expression:", err)
		return
	}

	expandedFields, err := ExpandCron(cronExpr)
	if err != nil {
		fmt.Println("Error expanding fields:", err)
		return
	}

	for label, values := range expandedFields {
		fmt.Printf("%s ", label)
		for _, value := range values {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	fmt.Printf("command %s\n", cronExpr.Command)
}
