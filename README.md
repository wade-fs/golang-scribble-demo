# 檔案結構
1. 採模組設計，先在本地目錄執行 go mod init db, 這樣整個套件稱之為 db，後面會看到好處
1. 將 scribble 放在 scribble/ 下，使用上因為模組設計，使用 import "db/scribble" 就可以引入
1. Examples/ 下都是些範例，使用時複製到本目錄的 main.go, 這樣就可以 go build 產生 db 執行檔
1. 目前 Examples/ 有幾個範例:
    1. simple: 最簡單的用法, 採用 interface{}, 雖然萬用，但是要對 [reflect](#1) 熟悉
    1. struct: 基本的 struct 用法，我最喜歡的使用方式
    1. array: 對陣列物件的處理
    1. readall: 每個成員分別在一個 json 檔案中，由 ReadAll() 讀取，比較麻煩
    1. write: 對資料的寫入比 Read() + ReadAll() 還簡單
1. 編譯與執行:  
  go build && db

# 參考
[1] [https://pkg.go.dev/reflect](https://pkg.go.dev/reflect)
