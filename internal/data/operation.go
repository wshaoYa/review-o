package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "review-o/api/review/v1"
	"review-o/internal/biz"
)

type operationRepo struct {
	data *Data
	log  *log.Helper
}

// NewOperationRepo .
func NewOperationRepo(data *Data, logger log.Logger) biz.OperationRepo {
	return &operationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// AuditReview 审核评价
func (o operationRepo) AuditReview(ctx context.Context, param *biz.AuditReviewParam) error {
	_, err := o.data.rc.AuditReview(ctx, &v1.AuditReviewRequest{
		ReviewID:  param.ReviewID,
		Status:    int32(param.Status),
		OpUser:    param.OpUser,
		OpReason:  param.OpReason,
		OpRemarks: &param.OpRemarks,
	})
	return err
}

func (o operationRepo) AuditAppeal(ctx context.Context, param *biz.AuditAppealParam) error {
	_, err := o.data.rc.AuditAppeal(ctx, &v1.AuditAppealRequest{
		AppealID:  param.AppealID,
		ReviewID:  param.ReviewID,
		Status:    int32(param.Status),
		OpUser:    param.OpUser,
		OpRemarks: &param.OpRemarks,
	})
	return err
}
