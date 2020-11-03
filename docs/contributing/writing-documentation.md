# Writing Documentation for PTerm

> Documentation is important to show users how to do the stuff they want to do.

## Documenting code with comments

Every exported function in the source code of PTerm must be documented with comments.\
The comment must start with the name of the function.\

**Example:**

```go
// HelloWorld appends a new line and the string "Hello, World!" to the input and returns it.
func HelloWorld(s string) string {
	return s + "\nHello, World!"
}
```

## Editing Documentation on the PTerm website

> [!NOTE]
> The documentation of the official PTerm website is written in [Markdown](https://en.wikipedia.org/wiki/Markdown).

To edit the official documentation of PTerm, you can submit a Pull Request, which changes the markdown files in [`./docs`](https://github.com/pterm/pterm/tree/master/docs).

## Special Formatting

> [!TIP]
> We have added some functionality to the standard Markdown parser to have more layout options.

### Tabbed content

<!-- tabs:start -->

#### ** English **

Hello!

#### ** French **

Bonjour!

#### ** German **

Hallo!

<!-- tabs:end -->

```markdown
<!-- tabs:start -->

#### ** English **

Hello!

#### ** French **

Bonjour!

#### ** German **

Hallo!

<!-- tabs:end -->
```

### Alerts

> [!NOTE]
> This is a note

```markdown
> [!NOTE]
> This is a note
```

> [!TIP]
> This is a tip

```markdown
> [!TIP]
> This is a tip
```

> [!WARNING]
> This is a warning

```markdown
> [!WARNING]
> This is a warning
```

> [!ATTENTION]
> This needs some attention.

```markdown
> [!ATTENTION]
> This needs some attention.
```

