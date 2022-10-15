package testt

import (
	"reflect"
	"testing"
	"time"
)

func WantError(tb testing.TB, err error) {
	tb.Helper()
	if err == nil {
		tb.Fail()
	}
}

func NoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

func MustEqual(tb testing.TB, have, want any) {
	tb.Helper()
	if !reflect.DeepEqual(have, want) {
		tb.Fatalf("\nhave: %+v\nwant: %+v\n", have, want)
	}
}

const pollRate = 20

func Poll(tb testing.TB, d time.Duration, want any, have func() any) {
	tb.Helper()

	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		// waiting a bit before 1st try
		time.Sleep(d / pollRate)

		if reflect.DeepEqual(have(), want) {
			return
		}
	}

	MustEqual(tb, have(), want)
}
