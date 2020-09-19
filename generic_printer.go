package pterm

type GenericPrinter interface {
	Sprint(a ...interface{}) string
	Sprintln(a ...interface{}) string
	Sprintf(format string, a ...interface{}) string
	Print(a ...interface{}) GenericPrinter
	Println(a ...interface{}) GenericPrinter
	Printf(format string, a ...interface{}) GenericPrinter
}
