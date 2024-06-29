package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLru(test *testing.T) {
	//arrange
	testTable := []struct {
		cap      int
		key      int
		value    int
		expected int
	}{
		{1, 0, 1, 1},
		{2, 1, 2, 2},
	}

	for _, testCase := range testTable {
		l := newLRUCache(testCase.cap)
		_ = l.setLRU(testCase.key, testCase.value)
		expected := testCase.expected
		//act
		result, _ := l.getLRU(testCase.key)
		test.Logf("expected %d, got %d", expected, result)
		//assert
		assert.Equal(test, expected, result)
	}

}
