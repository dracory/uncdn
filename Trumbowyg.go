package uncdn

import "github.com/dracory/uncdn/templates"

func Trumbowyg2260Css() string {
	return templates.ToString("vend/trumbowyg/2.26.0/trumbowyg.min.css")
}

func Trumbowyg2260Js() string {
	return templates.ToString("vend/trumbowyg/2.26.0/trumbowyg.min.js")
}
