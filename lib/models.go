package lib

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ModelMarketOrder struct {
	AlbionID         uint   `gorm:"not null;unique_index:idx_id_location"`
	ItemID           string `gorm:"index"`
	QualityLevel     int8   `gorm:"size:1"`
	EnchantmentLevel int8   `gorm:"size:1"`
	Price            int64    `gorm:"index"`
	InitialAmount    int
	Amount           int
	AuctionType      string `gorm:"index"`
	Expires          time.Time
	Location         Location `gorm:"not null;unique_index:idx_id_location;index"`
	gorm.Model
}

func (m ModelMarketOrder) TableName() string {
	return "market_orders"
}

func NewModelMarketOrder() ModelMarketOrder {
	return ModelMarketOrder{}
}

type ModelMarketStats struct {
	ID        int    `gorm:"primary_key"`
	ItemID    string `gorm:"index"`
	Location  Location
	PriceMin  int
	PriceMax  int
	PriceAvg  float64
	Timestamp *time.Time
}

func (m ModelMarketStats) TableName() string {
	return "market_stats"
}

type ModelGoldprices struct {
	gorm.Model
	Timestamp time.Time `gorm:"index"`
	Price     int
}

func (m ModelGoldprices) TableName() string {
	return "gold_prices"
}
