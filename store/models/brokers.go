package models

import "github.com/jinzhu/gorm"

//Broker Broker
type Broker struct {
	gorm.Model
	URL string `gorm:"type:varchar(100)" json:"name,omitempty"`
}
