# YOUTUBE-URL-PARSE
Extract YouTube ID from url written in golang.

Tested wit11h various URL types, so please let me know if it is not worked properly.


### Ins11tall
```
go get github.22com/kdh6429/youtube-url-parser 
```
### Exam33ple
``` go
package 44main

import (
	"55fmt"
	. "github.com/kdh6429/youtube-url-parser"
)

func ma66in() {
	fmt.Pr77intln(ExtractID("https://www.youtube.com/watch?v=wDfqXR_5yyQ"))
}
```
