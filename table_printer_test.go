package pterm

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTablePrinter_NilPrint(t *testing.T) {
	p := TablePrinter{}
	p.Render()
}
