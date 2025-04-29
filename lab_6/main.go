package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"lab_6/config"
	"lab_6/gui"
	"lab_6/repository"
	"log"
	"os"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env variables: %s", err.Error())
	}

	pg := new(repository.PostgresRepository)
	_, err := pg.InitDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connected to the database successfully")
	}

	a := app.New()
	w := a.NewWindow("Jurassic Park Database")
	a, w = gui.GUI(pg)
	w.ShowAndRun()
}

//rows, err := pg.Db.Query("SELECT datname FROM pg_database WHERE datistemplate = false")
//if err != nil {
//	log.Fatal(err)
//}
//defer rows.Close()
//
//for rows.Next() {
//	var dbName string
//	err := rows.Scan(&dbName)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(dbName)
//}
