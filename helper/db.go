package helper

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB, err error) {
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
