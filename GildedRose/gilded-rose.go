package main

import "math"

type Item struct {
	name            string
	sellIn, quality int
}

const standardQualityChange = 1

type QualityStrategy interface {
	ApplyQualityChanges(item *Item)
}

func qualityStrategyFor(item *Item) (strategy QualityStrategy) {
	if isAgedBrie(item) {
		strategy = AgedBrieQualityStrategy{}
	} else if isBackstagePass(item) {
		strategy = BackstagePassesQualityStrategy{}
	} else if isConjuredItem(item) {
		strategy = ConjuredItemQualityStrategy{}
	} else if isLegendaryItem(item) {
		strategy = LegendaryItemQualityStrategy{}
	} else {
		strategy = NormalItemQualityStrategy{}
	}
	return
}

func isAgedBrie(item *Item) bool {
	return item.name == "Aged Brie"
}

func isBackstagePass(item *Item) bool {
	return item.name == "Backstage passes to a TAFKAL80ETC concert"
}

func isConjuredItem(item *Item) bool {
	return item.name == "Conjured Mana Cake"
}

func isLegendaryItem(currentItem *Item) bool {
	return currentItem.name == "Sulfuras, Hand of Ragnaros"
}

type AgedBrieQualityStrategy struct{}

func (a AgedBrieQualityStrategy) ApplyQualityChanges(item *Item) {
	qualityChangesForAgedBrie(item)
}

type BackstagePassesQualityStrategy struct{}

func (b BackstagePassesQualityStrategy) ApplyQualityChanges(item *Item) {
	qualityChangesForBackstagePasses(item)
}

type ConjuredItemQualityStrategy struct {}

func (c ConjuredItemQualityStrategy) ApplyQualityChanges(item *Item) {
	qualityChangesForConjuredItems(item)
}

type LegendaryItemQualityStrategy struct {}

func (l LegendaryItemQualityStrategy) ApplyQualityChanges(item *Item) {
	qualityChangesForLegendaryItems(item)
}

type NormalItemQualityStrategy struct {}

func (n NormalItemQualityStrategy) ApplyQualityChanges(item *Item) {
	qualityChangesForNormalItems(item)
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		currentItem := items[i]

		qualityStrategy := qualityStrategyFor(currentItem)
		qualityStrategy.ApplyQualityChanges(currentItem)

		if !isLegendaryItem(currentItem) {
			capToStandardMaximumQuality(currentItem)
		}

		// letTimePassBy
		if isLegendaryItem(currentItem) {
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
