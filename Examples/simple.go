package main

import (
	"fmt"
	"db/scribble"
)

func main() {
	db,err := scribble.New("Db")	// "db" 是指資料庫路徑，放本目錄下 Db/
	if err != nil {	// go 的習慣是要立刻先處理錯誤
		fmt.Printf(`scribble.New("db"): %s\n`, err.Error())
		return
	}

	var family map[string]interface{}
	// err = db.Read("", "family", &family) // 讀 Db/family.json
	err = db.Read("family", "family", &family) // 讀 Db/family/family.json
	fmt.Printf("%+v", family)
}
