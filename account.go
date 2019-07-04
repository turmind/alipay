package alipay

// AuthAccount https://docs.open.alipay.com/218/105327/
func (a *Client) AuthAccount(param AccountAuth) (results string, err error) {
	p, err := a.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}
