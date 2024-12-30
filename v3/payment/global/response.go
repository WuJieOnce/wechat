package global

type Prepaid struct {
	/**
	 *【预支付交易会话标识】预支付交易会话标识，
	 * JSAPI调起支付时需要使用的参数，
	 * 有效期为2小时，
	 * 失效后需要重新请求该接口以获取新的prepay_id。
	 */
	PrepayId string `json:"prepay_id"`
}
