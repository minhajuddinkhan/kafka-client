package models

import "github.com/jinzhu/gorm"

//TLS TLS
type TLS struct {
	gorm.Model
	CertPath string `gorm:"type:varchar(100)" json:"certPath,omitempty"`
	KeyPath  string `gorm:"type:varchar(100)" json:"keyPath,omitempty"`
}
