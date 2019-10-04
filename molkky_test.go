package molkky

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMolkkyStartWithScoreZero(t *testing.T) {
	m := startMolkky()
	// given
	expected := 0
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

func TestMolkkyScoreWithNoFallenQuellsl(t *testing.T) {
	m := startMolkky()
	m.setFallenkeels([]int{})
	// given
	expected := 0
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

func TestMolkkyScoreTwoPointsIfTwoFallenQuells(t *testing.T) {
	m := startMolkky()
	m.setFallenkeels([]int{3, 12})
	// given
	expected := 2
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

func TestMolkkyScoreQuellPointsIfOnlyOneQuell(t *testing.T) {
	m := startMolkky()
	m.setFallenkeels([]int{3})
	// given
	expected := 3
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

func TestMolkkyScoreWithInvalidValue(t *testing.T) {
	m := startMolkky()
	m.setFallenkeels([]int{14})
	// given
	expected := 0
	// than
	actual := m.getScore()
	// that
	assert.Equal(t, expected, actual)
}

type keelValueTest struct {
	n        int
	expected bool
}

func TestKeelValueIsBetweenOneAndTwelve(t *testing.T) {
	tests := []keelValueTest{
		{1, true},
		{5, true},
		{12, true},
		{0, false},
		{-2, false},
		{13, false},
	}
	m := startMolkky()
	for _, keel := range tests {
		actual := m.isValidKeel(keel.n)
		if actual != keel.expected {
			t.Errorf("isValidkeel(%d): expected %t, actual %t", keel.n, keel.expected, actual)
		}
	}
}

type keelsValueTest struct {
	n        []int
	expected bool
}

func TestKeelsValueAreValid(t *testing.T) {
	tests := []keelsValueTest{
		{[]int{1}, true},
		{[]int{1, 5, 12}, true},
		{[]int{-1}, false},
		{[]int{0}, false},
		{[]int{0, 13}, false},
	}
	m := startMolkky()
	for _, keel := range tests {
		actual := m.isValidKeels(keel.n)
		if actual != keel.expected {
			t.Errorf("isValidkeels(%v): expected %t, actual %t", keel.n, keel.expected, actual)
		}
	}
}
