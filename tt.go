package testt

import (
	"testing"
	"time"
)

type TT struct {
	tb testing.TB
}

func New(tb testing.TB) *TT {
	return &TT{tb: tb}
}

func (tt *TT) WantError(err error) {
	tt.tb.Helper()
	WantError(tt.tb, err)
}

func (tt *TT) NoError(err error) {
	tt.tb.Helper()
	NoError(tt.tb, err)
}

func (tt *TT) MustEqual(have, want any) {
	tt.tb.Helper()
	MustEqual(tt.tb, have, want)
}

func (tt *TT) Poll(d time.Duration, want any, have func() any) {
	tt.tb.Helper()
	Poll(tt.tb, d, want, have)
}

// NoAllocs fails if f allocates.
func (tt *TT) NoAllocs(f func()) {
	tt.tb.Helper()
	NoAllocs(tt.tb, f)
}
