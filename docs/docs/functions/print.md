# Print Methods

> The global exposted print functions are listed here. Use them as an alternative for the print functions in`fmt`, because they support colors and various other checks.

> [!TIP]
> Any of the above functions can be used in a string context by putting an `S` in front of the function name. Example: `Sprint(...) `.

|Function Name|Description|
|-------------|-----------|
|[Print](https://pkg.go.dev/github.com/pterm/pterm#Print)|Print formats using the default formats for its operands and writes to standard output. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.|
|[Printf](https://pkg.go.dev/github.com/pterm/pterm#Printf)|Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.|
|[Println](https://pkg.go.dev/github.com/pterm/pterm#Println)|Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.|
|[Printo](https://pkg.go.dev/github.com/pterm/pterm#Printo)|Printo overrides the current line in a terminal. If the current line is empty, the text will be printed like with `pterm.Print`. Example: pterm.Printo("Hello, World") time.Sleep(time.Second) pterm.Printo("Hello, Earth!")|