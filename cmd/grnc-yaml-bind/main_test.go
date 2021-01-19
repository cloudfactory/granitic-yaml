package main

import (
	"path/filepath"
	"testing"
)

func TestValidManifestParsing(t *testing.T) {

	mfp := filepath.Join("testdata", "valid-manifest.yml")

	b := new(YamlDefinitionLoader)

	m, err := b.FacilityManifest(mfp)

	if m == nil {
		t.Errorf("Unexpected nil")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if m.ExternalFacilities == nil || len(m.ExternalFacilities) == 0 {
		t.Errorf("Expected definitions")
	}

	pm := m.ExternalFacilities["FacilityA"]

	if pm == nil {
		t.Errorf("Expected a definition")
	}
}
