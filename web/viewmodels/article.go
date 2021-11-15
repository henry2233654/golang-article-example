package viewmodels

import (
	"time"
)

type Article struct {
	ID        uint      `json:"id"`         // 流水號
	CreatedAt time.Time `json:"created_at"` // 建檔日期
	UpdatedAt time.Time `json:"updated_at"` // 修改日期
	No        string    `json:"no"`         // 文章代號(unique)
	Name      string    `json:"name"`       // 文章名稱
	QuickCode string    `json:"quick_code"` // 簡碼
}
