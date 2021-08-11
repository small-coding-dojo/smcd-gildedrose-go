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
		{Desc: "At the end of each day our system lowers both values for every item", Input: Item{"foo", 10, 100}, Expected: Item{"foo", 9, 99}},
		{Desc: "The Quality of an item is never negative", Input: Item{"foo", 10, 0}, Expected: Item{"foo", 9, 0}},
		{Desc: "Once the **sell by date** has passed, Quality degrades twice as fast", Input: Item{"foo", -1, 10}, Expected: Item{"foo", -2, 8}},
		{Desc: "Edge case: SellIn is 0 and Quality degrades twice as fast", Input: Item{"foo", 0, 2}, Expected: Item{"foo", -1, 0}},
		{Desc: "Edge case: Quality degrades twice as fast AND quality of an item is never negative", Input: Item{"foo", -1, 1}, Expected: Item{"foo", -2, 0}},
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

//TODO: discuss undocumented requirements!
func Test_Aged_Brie_specifications(t *testing.T) {

	testData := []TestStruct{
		{Desc: "\"Aged Brie\" actually increases in Quality the older it gets", Input: Item{"Aged Brie", 10, 10}, Expected: Item{"Aged Brie", 9, 11}},
		{Desc: "undocumented Requirement?", Input: Item{"Aged Brie", -1, 10}, Expected: Item{"Aged Brie", -2, 12}},
		{Desc: "Quality of an item is never more than 50", Input: Item{"Aged Brie", 10, 50}, Expected: Item{"Aged Brie", 9, 50}},
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

//"Sulfuras", being a legendary item, never has to be sold or decreases in Quality
func Test_Sulfuras_specifications(t *testing.T) {

	testData := []TestStruct{
		{Desc: "Sulfuras never has to be sold", Input: Item{"Sulfuras, Hand of Ragnaros", 10, 80}, Expected: Item{"Sulfuras, Hand of Ragnaros", 10, 80}},
		{Desc: "Sulfuras quality is always 80", Input: Item{"Sulfuras, Hand of Ragnaros", 10, 80}, Expected: Item{"Sulfuras, Hand of Ragnaros", 10, 80}},
		{Desc: "Sulfuras never decreases in quality", Input: Item{"Sulfuras, Hand of Ragnaros", 10, 80}, Expected: Item{"Sulfuras, Hand of Ragnaros", 10, 80}},
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
		},
	}

	CheckEffectsOnItemAfterOneDay(t, testData)
}

// "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
// Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
// Quality drops to 0 after the concert

func CheckEffectsOnItemAfterOneDay(t *testing.T, testData []TestStruct) {
	for _, record := range testData {
		t.Run(record.Desc, func(t *testing.T) {
			items := []*Item{&record.Input}
			UpdateQuality(items)
			assert.Equal(t, &record.Expected, items[0])
		})
	}
}
