# Overriding Default Printers

You can override the default printers to customize them:

```go
pterm.DefaultSection = *pterm.DefaultSection.WithLevel(2)
pterm.DefaultSection.Println("Hello, World!") // -> ## Hello, World!

pterm.DefaultSection = *pterm.DefaultSection.WithLevel(3)
pterm.DefaultSection.Println("Hello, World!") // -> ### Hello, World!
```
