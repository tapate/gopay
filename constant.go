package go_pay

const (
	//URL
	WX_PayUrl        = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	WX_PayUrl_SanBox = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"

	//支付类型
	WX_PayType_Mini   = "JSAPI"
	WX_PayType_JsApi  = "JSAPI"
	WX_PayType_App    = "APP"
	WX_PayType_H5     = "MWEB"
	WX_PayType_Native = "NATIVE"

	//签名方式
	WX_SignType_MD5         = "MD5"
	WX_SignType_HMAC_SHA256 = "HMAC-SHA256"

	//Debug数据
	SecretKey = "GFDS8j98rewnmgl45wHTt980jg543wmg"
	AppID     = "wxdaa2ab9ef87b5497"
	MchID     = "1368139502"
	OpenID    = "o0Df70H2Q0fY8JXh1aFPIRyOBgu8"
)
