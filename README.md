# Toolbox

`Setenv` sets the value of the environment variable named by the key. It returns an error, if any.



`Getenv` retrieves the value of the environment variable named by the key. It returns the value, which will be "" The second argument is the default value when the environment variable is not read.

```go
package main

import (
	"fmt"

	"github.com/ggorockee/toolbox

func main() {
	_ = toolbox.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%s lives in %s.\n", toolbox.Getenv("NAME", "gopher"), toolbox.Getenv("BURROW"))
	// gopher lives in /usr/gopher.

}
```

