# Mimetypes

Get the extension of a file from its content type without external dependencies

# Usage

```
go get github.com/NovikovRoman/mimetypes
```

```go
package main

import (
	"fmt"
	"github.com/NovikovRoman/mimetypes"
	"net/http"
)

func main() {
	var resp *http.Response

	// … get response

	// Get the content-type
	contentType := resp.Header.Get("Content-Type")
	// An empty string is returned if the content type is not found
	extension := mimetypes.ExtensionByType(contentType)

	fmt.Println(extension)
}
```
