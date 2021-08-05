package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Positive_SellIn_reduces_quality_by_one(t *testing.T) {
	var items = []*Item{
		&Item{"foo", 10, 100},
	}
	var expected = []*Item{
		&Item{"foo", 9, 99},
	}
	UpdateQuality(items)

	assert.Equal(t, expected, items)
}
