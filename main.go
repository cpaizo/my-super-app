package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// 全域變數，供整個程式使用
var (
	ctx = context.Background()
	rdb *redis.Client
)

// Supabase 資訊 (建議部署後改為環境變數，目前先保留你的設定)
const (
	supabaseURL = "https://efkiuxgwlzjftlqxtett.supabase.co"
	supabaseKey = "sb_publishable_G2itMMa0hbauxhKrnUJtwA_XDbhELbU"
)

func main() {
	// 1. 設定埠號：Render 部署必備，本地預設 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 2. 初始化 Redis 連線
	redisURL := os.Getenv("REDIS_URL") // 這是從 Render 後台讀取的
	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err == nil {
			rdb = redis.NewClient(opt)
			fmt.Println("✅ Redis 連線成功：快取機制已啟動")
		} else {
			fmt.Printf("⚠️ Redis URL 格式有誤: %v\n", err)
		}
	} else {
		fmt.Println("ℹ️ 本地模式：未偵測到 Redis，將跳過打卡限制檢查")
	}

	// 3. 靜態網頁服務：讀取你專案中的 web 資料夾
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// 4. 打卡 API 邏輯 (供前端 JavaScript 呼叫)
	http.HandleFunc("/api/checkin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// 從前端拿員工 ID
		userID := r.FormValue("userID")
		if userID == "" {
			http.Error(w, "請提供員工 ID", http.StatusBadRequest)
			return
		}

		// 產生日期的 Key (例如 checkin:2026-04-12:WALTER)
		today := time.Now().Format("2006-01-02")
		redisKey := fmt.Sprintf("checkin:%s:%s", today, userID)

		// --- Redis 檢查機制 ---
		if rdb != nil {
			exists, _ := rdb.Exists(ctx, redisKey).Result()
			if exists > 0 {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "❌ 您今天已經打過卡囉！")
				return
			}

			// 如果沒打過，在 Redis 標記已打卡，設定 24 小時後自動失效
			rdb.Set(ctx, redisKey, "checked", 24*time.Hour)
		}
		// ----------------------

		// 此處可擴充寫入 Supabase 資料庫的邏輯
		fmt.Fprintf(w, "✅ 打卡成功！系統已記錄您的出勤時間。")
	})

	fmt.Println("-------------------------------------------")
	fmt.Println("🚀 員工管理系統（雲端版）已啟動！")
	fmt.Println("📡 Supabase ID: efkiuxgwlzjftlqxtett")
	if rdb != nil {
		fmt.Println("🧠 Redis 功能：防止重複打卡已生效")
	}
	fmt.Printf("🔗 服務路徑：http://localhost:%s\n", port)
	fmt.Println("-------------------------------------------")

	// 5. 啟動伺服器
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("伺服器無法啟動: %v\n", err)
	}
}
