//go:generate packr2

package web

import "github.com/gobuffalo/packr/v2"

func Box() *packr.Box {
	return packr.New("web", "./")
}
