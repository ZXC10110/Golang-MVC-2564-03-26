package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zxc10110/mvc_08/Models"
)

func UploadPicture(c *gin.Context) {
	var pic Models.Picture
	c.BindJSON(&pic)
	//reach 1000-100000
	if pic.Reach < 1000 || pic.Reach > 100000 {
		c.JSON(http.StatusBadRequest, "ค่า Reach ต้องมีค่าตั้งแต่ 1000-100000 เท่านั้น!")
		return
	} else {
		//(like + share + comment) condition
		countQty := pic.LikeQty + pic.ShareQty + pic.ShareQty
		if countQty >= pic.Reach { // >= reach
			c.JSON(http.StatusBadRequest, "ยอด (Like + Share + Comment) ต้องไม่มากกว่าค่า Reach")
			return
		} else if countQty > 100000 || countQty < 1 {
			c.JSON(http.StatusBadRequest, "ยอด (Like + Share + Comment) ต้องอยู่ระหว่าง 1-100000")
			return
		} else {
			//create picture
			er := Models.CreatePicture(&pic)
			if er != nil {
				c.JSON(http.StatusBadRequest, "ไม่พบข้อมูลหรือหมายเลขไอดีของภาพมีการซ้ำกันไม่ได้")
				return
			}
			c.JSON(http.StatusOK, pic)
		}
	}

}

func ViralPicture(c *gin.Context) {

	//calculate score
	viral, err := Models.ViralScore()
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูลภาพ")
		return
	}

	//calculate score
	notViral, errr := Models.NotViralScore()
	if errr != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูลภาพ")
		return
	}

	//map score
	resultViral := Models.ResultPicture{
		Status:  " ภาพเป็นไวรัล",
		Picture: viral,
	}

	//map score
	resultNoteViral := Models.ResultPicture{
		Status:  " ภาพไม่เป็นไวรัล",
		Picture: notViral,
	}

	//map output
	result := Models.Result{
		Viral:    resultViral,
		NotViral: resultNoteViral,
	}
	//print output
	c.JSON(http.StatusOK, result)
}
