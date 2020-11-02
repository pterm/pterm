# Code Review Checks

```markdown 
<!-- Values for "Passed": [ N/A | NO | YES] -->
| Description                                                           | Passed |
|-----------------------------------------------------------------------|:------:|
| Every `WithXXX(...)` function returns a pointer of the parent struct. | NO     |
| Test contains nil check.                                              | NO     |
| Printers implement the right interface.                               | NO     |
| Variable and function names describe what they do.                    | NO     |
| Printer contains theme support.                                       | NO     |
| Styles in structs must be pointers.                                   | NO     |
| Printer tests uses correct test template.                             | NO     |
| Commit messages follow conventional commit style. `If NO -> Squash`   | NO     |
| Possible errors are declared in the `errors.go` file.                 | NO     |
| Theme styles do not return pointers                                   | NO     |
```
