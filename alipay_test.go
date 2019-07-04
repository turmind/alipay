package alipay_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/smartwalle/alipay"
)

var (
	appID = ""

	pID = ""

	// RSA2(SHA256)
	// 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取
	aliPublicKey = ""

	// 必须，上一步中使用 RSA签名验签工具 生成的私钥
	privateKey = ""
)

var client *alipay.Client

func init() {
	var err error
	client, err = alipay.New(appID, pID, aliPublicKey, privateKey, true)

	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		os.Exit(-1)
	}
}

func TestAccountAuth(t *testing.T) {
	t.Log("========== TestAccountAuth ==========")
	var p = alipay.AccountAuth{}
	p.TargeID = "1234567899"

	url, err := client.AuthAccount(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(url)
}

func TestTradeAppPay(t *testing.T) {
	t.Log("========== TestTradeAppPay ==========")
	var p = alipay.TradeAppPay{}
	p.NotifyURL = "http://zonepp.com:3090"
	p.TotalAmount = "0.01"
	p.TimeoutExpress = "30m"
	p.ProductCode = "QUICK_MSECURITY_PAY"
	p.Subject = "1"
	p.Body = "我是测试数据"
	p.OutTradeNo = "testgo03"

	url, err := client.TradeAppPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(url)
}
