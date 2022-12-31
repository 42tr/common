package api

import (
	"common/db"
	"encoding/base64"
	"io/ioutil"
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
	var url, base string
	if req.Type == 1 {
		url = req.Data
		res, err := http.Get(url)
		if err != nil {
			rsp.Code = -1
			rsp.Msg = err.Error()
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			rsp.Code = -1
			rsp.Msg = err.Error()
			c.JSON(http.StatusBadRequest, rsp)
			return
		}

		base = base64.StdEncoding.EncodeToString(data)
	} else {
		base = req.Data
	}

	id, err := db.SaveImg(url, base)
	if err != nil {
		rsp.Code = -1
		rsp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp.Data = id
	c.JSON(http.StatusOK, rsp)
}

func GetImg(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	img, err := db.GetImg(uint(id))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if img.Url != "" {
		res, err := http.Get(img.Url)
		if err != nil {
			goto local
		}
		defer res.Body.Close()
		headers := res.Header["Content-Type"]
		if len(headers) == 0 || headers[0] != "image/jpeg" {
			goto local
		}
		c.Redirect(http.StatusMovedPermanently, img.Url)
		return
	}
local:
	imageBuffer, err := base64.StdEncoding.DecodeString(img.Base64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.Writer.WriteString(string(imageBuffer))
}
