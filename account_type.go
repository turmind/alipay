package alipay

// Account 阿里账号相关
type Account struct {
}

// AccountAuth 阿里账号授权相关
type AccountAuth struct {
	TargeID string
}

// APIName 获取apiname
func (a AccountAuth) APIName() string {
	return "com.alipay.account.auth"
}

// Params 获取账号授权参数
func (a AccountAuth) Params() map[string]string {
	var m = make(map[string]string)
	m["apiname"] = "com.alipay.account.auth"
	m["methodname"] = "alipay.open.auth.sdk.code.get"
	m["app_name"] = "mc"
	m["biz_type"] = "openservice"
	m["product_id"] = "APP_FAST_LOGIN"
	m["scope"] = "kuaijie"
	m["target_id"] = a.TargeID
	m["auth_type"] = "AUTHACCOUNT"
	return m
}

func (a AccountAuth) ExtJSONParamName() string {
	return ""
}

func (a AccountAuth) ExtJSONParamValue() string {
	return marshal(a)
}
