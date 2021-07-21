package view

import (
	"github.com/cbroglie/mustache"
)

const templateRoot string = "view/templates/"

type Template string
const (
	Home Template = "home"
)

type Vars map[string]interface{}

func Render(t Template, context map[string]interface{}) (string, error) {
	mustache.AllowMissingVariables = false
	return mustache.RenderFileInLayout(templateRoot + string(t) + ".mustache", templateRoot + "layout.mustache", context)
}
