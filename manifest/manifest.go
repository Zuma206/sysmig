package manifest

import (
	_ "embed"
	"encoding/json"
)

type ManifestFile struct {
	Version string `json:"version"`
}

//go:embed manifest.json
var manifest []byte

// Unmarshalled manifest.json
var Manifest ManifestFile

func init() {
	json.Unmarshal(manifest, &Manifest)
}
