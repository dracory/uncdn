package uncdn

import (
	"strings"
	"testing"
)

func TestBootstrapCss538(t *testing.T) {
	output := BootstrapCss538()
	expected := ".container"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs538(t *testing.T) {
	output := BootstrapJs538()
	expected := "data-bs-no-jquery"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCeruleanCss538(t *testing.T) {
	table := map[string]func() string{
		"bootstrap": BootstrapCeruleanCss538,
		"cerulean":  BootstrapCeruleanCss538,
		"cosmo":     BootstrapCosmoCss538,
		"cyborg":    BootstrapCyborgCss538,
		"darkly":    BootstrapDarklyCss538,
		"flatly":    BootstrapFlatlyCss538,
		"journal":   BootstrapJournalCss538,
		"litera":    BootstrapLiteraCss538,
		"lumen":     BootstrapLumenCss538,
		"lux":       BootstrapLuxCss538,
		"materia":   BootstrapMateriaCss538,
		"minty":     BootstrapMintyCss538,
		"morph":     BootstrapMorphCss538,
		"pulse":     BootstrapPulseCss538,
		"quartz":    BootstrapQuartzCss538,
		"sandstone": BootstrapSandstoneCss538,
		"simplex":   BootstrapSimplexCss538,
		"sketchy":   BootstrapSketchyCss538,
		"slate":     BootstrapSlateCss538,
		"solar":     BootstrapSolarCss538,
		"spacelab":  BootstrapSpacelabCss538,
		"superhero": BootstrapSuperheroCss538,
		"united":    BootstrapUnitedCss538,
		"vapor":     BootstrapVaporCss538,
		"yeti":      BootstrapYetiCss538,
		"zephyr":    BootstrapZephyrCss538,
	}
	for key := range table {
		output := table[key]()
		expected := ".container"
		if !strings.Contains(output, expected) {
			t.Error("Does not contain '" + expected + "', Output:" + output)
		}
	}
}
