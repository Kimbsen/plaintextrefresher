# plaintextrefresher


```
package main

import (
	"fmt"
	"github.com/Kimbsen/plaintextrefresher"
	"net/http"
)

func main() {
	http.HandleFunc("/", plaintextrefresher.Handle("/stats"))
	count := 0
	http.HandleFunc("/stats", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(rw, "Hello World %d", count)
		count++
	})
	http.ListenAndServe(":8080", nil)
}

```
