# Strike A Match
A fuzzy string matching algorithm in Go, as described at <http://www.catalysoft.com/articles/strikeamatch.html>
## Installation
  $ go get github.com/ivanfoong/strikeamatch
## Unit Test
  $ cd $GOPATH/src/github.com/ivanfoong/strikeamatch
  $ go test
## Usage
```go
package main

import (
  "fmt"
  "github.com/ivanfoong/strikeamatch"
)

func main() {
  text1 := "Web Database Applications with PHP & MySQL"
  text2 := "Web Database Applications"
  fmt.Printf("`%s` compared to `%s` has a score of %.2f\n", text1, text2, strikeamatch.CompareString(text1, text2))
  // prints `Web Database Applications with PHP & MySQL` compared to `Web Database Applications` has a score of 0.82
}
```
## Changelog
* 1.0.0 - Added CompareString method
## License
MIT License