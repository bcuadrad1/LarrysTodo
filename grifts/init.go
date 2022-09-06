package grifts

import (
	"github.com/bcuadrad1/LarrysTodo/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
