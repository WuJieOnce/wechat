package payment

type Config struct {
	/**
	 * 应用ID:
	 * 微信开放平台审核通过的应用APPID
	 * 请登录open.weixin.qq.com查看，注意与公众号的APPID不同
	 */
	Appid string
	/**
	 * 应用Secret:
	 * 微信开放平台审核通过的应用AppSecret
	 * 请登录open.weixin.qq.com查看，注意与公众号的APPID不同
	 */
	AppSecret string
	/**
	 * 商户号:
	 * 微信支付分配的商户号
	 */
	MchId string
	/**
	 * 通知地址:
	 * 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
	 * 公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http。
	 */
	NotifyURL string
}
