package inputmodels

type CreateArticle struct {
	No        string `json:"no"`         // 文章代號(unique)
	Name      string `json:"name"`       // 文章名稱
	QuickCode string `json:"quick_code"` // 簡碼
}

type UpdateArticle struct {
	ID        uint   `json:"id" validate:"required"` // 流水號
	No        string `json:"no"`                     // 文章代號(unique)
	Name      string `json:"name"`                   // 文章名稱
	QuickCode string `json:"quick_code"`             // 簡碼
}
