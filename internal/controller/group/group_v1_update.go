package group

import (
	"context"
	"sort"

	"github.com/iimeta/fastapi-admin/v2/api/group/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	for _, rule := range req.TimeRules {
		rule.Discount = util.Round(rule.Discount/100, 2)
		sort.Slice(rule.Days, func(i, j int) bool {
			if rule.Days[i] == 0 && rule.Days[j] != 0 {
				return false
			}
			if rule.Days[j] == 0 && rule.Days[i] != 0 {
				return true
			}
			return rule.Days[i] < rule.Days[j]
		})
	}

	err = service.Group().Update(ctx, req.GroupUpdateReq)

	return
}
