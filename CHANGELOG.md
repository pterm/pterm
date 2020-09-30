<a name="unreleased"></a>
## [Unreleased]

### Ci
- add golangci linting
- make example source code expandable in main readme
- add multi threaded animation generation
- add dependabot

### Docs
- add unstable notice
- add code of conduct

### Feat
- add `UpdateText` to spinner

### Fix
- spinners spin evenly when multiple spinners are started

### Perf
- improve spinner performance

### Style
- clean up code


<a name="v0.1.0"></a>
## [v0.1.0] - 2020-09-28
### Ci
- remove go dep
- add go testing
- update changelog config

### Docs
- add symbols list (codepage 437)
- fix doc link

### Feat
- add spinners
- shorten printer names and add builder methods to printers
- add `Printo` to override printed text
- add `FullWidth` to `HeaderPrinter`
- add terminal size detection

### Refactor
- bump version to "v0.1.0"
- consistent example code for `Printo`
- better comments for `Printo`
- simplify `HeaderPrinter`

### BREAKING CHANGE

printer names changed

removed `Header` and put it's content directly into `HeaderPrinter`


<a name="v0.0.1"></a>
## [v0.0.1] - 2020-09-21
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

### Docs
- put documentation flag in seperate line
- add documentation badge
- update demo
- update header example
- fix release badge style
- add readme
- **readme:** put demo animation under header
- **readme:** fix readme animation href

### Feat
- add aliases to default printers
- add header example
- integrate ci
- add `HeaderPrinter`
- add exported version variable
- add example `override-default-printer`
- change prefix text color to `LightWhite`

### Fix
- header should now work in CI

### Refactor
- bump version to "v0.0.1"
- refactor project
- add comments to functions


<a name="v0.0.0"></a>
## v0.0.0 - 2020-09-18
### Feat
- add changelog template
- configs
- initial commit


[Unreleased]: https://github.com/pterm/pterm/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/pterm/pterm/compare/v0.0.1...v0.1.0
[v0.0.1]: https://github.com/pterm/pterm/compare/v0.0.0...v0.0.1
