package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

type Model struct {
	ID    int `xorm:"'id' bigserial notnull pk"`
	Name  string
	Email string
}

func (Model) TableName() string {
	return "models"
}

var (
	Db        *pgxpool.Pool
	DbGorm    *gorm.DB
	DbGoPg    *pg.DB
	Xorm      *xorm.Engine
	Ctx       = context.Background()
	Entries   = []interface{}{}
	Constr    = "postgres://postgres:postgres@localhost:5434/pg_performance_test?sslmode=disable"
	Condition = []string{"pgx", "gorm", "gorm-raw", "go-pg", "go-pg-raw", "xorm", "xorm-raw"}
)

func main() {

	r := gin.Default()

	r.GET("/pgx", func(c *gin.Context) { Insert(Condition[0]); c.JSON(http.StatusOK, gin.H{"data": "success"}) })
	r.GET("/gorm", func(c *gin.Context) { Insert(Condition[1]); c.JSON(http.StatusOK, gin.H{"data": "success"}) })
	r.GET("/gorm-raw", func(c *gin.Context) { Insert(Condition[2]); c.JSON(http.StatusOK, gin.H{"data": "success"}) })
	r.GET("/go-pg", func(c *gin.Context) { Insert(Condition[3]); c.JSON(http.StatusOK, gin.H{"data": "success"}) })
	r.GET("/go-pg-raw", func(c *gin.Context) { Insert(Condition[4]); c.JSON(http.StatusOK, gin.H{"data": "success"}) })
	r.GET("/xorm", func(c *gin.Context) { Insert(Condition[5]); c.JSON(http.StatusOK, gin.H{"data": "success"}) })
	r.GET("/xorm-raw", func(c *gin.Context) { Insert(Condition[6]); c.JSON(http.StatusOK, gin.H{"data": "success"}) })

	r.Run()

}

func init() {
	newDBPgx()
	newDBGorm()
	newDBGoPG()
	newDBXorm()

	setup()
}

func newDBPgx() {
	var err error
	Db, err = pgxpool.Connect(Ctx, Constr)
	if err != nil {
		panic("failed to connect database")
	}
}

func newDBGorm() {
	var err error
	DbGorm, err = gorm.Open(postgres.Open(Constr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func newDBGoPG() {
	DbGoPg = pg.Connect(&pg.Options{
		Addr:      ":5434",
		User:      "postgres",
		Password:  "postgres",
		Database:  "pg_performance_test",
		TLSConfig: nil,
	})

}

func newDBXorm() {
	var err error
	Xorm, err = xorm.NewEngine("postgres", Constr)
	if err != nil {
		panic("failed to connect database")
	}
}

func setup() {
	sqldrop := `DROP TABLE IF EXISTS models`
	_, err := Db.Exec(Ctx, sqldrop)
	if err != nil {
		panic(err)
	}

	sqlCreateTable := `CREATE TABLE models (id SERIAL PRIMARY KEY, name VARCHAR(50), email VARCHAR(150));`

	_, err = Db.Exec(Ctx, sqlCreateTable)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create models hypertable: %v\n", err)
		panic(err)
	}
}

func insertStartData() {
	// Add 10,000 fake entries
	for i := 0; i < 2000; i++ {
		testInsert := `INSERT INTO models (name,email) VALUES ($1, $2);`

		_, err := Db.Exec(context.Background(), testInsert, gofakeit.Name(), gofakeit.Email())
		if err != nil {
			panic(err)
		}
	}
}

func cleanup() {
	_, err := Db.Exec(Ctx, `DROP TABLE models`)
	if err != nil {
		panic(err)
	}
}

func Insert(condition string) {
	const insertQuery1 = `INSERT INTO models (name,email) VALUES ($1, $2);`
	const insertQuery2 = `INSERT INTO models (name,email) VALUES (?, ?);`
	switch condition {
	case "pgx":
		_, err := Db.Exec(Ctx, insertQuery1, gofakeit.Name(), gofakeit.Email())
		if err != nil {
			panic(err)
		}
	case "gorm":
		tx := DbGorm.Create(&Model{Name: gofakeit.Name(), Email: gofakeit.Email()})
		if tx.Error != nil {
			panic(tx.Error)
		}
	case "gorm-raw":
		tx := DbGorm.Exec(insertQuery2, gofakeit.Name(), gofakeit.Email())
		if tx.Error != nil {
			panic(tx.Error)
		}
	case "go-pg":
		_, err := DbGoPg.Model(&Model{Name: gofakeit.Name(), Email: gofakeit.Email()}).Insert()
		if err != nil {
			panic(err)
		}
	case "go-pg-raw":
		_, err := DbGoPg.Exec(insertQuery2, gofakeit.Name(), gofakeit.Email())
		if err != nil {
			panic(err)
		}
	case "xorm":
		_, err := Xorm.Insert(&Model{Name: gofakeit.Name(), Email: gofakeit.Email()})
		if err != nil {
			panic(err)
		}
	case "xorm-raw":
		_, err := Xorm.Exec(insertQuery2, gofakeit.Name(), gofakeit.Email())
		if err != nil {
			panic(err)
		}
	}

}
