package gobpmn_reader_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	gobpmn_reader "github.com/deemount/gobpmnReader"
)

var (
	// Defaults
	DefaultFiletestPath     = "testdata"
	DefaultFiletestNameBPMN = "test.bpmn"
	DefaultFiletestNameJSON = "test.json"
	DefaultTestBpmnChecksum = [32]uint8{0xd2, 0xf8, 0xfe, 0xac, 0x15, 0xa5, 0xcd, 0xc6, 0x4f, 0x69, 0x34, 0xf, 0x19, 0xcd, 0x98, 0x71, 0x43, 0x48, 0x3b, 0x3c, 0xe, 0xde, 0xaa, 0xea, 0x96, 0x3f, 0xda, 0x7d, 0x5a, 0x97, 0xbd, 0x1b}
)

// TestReadFileAndUnmarshal tests the ReadFileAndUnmarshal function
func TestReadFileAndUnmarshal(t *testing.T) {

	// create a new reader
	reader := gobpmn_reader.Reader{}

	// read a file
	_, err := reader.ReadFileAndUnmarshal(DefaultFiletestPath + "/" + DefaultFiletestNameBPMN)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if assert.Equal(t, DefaultTestBpmnChecksum, reader.Checksum) {
		t.Logf("Expected checksum %#v", DefaultTestBpmnChecksum)
	}

}
