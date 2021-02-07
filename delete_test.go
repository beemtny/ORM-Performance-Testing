package main

import (
	"fmt"
	"testing"
)

const deleteQuery1 = `DELETE FROM models WHERE id = $1`
const deleteQuery2 = `DELETE FROM models WHERE id = ?`

func BenchmarkDeleteTest(b *testing.B) {
	insertStartData()

	defer Db.Close()
	dbGorm, _ := DbGorm.DB()
	defer dbGorm.Close()
	defer DbGoPg.Close()
	defer Xorm.Close()
	defer cleanup()

	b.Run(fmt.Sprint("DELETE TEST (pgx)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Db.Exec(Ctx, deleteQuery1, i+1)
			if err != nil {
				panic(err)
			}
		}
	})

	setup()
	insertStartData()
	b.Run(fmt.Sprint("DELETE TEST (gorm)"), func(b *testing.B) {
		var model Model
		for i := 0; i < b.N; i++ {
			tx := DbGorm.Unscoped().Delete(&model, "id = ?", i+1)
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
	})

	setup()
	insertStartData()
	b.Run(fmt.Sprint("DELETE TEST (gorm raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tx := DbGorm.Exec(deleteQuery2, i+1)
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
	})

	setup()
	insertStartData()
	b.Run(fmt.Sprint("DELETE TEST (go-pg)"), func(b *testing.B) {
		var model Model
		for i := 0; i < b.N; i++ {
			_, err := DbGoPg.Model(&model).Where("id = ?", i+1).Delete()
			if err != nil {
				panic(err)
			}
		}
	})

	setup()
	insertStartData()
	b.Run(fmt.Sprint("DELETE TEST (go-pg raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DbGoPg.Exec(deleteQuery2, i+1)
			if err != nil {
				panic(err)
			}
		}
	})

	setup()
	insertStartData()
	b.Run(fmt.Sprint("DELETE TEST (xorm)"), func(b *testing.B) {
		model := Model{}
		for i := 0; i < b.N; i++ {
			_, err := Xorm.Where("id = ?", i+1).Delete(&model)
			if err != nil {
				panic(err)
			}
		}
	})

	setup()
	insertStartData()
	b.Run(fmt.Sprint("DELETE TEST (xorm raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Xorm.Exec(deleteQuery2, i+1)
			if err != nil {
				panic(err)
			}
		}
	})
}
