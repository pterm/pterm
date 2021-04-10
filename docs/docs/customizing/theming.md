# Theming

> PTerm lets you create and modify themes, so that you can choose the color theme of your CLI tool.
> All printers will use the `pterm.ThemeDefault`. You can overwrite it to change the colors of each printer.

## Modifying the default theme

> [!NOTE]
> By modifying the default theme, you globally change the colors for every printer in pterm.

```go
pterm.DefaultSection.Println("Hello, World!") // This will be yellow, as the default theme is used
pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgRed) // Here we change the default style for sections
pterm.DefaultSection.Println("Hello, World!") // This will now print in red
```

## Switching themes

> [!NOTE]
> You can also create multiple themes and switch between them while the program is running.

```go
darkTheme := pterm.Theme{...} // Create dark theme
lightTheme := pterm.Theme{...} // Create light theme

pterm.ThemeDefault = darkTheme // Switch to dark theme
pterm.DefaultSection.Println("Hello, World!") // This will be yellow, as the default theme is used

pterm.ThemeDefault = lightTheme // Switch to light theme
pterm.DefaultSection.Println("Hello, World!") // This will now print in red
```

