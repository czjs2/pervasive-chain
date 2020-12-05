package statecode

type ErrorInfo map[string]interface{}

var errorMap = map[int]string{
	Success:           "成功",
	UnknownError:      "未知错误",
	Fail:              "失败",
	TokenExpiredERROR: "token失效",
	ParamError:        "参数不对",
}

func StateInfo(code int) interface{} {
	str, ok := errorMap[code]
	if ok {
		return ErrorInfo{"code": code, "message": str}
	}
	return ErrorInfo{"code":UnknownError,"message":errorMap[UnknownError]}
}

func CodeInfo(code int) string {
	str, ok := errorMap[code]
	if ok {
		return str
	}
	return errorMap[UnknownError]
}
