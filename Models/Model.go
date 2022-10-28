package Models

func (b *Picture) TableName() string {
	return "picture"
}

type Picture struct {
	PicNo      int `json:"pic_no"`
	Reach      int `json:"reach"`
	LikeQty    int `json:"like_qty"`
	ShareQty   int `json:"share_qty"`
	CommentQty int `json:"comment_qty"`
}

type ViralPicture struct {
	Score      float64 `json:"score"`
	PicNo      int     `json:"pic_no"`
	Reach      int     `json:"reach"`
	LikeQty    int     `json:"like_qty"`
	ShareQty   int     `json:"share_qty"`
	CommentQty int     `json:"comment_qty"`
}

type ResultPicture struct {
	Status  string `jspn:"status"`
	Picture []ViralPicture
}

type Result struct {
	Viral    ResultPicture
	NotViral ResultPicture
}
