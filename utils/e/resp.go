package e

const (
	//成功
	RESP_SUCC = 200

	//客户端错误
	RESP_SIGNATURE_ERROR = 403
	RESP_NOT_FOUND       = 404
	RESP_METHOD_ERR      = 405
	RESP_PARAMS_ERR      = 406

	//服务端错误
	RESP_SYSTEM_ERR = 500

	//异常
	RESP_UNKNOW_ERR  = 600
	RESP_NETWORK_ERR = 601
)

var respMap = map[int]string{
	RESP_SUCC:            "success",
	RESP_SIGNATURE_ERROR: "签名失败",
	RESP_NOT_FOUND:       "数据不存在",
	RESP_METHOD_ERR:      "请求方式错误",
	RESP_PARAMS_ERR:      "参数错误",
	RESP_SYSTEM_ERR:      "系统繁忙",
	RESP_UNKNOW_ERR:      "未知错误",
	RESP_NETWORK_ERR:     "网络异常",
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResp() *response {
	return &response{
		Code: RESP_SUCC,
	}
}

func getMessage(code int) string {
	msg, exists := respMap[code]
	if exists {
		return msg
	}
	return ""
}

func (r *response) Output(code int, msg string, data interface{}) {
	if msg == "" {
		msg = getMessage(code)
	}
	r.Code = code
	r.Message = msg
	r.Data = data
}
