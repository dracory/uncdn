package uncdn

import "github.com/dracory/uncdn/templates"

func Raw(path string) string {
	return templates.ToString("vend/" + path)
}
