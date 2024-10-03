package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// OperationRepo 运营审核repo
type OperationRepo interface {
	AuditReview(context.Context, *AuditReviewParam) error // O端 审核评价
	AuditAppeal(context.Context, *AuditAppealParam) error // O端 审核申诉
}

// OperationUsecase 运营审核usecase
type OperationUsecase struct {
	repo OperationRepo
	log  *log.Helper
}

// NewOperationUsecase 运营审核usecase构造函数
func NewOperationUsecase(repo OperationRepo, logger log.Logger) *OperationUsecase {
	return &OperationUsecase{repo: repo, log: log.NewHelper(logger)}
}

// AuditReview 运营审核评价
func (uc *OperationUsecase) AuditReview(ctx context.Context, param *AuditReviewParam) error {
	return uc.repo.AuditReview(ctx, param)
}

// AuditAppeal 运营审核申诉
func (uc *OperationUsecase) AuditAppeal(ctx context.Context, param *AuditAppealParam) error {
	return uc.repo.AuditAppeal(ctx, param)
}
