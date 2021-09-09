package main

import "math"

type Item struct {
	name            string
	sellIn, quality int
}

const standardQualityChange = 1

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		currentItem := items[i]

		if currentItem.name == "Aged Brie" {
			qualityChangesForAgedBrie(currentItem)
		} else if currentItem.name == "Backstage passes to a TAFKAL80ETC concert" {
			qualityChangesForBackstagePasses(currentItem)
		} else if currentItem.name == "Conjured Mana Cake" {
			qualityChangesForConjuredItems(currentItem)
		} else if currentItem.name == "Sulfuras, Hand of Ragnaros" {
			qualityChangesForLegendaryItems(currentItem)
		} else {
			qualityChangesForNormalItems(currentItem)
		}

		if currentItem.name != "Sulfuras, Hand of Ragnaros" {
			capToStandardMaximumQuality(currentItem)
		}

		// letTimePassBy
		if currentItem.name == "Sulfuras, Hand of Ragnaros" {
			agingForLegendaryItems(currentItem)
		} else {
			agingForNonLegendaryItems(currentItem)
		}
	}

}

func capToStandardMaximumQuality(item *Item) {
	item.quality = int(math.Min(50, float64(item.quality)))
}

func agingForLegendaryItems(item *Item) {

}

func agingForNonLegendaryItems(item *Item) {
	item.sellIn -= 1
}

func qualityChangesForAgedBrie(item *Item) {
	item.quality += standardQualityChange
	if item.sellIn <= 0 {
		item.quality += standardQualityChange
	}
}

func qualityChangesForNormalItems(item *Item) {
	if item.quality > 0 {
		item.quality -= standardQualityChange
	}
	if (item.sellIn <= 0) && (item.quality > 0) {
		item.quality -= standardQualityChange
	}

}

func qualityChangesForConjuredItems(item *Item) {
	if item.quality > 0 {
		item.quality -= 2
	}
	if (item.sellIn <= 0) && (item.quality > 0) {
		item.quality -= 2
	}
}

func qualityChangesForBackstagePasses(item *Item) {
	item.quality += standardQualityChange
	if item.sellIn <= 10 {
		item.quality += standardQualityChange
	}
	if item.sellIn <= 5 {
		item.quality += standardQualityChange
	}

	if item.sellIn <= 0 {
		item.quality = 0
	}
}

func qualityChangesForLegendaryItems(item *Item) {

}
