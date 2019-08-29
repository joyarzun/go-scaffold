package iohelpers_test

import (
	"os"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/stretchr/testify/assert"
)

func Test_Read_Success_ShouldReadContentOfFile(t *testing.T) {
	reader, err := os.Open("testdata/file_to_read.txt")
	assert.Nil(t, err)
	defer reader.Close()
	assert.Equal(t, "file-to-read-content\n", iohelpers.Read(reader))
}
