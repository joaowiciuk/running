package running

import (
	"testing"
	"time"
)

func TestR_Run(t *testing.T) {
	type fields struct {
		runningTimes map[int]time.Duration
		N            int
	}
	type args struct {
		input     Sizer
		algorithm func(interface{}) interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &R{
				runningTimes: tt.fields.runningTimes,
				N:            tt.fields.N,
			}
			r.Run(tt.args.input, tt.args.algorithm)
		})
	}
}
