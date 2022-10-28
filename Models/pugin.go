package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zxc10110/mvc_08/Config"
)

//create Picture ... Insert New Picture
func CreatePicture(pic *Picture) (err error) {
	if err = Config.DB.Create(pic).Error; err != nil {
		return err
	}
	return
}

//viralScore ... calculate viral score
func ViralScore() (vir []ViralPicture, err error) {
	if err = Config.DB.Select("((like_qty + share_qty + comment_qty) / reach)as score, pic_no, reach,like_qty, share_qty, comment_qty").
		Table("Test.picture").
		Having("score > 0.1").
		//Group("pic_no").
		Find(&vir).Error; err != nil {
		return
	}
	return
}

//notViralScore ... calculate not viral score
func NotViralScore() (vir []ViralPicture, err error) {
	if err = Config.DB.Select("((like_qty + share_qty + comment_qty) / reach)as score, pic_no, reach,like_qty, share_qty, comment_qty").
		Table("Test.picture").
		Having("score < 0.1").
		//Group("pic_no").
		Find(&vir).Error; err != nil {
		return
	}
	return
}
