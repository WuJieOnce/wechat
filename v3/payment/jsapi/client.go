package jsapi

import (
	"encoding/json"
	"fmt"
	"github.com/WuJieOnce/wechat/utils"
	"github.com/WuJieOnce/wechat/v3/payment/global"
	"time"
)

type Client struct {
	Config    *global.Config
	Body      *Body
	Signature string
	Url       string
}

func (c *Client) SetConfig(conf *global.Config) *Client {
	c.Url = "https://api.mch.weixin.qq.com"
	c.Config = conf
	return c
}

func (c *Client) SetBody(body *Body) *Client {
	c.Body = body
	return c
}

// PlaceOrder 发起下单
func (c *Client) PlaceOrder() (*CallsUpPayment, error) {
	// 获取当前时间戳（单位：秒）
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	// 生成随机字符串
	nonceStr := utils.GenerateNonceStr(32)
	// 请求地址
	url := "/v3/pay/transactions/jsapi"

	marshal, err := json.Marshal(c.Body)
	if err != nil {
		return nil, err
	}

	// 拼接待签名字符串
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", "POST", url, timestamp, nonceStr, string(marshal))
	sign, err := utils.SignWithPrivateKey(message, c.Config.PrivateKey)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf(`Authorization: WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",signature="%s",timestamp="%s",serial_no="%s"`, c.Config.MchId, nonceStr, sign, timestamp, c.Config.SerialNo)
	headers["Accept"] = "application/json"
	resp, err := utils.PostJSON(fmt.Sprintf("%s%s", c.Url, url), c.Body, headers)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, resp.Err
	}

	prepaid := Prepaid{}
	if err = json.Unmarshal(resp.Body, &prepaid); err != nil {
		return nil, err
	}

	prepayId := fmt.Sprintf("prepay_id=%s", prepaid.PrepayId)
	message = fmt.Sprintf("%s\n%s\n%s\n%s\n", c.Config.Appid, timestamp, nonceStr, prepayId)
	sign, err = utils.SignWithPrivateKey(message, c.Config.PrivateKey)
	if err != nil {
		return nil, err
	}
	response := CallsUpPayment{
		AppId:     c.Config.Appid,
		TimeStamp: timestamp,
		NonceStr:  nonceStr,
		Package:   prepayId,
		SignType:  "RSA",
		PaySign:   sign,
	}
	return &response, nil
}

// QueryByTransactionId 微信支付订单号查询订单
func (c *Client) QueryByTransactionId(transactionId string) (*QueryOrderResp, error) {
	url := fmt.Sprintf("/v3/pay/transactions/id/%s?mchid=%s", transactionId, c.Config.MchId)
	// 获取当前时间戳（单位：秒）
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	// 生成随机字符串
	nonceStr := utils.GenerateNonceStr(32)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n\n", "GET", url, timestamp, nonceStr)
	sign, err := utils.SignWithPrivateKey(message, c.Config.PrivateKey)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",signature="%s",timestamp="%s",serial_no="%s"`, c.Config.MchId, nonceStr, sign, timestamp, c.Config.SerialNo),
	}
	resp, err := utils.Get(fmt.Sprintf("%s%s", c.Url, url), headers)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, resp.Err
	}
	response := QueryOrderResp{}
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// QueryOrderByMerchantOrder 商户订单号查询订单
func (c *Client) QueryOrderByMerchantOrder() {
	//TODO implement me
	panic("implement me")
}

// CloseOrder 关闭订单
func (c *Client) CloseOrder() {
	//TODO implement me
	panic("implement me")
}

// Notification 支付成功回调通知
func (c *Client) Notification() {
	//TODO implement me
	panic("implement me")
}

// RefundOrder 退款申请
func (c *Client) RefundOrder() {
	//TODO implement me
	panic("implement me")
}

// QueryASingleRefund 查询单笔退款（通过商户退款单号）
func (c *Client) QueryASingleRefund() {
	//TODO implement me
	panic("implement me")
}

// AbnormalRefund 发起异常退款
func (c *Client) AbnormalRefund() {
	//TODO implement me
	panic("implement me")
}

// NotificationByRefund 退款结果回调通知
func (c *Client) NotificationByRefund() {
	//TODO implement me
	panic("implement me")
}

// TransactionStatement 申请交易账单
func (c *Client) TransactionStatement() {
	//TODO implement me
	panic("implement me")
}

// FundsBill 申请资金账单
func (c *Client) FundsBill() {
	//TODO implement me
	panic("implement me")
}

// DownloadBill 下载账单
func (c *Client) DownloadBill() {
	//TODO implement me
	panic("implement me")
}
