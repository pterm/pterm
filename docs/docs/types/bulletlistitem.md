# BulletListItem

<!-- 
Replace all of the following strings with the current printer.
        bulletlistitem BulletListItem BulletListItemPrinter DefaultBulletListItem
-->

## Usage

### Basic usage

```go
pterm.BulletListItem{}.WithText("Hello, World").WithLevel(0)
```

### Functions

|Function|Description|
|--------|-----------|
|[NewBulletListItemFromString(text string, padding string)](https://pkg.go.dev/github.com/pterm/pterm#TemplatePrinter.NewBulletListItemFromString)|NewBulletListItemFromString returns a BulletListItem with a Text. The padding is counted in the Text to define the Level of the ListItem.|

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultBulletListItem.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultBulletListItem.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

|Option|Type|
|------|----|
|[Text](https://pkg.go.dev/github.com/pterm/pterm#BulletListItemPrinter.WithText)|string|
|[TextStyle](https://pkg.go.dev/github.com/pterm/pterm#BulletListItemPrinter.WithTextStyle)|[*Style](https://pkg.go.dev/github.com/pterm/pterm#Style)|
|[Bullet](https://pkg.go.dev/github.com/pterm/pterm#BulletListItemPrinter.WithBullet)|string|
|[BulletStyle](https://pkg.go.dev/github.com/pterm/pterm#BulletListItemPrinter.WithBulletStyle)|[*Style](https://pkg.go.dev/github.com/pterm/pterm#Style)|
|[Level](https://pkg.go.dev/github.com/pterm/pterm#BulletListItemPrinter.WithLevel)|int|