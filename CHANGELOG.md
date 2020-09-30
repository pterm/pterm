<a name="unreleased"></a>
## [Unreleased]

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
- change `WithXXX` to `SetXXX`
- removed `Println` aliases

### Documentation Changes
- add code of conduct
- add unstable notice
- **examples:** update demo example
- **examples:** update demo example

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


[Unreleased]: https://github.com/pterm/pterm/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/pterm/pterm/compare/v0.0.1...v0.1.0
[v0.0.1]: https://github.com/pterm/pterm/compare/v0.0.0...v0.0.1
