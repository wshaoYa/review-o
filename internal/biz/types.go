package biz

// AuditReviewParam 审核评价的参数
type AuditReviewParam struct {
	ReviewID  int64
	Status    int
	OpReason  string
	OpRemarks string
	OpUser    string
}

// AuditAppealParam 审核申诉的参数
type AuditAppealParam struct {
	AppealID  int64
	ReviewID  int64
	StoreID   int64
	Status    int
	OpReason  string
	OpRemarks string
	OpUser    string
}
