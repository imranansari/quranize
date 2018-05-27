package quran

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuranSimpleClean(t *testing.T) {
	expected := 114
	actual := len(NewQuranSimpleClean().Suras)
	assert.Equal(t, expected, actual)
}

func TestNewQuranSimpleEnhanced(t *testing.T) {
	expected := 114
	actual := len(NewQuranSimpleEnhanced().Suras)
	assert.Equal(t, expected, actual)
}

func TestGetAyaFound(t *testing.T) {
	text, err := quranTest.GetAya(1, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetAyaSuraNotFound(t *testing.T) {
	_, err := quranTest.GetAya(0, 0)
	assert.Error(t, err)
}

func TestGetAyaAyaNotFound(t *testing.T) {
	_, err := quranTest.GetAya(1, 0)
	assert.Error(t, err)
}

func TestGetSuraNameFound(t *testing.T) {
	text, err := quranTest.GetSuraName(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetSuraNameNotFound(t *testing.T) {
	_, err := quranTest.GetSuraName(0)
	assert.Error(t, err)
}
