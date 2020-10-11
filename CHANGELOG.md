<a name="unreleased"></a>
## [Unreleased]

### Chore
- **deps:** update deps

### Documentation Changes

### Features
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


[Unreleased]: https://github.com/pterm/pterm/compare/v0.5.0...HEAD
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
