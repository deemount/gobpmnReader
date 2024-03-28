package gobpmn_reader

import (
	"crypto/sha256"
	"encoding/xml"
	"os"

	"github.com/deemount/gobpmnModels/pkg/core"
)

// Reader ...
type Reader struct {
	Def      core.TDefinitions
	Checksum [32]byte
}

// ReadFileAndUnmarshal reads a given BPMN file by file argument, sets the checksum
// and unmarshals it into the core.TDefinitions struct
func (reader *Reader) ReadFileAndUnmarshal(file string) (*core.TDefinitions, error) {

	var err error

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// create checksum
	reader.Checksum = sha256.Sum256(data)

	return reader.Unmarshal(data)

}

// Unmarshal unmarshals a given byte array into the core.TDefinitions struct
func (reader *Reader) Unmarshal(data []byte) (*core.TDefinitions, error) {

	err := xml.Unmarshal(data, &reader.Def)
	if err != nil {
		return nil, err
	}

	return &reader.Def, nil

}
