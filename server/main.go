package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"traQer/handler"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/robfig/cron/v3"
	"github.com/traPtitech/go-traq"
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("MARIADB_USER") == "" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(os.Getenv("MARIADB_USER"))
	fmt.Println("aa")
	conf := mysql.Config{
		User:                 os.Getenv("MARIADB_USER"),
		Passwd:               os.Getenv("MARIADB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MARIADB_HOSTNAME") + ":" + os.Getenv("MARIADB_PORT"),
		DBName:               os.Getenv("MARIADB_DATABASE"),
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conntected")

	db := _db
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, os.Getenv("TRAQ_TOKEN"))
	h := handler.NewHandler(db, handler.NewBot(), client, auth, time.Now().UTC(), "")

	c := cron.New() //定時実行用
	e := echo.New()
	go h.BotHandler()

	//再起動でデータ取得
	//ハンドラに情報を持たせる
	h.MessageCountsBind()

	//SELECT EXISTS (SELECT * FROM `messagecounts`)
	if false {
		h.GetUserPostCount()
	}
	//cron動作確認
	c.AddFunc("* * * * *", func() { log.Println("cron is running") })
	//1日毎に全ユーザ読み込みを行う(データの補正,午前4時に実施 ただしNSはUTC)
	c.AddFunc("0 19 * * *", h.GetUserPostCount)
	//5分ごとに差分読み取りを行う
	c.AddFunc("0-59/5 * * * *", h.SearchMessagesRunner)

	c.Start()

	time.Sleep(time.Second * 2) //cronスタート用

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error { return c.String(http.StatusOK, "pong") })
	e.GET("/alter", func(c echo.Context) error { return c.String(http.StatusOK, "pong") })
	e.GET("/messages", h.GetMessageCounts)

	e.Logger.Fatal(e.Start(":8080"))

}
