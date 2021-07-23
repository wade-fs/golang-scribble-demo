package main

import (
	"fmt"

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

	var members []Member
	err = db.Read("family", "members", &members) // 讀 Db/family/members.json
	for _,m := range members {
		fmt.Println(describe(m))
		if len(m.Children) > 0 {
			fmt.Printf("有 %d 位小孩:\n", len(m.Children))
			for _,child := range m.Children {
				fmt.Println("\t", describe(findChild(child.Name, members)))
			}
		}
	}
	m := Member{
		Name: "私生子",
		Birth: "01/02/03",
		Service: "無法告知",
	} // 缺 Children
	members = append(members, m)

	// s,_ := json.MarshalIndent(members, "", "\t")
	// fmt.Printf("%s\n", s)
	db.Write("family", "illegitimate", members)
}
