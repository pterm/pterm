# Code Review Checks

> [!NOTE]
> This page is for people who review pull requests inside the official PTerm repository on GitHub.

Every pull request made in the official PTerm repository has to be checked manually, before it can be merged.
By doing so, we ensure that PTerm can always deliver the highest possible quality of code and functionality. 

To review a pull request, the following text must be appended as a comment to an `approve` or `request changes`. 

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
| Documentation in `./docs` is updated correctly.                       | NO     |
```
