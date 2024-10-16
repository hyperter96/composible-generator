// Biz Layer for DoctorService
package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
  
)

// DoctorGlobalRequestParam struct represents the request parameters for DoctorGlobal API
type DoctorGlobalRequestParam struct {
  Id string
  OrgId *int32
  BeginDate string
  EndDate string
  Page int32
  PageSize int32
  Desc *bool
}

// DoctorDeptRequestParam struct represents the request parameters for DoctorDept API
type DoctorDeptRequestParam struct {
  Id string
  OrgId *int32
  BeginDate string
  EndDate string
  Page int32
  PageSize int32
  Desc *bool
  DeptId int32
}


// DoctorGlobalReplyParam struct represents the reply parameters for DoctorGlobal API
type DoctorGlobalReplyParam struct {
  List []*DoctorGlobalReplyDoctors
  Count int32
}

// DoctorDeptReplyParam struct represents the reply parameters for DoctorDept API
type DoctorDeptReplyParam struct {
  List []*DoctorDeptReplyDoctors
  Count int32
}



type DoctorDeptReplyDoctors struct {
  ActualValue  float64
  TargetValue  float64
  Diff  float64
  DoctorId  int32
  DoctorName  string
}

type DoctorGlobalReplyDoctors struct {
  ActualValue  float64
  TargetValue  float64
  DoctorName  string
  Diff  float64
  DoctorId  int32
}


type DoctorServiceRepo interface {
  DoctorGlobal(ctx context.Context, param *DoctorGlobalRequestParam) (*DoctorGlobalReplyParam, error)
  DoctorDept(ctx context.Context, param *DoctorDeptRequestParam) (*DoctorDeptReplyParam, error)
}

type DoctorServiceUsecase struct {
	repo DoctorServiceRepo
	log  *log.Helper
}

func NewDoctorServiceUsecase(repo DoctorServiceRepo, logger log.Logger) *DoctorServiceUsecase {
	return &DoctorServiceUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/doctor_service")),
	}
}


func (uc *DoctorServiceUsecase) DoctorGlobal(ctx context.Context, param *DoctorGlobalRequestParam) (*DoctorGlobalReplyParam, error) {
	return uc.repo.DoctorGlobal(ctx, param)
}

func (uc *DoctorServiceUsecase) DoctorDept(ctx context.Context, param *DoctorDeptRequestParam) (*DoctorDeptReplyParam, error) {
	return uc.repo.DoctorDept(ctx, param)
}

