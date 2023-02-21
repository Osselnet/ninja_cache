package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var c *cache = NewCache()

func TestSet(t *testing.T) {
	testTable := []struct {
		key string
		val string
	} {
		{
			key: "userId1",
			val: "11",
		},
		{
			key: "userId2",
			val: "12",
		},
		{
			key: "userId3",
			val: "13",
		},
		{
			key: "userId4",
			val: "14",
		},
		{
			key: "userId5",
			val: "15",
		},
	}

	t.Run("Get", func(t *testing.T) {
		for _, testCase := range testTable {
			c.Set(testCase.key, testCase.val, time.Millisecond * 100)
		}	
	})

	t.Run("wait for the key to expire", func(t *testing.T) {
		t.Parallel()

		for _, testCase := range testTable {
			result, _ := c.Get(testCase.key)
	
			assert.Equal(t, testCase.val, result, fmt.Sprintln("Expected value to be nil"))
		}
	
	})
}