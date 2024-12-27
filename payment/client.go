package payment

type Client struct {
	Config
}

// NewClient 创建客户端
func NewClient(config *Config) *Client {
	return new(Client)
}

// SetConfig 设置配置信息
func (c *Client) SetConfig(config *Config) *Client {
	c.Config = *config
	return c
}

// SetBody 设置商品信息
func (c *Client) SetBody() *Client {
	return c
}

// TradeAppPay 发起APP支付
func (c *Client) TradeAppPay() *Client {
	return c
}

// TradeJsApiPay 发起JSAPI支付
func (c *Client) TradeJsApiPay() *Client {
	return c
}

// TradeH5Pay 发起H5支付
func (c *Client) TradeH5Pay() *Client {
	return c
}

// TradeNativePay 发起Native支付
func (c *Client) TradeNativePay() *Client {
	return c
}

// TradeMiniProgramPay 发起小程序支付
func (c *Client) TradeMiniProgramPay() *Client {
	return c
}

// TradePaymentCodePay 发起付款码支付
func (c *Client) TradePaymentCodePay() *Client {
	return c
}

// TradeFacePaymentPay 发起刷脸支付
func (c *Client) TradeFacePaymentPay() *Client {
	return c
}
