# BBChat::server::Golang

This is a boilerplate template for new GO projects. 
**If you want to add or improve this boilerplate, create a PR**

## Folders
Much more information (and more standard folders) can be found here: [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

* `build` - Packaging and Continuous Integration. (See link above!)
* `cmd` - Main applications for this project, the folder should match the name of the executable.
* `internal` - Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself.
* `pkg` - Library code that's ok to use by external applications. Other projects will import these libraries expecting them to work, so think twice before you put something here.
* `scripts` - Scripts to perform various build, install, analysis, etc operations. These scripts keep the root level Makefile small and simple.


