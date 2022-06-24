# homograph
## Installation

```
go get github.com/shota-k2018/homograph
```

## Example Usage

```go
package main

import (
	"github.com/shota-k2018/homograph"
	"log"
)

func main() {
	// If you want to minimize the number of patterns generated, set the second argument to true
	log.Printf("gen: %+v", homograph.Generate("test"), false)
}
```
