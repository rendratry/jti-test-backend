package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	} else {
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}
}
