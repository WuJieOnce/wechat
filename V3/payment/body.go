package payment

import "time"

type Body struct {
	/**
	 * 【公众账号ID】
	 * APPID是商户公众号唯一标识，在公众平台申请
	 */
	Appid string `json:"appid"`
	/**
	 * 【商户号】
	 * 是由微信支付系统生成并分配给每个商户的唯一标识符
	 */
	MchId string `json:"mch_id"`
	/**
	 * 【商品描述】
	 * 商品信息描述，用户微信账单的商品字段中可见(可参考JSAPI支付示例说明-账单示意图)，商户需传递能真实代表商品信息的描述，不能超过127个字符。
	 */
	Description string `json:"description"`
	/**
	 * 【商户订单号】
	 * 商户系统内部订单号，要求6-32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。
	 */
	OutTradeNo string `json:"out_trade_no"`
	/**
	 * 【支付结束时间】
	 * 1、定义：支付结束时间是指用户能够完成该笔订单支付的最后时限，并非订单关闭的时间。
	 * 超过此时间后，用户将无法对该笔订单进行支付。如需关闭订单，请调用关闭订单API接口。
	 * 2、格式要求：支付结束时间需遵循rfc3339标准格式：yyyy-MM-DDTHH:mm:ss+TIMEZONE。yyyy-MM-DD 表示年月日；
	 * T 字符用于分隔日期和时间部分；HH:mm:ss 表示具体的时分秒；TIMEZONE 表示时区（例如，+08:00 对应东八区时间，即北京时间）。
	 * 示例：2015-05-20T13:29:35+08:00 表示北京时间2015年5月20日13点29分35秒。
	 * 3、注意事项：
	 * time_expire 参数仅在用户首次下单时可设置，且不允许后续修改，尝试修改将导致错误。
	 * 若用户实际进行支付的时间超过了订单设置的支付结束时间，商户需使用新的商户订单号下单，生成新的订单供用户进行支付。
	 * 若未超过支付结束时间，则可使用原参数重新请求下单接口，以获取当前订单最新的prepay_id 进行支付。
	 * 支付结束时间不能早于下单时间后1分钟，若设置的支付结束时间早于该时间，系统将自动调整为下单时间后1分钟作为支付结束时间。
	 */
	TimeExpire string `json:"time_expire,omitempty"`
	/**
	 * 【商户数据包】
	 * 商户在创建订单时可传入自定义数据包，该数据对用户不可见，用于存储订单相关的商户自定义信息，其总长度限制在128字符以内。
	 * 支付成功后查询订单API和支付成功回调通知均会将此字段返回给商户，并且该字段还会体现在交易账单。
	 */
	Attach string `json:"attach,omitempty"`
	/**
	 * 【商户回调地址】
	 * 商户接收支付成功回调通知的地址，需按照notify_url填写注意事项规范填写。
	 */
	NotifyUrl string `json:"notify_url"`
	/**
	 * 【订单优惠标记】
	 * 代金券在创建时可以配置多个订单优惠标记，标记的内容由创券商户自定义设置。
	 * 详细参考：创建代金券批次API。如果代金券有配置订单优惠标记，则必须在该参数传任意一个配置的订单优惠标记才能使用券。
	 * 如果代金券没有配置订单优惠标记，则可以不传该参数。
	 * 示例：如有两个活动，活动A设置了两个优惠标记：WXG1、WXG2；
	 * 活动B设置了两个优惠标记：WXG1、WXG3；
	 * 下单时优惠标记传WXG2，则订单参与活动A的优惠；
	 * 下单时优惠标记传WXG3，则订单参与活动B的优惠；
	 * 下单时优惠标记传共同的WXG1，则订单参与活动A、B两个活动的优惠；
	 */
	GoodsTag string `json:"goods_tag,omitempty"`
	/**
	 * 【电子发票入口开放标识】
	 * 传入true时，支付成功消息和支付详情页将出现开票入口。
	 * 需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。
	 * 详细参考：电子发票介绍
	 * true：是
	 * false：否
	 */
	SupportFapiao bool `json:"support_fapiao,omitempty"`
	// 【订单金额】订单金额信息
	Amount *Amount `json:"amount"`
	// 【支付者信息】支付者信息 JSAPI/小程序 支付必填
	Payer *Payer `json:"payer,omitempty"`
	// 【优惠功能】 优惠功能
	Detail *Detail `json:"detail,omitempty"`
	// 【场景信息】 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 【结算信息】 结算信息
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
}

// Amount 【订单金额】订单金额信息
type Amount struct {
	/**
	 * 【总金额】
	 * 订单总金额，单位为分，整型。
	 * 示例：1元应填写 100
	 */
	Total int64 `json:"total"`
	/**
	 * 【货币类型】
	 * 固定传：CNY，代表人民币。
	 */
	Currency string `json:"currency,omitempty"`
	/**
	 * 【用户支付金额】用户支付金额，整型，单位为分。
	 * 指使用优惠券的情况下，这里等于总金额-优惠券金额
	 */
	PayerTotal int64 `json:"payer_total,omitempty"`
	/**
	 * 【用户支付币种】 固定返回：CNY，代表人民币。
	 */
	PayerCurrency string `json:"payer_currency,omitempty"`
}

// Payer 【支付者信息】支付者信息
type Payer struct {
	/**
	 * 【用户标识】
	 * 用户在商户appid下的唯一标识。
	 * 下单前需获取到用户的OpenID，详见OpenID获取详见。
	 */
	Openid string `json:"openid"`
}

// GoodsDetail 【单品列表】 单品列表信息,条目个数限制：【1，6000】
type GoodsDetail struct {
	// 【商户侧商品编码】 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 【微信支付商品编码】 微信支付定义的统一商品编号（没有可不传）
	WechatPayGoodsId string `json:"wechat_pay_goods_id,omitempty"`
	// 【商品名称】 商品的实际名称
	GoodsName string `json:"goods_name,omitempty"`
	// 【商品数量】 用户购买的数量
	Quantity int64 `json:"quantity"`
	// 【商品单价】整型，单位为：分。如果商户有优惠，需传输商户优惠后的单价(例如：用户对一笔100元的订单使用了商场发的纸质优惠券100-50，则活动商品的单价应为原单价-50)
	UnitPrice int64 `json:"unit_price"`
	// 【商品备注】 商品备注。创券商户在商户平台创建单品券时，若设置了商品备注则会返回。
	GoodsRemark string `json:"goods_remark,omitempty"`
	// 【商品优惠金额】 商品优惠金额。
	DiscountAmount int64 `json:"discount_amount"`
	// 【商品编码】 商品编码。
	GoodsId string `json:"goods_id"`
}

// Detail 【优惠功能】 优惠功能
type Detail struct {
	/**
	 * 【订单原价】
	 * 1、商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。
	 * 2、当订单原价与支付金额不相等，则不享受优惠。
	 * 3、该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	 */
	CostPrice int64 `json:"cost_price,omitempty"`
	// 【商品小票ID】 商家小票ID
	InvoiceId string `json:"invoice_id,omitempty"`
	// 【单品列表】 单品列表信息,条目个数限制：【1，6000】
	GoodsDetail *GoodsDetail `json:"goods_detail,omitempty"`
}

// StoreInfo 【优惠功能】 优惠功能
type StoreInfo struct {
	// 【门店编号】商户侧门店编号，总长度不超过32字符。必填 string(32)
	Id string `json:"id"`
	// 【门店名称】 商户侧门店名称选填 string(256)
	Name string `json:"name"`
	// 【地区编码】 地区编码，详细请见省市区编号对照表。选填 string(32)
	AreaCode string `json:"area_code"`
	// 【详细地址】 详细的商户门店地址选填 string(512)
	Address string `json:"address"`
}

// SceneInfo 【场景信息】 场景信息
type SceneInfo struct {
	// 【用户终端IP】 用户的客户端IP，支持IPv4和IPv6两种格式的IP地址。
	PayerClientIp string `json:"payer_client_ip"`
	// 【商户端设备号】 商户端设备号（门店号或收银设备ID）。
	DeviceId string `json:"device_id,omitempty"`
	// 【商户门店信息】 商户门店信息
	StoreInfo StoreInfo `json:"store_info,omitempty"`
}

// SettleInfo 【结算信息】 结算信息
type SettleInfo struct {
	/**
	 * 【分账标识】订单的分账标识在下单时设置，传入true表示在订单支付成功后可进行分账操作。以下是详细说明：
	 * 需要分账（传入true）：
	 * 订单收款成功后，资金将被冻结并转入基本账户的不可用余额。商户可通过请求分账API，将收款资金分配给其他商户或用户。
	 * 完成分账操作后，可通过接口解冻剩余资金，或在支付成功30天后自动解冻。
	 * 不需要分账（传入false或不传，默认为false）：
	 * 订单收款成功后，资金不会被冻结，而是直接转入基本账户的可用余额。
	 */
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

type Notice struct {
	Id           string    `json:"id"`
	CreateTime   time.Time `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     *Resource `json:"resource"`
}

type Resource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type PromotionDetail struct {
	Amount              int            `json:"amount"`
	WechatpayContribute int            `json:"wechatpay_contribute"`
	CouponId            string         `json:"coupon_id"`
	Scope               string         `json:"scope"`
	MerchantContribute  int            `json:"merchant_contribute"`
	Name                string         `json:"name"`
	OtherContribute     int            `json:"other_contribute"`
	Currency            string         `json:"currency"`
	StockId             string         `json:"stock_id"`
	GoodsDetail         []*GoodsDetail `json:"goods_detail"`
}

type Notification struct {
	TransactionId   string             `json:"transaction_id"`
	Amount          *Amount            `json:"amount"`
	Mchid           string             `json:"mchid"`
	TradeState      string             `json:"trade_state"`
	BankType        string             `json:"bank_type"`
	PromotionDetail []*PromotionDetail `json:"promotion_detail"`
	SuccessTime     time.Time          `json:"success_time"`
	Payer           *Payer             `json:"payer"`
	OutTradeNo      string             `json:"out_trade_no"`
	Appid           string             `json:"appid"`
	TradeStateDesc  string             `json:"trade_state_desc"`
	TradeType       string             `json:"trade_type"`
	Attach          string             `json:"attach"`
	SceneInfo       *SceneInfo         `json:"scene_info"`
}
