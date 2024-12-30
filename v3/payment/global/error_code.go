package global

type ErrorBase struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Solution string `json:"solution"`
}

func (e *ErrorBase) Error() string {
	return e.Message
}

var ErrorMap = map[string]*ErrorBase{
	"APPID_MCHID_NOT_MATCH": NewErrorBase(400, "AppID和mch_id不匹配", "请确认AppID和mch_id是否匹配，查询指引参考：查询商户号绑定的APPID"),
	"INVALID_REQUEST":       NewErrorBase(400, "无效请求", "请根据接口返回的详细信息检查"),
	"MCH_NOT_EXISTS":        NewErrorBase(400, "商户号不存在", "请检查商户号是否正确，商户号获取方式请参考普通商户模式开发必要参数说明"),
	"PARAM_ERROR":           NewErrorBase(400, "参数错误", "请根据接口返回的错误描述检查参数，参数需按API文档字段填写说明填写"),
	"SIGN_ERROR":            NewErrorBase(401, "签名错误", "请检查签名参数和方法是否都符合签名算法要求，参考：如何生成签名"),
	"NO_AUTH":               NewErrorBase(403, "商户无权限", "请商户前往商户平台申请此接口相关权限，参考：权限申请"),
	"OUT_TRADE_NO_USED":     NewErrorBase(403, "商户订单号重复", "请核实商户订单号是否重复提交"),
	"FREQUENCY_LIMITED":     NewErrorBase(429, "频率超限", "请求频率超限，请降低请求接口频率"),
	"SYSTEM_ERROR":          NewErrorBase(500, "系统错误", "系统异常，请用相同参数重新调用"),
}

func NewError(err string) *ErrorBase {
	if e, ok := ErrorMap[err]; ok {
		return e
	}
	return NewErrorBase(400, err, "")
}

func NewErrorBase(code int, message string, solution string) *ErrorBase {
	return &ErrorBase{Code: code, Message: message, Solution: solution}
}
