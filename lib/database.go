package lib

import (
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}
