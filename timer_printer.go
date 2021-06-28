package pterm

import (
	"time"
)

// DefaultTimer is the default TimerPrinter.
var DefaultTimer = TimerPrinter{
	UpdateInterval: time.Second,
	RoundingFactor: time.Second,
}

// TODO: Add theme
// TODO: Add docs
// TODO: Add styling
// TODO: Add `With` methods
// TODO: Add template output
type TimerPrinter struct {
	RoundingFactor time.Duration
	UpdateInterval time.Duration
	StartTime      time.Time
	IsActive       bool
}

// Start the TimerPrinter.
func (s TimerPrinter) Start(text ...interface{}) (*TimerPrinter, error) {
	s.IsActive = true
	s.StartTime = time.Now()

	go func() {
		for s.IsActive {
			Printo(time.Since(s.StartTime).Round(s.RoundingFactor))
			time.Sleep(s.UpdateInterval)
		}
	}()

	return &s, nil
}

// Stop the TimerPrinter.
func (s *TimerPrinter) Stop() error {
	s.IsActive = false
	Println()

	return nil
}

func (s *TimerPrinter) GenericStart() (*LivePrinter, error) {
	_, _ = s.Start()
	lp := LivePrinter(s)
	return &lp, nil
}

func (s *TimerPrinter) GenericStop() (*LivePrinter, error) {
	_ = s.Stop()
	lp := LivePrinter(s)
	return &lp, nil
}
