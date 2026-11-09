package uncdn

import (
	"compress/gzip"
	"html"
	"net/http"
	"strings"
)

type cdnController struct{}

// CdnHandler returns a handler that serves the requested CDN assets
// (combined, gzipped JS or CSS).
func CdnHandler() func(w http.ResponseWriter, r *http.Request) {
	return newCdnController().Handler
}

// newCdnController creates the internal controller backing CdnHandler.
func newCdnController() *cdnController {
	return &cdnController{}
}

func (c cdnController) Handler(w http.ResponseWriter, r *http.Request) {
	required, extension := c.findRequiredAndExtension(r)

	if len(required) < 1 {
		c.writeHTML(w, "Nothing requested")
		return
	}
	if extension == "" {
		c.writeHTML(w, "No extension provided")
		return
	}
	if extension != "css" && extension != "js" {
		if _, err := w.Write([]byte("Extension " + html.EscapeString(extension) + " not supported")); err != nil { // #nosec G705 -- extension is escaped
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
		return
	}

	if extension == "js" {
		c.writeGzipJSResponse(w, r, c.compileJS(required))
		return
	}
	if extension == "css" {
		c.writeGzipCSSResponse(w, r, c.compileCSS(required))
		return
	}

	c.writeHTML(w, "Extension "+extension+" not found")
}

func (c cdnController) writeHTML(w http.ResponseWriter, body string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(body)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (c cdnController) writeGzip(
	w http.ResponseWriter,
	contentType string,
	body string,
) error {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	defer func() { _ = gz.Close() }()

	if _, err := gz.Write([]byte(body)); err != nil {
		return err
	}

	return nil
}

func (c cdnController) compileJS(required []string) string {
	js := []string{}
	for _, item := range required {
		js = append(js, c.findJS(item))
	}
	return strings.Join(js, "\n;\n")
}

func (c cdnController) compileCSS(required []string) string {
	css := []string{}
	for _, item := range required {
		css = append(css, c.findCss(item))
	}
	return strings.Join(css, "\n\n")
}

func (c cdnController) findJS(required string) string {
	if required == "jq360" {
		return Jquery360()
	}

	if required == "bs523" {
		return BootstrapJs523()
	}

	if required == "vue3" {
		return VueJs3()
	}

	if required == "web260" {
		return WebJs260()
	}

	if required == "ntf" {
		return NotifyJs()
	}

	if required == "swal" {
		return Sweetalert2_11432()
	}

	return ""
}

func (c cdnController) findCss(required string) string {
	css := map[string]func() string{
		"bs523":          BootstrapCss523,
		"bs523cerulean":  BootstrapCeruleanCss523,
		"bs523cosmo":     BootstrapCosmoCss523,
		"bs523cyborg":    BootstrapCyborgCss523,
		"bs523darkly":    BootstrapDarklyCss523,
		"bs523flatly":    BootstrapFlatlyCss523,
		"bs523journal":   BootstrapJournalCss523,
		"bs523litera":    BootstrapLiteraCss523,
		"bs523lumen":     BootstrapLumenCss523,
		"bs523lux":       BootstrapLuxCss523,
		"bs523materia":   BootstrapMateriaCss523,
		"bs523minty":     BootstrapMintyCss523,
		"bs523morph":     BootstrapMorphCss523,
		"bs523pulse":     BootstrapPulseCss523,
		"bs523quartz":    BootstrapQuartzCss523,
		"bs523sandstone": BootstrapSandstoneCss523,
		"bs523simplex":   BootstrapSimplexCss523,
		"bs523sketchy":   BootstrapSketchyCss523,
		"bs523slate":     BootstrapSlateCss523,
		"bs523solar":     BootstrapSolarCss523,
		"bs523spacelab":  BootstrapSpacelabCss523,
		"bs523superhero": BootstrapSuperheroCss523,
		"bs523united":    BootstrapUnitedCss523,
		"bs523vapor":     BootstrapVaporCss523,
		"bs523yeti":      BootstrapYetiCss523,
		"bs523zephyr":    BootstrapZephyrCss523,
	}
	if style, ok := css[required]; ok {
		return style()
	}
	return ""
}

func (c cdnController) findRequiredAndExtension(req *http.Request) (required []string, extension string) {
	uriParts := strings.Split(strings.Trim(req.RequestURI, "/"), "/")
	if len(uriParts) < 2 {
		return []string{}, ""
	}

	name := uriParts[1]
	if name == "" {
		return []string{}, ""
	}

	nameParts := strings.Split(name, ".")
	if len(nameParts) < 2 {
		return []string{}, ""
	}

	requiredArray := strings.Split(nameParts[0], "-")
	if requiredArray == nil {
		return []string{}, ""
	}

	return requiredArray, nameParts[1]
}

func (c cdnController) writeGzipJSResponse(w http.ResponseWriter, _ *http.Request, content string) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer func() { _ = gz.Close() }()
	_, _ = gz.Write([]byte(content))
}

func (c cdnController) writeGzipCSSResponse(w http.ResponseWriter, _ *http.Request, content string) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer func() { _ = gz.Close() }()
	_, _ = gz.Write([]byte(content))
}
