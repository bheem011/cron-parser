package main

import (
	"cron-parser/parser"

	"reflect"
	"testing"
)

func Test_parseCron(t *testing.T) {
	tests := []struct {
		name string
		args string
		data parser.CronOutput
		Err  bool
	}{
		{
			name: "test1",
			args: "*/15 1-10/2 1,15 1 1-5 /usr/bin/find",
			data: parser.CronOutput{
				Minute:     []string{"0", "15", "30", "45"},
				Hour:       []string{"1", "3", "5", "7", "9"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1"},
				DayOfWeek:  []string{"1", "2", "3", "4", "5"},
				Command:    []string{"/usr/bin/find"},
			},
			Err: false,
		},
		{
			name: "test2",
			args: "90/15 0 1,15 * 1-5 /usr/bin/find",
			data: parser.CronOutput{},
			Err:  true,
		},
		{
			name: "test2",
			args: "*/15 26 1,15 * 1-5 /usr/bin/find",
			data: parser.CronOutput{},
			Err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedData, err := parser.ParseCron(tt.args)
			if (err != nil) != tt.Err {
				t.Errorf("parseCron() error = %v, wantErr %v", err, tt.Err)
				return
			}
			if !reflect.DeepEqual(parsedData, tt.data) {
				t.Errorf("parseCron() = %v, want %v", parsedData, tt.data)
			}
		})
	}
}
