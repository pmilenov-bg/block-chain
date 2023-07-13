package models

import (
	"testing"
)

func TestExportToString(t *testing.T) {
	strData := "Hello, World!"
	blockData := BlockData{stringData: strData}

	result := blockData.ExportToString()

	if result != strData {
		t.Errorf("Expected ExportToString() to return '%s', but got '%s'", strData, result)
	}
}

func TestCreateBlockData(t *testing.T) {
	strData := "Hello, World!"

	blockData := CreateBlockData(strData)

	if blockData.stringData != strData {
		t.Errorf("Expected CreateBlockData() to set stringData to '%s', but got '%s'", strData, blockData.stringData)
	}
}
