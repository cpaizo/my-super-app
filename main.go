package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 設定網頁檔案的路徑
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	fmt.Println("🚀 員工系統伺服器已啟動！")
	fmt.Println("請打開瀏覽器並輸入：http://localhost:8080")

	// 監聽 8080 埠號
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("啟動失敗：", err)
	}
}
