# Printer Templates

## Live Printer

```go
package pterm

import (

)

// DefaultTemplate is the default TemplatePrinter.
var DefaultTemplate = TemplatePrinter {
	
}


type TemplatePrinter struct {

}

// Start the TemplatePrinter.
func (s TemplatePrinter) Start(text ...interface{}) (*TemplatePrinter, error) {
	
	return &s, nil
}

// Stop the TemplatePrinter.
func (s *TemplatePrinter) Stop() error {
	
	return nil
}

func (s *TemplatePrinter) GenericStart() (*LivePrinter, error) {
	_, _ = s.Start()
	lp := LivePrinter(s)
	return &lp, nil
}

func (s *TemplatePrinter) GenericStop() (*LivePrinter, error) {
	_ = s.Stop()
	lp := LivePrinter(s)
	return &lp, nil
}
```