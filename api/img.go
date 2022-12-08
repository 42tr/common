package api

import (
	"common/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SaveImg(c *gin.Context) {
	var req SaveImgReq
	var rsp SaveImgRsp

	if err := c.ShouldBindJSON(&req); err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	id, err := db.SaveImg(req.Data)
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp.Data.ID = id
	c.JSON(http.StatusOK, rsp)
}

func GetImg(c *gin.Context) {
	idStr := c.Param("id")
	var rsp GetImgRsp

	id, err := strconv.Atoi(idStr)
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	base64, err := db.GetImg(uint(id))
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp.Data = base64
	c.JSON(http.StatusOK, rsp)
}
