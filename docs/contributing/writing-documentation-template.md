# Documentation Template

```markdown
    # TemplatePrinter
    
    <!-- 
    Replace all of the following strings with the current printer.
         template Template TemplatePrinter DefaultTemplate
    -->
    
    ![TemplatePrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/template/animation.svg)
    
    <p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/template/main.go" target="_blank">(Show source of demo)</a></p>
    
    
    ## Usage
    
    ### Basic usage
    
    ```go
    pterm.DefaultTemplate.AddMethodsHere()
    ```
    <!-- Delete this section if the printer does not expose functions other than the default output functions -->
    ### Functions
    
    |Function|Description|
    |--------|-----------|
    |[FunctionName](https://pkg.go.dev/github.com/pterm/pterm#TemplatePrinter.FunctionName)|Description of function|
    
    ### Options
    
    > To make a copy with modified options you can use:
    > `pterm.DefaultTemplate.WithOptionName(option)`
    >
    > To change multiple options at once, you can chain the functions:
    > `pterm.DefaultTemplate.WithOptionName(option).WithOptionName2(option2)...`
    
    > [!TIP]
    > Click the options and types to show the documentation on _pkg.go.dev_
    
    |Option|Type|
    |------|----|
    |[OptionName](https://pkg.go.dev/github.com/pterm/pterm#TemplatePrinter.OptionName)|[TypeName](https://pkg.go.dev/github.com/pterm/pterm#TypeName)|
    
    ### Output functions
    <!-- Remove comment of the correct interface -->
    
    <!--
    > This printer implements the interface [`TextPrinter`](https://github.com/pterm/pterm/blob/master/interface_text_printer.go)
    
    |Function|Description|
    |------|---------|
    |Sprint(a ...interface{})|Returns a string|
    |Sprintln(a ...interface{})|Returns a string with a new line at the end|
    |Sprintf(format string, a ...interface{})|Returns a string, formatted according to a format specifier|
    |Sprintfln(format string, a ...interface{})|Returns a string, formatted according to a format specifier with a new line at the end|
    |Print(a ...interface{})|Prints to the terminal|
    |Println(a ...interface{})|Prints to the terminal with a new line at the end|
    |Printf(format string, a ...interface{})|Prints to the terminal, formatted according to a format specifier|
    |Printfln(format string, a ...interface{})|Prints to the terminal, formatted according to a format specifier with a new line at the end|
    -->
    
    <!--
    > This printer implements the interface [`RenderablePrinter`](https://github.com/pterm/pterm/blob/master/interface_renderable_printer.go)
    
    |Function|Description|
    |------|---------|
    |Render()|Prints to Terminal|
    |Srender()|Returns a string|
    -->
    
    <!--
    > This printer implements the interface [`LivePrinter`](https://github.com/pterm/pterm/blob/master/interface_live_printer.go)
    
    |Function|Description|
    |------|---------|
    |Start()|Returns itself and possible errors|
    |Stop()|Returns itself and possible errors|
    |GenericStart()|Returns the started LivePrinter and possible errors|
    |GenericStop()|Returns the stopped LivePrinter and possible errors|
    
    > [!NOTE]
    > The generic start and stop methods are only used to implement the printer into the interface.
    > Use the normal `Start()` and `Stop()` methods if possible.
    -->
    
    ## Related
    - [Override default printers](../customizing/override-default-printer.md)
```
