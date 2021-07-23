package main

import (
	"fmt"
	"encoding/json"

	"db/scribble"
)

type Child struct {
	Name	string
}

type Member struct {
    Name		string
	Birth		string
    Service		string
    Children	[]Child
}

func describe(m Member) string {
	return fmt.Sprintf("%s, 生日是%s, 服務於%s", m.Name, m.Birth, m.Service)
}

func findChild(name string, members []Member) Member {
	for _,m := range members {
		if m.Name == name {
			return m
		}
	}
	return Member{}
}

func main() {
	db,err := scribble.New("Db")	// "db" 是指資料庫路徑，放本目錄下 Db/
	if err != nil {	// go 的習慣是要立刻先處理錯誤
		fmt.Printf(`scribble.New("db"): %s\n`, err.Error())
		return
	}

	records,err := db.ReadAll("members") // 讀 Db/members/*.json
	if err != nil {
		fmt.Printf("Cannot read all from Db/members/*\n")
		return
	}
	members := []Member{}
	for _,r := range records {
		m := Member{}
		if err := json.Unmarshal([]byte(r), &m); err != nil {
			fmt.Printf("%s: %s\n", r, err.Error())
			continue
		}
		members = append(members, m)
	}
	for _,m := range members {
		fmt.Println(describe(m))
		if len(m.Children) > 0 {
			fmt.Printf("有 %d 位小孩:\n", len(m.Children))
			for _,child := range m.Children {
				fmt.Println("\t", describe(findChild(child.Name, members)))
			}
		}
	}

	s,_ := json.MarshalIndent(members, "", "\t")
	fmt.Printf("%s\n", s)
}
