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

		qualityStrategy := qualityStrategyFor(currentItem)
		qualityStrategy.ApplyChangesForOneDay(currentItem)
	}
}

type QualityStrategy interface {
	IsApplicableFor(item *Item) bool
	ApplyChangesForOneDay(item *Item)
}

func qualityStrategyFor(item *Item) QualityStrategy {

	strategies := []QualityStrategy{
		AgedBrieQualityStrategy{},
		BackstagePassesQualityStrategy{},
		ConjuredItemQualityStrategy{},
		LegendaryItemQualityStrategy{},
	}

	for _, qualityStrategy := range strategies {
		if qualityStrategy.IsApplicableFor(item) {
			return qualityStrategy
		}
	}

	return NormalItemQualityStrategy{}
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

func (a AgedBrieQualityStrategy) IsApplicableFor(item *Item) bool {
	return isAgedBrie(item)
}

func (a AgedBrieQualityStrategy) ApplyChangesForOneDay(item *Item) {
	item.quality += standardQualityChange
	if item.sellIn <= 0 {
		item.quality += standardQualityChange
	}
	capToStandardMaximumQuality(item)
	applyStandardAging(item)
}

type BackstagePassesQualityStrategy struct{}

func (b BackstagePassesQualityStrategy) IsApplicableFor(item *Item) bool {
	return isBackstagePass(item)
}

func (b BackstagePassesQualityStrategy) ApplyChangesForOneDay(item *Item) {

	qualityUpdate := 0

	if item.sellIn <= 5 {
		item.quality += 2 * standardQualityChange
	} else if item.sellIn <= 10 {
		item.quality += standardQualityChange
	}
	item.quality += standardQualityChange

	if item.sellIn <= 0 {
		item.quality = 0
	}

	item.quality += qualityUpdate

	capToStandardMaximumQuality(item)
	applyStandardAging(item)
}

type ConjuredItemQualityStrategy struct{}

func (c ConjuredItemQualityStrategy) IsApplicableFor(item *Item) bool {
	return isConjuredItem(item)
}

func (c ConjuredItemQualityStrategy) ApplyChangesForOneDay(item *Item) {
	if item.quality > 0 {
		item.quality -= 2
	}
	if (item.sellIn <= 0) && (item.quality > 0) {
		item.quality -= 2
	}
	capToStandardMaximumQuality(item)
	applyStandardAging(item)
}

type LegendaryItemQualityStrategy struct{}

func (l LegendaryItemQualityStrategy) IsApplicableFor(item *Item) bool {
	return isLegendaryItem(item)
}

func (l LegendaryItemQualityStrategy) ApplyChangesForOneDay(item *Item) {

}

type NormalItemQualityStrategy struct{}

func (n NormalItemQualityStrategy) IsApplicableFor(item *Item) bool {
	return true
}

func (n NormalItemQualityStrategy) ApplyChangesForOneDay(item *Item) {
	if item.quality > 0 {
		item.quality -= standardQualityChange
	}
	if (item.sellIn <= 0) && (item.quality > 0) {
		item.quality -= standardQualityChange
	}
	capToStandardMaximumQuality(item)
	applyStandardAging(item)
}

func capToStandardMaximumQuality(item *Item) {
	item.quality = int(math.Min(50, float64(item.quality)))
}

func applyStandardAging(item *Item) {
	item.sellIn -= 1
}
