package migration0

import (
	"time"

	"gorm.io/gorm"
)

type article struct {
	ID        uint           `json:"id" mapstructure:"id" gorm:"primaryKey"`                                       // 流水號
	CreatedAt time.Time      `json:"created_at" mapstructure:"created_at" gorm:"Column:created_at"`                // 建檔日期
	UpdatedAt time.Time      `json:"updated_at" mapstructure:"updated_at" gorm:"Column:updated_at"`                // 修改日期
	DeletedAt gorm.DeletedAt `json:"-" mapstructure:"deleted_at" gorm:"Column:deleted_at;index"`                   // 刪除日期
	No        string         `json:"no" mapstructure:"no" gorm:"Column:no;index:,unique,where:deleted_at IS NULL"` // 文章代號(unique)
	Name      string         `json:"name" mapstructure:"name" gorm:"Column:name"`                                  // 文章名稱
	QuickCode string         `json:"quick_code" mapstructure:"quick_code" gorm:"Column:quick_code"`                // 簡碼
}
