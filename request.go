package lubantop

import "lubantop/requests"

func TaobaoTbkItemInfoGetRequest() *requests.TaobaoTbkItemInfoGet {
	var r requests.TaobaoTbkItemInfoGet
	return r.New()
}

func TaobaoTbkOrderDetailsGetRequest() *requests.TaobaoTbkOrderDetailsGet {
	var r requests.TaobaoTbkOrderDetailsGet
	return r.New()
}

func TaobaoTbkRelationRefundRequest() *requests.TaobaoTbkRelationRefund {
	var r requests.TaobaoTbkRelationRefund
	return r.New()
}
