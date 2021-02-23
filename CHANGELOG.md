<a name="unreleased"></a>
## [Unreleased]

### Bug Fixes
- correct `pterm.Println()` behaviour to fit to `fmt.Println()`
- change terminal package import path to updated version

### Chore
- **deps:** update dependencies

### Documentation Changes


<a name="v0.12.8"></a>
## [v0.12.8] - 2020-12-11
### Bug Fixes
- revert change horizontal string change

### Chore
- **intellij:** update conventionalCommit.xml

### Code Refactoring
- remove analytics
- **boxprinter:** change from `RenderablePrinter` to `TextPrinter`
- **boxprinter:** return theme when style is nil
- **boxprinter:** change `DefaultBox` top and bottom padding to 0
- **boxprinter:** fix spacing between boxes and in boxes
- **boxprinter:** refactor code
- **panelprinter:** optional border for `Panel`
- **panelprinter:** add `BoxPrinter` to surround panels with a fully custom box

### Documentation Changes
- **boxprinter:** update `BoxPrinter` example
- **boxprinter:** update documentation to printer
- **boxprinter:** add documentation
- **examples:** add example for `BoxPrinter`
- **examples:** change example for `BoxPrinter`
- **pterm-sh:** add `BoxPrinter`
- **pterm-sh:** update sidebar
- **readme:** update readme with new main features

### Features
- **boxprinter:** add `WithHorizontalString` to `BoxPrinter`
- **boxprinter:** add `BoxPrinter`
- **panel:** add optional border for `Panel`
- **panelprinter:** add theme support to `PanelPrinter`
- **theme:** add `BoxStyle` and `BoxTextStyle`
- **theme:** add optional theme for border in `Panel`

### Style
- fix go vet issue
- **panelprinter:** bottom padding in `PanelPrinter`

### Test
- **boxprinter:** add test
- **boxprinter:** test multiple lines in one box
- **boxprinter:** add tests for `BoxPrinter`
- **panelprinter:** add tests for adding box printer
- **panelprinter:** add tests for optional border for `Panel`
- **theme:** add tests for `BoxStyle` and `BoxTextStyle`


<a name="v0.12.7"></a>
## [v0.12.7] - 2020-11-24
### Bug Fixes
- center bars over colored labels in `BarChart`

### Documentation Changes
- change Template to BarChart
- fix readme formatting
- add buymeacoffee link to readme
- **contributing:** update contributing file
- **funding:** add buymeacoffee link to FUNDING.yml
- **pterm-sh:** add docs for `BarChartPrinter`
- **pterm-sh:** remove empty codeblock
- **readme:** add emoji to support section
- **readme:** add spacing between text

### Features
- add values to chart
- add horizontal `BarChartPrinter`
- add `BarChartPrinter`
- add `BarChartPrinter`
- add `BarChartPrinter`
- **theme:** add theme support to `BarChart`

### Test
- add tests to `BarChartPrinter`


<a name="v0.12.6"></a>
## [v0.12.6] - 2020-11-17
### Bug Fixes
- disabling output works as expected now ([#149](https://github.com/pterm/pterm/issues/149))

### Documentation Changes


<a name="v0.12.5"></a>
## [v0.12.5] - 2020-11-17
### Bug Fixes
- fix `PrefixPrinter` with multiple trailing newline endings.

### Documentation Changes


<a name="v0.12.4"></a>
## [v0.12.4] - 2020-11-17
### Bug Fixes
- fix `Printf` of `PrefixPrinter`

### Ci
- count tests from all packages
- read stderr to count unit tests
- try to repeat unittest count

### Documentation Changes
- **contributing:** change examples to return error ([#145](https://github.com/pterm/pterm/issues/145))

### Revert
- "ci: try to repeat unittest count"

### Style
- remove commented code


<a name="v0.12.3"></a>
## [v0.12.3] - 2020-11-12
### Chore
- **deps:** update dependencies

### Ci
- disable ubuntu and macOS tests
- disable parallel unit test counting
- fix generation of readmes
- disable bodyclose
- add stats
- disable parallel tests
- add timeout to unittests
- add json output on unittest count fail

### Documentation Changes

### Test
- reduce tests
- different test logic for rgb printing
- add better test names for `RGB` tests


<a name="v0.12.2"></a>
## [v0.12.2] - 2020-11-05
### Bug Fixes
- fix internal `GetStringMaxWidth` max width

### Chore
- **intellij:** add conventional commit linting defaults

### Ci
- only count tests in pterm package
- change svg font to JetBrains Mono

### Code Refactoring
- use `pterm.Sprint` to print

### Documentation Changes
- **code:** fix `Printo` comment
- **panel:** update example
- **pterm-sh:** remove variables for now
- **pterm-sh:** add print methods documentation
- **pterm-sh:** move bulletlistitem.md into types directory
- **pterm-sh:** move printer files in own directory
- **pterm-sh:** add better structure to the documentation
- **pterm-sh:** add functions with descriptions
- **pterm-sh:** update sidebar
- **pterm-sh:** add panel-printer
- **pterm-sh:** update sidebar
- **pterm-sh:** sort sidebar
- **pterm-sh:** rename basictext.md to basictext-printer.md
- **pterm-sh:** update sidebar
- **pterm-sh:** add basictext
- **pterm-sh:** make writing-documentation-template.md to a codeblock again
- **pterm-sh:** add custom functions to writing-documentation-template.md
- **pterm-sh:** change basic usage of progressbar
- **pterm-sh:** rename printer documentation files
- **pterm-sh:** update sidebar
- **pterm-sh:** add tree
- **pterm-sh:** add table
- **pterm-sh:** update sidebar
- **pterm-sh:** change `bool` to `...bool`
- **pterm-sh:** add spinner
- **pterm-sh:** update sidebar
- **pterm-sh:** update sidebar
- **pterm-sh:** add section
- **pterm-sh:** add progressbar
- **pterm-sh:** update sidebar
- **pterm-sh:** update sidebar
- **pterm-sh:** add paragraph
- **pterm-sh:** change to same naming convention
- **pterm-sh:** add header
- **pterm-sh:** update sidebar
- **pterm-sh:** update sidebar
- **pterm-sh:** add center
- **pterm-sh:** add bulletlistitem
- **pterm-sh:** add bulletlist
- **pterm-sh:** update sidebar
- **pterm-sh:** update sidebar
- **pterm-sh:** add override-default-printer.md
- **pterm-sh:** add `BigText` documentation
- **pterm-sh:** add documentation on how to write documentation
- **pterm-sh:** change style.css
- **pterm-sh:** disbale service worker
- **pterm-sh:** add favicon
- **pterm-sh:** fix some styling issues
- **pterm-sh:** disable share buttons
- **pterm-sh:** add quick-start.md
- **pterm-sh:** fix link in intro
- **pterm-sh:** add readme
- **pterm-sh:** add review check description
- **pterm-sh:** update documentation intro
- **pterm-sh:** write writing-documentation.md
- **pterm-sh:** add flat style to alerts
- **pterm-sh:** fix coverpage link to documentation
- **pterm-sh:** add writing-documentation.md
- **pterm-sh:** update sidebar
- **pterm-sh:** add contributing.md, intro.md and review-checks.md
- **pterm-sh:** update coverpage
- **pterm-sh:** change `|` to `│`
- **pterm.sh:** update pterm logo
- **readme:** add link to pterm documentation
- **readme:** add links to pterm website

### Features
- color each line separately when using multi line input

### Test
- **basictext:** proxy print functions to DevNull
- **progressbar:** proxy print functions to DevNull


<a name="v0.12.1"></a>
## [v0.12.1] - 2020-11-04
### Bug Fixes
- **panel:** Fix output when input is colored

### Ci
- fix unit test count glitch

### Documentation Changes

### Performance Improvements
- **header:** calculate margin faster


<a name="v0.12.0"></a>
## [v0.12.0] - 2020-11-04
### Bug Fixes
- **panel:** add invalid check for `padding` in `panel`

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

### Documentation Changes
- **tree:** change `Tree` to `TreePrinter`

### Features
- **panel:** add an option to make a padding beneath `panel`
- **panel:** add an option to make columns the same length

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
### Bug Fixes
- **centerprinter:** make centerprinter func return pointer

### Chore
- **deps:** bump github.com/gookit/color from 1.3.1 to 1.3.2

### Documentation Changes
- add docs website
- **review:** add review docs check
- **review:** update review checks
- **review:** add review check file

### Features
- add `PanelPrinter`

### BREAKING CHANGE

make centerprinter func `WithCenterEachLineSeparately` return a pointer of centerprinter


<a name="v0.10.1"></a>
## [v0.10.1] - 2020-11-02
### Documentation Changes
- **examples:** fix `disable-color` example
- **examples:** add center examples
- **examples:** fix examples
- **examples:** fix demo
- **readme:** add center to features

### Features
- add `CenterPrinter`


<a name="v0.10.0"></a>
## [v0.10.0] - 2020-11-01
### Chore
- **deps:** update dependencies

### Ci
- change font of animation SVGs to `consolas`
- change font of animation SVGs

### Code Refactoring
- ignore errors where no errors can occur
- **theme:** change `ListTextStyle` to `BulletListTextStyle` ([#104](https://github.com/pterm/pterm/issues/104))
- **theme:** change `ProgressbarBarStyle` to `FgCyan` ([#106](https://github.com/pterm/pterm/issues/106))
- **theme:** change white to default color in `Theme` ([#103](https://github.com/pterm/pterm/issues/103))

### Documentation Changes
- **examples:** update `disable-color` example ([#107](https://github.com/pterm/pterm/issues/107))
- **examples:** change color of `BigLetter` P of PTerm to `FgLightCyan` ([#105](https://github.com/pterm/pterm/issues/105))
- **examples:** change length of tree example in demo

### Features
- make printers return errors
- add `DisableOutput()` and `EnableOutput()` ([#108](https://github.com/pterm/pterm/issues/108))

### BREAKING CHANGE

Interface of `RenderablePrinter` and `LivePrinter` changed.

The global variable `DisableOutput` was renamed to `Output`.


<a name="v0.9.3"></a>
## [v0.9.3] - 2020-10-31
### Chore
- approach for an interactive solution
- **deps:** update dependencies

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

### Documentation Changes
- **examples:** add `Tree` to demo
- **examples:** refactor `Tree` example
- **examples:** refactor `treelist` example
- **examples:** add example for `Tree`
- **examples:** add example for `Tree`
- **examples:** add `Tree` to demo
- **examples:** refactor `treelist` example
- **examples:** refactor `treelist` example
- **readme:** add `Tree` to features

### Features
- add a levelList converter for TreeListPrinter
- add `TreeListPrinter` as a renderable printer
- add `TreeListPrinter` as a renderable printer
- **theme:** add theme support for `Tree`

### Test
- **tree:** add `Tree` tests


<a name="v0.9.2"></a>
## [v0.9.2] - 2020-10-29
### Documentation Changes
- **examples:** add `disable-color` example

### Features
- add option to disable and enable colors


<a name="v0.9.1"></a>
## [v0.9.1] - 2020-10-27
### Code Refactoring
- make the prefix of `Info` smaller again

### Documentation Changes


<a name="v0.9.0"></a>
## [v0.9.0] - 2020-10-27
### Bug Fixes
- progressbar disappears when done and something is printed after

### Chore
- **deps:** update dependencies

### Ci
- replace log panic with fatal
- proxy cmd stdout to os stdout

### Code Refactoring
- remove `UpdateDelay` from `Progressbar`
- change `NewList` to `NewBulletList`
- change `NewList` to `NewBulletList`
- deprecate `UpdateDelay` in `Progressbar`

### Documentation Changes
- **examples:** update `print-with-color` example
- **examples:** remove 2 lines from max terminal max height in fade
- **examples:** update `BigTextPrinter` example
- **examples:** update `Theme` example
- **examples:** update `BulletListPrinter` example
- **examples:** change color fade multiple example ([#85](https://github.com/pterm/pterm/issues/85))
- **examples:** update `BulletListPrinter` custom example
- **examples:** add more delay between table and end
- **examples:** redo demo ([#90](https://github.com/pterm/pterm/issues/90))
- **examples:** update `Table` example
- **examples:** update `Spinner` example
- **examples:** update `DefaultSection` example
- **examples:** update `DefaultProgressbar` example
- **examples:** update `HeaderPrinter` example
- **examples:** update `print-color-rgb` example
- **examples:** update `print-color-fade-multiple` example
- **examples:** update `print-color-fade` example
- **examples:** update `BasicTextPrinter` example
- **examples:** update `ParagraphPrinter` custom example
- **examples:** update `HeaderPrinter` custom example
- **examples:** update `override-default-printers` example
- **examples:** update `ParagraphPrinter` example
- **godoc:** add package description
- **readme:** add features list to README ([#84](https://github.com/pterm/pterm/issues/84))
- **readme:** add link to examples location
- **readme:** shrink features list
- **readme:** remove one space between demo and header
- **readme:** use real emoji for features list
- **readme:** add prefix to features
- **readme:** add source link to demo

### Features
- add `Debug` `PrefixPrinter`
- add support for enabling and disabling debug messages

### Test
- add debugger tests to `PrefixPrinter`
- add progressbar tests

### BREAKING CHANGE

Removed `UpdateDelay` from `Progressbar`. It's no longer used. The Progressbar automatically updates on every change to the current value.

Changed `NewList` to `NewBulletList`.


<a name="v0.8.1"></a>
## [v0.8.1] - 2020-10-26
### Code Refactoring
- refactor doc
- refactor code

### Documentation Changes
- change bug emoji
- add issue templates
- update bug issue template
- **examples:** add multiple color fade example
- **examples:** add TrueColor info to fade examples
- **examples:** add space to fade examples
- **examples:** shrink fade TrueColor info
- **readme:** add short description

### Features
- add fade from one RGB over several RGBs to another RGB


<a name="v0.8.0"></a>
## [v0.8.0] - 2020-10-24
### Ci
- add restarting text
- more logs in CI
- remove goroutines to test efficiency
- escape coverage filepath
- fix codecov
- upload coverage report

### Code Refactoring
- make `BigTextPrinter` release ready
- change `LineCharacter` to `BarCharacter` ([#70](https://github.com/pterm/pterm/issues/70))

### Documentation Changes
- **contributing:** update contributing guides
- **examples:** color fade demo, fade from cyan to magenta ([#68](https://github.com/pterm/pterm/issues/68))
- **examples:** refactor demo
- **examples:** add header to demo
- **readme:** add test coverage to readme
- **readme:** change unit tests badge color to magenta
- **readme:** add coverage badge
- **readme:** update readme
- **readme:** add Q&A to readme

### Features
- add `BigTextPrinter` ([#75](https://github.com/pterm/pterm/issues/75))
- use level of section printer
- add `BulletListPrinter` ([#67](https://github.com/pterm/pterm/issues/67))

### Test
- test that `%s` won't fail to print

### BREAKING CHANGE

Changed `LineCharacter` to `BarCharacter`.


<a name="v0.7.0"></a>
## [v0.7.0] - 2020-10-20
### Bug Fixes
- make theme accept pointer styles
- make Spinner accept pointer Style
- make WithMessageStyle accept Style pointer
- add nil check to SectionPrinter Style
- section printer Style to pointer

### Ci
- better error handling in CI System ([#61](https://github.com/pterm/pterm/issues/61))
- disable paramTypeCombine
- dont check for paramTypeCombine
- put unit test count into readme

### Code Refactoring
- use log output
- remove obsolete if
- fit progressbar to new percentage calculation method
- make fatal panic
- rename parameters
- don't show empty line when removing a progressbar

### Documentation Changes
- comment WithBoolean
- less to-do checks
- add unit test count badge
- document test utils
- **contributing:** update test template
- **contributing:** add test template to contributing guides
- **contributing:** add printer nil check example

### Features
- **progressbar:** add RemoveWhenDone

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

### Tests
- add internal percentage tests
- add BasicTextPrinter tests
- test PrefixPrinter special cases


<a name="v0.6.1"></a>
## [v0.6.1] - 2020-10-20
### Bug Fixes
- fix RGB methods

### Documentation Changes
- **examples:** fix override-default-printer example


<a name="v0.6.0"></a>
## [v0.6.0] - 2020-10-19
### Ci
- disable gocritic in test files
- disable some checks for test files
- don't lint test files
- don't check for pointer returns
- don't check for pointer returns

### Code Refactoring
- make printers accept pointers to styles
- remove emojis to comply with cross-platform policy
- change LivePrinter interface to pointer output
- change TextPrinter interface to pointer output

### Documentation Changes
- add to do list to PR template ([#52](https://github.com/pterm/pterm/issues/52))
- add documentation to theme

### Features
- add BasicTextPrinter
- add theme support to section and table printer
- add theme support to spinner
- add theme support to headers
- add template support for progressbars
- add default theme

### Style
- code styling

### Test
- **benchmark:** fix spinner benchmark

### BREAKING CHANGE

All printers only accept pointers as any `Style` attribute.

LivePrinter now requires to return a pointer.

TextPrinter now requires to return a pointer.


<a name="v0.5.1"></a>
## [v0.5.1] - 2020-10-14
### Chore
- **deps:** update deps

### Ci
- add codeowners file
- update animations
- write breaking changes into unreleased changelog version
- disable cursor in animations

### Code Refactoring
- declare function name as `WithCSVReader`

### Documentation Changes
- **contributing:** change header name
- **examples:** add every ANSI color to examples
- **readme:** put header under demo
- **readme:** update readme
- **readme:** add goal to readme
- **readme:** center header
- **readme:** fix header links
- **readme:** add space between demo and header
- **readme:** add space between demo and header
- **readme:** add easy to use section
- **readme:** style readme
- **readme:** replace emoji strings with actual emojis

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


<a name="v0.5.0"></a>
## [v0.5.0] - 2020-10-08
### Code Refactoring
- rename spinner_printer.go to spinner.go
- rename `GenericPrinter` to `TextPrinter`

### Documentation Changes
- **demo:** add `SectionPrinter` to demo ([#25](https://github.com/pterm/pterm/issues/25))

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

### BREAKING CHANGE

The `GenericPrinter` is now called `TextPrinter`.


<a name="v0.4.1"></a>
## [v0.4.1] - 2020-10-07
### Documentation Changes
- **examples:** fix examples


<a name="v0.4.0"></a>
## [v0.4.0] - 2020-10-07
### Chore
- **deps:** bump github.com/stretchr/testify from 1.3.0 to 1.6.1 ([#22](https://github.com/pterm/pterm/issues/22))
- **intellij:** add config for commit body styling

### Code Refactoring
- use `Style` instead of colors
- refactor function parameters to fit expectation
- rename `RemoveColors` to `RemoveColorFromString`

### Documentation Changes
- **examples:** restructure `examples` directory
- **examples:** rename progressbar to progressbar-default
- **examples:** add default and custom `HeaderPrinter` examples

### Features
- add `Add` to `Style`
- add options shorthands to `SectionPrinter`

### Style
- restyle color fade table
- format project

### Test
- ignore writer close errors in stdout capture

### BREAKING CHANGE

use `Style` instead of colors

Refactor function parameters to fit expectation.
Affects: `WithStyle(colors -> style)`,  `WithScope(string, colors -> scope)`

rename `RemoveColors` to `RemoveColorFromString`


<a name="v0.3.2"></a>
## [v0.3.2] - 2020-10-06
### Bug Fixes
- fix `Sprintf` function of `HeaderPrinter`

### Documentation Changes

### Features
- add `SectionPrinter`

### Test
- add tests for `HeaderPrinter` and `SectionPrinter`


<a name="v0.3.1"></a>
## [v0.3.1] - 2020-10-06
### Ci
- only test pterm package
- test all packages

### Code Refactoring
- set default `BarFiller` to space
- move tests directly into `pterm` module

### Documentation Changes

### Features
- add `BarFiller` to `Progressbar`

### Test
- fix import cycle
- change to inbuilt `SetDefaultOutput` option
- add more benchmarks
- add benchmarks
- add tests to `GenericPrinter` and default print methods


<a name="v0.3.0"></a>
## [v0.3.0] - 2020-10-05
### Bug Fixes
- fix `WithXYZ(b ...bool)` to detect booleans correctly

### Code Refactoring
- remove `Version` constant
- change `WithXXX(b bool)` to `WithXXX(b ...bool)`
- change `SetXXX` to `WithXXX`
- change `Header` to `DefaultHeader`

### Documentation Changes
- **demo:** add `Progressbar` to demo

### BREAKING CHANGE

remove `Version` constant

rename `SetXXX` to `WithXXX`

rename `Header` to `DefaultHeader`


<a name="v0.2.4"></a>
## [v0.2.4] - 2020-10-04
### Bug Fixes
- `Printf` works again

### Documentation Changes


<a name="v0.2.3"></a>
## [v0.2.3] - 2020-10-04
### Code Refactoring
- remove goroutine from `Progressbar`

### Documentation Changes

### Features
- automatically print above `Progressbar`


<a name="v0.2.2"></a>
## [v0.2.2] - 2020-10-04
### Documentation Changes

### Features
- add `Fatal` printer


<a name="v0.2.1"></a>
## [v0.2.1] - 2020-10-04
### Ci
- remove commit to tag
- remove automatic releases
- disable maligned

### Code Refactoring
- bump version to "v0.2.1"

### Documentation Changes
- add sponsor button
- add pull_request_template
- **demo:** fix updated timestamp
- **readme:** `installation`, `documentation` and `contributing` sections
- **reamde:** add emojis to sections
- **reamde:** update examples section link

### Features
- make progressbar configurable
- add percentage helper
- add `RemoveColors`
- add `Progressbar` ([#5](https://github.com/pterm/pterm/issues/5))
- add `Progressbar`
- add fatal to `PrefixPrinter` ([#4](https://github.com/pterm/pterm/issues/4))
- **progressbar:** fade percentage color according to value

### Style
- format code


<a name="v0.2.0"></a>
## [v0.2.0] - 2020-09-30
### Bug Fixes
- spinners spin evenly when multiple spinners are started

### Chore
- **deps:** update deps

### Ci
- update changelog generation
- run golangci-lint on every push and pull-request
- add golangci linting
- make example source code expandable in main readme
- add multi threaded animation generation
- add dependabot

### Code Refactoring
- bump version to "v0.2.0"
- change `WithXXX` to `SetXXX`
- removed `Println` aliases

### Documentation Changes
- add CONTRIBUTING.md
- add code of conduct
- add unstable notice
- **demo:** add space between intro and spinner
- **examples:** update demo example
- **examples:** update demo example
- **examples:** break `demo` example into two parts

### Features
- change style of `Description` printer
- add color in color support
- add `RemoveWhenDone` to `Spinner`
- add multiline support to `PrefixPrinter`
- add `UpdateText` to spinner

### Performance Improvements
- improve spinner performance

### Style
- clean up code

### BREAKING CHANGE

every `WithXXX` is renamed to `SetXXX`

remove `GetFormattedMessage` from `PrefixPrinter`

removed `Println` aliases


<a name="v0.1.0"></a>
## [v0.1.0] - 2020-09-28
### Ci
- remove go dep
- add go testing
- update changelog config

### Code Refactoring
- bump version to "v0.1.0"
- consistent example code for `Printo`
- better comments for `Printo`
- simplify `HeaderPrinter`

### Documentation Changes
- add symbols list (codepage 437)
- fix doc link

### Features
- add spinners
- shorten printer names and add builder methods to printers
- add `Printo` to override printed text
- add `FullWidth` to `HeaderPrinter`
- add terminal size detection

### BREAKING CHANGE

printer names changed

removed `Header` and put it's content directly into `HeaderPrinter`


<a name="v0.0.1"></a>
## [v0.0.1] - 2020-09-21
### Bug Fixes
- header should now work in CI

### Chore
- **gitignore:** ignore experimenting directory
- **idea:** unexclude experimenting directory
- **idea:** exclude experimenting directory

### Ci
- remove animationdata again
- fix sleep timing of animations
- update commit-to-tag.yml
- add release system
- automatically add example code to readme files

### Code Refactoring
- bump version to "v0.0.1"
- refactor project
- add comments to functions

### Documentation Changes
- put documentation flag in seperate line
- add documentation badge
- update demo
- update header example
- fix release badge style
- add readme
- **readme:** put demo animation under header
- **readme:** fix readme animation href

### Features
- add aliases to default printers
- add header example
- integrate ci
- add `HeaderPrinter`
- add exported version variable
- add example `override-default-printer`
- change prefix text color to `LightWhite`


<a name="v0.0.0"></a>
## v0.0.0 - 2020-09-18
### Features
- add changelog template
- configs
- initial commit


[Unreleased]: https://github.com/pterm/pterm/compare/v0.12.8...HEAD
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
