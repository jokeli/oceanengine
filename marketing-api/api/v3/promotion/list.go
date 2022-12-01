package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// List 获取广告列表
func List(clt *core.SDKClient, accessToken string, req *promotion.ListRequest) (*promotion.ListResponseData, error) {
	var resp promotion.ListResponse
	if err := clt.Get("v3.0/promotion/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
