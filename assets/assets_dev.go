// +build !release

package assets

import "net/http"

var HTTP http.FileSystem = http.Dir("./res")
