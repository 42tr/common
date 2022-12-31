package api

type SaveImgReq struct {
	// 1 url, 2 base64
	Type int    `json:"type"`
	Data string `json:"data"`
}

type SaveImgRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}
