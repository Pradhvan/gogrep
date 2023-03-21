package iohandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var directory = "testdata"
var totalFilesInDir = 3

// Using a pre-existing data directory and
// not programitically creating one as setup & teardown
//because in the refactor issue #10
// mocking would take care of it.
//link: https://github.com/Pradhvan/gogrep/issues/10

func TestListFilesInDir(t *testing.T) {
	assert.DirExists(t, directory)
	filepaths, err := ListFilesInDir(directory)
	assert.EqualValues(t, nil, err)
	assert.Equal(t, len(filepaths), totalFilesInDir)
}
