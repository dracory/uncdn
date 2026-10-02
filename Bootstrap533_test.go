package uncdn

import (
	"strings"
	"testing"
)

func TestBootstrapCss533(t *testing.T) {
	output := BootstrapCss533()
	expected := ".container"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs533(t *testing.T) {
	output := BootstrapJs533()
	expected := "data-bs-no-jquery"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCeruleanCss533(t *testing.T) {
	table := map[string]func() string{
		"bootstrap": BootstrapCeruleanCss533,
		"cerulean":  BootstrapCeruleanCss533,
		"cosmo":     BootstrapCosmoCss533,
		"cyborg":    BootstrapCyborgCss533,
		"darkly":    BootstrapDarklyCss533,
		"flatly":    BootstrapFlatlyCss533,
		"journal":   BootstrapJournalCss533,
		"litera":    BootstrapLiteraCss533,
		"lumen":     BootstrapLumenCss533,
		"lux":       BootstrapLuxCss533,
		"materia":   BootstrapMateriaCss533,
		"minty":     BootstrapMintyCss533,
		"morph":     BootstrapMorphCss533,
		"pulse":     BootstrapPulseCss533,
		"quartz":    BootstrapQuartzCss533,
		"sandstone": BootstrapSandstoneCss533,
		"simplex":   BootstrapSimplexCss533,
		"sketchy":   BootstrapSketchyCss533,
		"slate":     BootstrapSlateCss533,
		"solar":     BootstrapSolarCss533,
		"spacelab":  BootstrapSpacelabCss533,
		"superhero": BootstrapSuperheroCss533,
		"united":    BootstrapUnitedCss533,
		"vapor":     BootstrapVaporCss533,
		"yeti":      BootstrapYetiCss533,
		"zephyr":    BootstrapZephyrCss533,
	}
	for key := range table {
		output := table[key]()
		expected := ".container"
		if !strings.Contains(output, expected) {
			t.Error("Does not contain '" + expected + "', Output:" + output)
		}
	}
}
