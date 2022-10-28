package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/zxc10110/mvc_08/Config"
	Routes "github.com/zxc10110/mvc_08/Controllers"
	"github.com/zxc10110/mvc_08/Models"
)

var err error

func main() {
	fmt.Println("Run Swagger : http://localhost:8080/swagger/index.html")
	fmt.Println()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Picture{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
