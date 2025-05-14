package models

import "gorm.io/gorm"

type Ders struct {
	gorm.Model
	Ad       string `json:"ad"`
	Ogretmen string `json:"ogretmen"`
	BasSaat  string `json:"bas_saat"` // Örn. "09:00"
	BitSaat  string `json:"bit_saat"` // Örn. "10:30"
	Gun      string `json:"gun"`      // Örn. "Pazartesi"
}
