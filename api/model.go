package api

type SaveImgReq struct {
	Data string `json:"data"`
}

type IDData struct {
	ID uint `json:"id"`
}

type SaveImgRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data IDData `json:"data"`
}

type GetImgRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
