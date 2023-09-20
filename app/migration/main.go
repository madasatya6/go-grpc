package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func init() {
	godotenv.Load()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide action!")
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "migration",
	}

	switch os.Args[1] {
	case "up":
		m, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Applied %d migrations!\n", m)
	case "down":
		m, err := migrate.ExecMax(db, "postgres", migrations, migrate.Down, 1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Applied %d migrations!\n", m)
	case "status":
		m, err := migrate.GetMigrationRecords(db, "postgres")
		if err != nil {
			panic(err)
		}
		for _, record := range m {
			fmt.Printf("%s - %s\n", record.Id, record.AppliedAt)
		}
	}
}
