# Packages
#### Every Go program is made up of packages.

We had package main at the top of all the programs that has been written so far.

A package named "main" has an entrypoint at the main() function. A main package is compiled into an executable program.

A package by any other name is a "library package". Libraries have no entry point. Libraries simply export functionality that can be used by other packages. For example:

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```
This program is an executable. It is a "main" package and imports from the fmt and math/rand library packages.

# GO init
```
go mod init {REMOTE}/{USERNAME}/hellogo
			github.com/username/repo name
```

# Running a GO Program
```
go run main.go
```

# Building a go program
Execute the command inside the directory which you want to build
```
go build
```

# Running the built program
Instead of hellogo in the command, use the directory name or whatever the name of the executable(.exe) file that was created
```
./hellogo
```

# Installing the program
Run this inside the directory to install the package's(program) compiled binary in the GOBIN directory.
```
go install
```
After installing, just type the program name anywhere in cmd to run the program.
```
hellogo
```