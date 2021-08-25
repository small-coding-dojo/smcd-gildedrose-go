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
			qualityChangesFroOverdueAgedBrie(currentItem)
		} else if currentItem.name == "Backstage passes to a TAFKAL80ETC concert" {
			qualityChangesForBackstagePasses(currentItem)
		} else if currentItem.name == "Conjured Mana Cake" {
			qualityChangesForConjuredItems(currentItem)
		} else if currentItem.name == "Sulfuras, Hand of Ragnaros" {
			qualityChangesForLegendaryItems(currentItem)
		} else {
			qualityChangesForNormalItems(currentItem)
		}

		if currentItem.sellIn <= 0 {
			if currentItem.name == "Aged Brie" {
				// this is left blank intentionally, due to refactoring
			} else if currentItem.name == "Backstage passes to a TAFKAL80ETC concert" {
				currentItem.quality = 0
			} else if currentItem.quality > 0 {
				if currentItem.name == "Conjured Mana Cake" {
					currentItem.quality -= 2
				} else if currentItem.name != "Sulfuras, Hand of Ragnaros" {
					currentItem.quality -= standardQualityChange
				}
			}
		}

		if currentItem.name != "Sulfuras, Hand of Ragnaros" {
			currentItem.quality = capToMaxStandardQuality(currentItem.quality)
		}

		// letTimePassBy
		if currentItem.name == "Sulfuras, Hand of Ragnaros" {
			agingForLegendaryItems(currentItem)
		} else {
			agingForNonLegendaryItems(currentItem)
		}


	}

}

func agingForLegendaryItems(item *Item) {
	
}

func agingForNonLegendaryItems(currentItem *Item) {
	currentItem.sellIn -= 1
}

func qualityChangesForAgedBrie(item *Item) {
	item.quality += standardQualityChange
}

func qualityChangesFroOverdueAgedBrie(item *Item) {
	if item.sellIn <= 0 {
		item.quality += standardQualityChange
	}
}

func qualityChangesForNormalItems(item *Item) {
	if item.quality > 0 {
		item.quality -= standardQualityChange
	}
}

func qualityChangesForConjuredItems(item *Item) {
	if item.quality > 0 {
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
}

func qualityChangesForLegendaryItems(item *Item){

}

func capToMaxStandardQuality(quality int) int {
	return int(math.Min(50, float64(quality)))
}
