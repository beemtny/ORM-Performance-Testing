package main

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
)

const insertQuery1 = `INSERT INTO models (name,email) VALUES ($1, $2);`
const insertQuery2 = `INSERT INTO models (name,email) VALUES (?, ?);`

func BenchmarkInsertTest(b *testing.B) {

	defer Db.Close()
	dbGorm, _ := DbGorm.DB()
	defer dbGorm.Close()
	defer DbGoPg.Close()
	defer Xorm.Close()
	defer cleanup()

	// Benchmark pgx
	b.Run(fmt.Sprint("INSERT TEST (pgx)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Db.Exec(Ctx, insertQuery1, gofakeit.Name(), gofakeit.Email())
			if err != nil {
				panic(err)
			}
		}
	})

	// Benchmark gorm
	b.Run(fmt.Sprint("INSERT TEST (gorm)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tx := DbGorm.Create(&Model{Name: gofakeit.Name(), Email: gofakeit.Email()})
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
	})

	// Benchmark gorm-raw
	b.Run(fmt.Sprint("INSERT TEST (gorm raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tx := DbGorm.Exec(insertQuery2, gofakeit.Name(), gofakeit.Email())
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
	})

	// Benchmark go-pg
	b.Run(fmt.Sprint("INSERT TEST (go-pg)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DbGoPg.Model(&Model{Name: gofakeit.Name(), Email: gofakeit.Email()}).Insert()
			if err != nil {
				panic(err)
			}
		}
	})

	// Benchmark go-pg-raw
	b.Run(fmt.Sprint("INSERT TEST (go-pg raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DbGoPg.Exec(insertQuery2, gofakeit.Name(), gofakeit.Email())
			if err != nil {
				panic(err)
			}
		}
	})

	// Benchmark xorm
	b.Run(fmt.Sprint("INSERT TEST (xorm)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Xorm.Insert(&Model{Name: gofakeit.Name(), Email: gofakeit.Email()})
			if err != nil {
				panic(err)
			}
		}
	})

	// Benchmark xorm raw
	b.Run(fmt.Sprint("INSERT TEST (xorm raw)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Xorm.Exec(insertQuery2, gofakeit.Name(), gofakeit.Email())
			if err != nil {
				panic(err)
			}
		}
	})

}
