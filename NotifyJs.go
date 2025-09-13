package uncdn

import "github.com/dracory/uncdn/templates"

func NotifyJs() string {
	return templates.ToString("vend/notifyjs/notifyjs.min.js")
}
