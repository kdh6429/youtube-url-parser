# YOUTUBE-URL-PARSE
Extract YouTube ID from url written in golang.

Tested with various URL types, so please let me know if it is not worked properly.


### Install
```
go get github.com/kdh6429/youtube-url-parser 
```
### Example
``` go
package main

import (
	"fmt"
	. "github.com/kdh6429/youtube-url-parser"
)

func main() {
	fmt.Println(ExtractID("https://www.youtube.com/watch?v=wDfqXR_5yyQ"))
}
```
