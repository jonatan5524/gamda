
Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function and example function that demonstrates the issue.

Actions to take when contributing:
1. make sure there is an existing issue for your PR, if one does not exist create one, if one exists ask to work on it.
2. fork the project and clone your fork.
3. create a local branch:
  ` git checkout -b <branch_name>`
5. if adding a function make sure to classify its field and add it to the matching folder, for example, Filter is a function related to slices operations, so the Filter function can be found there.
7. after adding a new function make sure there is testing covering it, we are trying to achieve 100% coverage!
8. make sure to add an example function in the test file for documentation
9. make sure to document the new function for documentation
10. before committing run `make` which runs: unit tests, race test, coverage test, go vet, gofmt
11. when you commit, follow angular [commit message header format](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#commit-message-header), our scope is the folders in the codebase.
13. push your fork:
  `git push origin <branch>`
12. open a pull request, make sure to link the issue related to your work with a close keyword: [Linking a pull request to an issue
](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue)
