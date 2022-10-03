<a name="unreleased"></a>
## [Unreleased]


<a name="v0.12.48"></a>
## [v0.12.48] - 2022-10-02
### Features
- custom select/confirm key for interactive printer
- add flag to disable filter/search for interactive printer


<a name="v0.12.47"></a>
## [v0.12.47] - 2022-09-19
### Features
- adding interactive continue printer

### Bug Fixes
- typo
- append the selected value to the prompt

### Code Refactoring
- ignore invalid custom handles
- initiazile handles on getSuffix
- comment format
- address renaming PR comments
- show full handles by default
- use a map for the options

### Reverts
- refactor: use a map for the options


<a name="v0.12.46"></a>
## [v0.12.46] - 2022-09-05
### Features
- **putils:** add `CenterText` in putils

### Bug Fixes
- **textinput:** fixed overwriting the default values


<a name="v0.12.45"></a>
## [v0.12.45] - 2022-07-26
### Bug Fixes
- make sure the interactive printers can cleanup after Ctrl+C
- the interactive confirm answers should match the confirm/reject text

### Test
- add tests for custom answers


<a name="v0.12.44"></a>
## [v0.12.44] - 2022-07-22

<a name="v0.12.43"></a>
## [v0.12.43] - 2022-07-17
### Bug Fixes
- **spinner:** fix line didn't clear properly
- **table:** fixed column length calculation for Chinese strings


<a name="v0.12.42"></a>
## [v0.12.42] - 2022-06-21
### Features
- **input:** added text input printer


<a name="v0.12.41"></a>
## [v0.12.41] - 2022-04-12

<a name="v0.12.40"></a>
## [v0.12.40] - 2022-03-28
### Features
- added a custom writer for all printers


<a name="v0.12.39"></a>
## [v0.12.39] - 2022-03-18
### Features
- use fallback color in `BigTextPrinter` when `RGB` is not supported

### Test
- fix `BigTextPrinter` test
- removed `testdata`
- removed snapshot testing


<a name="v0.12.38"></a>
## [v0.12.38] - 2022-03-09
### Features
- added `NewLettersFromStringWithRGB`
- added `NewLettersFromStringWithRGB`

### Test
- **bigtext:** fix formatting error in bigtext test


<a name="v0.12.37"></a>
## [v0.12.37] - 2022-02-17
### Features
- **progressbar:** Add option to set the `MaxWidth` of the progressbar

### Test
- **progressbar:** added 100% test coverage for new features

### Code Refactoring
- **putils:** Improved styling


<a name="v0.12.36"></a>
## [v0.12.36] - 2022-02-01
### Features
- proposal horizontal table header row and row separators fixes
- proposal horizontal table header row and row separators changes
- proposal horizontal table header row and row separators changes
- proposal horizontal table header row and row separators changes
- proposal horizontal table header row and row separators changes
- proposal horizontal table header row and row separators changes
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators changes
- proposal horizontal table header row and row separators changes
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators
- proposal horizontal table header row and row separators


<a name="v0.12.35"></a>
## [v0.12.35] - 2022-02-01
### Code Refactoring
- fix linting
- regenerate snapshots


<a name="v0.12.34"></a>
## [v0.12.34] - 2022-01-16
### Bug Fixes
- **progressbar:** refresh progressbars on every PTerm print ([#302](https://github.com/pterm/pterm/issues/302))

### Test
- removed `AreaPrinter` test output
- **table:** changed mock reader from `os.Stdin` to `outBuf`


<a name="v0.12.33"></a>
## [v0.12.33] - 2021-10-24
### Features
- add `PrintOnErrorf` for every `TextPrinter` ([#279](https://github.com/pterm/pterm/issues/279))
- **coverage:** add unit test
- **progressbar:** made updating the progressbar title easier. ([#267](https://github.com/pterm/pterm/issues/267))
- **table:** increase test coverage
- **table:** revamp to follow practice
- **table:** add support for right data alignment

### Bug Fixes
- **idea:** revert unwanted config changes
- **linter:** do linter recommendation to delete fallthrough


<a name="v0.12.32"></a>
## [v0.12.32] - 2021-10-15
### Features
- added `AreaPrinter.Clear()`

### Bug Fixes
- progressbar method name
- **header:** fixed length calculation for Chinese strings

### Code Refactoring
- change bitSize size


<a name="v0.12.31"></a>
## [v0.12.31] - 2021-09-21
### Features
- **prefix:** added `LineNumberOffset` to `PrefixPrinter`


<a name="v0.12.30"></a>
## [v0.12.30] - 2021-08-16
### Bug Fixes
- **style:** resetting to previous color also resets attributes

### Code Refactoring
- adapt new testza function name


<a name="v0.12.29"></a>
## [v0.12.29] - 2021-07-19
### Features
- **putils:** add `PrintAverageExecutionTime`

### Test
- fix rgb error test
- fix internal test import cycle
- move tests into own package

### Code Refactoring
- replace `testify` with `testza`


<a name="v0.12.28"></a>
## [v0.12.28] - 2021-07-17
### Features
- **spinner:** add option to show a timer

### Bug Fixes
- **bar chart:** fix panic when rendering empty horizontal bar chart

### Test
- **spinner:** try to fix RawOutput text
- **spinner:** add raw output test

### Code Refactoring
- **spinner:** better raw output logic
- **spinner:** refactor


<a name="v0.12.27"></a>
## [v0.12.27] - 2021-07-05
### Bug Fixes
- **style:** fix multiline style coloring

### Test
- **style:** fix multiline style coloring
- **style:** fix multiline style coloring


<a name="v0.12.26"></a>
## [v0.12.26] - 2021-07-01
### Bug Fixes
- **spinner:** Override previous text in `UpdateText`


<a name="v0.12.25"></a>
## [v0.12.25] - 2021-07-01
### Features
- **table:** add `Boxed` option

### Test
- add tests for boxed `TablePrinter`


<a name="v0.12.24"></a>
## [v0.12.24] - 2021-06-13
### Features
- **boxprinter:** replace line breaks in title with space
- **boxprinter:** add title center position to `BoxPrinter`
- **boxprinter:** add title & title position to `BoxPrinter`
- **boxprinter:** add title & title position to `BoxPrinter`
- **putils:** add `TableDataFromSeparatedValues`
- **putils:** add `TableDataFromTSV`
- **putils:** add `TableDataFromCSV`
- **putils:** add function to convert TSV to `TableData`
- **putils:** add function to convert CSV to `TableData`

### Test
- add test for putils `TableData` generation
- **boxprinter:** add tests for title center position to `BoxPrinter`
- **boxprinter:** add tests for title & title position

### Code Refactoring
- **boxprinter:** prefix title positions with `Title`
- **putils:** add `rowSeparator` to `TableFromSeparatedValues`


<a name="v0.12.23"></a>
## [v0.12.23] - 2021-06-07
### Features
- Add util functions to create tables from slices of structs ([#217](https://github.com/pterm/pterm/issues/217))

### Bug Fixes
- **headerprinter:** don't panic if content width > terminal width

### Test
- **prefix:** `pterm.Error` default no line number shown

### Code Refactoring
- **prefix:** `pterm.Error` default no line number shown


<a name="v0.12.22"></a>
## [v0.12.22] - 2021-05-30
### Features
- make spinner update faster

### Performance Improvements
- improve performance of `SpinnerPrinter`


<a name="v0.12.21"></a>
## [v0.12.21] - 2021-05-30
### Features
- print lines above active spinners
- **putils:** add `DownloadFileWithProgressbar`

### Test
- clear active spinners after tests complete

### Code Refactoring
- **putils:** change internal variable name


<a name="v0.12.20"></a>
## [v0.12.20] - 2021-05-29
### Features
- force color output by default


<a name="v0.12.19"></a>
## [v0.12.19] - 2021-05-29
### Features
- add `PrintOnError` for all printers and interface
- **putils:** add `putils` package ([#206](https://github.com/pterm/pterm/issues/206))

### Bug Fixes
- **header:** fix multiline header

### Test
- add tests for all printers for `PrintOnError`

### Code Refactoring
- make `PrintOnError` return `*TextPrinter`
- **area:** better height calculation


<a name="v0.12.18"></a>
## [v0.12.18] - 2021-05-22
### Features
- add `AreaPrinter`
- **area:** add `Center` option
- **area:** add `Fullscreen` option
- **area:** add `GetContent` function
- **area:** add `AreaPrinter`

### Test
- **area:** fix tests for `AreaPrinter`
- **area:** add `AreaPrinter` tests

### Code Refactoring
- fix linting errors


<a name="v0.12.17"></a>
## [v0.12.17] - 2021-05-14
### Bug Fixes
- fix `pterm.Fatal.Printfln` not panicking
- **prefix:** fix `pterm.Fatal.Printfln` not panicking and had output in debug mode

### Test
- **prefix:** add tests for `Sprintfln` and `Printfln` function when in debug mode


<a name="v0.12.16"></a>
## [v0.12.16] - 2021-05-13
### Code Refactoring
- **prefix:** make `PrintOnError` accept multiple inputs


<a name="v0.12.15"></a>
## [v0.12.15] - 2021-05-13
### Features
- add raw output mode for `BarChart`
- add disable styling boolean option
- **bigtext:** add raw output mode
- **centerprinter:** add raw output mode
- **headerprinter:** add raw output mode
- **panelprinter:** add raw output mode
- **paragraph:** add raw output mode
- **prefix:** add `PrintIfError`
- **prefix:** add raw output mode
- **progressbar:** add raw output mode
- **spinner:** add raw output mode

### Bug Fixes
- **prefix:** fix `PrintOnError`

### Test
- add tests with `RawOutput` enabled
- add interface tests for `Color` and `RGB`
- added tests for `DisableStyling` and `EnableStyling`

### Code Refactoring
- correct behaviour of Enable-/DisableStyling
- fix variable names


<a name="v0.12.14"></a>
## [v0.12.14] - 2021-05-09
### Features
- **basic-text:** add `Sprintfln` and `Printfln` function
- **boxprinter:** add `Sprintfln` and `Printfln` function
- **centerprinter:** add `Sprintfln` and `Printfln` function
- **color:** add `Sprintfln` and `Printfln` function
- **header:** add `Sprintfln` and `Printfln` function
- **paragraph:** add `Sprintfln` and `Printfln` function
- **prefix:** add `Sprintfln` and `Printfln` function
- **print:** add `Sprintfln` and `Printfln` function
- **printer-interface:** add `Sprintfln` and `Printfln` to the interface
- **rgb:** add `Sprintfln` and `Printfln` function
- **section:** add `Sprintfln` and `Printfln` function

### Bug Fixes
- **header:** fix inline color in `Header`

### Test
- add tests for `Sprintfln` and `Printfln` function

### Code Refactoring
- refactor `Sprintfln` and `Printfln` func. for better performance

### Reverts
- ci: change color scheme for rendered examples


<a name="v0.12.13"></a>
## [v0.12.13] - 2021-04-10
### Bug Fixes
- **bigtext:** fix height of some characters [#180](https://github.com/pterm/pterm/issues/180)
- **color:** make color implement `TextPrinter`

### Test
- add interface tests

### Code Refactoring
- **examples:** center the intro of `demo`
- **examples:** add note to box printer


<a name="v0.12.12"></a>
## [v0.12.12] - 2021-03-01
### Features
- **prefixprinter:** Add option to show line number of caller

### Code Refactoring
- **examples:** Update `PrefixPrinter` example


<a name="v0.12.11"></a>
## [v0.12.11] - 2021-02-26
### Code Refactoring
- refactor print logic of `BoxPrinter`
- refactor print logic of `CenterPrinter`


<a name="v0.12.10"></a>
## [v0.12.10] - 2021-02-26
### Bug Fixes
- correct `pterm.Println()` behaviour to fit to `fmt.Println()`


<a name="v0.12.9"></a>
## [v0.12.9] - 2021-02-23
### Bug Fixes
- correct `pterm.Println()` behaviour to fit to `fmt.Println()`
- change terminal package import path to updated version


<a name="v0.12.8"></a>
## [v0.12.8] - 2020-12-11
### Features
- **boxprinter:** add `WithHorizontalString` to `BoxPrinter`
- **boxprinter:** add `BoxPrinter`
- **panel:** add optional border for `Panel`
- **panelprinter:** add theme support to `PanelPrinter`
- **theme:** add `BoxStyle` and `BoxTextStyle`
- **theme:** add optional theme for border in `Panel`

### Bug Fixes
- revert change horizontal string change

### Test
- **boxprinter:** add test
- **boxprinter:** test multiple lines in one box
- **boxprinter:** add tests for `BoxPrinter`
- **panelprinter:** add tests for adding box printer
- **panelprinter:** add tests for optional border for `Panel`
- **theme:** add tests for `BoxStyle` and `BoxTextStyle`

### Code Refactoring
- remove analytics
- **boxprinter:** change from `RenderablePrinter` to `TextPrinter`
- **boxprinter:** return theme when style is nil
- **boxprinter:** change `DefaultBox` top and bottom padding to 0
- **boxprinter:** fix spacing between boxes and in boxes
- **boxprinter:** refactor code
- **panelprinter:** optional border for `Panel`
- **panelprinter:** add `BoxPrinter` to surround panels with a fully custom box


<a name="v0.12.7"></a>
## [v0.12.7] - 2020-11-24
### Features
- add values to chart
- add horizontal `BarChartPrinter`
- add `BarChartPrinter`
- add `BarChartPrinter`
- add `BarChartPrinter`
- **theme:** add theme support to `BarChart`

### Bug Fixes
- center bars over colored labels in `BarChart`

### Test
- add tests to `BarChartPrinter`


<a name="v0.12.6"></a>
## [v0.12.6] - 2020-11-17
### Bug Fixes
- disabling output works as expected now ([#149](https://github.com/pterm/pterm/issues/149))


<a name="v0.12.5"></a>
## [v0.12.5] - 2020-11-17
### Bug Fixes
- fix `PrefixPrinter` with multiple trailing newline endings.


<a name="v0.12.4"></a>
## [v0.12.4] - 2020-11-17
### Bug Fixes
- fix `Printf` of `PrefixPrinter`


<a name="v0.12.3"></a>
## [v0.12.3] - 2020-11-12
### Test
- reduce tests
- different test logic for rgb printing
- add better test names for `RGB` tests


<a name="v0.12.2"></a>
## [v0.12.2] - 2020-11-05
### Features
- color each line separately when using multi line input

### Bug Fixes
- fix internal `GetStringMaxWidth` max width

### Test
- **basictext:** proxy print functions to DevNull
- **progressbar:** proxy print functions to DevNull

### Code Refactoring
- use `pterm.Sprint` to print


<a name="v0.12.1"></a>
## [v0.12.1] - 2020-11-04
### Bug Fixes
- **panel:** Fix output when input is colored

### Performance Improvements
- **header:** calculate margin faster


<a name="v0.12.0"></a>
## [v0.12.0] - 2020-11-04
### Features
- **panel:** add an option to make a padding beneath `panel`
- **panel:** add an option to make columns the same length

### Bug Fixes
- **panel:** add invalid check for `padding` in `panel`

### Test
- **bulletlist:** `BulletListItem` remove `Render` and `Srender`
- **bulletlist:** change `BulletList` to `BulletListPrinter`
- **panel:** add invalid check for `padding` in `panel`
- **panel:** add test for `WithBottomPadding`
- **panel:** add test for `WithSameColumnWidth` & multiple `panel`
- **panel:** add test for `WithSameColumnWidth`
- **progressbar:** change `Progressbar` to `ProgressbarPrinter`
- **progressbar:** change directory name `progressbar_test` to `progressbar_printer_test`
- **spinner:** change directory name `spinner_test` to `spinner_printer_test`
- **spinner:** change `Spinner` to `SpinnerPrinter`
- **table:** change `Table` to `TablePrinter`
- **tree:** change `Tree` to `TreePrinter`

### Code Refactoring
- make all printer names end with `Printer` ([#134](https://github.com/pterm/pterm/issues/134))
- **bulletlist:** remove `DefaultBulletListItem`
- **bulletlist:** `BulletListItem` remove `Render` and `Srender`
- **bulletlist:** `BulletListItem` is no renderable anymore
- **bulletlist:** change `BulletList` to `BulletListPrinter`
- **progressbar:** change `ActiveProgressbars` to `ActiveProgressbarPrinters`
- **progressbar:** change directory name `progressbar` to `progressbar_printer`
- **progressbar:** change `Progressbar` to `ProgressbarPrinter`
- **spinner:** change `Spinner` to `SpinnerPrinter`
- **spinner:** change directory name `spinner` to `spinner_printer`
- **table:** change `Table` to `TablePrinter`
- **tree:** change `Tree` to `TreePrinter`

### BREAKING CHANGE

Removed `DefaultBulletListItem`.

Change names of printers which didn't end with `Printer`. Every printer name ends with `Printer` now to fit into the new naming convention.

change `ActiveProgressbars` to `ActiveProgressbarPrinters`

`BulletListItem` is no renderable anymore, removed `Render` and `Srender`

`BulletListItem` is no renderable anymore, removed `Render` and `Srender`

`BulletListItem` is no renderable anymore

change `Tree` to `TreePrinter` to unify the naming scheme

change `Tree` to `TreePrinter` to unify the naming scheme

change `Table` to `TablePrinter` to unify the naming scheme

change `Table` to `TablePrinter` to unify the naming scheme

change `Spinner` to `SpinnerPrinter` to unify the naming scheme

change `Spinner` to `SpinnerPrinter` to unify the naming scheme

change `Progressbar` to `ProgressbarPrinter` to unify the naming scheme

change `Progressbar` to `ProgressbarPrinter` to unify the naming scheme

change `BulletList` to `BulletListPrinter` to unify the naming scheme

change `BulletList` to `BulletListPrinter` to unify the naming scheme


<a name="v0.11.0"></a>
## [v0.11.0] - 2020-11-03
### Features
- add `PanelPrinter`

### Bug Fixes
- **centerprinter:** make centerprinter func return pointer

### BREAKING CHANGE

make centerprinter func `WithCenterEachLineSeparately` return a pointer of centerprinter


<a name="v0.10.1"></a>
## [v0.10.1] - 2020-11-02
### Features
- add `CenterPrinter`


<a name="v0.10.0"></a>
## [v0.10.0] - 2020-11-01
### Features
- make printers return errors
- add `DisableOutput()` and `EnableOutput()` ([#108](https://github.com/pterm/pterm/issues/108))

### Code Refactoring
- ignore errors where no errors can occur
- **theme:** change `ListTextStyle` to `BulletListTextStyle` ([#104](https://github.com/pterm/pterm/issues/104))
- **theme:** change `ProgressbarBarStyle` to `FgCyan` ([#106](https://github.com/pterm/pterm/issues/106))
- **theme:** change white to default color in `Theme` ([#103](https://github.com/pterm/pterm/issues/103))

### BREAKING CHANGE

Interface of `RenderablePrinter` and `LivePrinter` changed.

The global variable `DisableOutput` was renamed to `Output`.


<a name="v0.9.3"></a>
## [v0.9.3] - 2020-10-31
### Features
- add a levelList converter for TreeListPrinter
- add `TreeListPrinter` as a renderable printer
- add `TreeListPrinter` as a renderable printer
- **theme:** add theme support for `Tree`

### Test
- **tree:** add `Tree` tests

### Code Refactoring
- clean up `Tree`
- **theme:** change `TreeTextStyle` to `FgDefault`
- **tree:** add Indent to control the spacing between levels and changed docs(examples)
- **tree:** add more spacing between levels
- **tree:** refactor `Tree` code and write tests for `Tree`
- **tree:** refactor `Tree` code
- **tree:** refactor `Tree` code
- **tree:** refactor `Tree` code
- **tree:** refactor `Tree` code and write tests for `Tree`


<a name="v0.9.2"></a>
## [v0.9.2] - 2020-10-29
### Features
- add option to disable and enable colors


<a name="v0.9.1"></a>
## [v0.9.1] - 2020-10-27
### Code Refactoring
- make the prefix of `Info` smaller again


<a name="v0.9.0"></a>
## [v0.9.0] - 2020-10-27
### Features
- add `Debug` `PrefixPrinter`
- add support for enabling and disabling debug messages

### Bug Fixes
- progressbar disappears when done and something is printed after

### Test
- add debugger tests to `PrefixPrinter`
- add progressbar tests

### Code Refactoring
- remove `UpdateDelay` from `Progressbar`
- change `NewList` to `NewBulletList`
- change `NewList` to `NewBulletList`
- deprecate `UpdateDelay` in `Progressbar`

### BREAKING CHANGE

Removed `UpdateDelay` from `Progressbar`. It's no longer used. The Progressbar automatically updates on every change to the current value.

Changed `NewList` to `NewBulletList`.


<a name="v0.8.1"></a>
## [v0.8.1] - 2020-10-26
### Features
- add fade from one RGB over several RGBs to another RGB

### Code Refactoring
- refactor doc
- refactor code


<a name="v0.8.0"></a>
## [v0.8.0] - 2020-10-24
### Features
- add `BigTextPrinter` ([#75](https://github.com/pterm/pterm/issues/75))
- use level of section printer
- add `BulletListPrinter` ([#67](https://github.com/pterm/pterm/issues/67))

### Test
- test that `%s` won't fail to print

### Code Refactoring
- make `BigTextPrinter` release ready
- change `LineCharacter` to `BarCharacter` ([#70](https://github.com/pterm/pterm/issues/70))

### BREAKING CHANGE

Changed `LineCharacter` to `BarCharacter`.


<a name="v0.7.0"></a>
## [v0.7.0] - 2020-10-20
### Features
- **progressbar:** add RemoveWhenDone

### Bug Fixes
- make theme accept pointer styles
- make Spinner accept pointer Style
- make WithMessageStyle accept Style pointer
- add nil check to SectionPrinter Style
- section printer Style to pointer

### Test
- add tests color and style
- add tests to root print functions
- add tests to progressbar
- add tests to terminal
- add tests to theme
- fix internal percentage test
- add tests to Spinner
- add tests for TablePrinter
- special tests for special statements
- complete PrefixPrinter tests
- add PrefixPrinter tests
- rename HeaderPrinter tests
- complete HeaderPrinter tests
- add ParagraphPrinter tests
- add HeaderPrinter tests
- make unit test system check different types
- add SectionPrinter tests
- implement test utils
- add rgb tests

### Code Refactoring
- use log output
- remove obsolete if
- fit progressbar to new percentage calculation method
- make fatal panic
- rename parameters
- don't show empty line when removing a progressbar


<a name="v0.6.1"></a>
## [v0.6.1] - 2020-10-20
### Bug Fixes
- fix RGB methods


<a name="v0.6.0"></a>
## [v0.6.0] - 2020-10-19
### Features
- add BasicTextPrinter
- add theme support to section and table printer
- add theme support to spinner
- add theme support to headers
- add template support for progressbars
- add default theme

### Test
- **benchmark:** fix spinner benchmark

### Code Refactoring
- make printers accept pointers to styles
- remove emojis to comply with cross-platform policy
- change LivePrinter interface to pointer output
- change TextPrinter interface to pointer output

### BREAKING CHANGE

All printers only accept pointers as any `Style` attribute.

LivePrinter now requires to return a pointer.

TextPrinter now requires to return a pointer.


<a name="v0.5.1"></a>
## [v0.5.1] - 2020-10-14
### Features
- add ability to disable output ([#44](https://github.com/pterm/pterm/issues/44))
- add `Srender` to `RenderPrinter` interface
- add csv table support ([#42](https://github.com/pterm/pterm/issues/42))
- add HEX to RGB converter in `RGB` ([#41](https://github.com/pterm/pterm/issues/41))
- add theme to generated animations
- add color fade example ([#38](https://github.com/pterm/pterm/issues/38))
- implement `TextPrinter` into `RGB`
- implement color fade to `Progressbar` ([#37](https://github.com/pterm/pterm/issues/37))
- add color fade function and `RBG` ([#34](https://github.com/pterm/pterm/issues/34))
- change `Section` style

### Code Refactoring
- declare function name as `WithCSVReader`


<a name="v0.5.0"></a>
## [v0.5.0] - 2020-10-08
### Features
- implement `LivePrinter` in `Spinner`
- add `BottomPadding` to `SectionPrinter`
- add `RenderPrinter` interface
- implement `LivePrinter` in `Progressbar`
- add `LivePrinter` interface
- add `TablePrinter` ([#27](https://github.com/pterm/pterm/issues/27))
- add `ParagraphPrinter` ([#24](https://github.com/pterm/pterm/issues/24))

### Test
- add `Print` equals `Sprint` tests for `GenericPrinter`
- add `Spinner` benchmarks

### Code Refactoring
- rename spinner_printer.go to spinner.go
- rename `GenericPrinter` to `TextPrinter`

### BREAKING CHANGE

The `GenericPrinter` is now called `TextPrinter`.


<a name="v0.4.1"></a>
## [v0.4.1] - 2020-10-07

<a name="v0.4.0"></a>
## [v0.4.0] - 2020-10-07
### Features
- add `Add` to `Style`
- add options shorthands to `SectionPrinter`

### Test
- ignore writer close errors in stdout capture

### Code Refactoring
- use `Style` instead of colors
- refactor function parameters to fit expectation
- rename `RemoveColors` to `RemoveColorFromString`

### BREAKING CHANGE

use `Style` instead of colors

Refactor function parameters to fit expectation.
Affects: `WithStyle(colors -> style)`,  `WithScope(string, colors -> scope)`

rename `RemoveColors` to `RemoveColorFromString`


<a name="v0.3.2"></a>
## [v0.3.2] - 2020-10-06
### Features
- add `SectionPrinter`

### Bug Fixes
- fix `Sprintf` function of `HeaderPrinter`

### Test
- add tests for `HeaderPrinter` and `SectionPrinter`


<a name="v0.3.1"></a>
## [v0.3.1] - 2020-10-06
### Features
- add `BarFiller` to `Progressbar`

### Test
- fix import cycle
- change to inbuilt `SetDefaultOutput` option
- add more benchmarks
- add benchmarks
- add tests to `GenericPrinter` and default print methods

### Code Refactoring
- set default `BarFiller` to space
- move tests directly into `pterm` module


<a name="v0.3.0"></a>
## [v0.3.0] - 2020-10-05
### Bug Fixes
- fix `WithXYZ(b ...bool)` to detect booleans correctly

### Code Refactoring
- remove `Version` constant
- change `WithXXX(b bool)` to `WithXXX(b ...bool)`
- change `SetXXX` to `WithXXX`
- change `Header` to `DefaultHeader`

### BREAKING CHANGE

remove `Version` constant

rename `SetXXX` to `WithXXX`

rename `Header` to `DefaultHeader`


<a name="v0.2.4"></a>
## [v0.2.4] - 2020-10-04
### Bug Fixes
- `Printf` works again


<a name="v0.2.3"></a>
## [v0.2.3] - 2020-10-04
### Features
- automatically print above `Progressbar`

### Code Refactoring
- remove goroutine from `Progressbar`


<a name="v0.2.2"></a>
## [v0.2.2] - 2020-10-04
### Features
- add `Fatal` printer


<a name="v0.2.1"></a>
## [v0.2.1] - 2020-10-04
### Features
- make progressbar configurable
- add percentage helper
- add `RemoveColors`
- add `Progressbar` ([#5](https://github.com/pterm/pterm/issues/5))
- add `Progressbar`
- add fatal to `PrefixPrinter` ([#4](https://github.com/pterm/pterm/issues/4))
- **progressbar:** fade percentage color according to value

### Code Refactoring
- bump version to "v0.2.1"


<a name="v0.2.0"></a>
## [v0.2.0] - 2020-09-30
### Features
- change style of `Description` printer
- add color in color support
- add `RemoveWhenDone` to `Spinner`
- add multiline support to `PrefixPrinter`
- add `UpdateText` to spinner

### Bug Fixes
- spinners spin evenly when multiple spinners are started

### Performance Improvements
- improve spinner performance

### Code Refactoring
- bump version to "v0.2.0"
- change `WithXXX` to `SetXXX`
- removed `Println` aliases

### BREAKING CHANGE

every `WithXXX` is renamed to `SetXXX`

remove `GetFormattedMessage` from `PrefixPrinter`

removed `Println` aliases


<a name="v0.1.0"></a>
## [v0.1.0] - 2020-09-28
### Features
- add spinners
- shorten printer names and add builder methods to printers
- add `Printo` to override printed text
- add `FullWidth` to `HeaderPrinter`
- add terminal size detection

### Code Refactoring
- bump version to "v0.1.0"
- consistent example code for `Printo`
- better comments for `Printo`
- simplify `HeaderPrinter`

### BREAKING CHANGE

printer names changed

removed `Header` and put it's content directly into `HeaderPrinter`


<a name="v0.0.1"></a>
## [v0.0.1] - 2020-09-21
### Features
- add aliases to default printers
- add header example
- integrate ci
- add `HeaderPrinter`
- add exported version variable
- add example `override-default-printer`
- change prefix text color to `LightWhite`

### Bug Fixes
- header should now work in CI

### Code Refactoring
- bump version to "v0.0.1"
- refactor project
- add comments to functions


<a name="v0.0.0"></a>
## v0.0.0 - 2020-09-18
### Features
- add changelog template
- configs
- initial commit


[Unreleased]: https://github.com/pterm/pterm/compare/v0.12.48...HEAD
[v0.12.48]: https://github.com/pterm/pterm/compare/v0.12.47...v0.12.48
[v0.12.47]: https://github.com/pterm/pterm/compare/v0.12.46...v0.12.47
[v0.12.46]: https://github.com/pterm/pterm/compare/v0.12.45...v0.12.46
[v0.12.45]: https://github.com/pterm/pterm/compare/v0.12.44...v0.12.45
[v0.12.44]: https://github.com/pterm/pterm/compare/v0.12.43...v0.12.44
[v0.12.43]: https://github.com/pterm/pterm/compare/v0.12.42...v0.12.43
[v0.12.42]: https://github.com/pterm/pterm/compare/v0.12.41...v0.12.42
[v0.12.41]: https://github.com/pterm/pterm/compare/v0.12.40...v0.12.41
[v0.12.40]: https://github.com/pterm/pterm/compare/v0.12.39...v0.12.40
[v0.12.39]: https://github.com/pterm/pterm/compare/v0.12.38...v0.12.39
[v0.12.38]: https://github.com/pterm/pterm/compare/v0.12.37...v0.12.38
[v0.12.37]: https://github.com/pterm/pterm/compare/v0.12.36...v0.12.37
[v0.12.36]: https://github.com/pterm/pterm/compare/v0.12.35...v0.12.36
[v0.12.35]: https://github.com/pterm/pterm/compare/v0.12.34...v0.12.35
[v0.12.34]: https://github.com/pterm/pterm/compare/v0.12.33...v0.12.34
[v0.12.33]: https://github.com/pterm/pterm/compare/v0.12.32...v0.12.33
[v0.12.32]: https://github.com/pterm/pterm/compare/v0.12.31...v0.12.32
[v0.12.31]: https://github.com/pterm/pterm/compare/v0.12.30...v0.12.31
[v0.12.30]: https://github.com/pterm/pterm/compare/v0.12.29...v0.12.30
[v0.12.29]: https://github.com/pterm/pterm/compare/v0.12.28...v0.12.29
[v0.12.28]: https://github.com/pterm/pterm/compare/v0.12.27...v0.12.28
[v0.12.27]: https://github.com/pterm/pterm/compare/v0.12.26...v0.12.27
[v0.12.26]: https://github.com/pterm/pterm/compare/v0.12.25...v0.12.26
[v0.12.25]: https://github.com/pterm/pterm/compare/v0.12.24...v0.12.25
[v0.12.24]: https://github.com/pterm/pterm/compare/v0.12.23...v0.12.24
[v0.12.23]: https://github.com/pterm/pterm/compare/v0.12.22...v0.12.23
[v0.12.22]: https://github.com/pterm/pterm/compare/v0.12.21...v0.12.22
[v0.12.21]: https://github.com/pterm/pterm/compare/v0.12.20...v0.12.21
[v0.12.20]: https://github.com/pterm/pterm/compare/v0.12.19...v0.12.20
[v0.12.19]: https://github.com/pterm/pterm/compare/v0.12.18...v0.12.19
[v0.12.18]: https://github.com/pterm/pterm/compare/v0.12.17...v0.12.18
[v0.12.17]: https://github.com/pterm/pterm/compare/v0.12.16...v0.12.17
[v0.12.16]: https://github.com/pterm/pterm/compare/v0.12.15...v0.12.16
[v0.12.15]: https://github.com/pterm/pterm/compare/v0.12.14...v0.12.15
[v0.12.14]: https://github.com/pterm/pterm/compare/v0.12.13...v0.12.14
[v0.12.13]: https://github.com/pterm/pterm/compare/v0.12.12...v0.12.13
[v0.12.12]: https://github.com/pterm/pterm/compare/v0.12.11...v0.12.12
[v0.12.11]: https://github.com/pterm/pterm/compare/v0.12.10...v0.12.11
[v0.12.10]: https://github.com/pterm/pterm/compare/v0.12.9...v0.12.10
[v0.12.9]: https://github.com/pterm/pterm/compare/v0.12.8...v0.12.9
[v0.12.8]: https://github.com/pterm/pterm/compare/v0.12.7...v0.12.8
[v0.12.7]: https://github.com/pterm/pterm/compare/v0.12.6...v0.12.7
[v0.12.6]: https://github.com/pterm/pterm/compare/v0.12.5...v0.12.6
[v0.12.5]: https://github.com/pterm/pterm/compare/v0.12.4...v0.12.5
[v0.12.4]: https://github.com/pterm/pterm/compare/v0.12.3...v0.12.4
[v0.12.3]: https://github.com/pterm/pterm/compare/v0.12.2...v0.12.3
[v0.12.2]: https://github.com/pterm/pterm/compare/v0.12.1...v0.12.2
[v0.12.1]: https://github.com/pterm/pterm/compare/v0.12.0...v0.12.1
[v0.12.0]: https://github.com/pterm/pterm/compare/v0.11.0...v0.12.0
[v0.11.0]: https://github.com/pterm/pterm/compare/v0.10.1...v0.11.0
[v0.10.1]: https://github.com/pterm/pterm/compare/v0.10.0...v0.10.1
[v0.10.0]: https://github.com/pterm/pterm/compare/v0.9.3...v0.10.0
[v0.9.3]: https://github.com/pterm/pterm/compare/v0.9.2...v0.9.3
[v0.9.2]: https://github.com/pterm/pterm/compare/v0.9.1...v0.9.2
[v0.9.1]: https://github.com/pterm/pterm/compare/v0.9.0...v0.9.1
[v0.9.0]: https://github.com/pterm/pterm/compare/v0.8.1...v0.9.0
[v0.8.1]: https://github.com/pterm/pterm/compare/v0.8.0...v0.8.1
[v0.8.0]: https://github.com/pterm/pterm/compare/v0.7.0...v0.8.0
[v0.7.0]: https://github.com/pterm/pterm/compare/v0.6.1...v0.7.0
[v0.6.1]: https://github.com/pterm/pterm/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/pterm/pterm/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/pterm/pterm/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/pterm/pterm/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/pterm/pterm/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/pterm/pterm/compare/v0.3.2...v0.4.0
[v0.3.2]: https://github.com/pterm/pterm/compare/v0.3.1...v0.3.2
[v0.3.1]: https://github.com/pterm/pterm/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/pterm/pterm/compare/v0.2.4...v0.3.0
[v0.2.4]: https://github.com/pterm/pterm/compare/v0.2.3...v0.2.4
[v0.2.3]: https://github.com/pterm/pterm/compare/v0.2.2...v0.2.3
[v0.2.2]: https://github.com/pterm/pterm/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/pterm/pterm/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/pterm/pterm/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/pterm/pterm/compare/v0.0.1...v0.1.0
[v0.0.1]: https://github.com/pterm/pterm/compare/v0.0.0...v0.0.1
