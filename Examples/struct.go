package main

import (
	"fmt"
	"db/scribble"
)

type Child struct {
	Name	string
}

type Family struct {
    Name		string
    Spouse		string
    Children	[]Child
}

func main() {
	db,err := scribble.New("Db")	// "db" 是指資料庫路徑，放本目錄下 Db/
	if err != nil {	// go 的習慣是要立刻先處理錯誤
		fmt.Printf(`scribble.New("db"): %s\n`, err.Error())
		return
	}

	var family Family
	// err = db.Read("", "family", &family) // 讀 Db/family.json
	err = db.Read("family", "family", &family) // 讀 Db/family/family.json
	fmt.Printf("I am %s. My spouse is %s.\nAnd I have children:\n", 
		family.Name, family.Spouse)
	for _,child := range family.Children {
		fmt.Printf("\t%s\n", child.Name)
	}
}
