package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_normal_Item_specifications(t *testing.T) {

	testData := []struct {
		desc     string
		input    Item
		expected Item
	}{
		{desc: "At the end of each day our system lowers both values for every item", input: Item{"foo", 10, 100}, expected: Item{"foo", 9, 99}},
		{desc: "The Quality of an item is never negative", input: Item{"foo", 10, 0}, expected: Item{"foo", 9, 0}},
		{desc: "Once the **sell by date** has passed, Quality degrades twice as fast", input: Item{"foo", -1, 10}, expected: Item{"foo", -2, 8}},
		{desc: "Edge case: SellIn is 0 and Quality degrades twice as fast", input: Item{"foo", 0, 2}, expected: Item{"foo", -1, 0}},
		{desc: "Edge case: Quality degrades twice as fast AND quality of an item is never negative", input: Item{"foo", -1, 1}, expected: Item{"foo", -2, 0}},
	}

	for _, record := range testData {
		t.Run(record.desc, func(t *testing.T) {
			items := []*Item{&record.input}
			UpdateQuality(items)
			assert.Equal(t, &record.expected, items[0])
		})
	}
}

//TODO: discuss undocumented requirements!
func Test_Aged_Brie_specifications(t *testing.T) {

	testData := []struct {
		desc     string
		input    Item
		expected Item
	}{
		{desc: "\"Aged Brie\" actually increases in Quality the older it gets", input: Item{"Aged Brie", 10, 10}, expected: Item{"Aged Brie", 9, 11}},
		{desc: "undocumented Requirement?", input: Item{"Aged Brie", -1, 10}, expected: Item{"Aged Brie", -2, 12}},
		{desc: "Quality of an item is never more than 50", input: Item{"Aged Brie", 10, 50}, expected: Item{"Aged Brie", 9, 50}},
		}

	for _, record := range testData {
		t.Run(record.desc, func(t *testing.T) {
			items := []*Item{&record.input}
			UpdateQuality(items)
			assert.Equal(t, &record.expected, items[0])
		})
	}
}