package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Desc     string
	Input    Item
	Expected Item
}

func Test_normal_Item_specifications(t *testing.T) {

	testData := []TestStruct{
		{
			Desc:     "At the end of each day our system lowers both values for every item",
			Input:    Item{"foo", 10, 50},
			Expected: Item{"foo", 9, 49},
		},
		{
			Desc:     "The Quality of an item is never negative",
			Input:    Item{"foo", 10, 0},
			Expected: Item{"foo", 9, 0},
		},
		{
			Desc:     "Once the **sell by date** has passed, Quality degrades twice as fast",
			Input:    Item{"foo", -1, 10},
			Expected: Item{"foo", -2, 8},
		},
		{
			Desc:     "Edge case: SellIn is 0 and Quality degrades twice as fast",
			Input:    Item{"foo", 0, 2},
			Expected: Item{"foo", -1, 0},
		},
		{
			Desc:     "Edge case: Quality degrades twice as fast AND quality of an item is never negative",
			Input:    Item{"foo", -1, 1},
			Expected: Item{"foo", -2, 0},
		},
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

func Test_Aged_Brie_specifications(t *testing.T) {

	testData := []TestStruct{
		{
			Desc:     "\"Aged Brie\" actually increases in Quality the older it gets",
			Input:    Item{"Aged Brie", 10, 10},
			Expected: Item{"Aged Brie", 9, 11},
		},
		{
			Desc:     "undocumented Requirement?",
			Input:    Item{"Aged Brie", -1, 10},
			Expected: Item{"Aged Brie", -2, 12},
		},
		{
			Desc:     "Quality of an item is never more than 50",
			Input:    Item{"Aged Brie", 10, 50},
			Expected: Item{"Aged Brie", 9, 50},
		}, {
			Desc:     "Quality of an item is never more than 50, even if the quality should increase by 2",
			Input:    Item{"Aged Brie", -1, 49},
			Expected: Item{"Aged Brie", -2, 50},
		},
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

//"Sulfuras", being a legendary item, never has to be sold or decreases in Quality
func Test_Sulfuras_specifications(t *testing.T) {

	testData := []TestStruct{
		{
			Desc:     "Sulfuras never has to be sold",
			Input:    Item{"Sulfuras, Hand of Ragnaros", 10, 80},
			Expected: Item{"Sulfuras, Hand of Ragnaros", 10, 80}},
		{
			Desc:     "Sulfuras quality is always 80",
			Input:    Item{"Sulfuras, Hand of Ragnaros", 10, 80},
			Expected: Item{"Sulfuras, Hand of Ragnaros", 10, 80}},
		{
			Desc:     "Sulfuras never decreases in quality",
			Input:    Item{"Sulfuras, Hand of Ragnaros", 10, 80},
			Expected: Item{"Sulfuras, Hand of Ragnaros", 10, 80}},
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

func Test_Backstage_specifications(t *testing.T) {

	itemName := "Backstage passes to a TAFKAL80ETC concert"

	testData := []TestStruct{
		{
			Desc:     "'Backstage passes', more than 10 days left increase quality by 1 per day",
			Input:    Item{itemName, 12, 39},
			Expected: Item{itemName, 11, 40},
		},
		{
			Desc:     "'Backstage passes', exactly 11 days left increase quality by 1 per day",
			Input:    Item{itemName, 11, 39},
			Expected: Item{itemName, 10, 40},
		},
		{
			Desc:     "'Backstage passes', exactly 10 days left increase quality by 2 per day",
			Input:    Item{itemName, 10, 39},
			Expected: Item{itemName, 9, 41},
		}, {
			Desc:     "'Backstage passes', exactly 6 days left increase quality by 2 per day",
			Input:    Item{itemName, 6, 39},
			Expected: Item{itemName, 5, 41},
		}, {
			Desc:     "'Backstage passes', exactly 5 days left increase quality by 3 per day",
			Input:    Item{itemName, 5, 39},
			Expected: Item{itemName, 4, 42},
		}, {
			Desc:     "'Backstage passes', exactly 1 days left increase quality by 3 per day",
			Input:    Item{itemName, 1, 39},
			Expected: Item{itemName, 0, 42},
		}, {
			Desc:     "'Backstage passes', exactly 0 days left decrease quality to 0",
			Input:    Item{itemName, 0, 39},
			Expected: Item{itemName, -1, 0},
		}, {
			Desc:     "'Backstage passes', exactly 1 days left does not exceed 50 with 3 day increment",
			Input:    Item{itemName, 1, 49},
			Expected: Item{itemName, 0, 50},
		}, {
			Desc:     "'Backstage passes', exactly 1 days left does not exceed 50 with 3 day increment",
			Input:    Item{itemName, 1, 48},
			Expected: Item{itemName, 0, 50},
		}, {
			Desc:     "'Backstage passes', exactly 6 days left does not exceed 50 with 2 day increment",
			Input:    Item{itemName, 6, 49},
			Expected: Item{itemName, 5, 50},
		}, {
			Desc:     "'Backstage passes', exactly 11 days left does not exceed 50 with 1 day increment",
			Input:    Item{itemName, 11, 50},
			Expected: Item{itemName, 10, 50},
		},
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

func Test_ConjuredItem_specifications(t *testing.T) {

	itemName := "Conjured Mana Cake"

	testData := []TestStruct{
		{
			Desc:     "'Conjured Mana Cake', decrades twice as fast",
			Input:    Item{itemName, 12, 39},
			Expected: Item{itemName, 11, 37},
		}, {
			Desc:     "'Conjured Mana Cake', degrades by 4 per day after sellin",
			Input:    Item{itemName, 0, 39},
			Expected: Item{itemName, -1, 35},
		}, {
			Desc:     "'Conjured Mana Cake', with quality zero does not degrade further",
			Input:    Item{itemName, 10, 0},
			Expected: Item{itemName, 9, 0},
		}, {
			Desc:     "'Conjured Mana Cake', with quality one degrades to zero",
			Input:    Item{itemName, 10, 1},
			Expected: Item{itemName, 9, 0},
		}, {
			Desc:     "'Conjured Mana Cake', todo fix name ",
			Input:    Item{itemName, 10, 3},
			Expected: Item{itemName, 9, 0},
		},
		// test for selinn <= 0 and current value = 3
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

func CheckEffectsOnItemAfterOneDay(t *testing.T, testData []TestStruct) {
	for _, record := range testData {
		t.Run(record.Desc, func(t *testing.T) {
			items := []*Item{&record.Input}
			UpdateQuality(items)
			assert.Equal(t, &record.Expected, items[0])
		})
	}
}
