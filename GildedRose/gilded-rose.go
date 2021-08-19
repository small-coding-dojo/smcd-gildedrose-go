package main

import "math"

type Item struct {
	name            string
	sellIn, quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		standardQualityChange := 1
		aging := 0

		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].quality > 0 {
				if items[i].name == "Conjured Mana Cake" {
					items[i].quality -= 2
				} else if items[i].name != "Sulfuras, Hand of Ragnaros" {
					items[i].quality -= standardQualityChange
				}
			}
		} else {
			items[i].quality += standardQualityChange
			if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
				if items[i].sellIn < 11 {
					items[i].quality += standardQualityChange
				}
				if items[i].sellIn < 6 {
					items[i].quality += standardQualityChange
				}
			}
		}

		items[i].quality += aging

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn -= 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].quality > 0 {
						if items[i].name == "Conjured Mana Cake" {
							items[i].quality -= 2
						} else if items[i].name != "Sulfuras, Hand of Ragnaros" {
							items[i].quality -= standardQualityChange
						}
					}
				} else {
					items[i].quality = 0
				}
			} else {
				items[i].quality += standardQualityChange
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].quality = capToMaxStandardQuality(items[i].quality)
		}
	}

}

func capToMaxStandardQuality(quality int) int {
	return int(math.Min(50, float64(quality)))
}
