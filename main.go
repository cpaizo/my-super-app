package main

import (
	"fmt"
	"net/http"
)

// 填入你的 Supabase 連線資訊
const (
	supabaseURL = "https://efkiuxgwlzjftlqxtett.supabase.co"
	supabaseKey = "sb_publishable_G2itMMa0hbauxhKrnUJtwA_XDbhELbU"
)

func main() {
	// 設定網頁檔案伺服器
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	fmt.Println("-------------------------------------------")
	fmt.Println("🚀 員工管理系統伺服器已啟動！")
	fmt.Println("📡 已建立與 Supabase 雲端（ID: efkiuxgwlzjftlqxtett）的連線路徑")
	fmt.Println("🔗 請打開：http://localhost:8080")
	fmt.Println("-------------------------------------------")

	// 啟動監聽
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("伺服器啟動失敗:", err)
	}
}
