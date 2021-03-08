package error2

const (
	// Unknown 未知错误
	Unknown = -1
	// Success 成功，没有错误且成功返回
	Success = 0
)

var codeTable = map[int]string{
	Unknown: "未知错误",
	Success: "成功",
}

func translation(code int) string {
	if text, ok := codeTable[code]; ok {
		return text
	}
	return "未知code"
}
