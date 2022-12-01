package project

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// Update 修改项目
func Update(clt *core.SDKClient, accessToken string, req *project.UpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.Post("v3.0/project/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
