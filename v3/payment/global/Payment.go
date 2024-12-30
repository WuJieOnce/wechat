package global

import (
	"crypto/rsa"
	"wechat/V3/payment"
)

type Payment interface {
	// LoadPrivateKey 加载私钥
	LoadPrivateKey(path string) error
	ParseRSAPrivateKey(content string) (*rsa.PrivateKey, error)
	// SetConfig 设置配置信息
	SetConfig(conf *Config) Payment
	// SetBody 设置订单信息
	SetBody(order payment.Transaction) Payment
	// InitiateAnOrder 发起下单
	InitiateAnOrder()
	// QueryOrderByWechatPaymentOrder 微信支付订单号查询订单
	QueryOrderByWechatPaymentOrder()
	// QueryOrderByMerchantOrder 商户订单号查询订单
	QueryOrderByMerchantOrder()
	// CloseOrder 关闭订单
	CloseOrder()
	// Notification 支付成功回调通知
	Notification()
	// RefundOrder 退款申请
	RefundOrder()
	// QueryASingleRefund 查询单笔退款（通过商户退款单号）
	QueryASingleRefund()
	// AbnormalRefund 发起异常退款
	AbnormalRefund()
	// NotificationByRefund 退款结果回调通知
	NotificationByRefund()
	// TransactionStatement 申请交易账单
	TransactionStatement()
	// FundsBill 申请资金账单
	FundsBill()
	// DownloadBill 下载账单
	DownloadBill()
}
