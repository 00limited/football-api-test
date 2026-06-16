package services

import "gorm.io/gorm"

func gormNotFound() error {
	return gorm.ErrRecordNotFound
}
