package implentacoes

import(
	"fmt"
	" path P denotes the package found in
the directory DIR/src/P for some DIR listed in the GOPATH
environment variable (For more details see: 'go help gopath').

If no import paths are given, the action applies to the
package in the current directory.

There are four reserved names for paths that should not be used
for packages to be built with the go tool:

- "main" denotes the top-level package in a stand-alone executable.

- "all" expands to all packages found in all the GOPATH
trees. For example, 'go list all' lists all the packages on the local
system. When using modules, "all" expands to all packages in
the main module and their dependencies, including dependencies
needed by tests of any of those.

- "std" is like all but expands to just the packages in the standard
Go library.

- "cmd" expands to the Go repository's commands and their
internal libraries.

Import paths beginning with "cmd/" only match source code in
the Go repository.

An import path is a pattern if it includes one or more "..." wildcards,
each of which can match any string, including the empty string and
strings containing slashes. Such a pattern expands to all package
directories found in the GOPATH trees with names matching the
patterns."
)

func main(){
	l := &ArrayList{}
	l.Init(2)
	l.Add(10)
	l.Add(6)
	l.Add(4)
	fmt.Println(l.Size())
}