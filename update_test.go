package main

import (
	"fmt"
	"testing"
)

const updateQuery1 = `UPDATE models SET name = $1 WHERE id = $2`
const updateQuery2 = `UPDATE models SET name = ? WHERE id = ?`

func BenchmarkUpdateTest(b *testing.B) {
	insertStartData()

	defer Db.Close()
	dbGorm, _ := DbGorm.DB()
	defer dbGorm.Close()
	defer DbGoPg.Close()
	defer Xorm.Close()
	defer cleanup()

	b.Run(fmt.Sprint("UPDATE TEST (pgx)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Db.Exec(Ctx, updateQuery1, "aaaa", i+1)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run(fmt.Sprint("UPDATE TEST (gorm)"), func(b *testing.B) {
		var model Model
		for i := 0; i < b.N; i++ {
			tx := DbGorm.Model(&model).Where("id = ?", i+1).Update("name", "bbbb")
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
	})

	b.Run(fmt.Sprint("UPDATE TEST (gorm raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tx := DbGorm.Exec(updateQuery2, "cccc", i+1)
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
	})

	b.Run(fmt.Sprint("UPDATE TEST (go-pg)"), func(b *testing.B) {
		model := &Model{Name: "dddd"}
		for i := 0; i < b.N; i++ {
			_, err := DbGoPg.Model(model).Column("name").Where("id = ?", i+1).Update()
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run(fmt.Sprint("UPDATE TEST (go-pg raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DbGoPg.Exec(updateQuery2, "eeee", i+1)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run(fmt.Sprint("UPDATE TEST (xorm)"), func(b *testing.B) {
		model := &Model{Name: "ffff"}
		for i := 0; i < b.N; i++ {
			_, err := Xorm.Where("id = ?", i+1).Update(model)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run(fmt.Sprint("UPDATE TEST (xorm raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Xorm.Exec(updateQuery2, "gggg", i+1)
			if err != nil {
				panic(err)
			}
		}
	})
}
