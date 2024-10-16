package service

import (
	"context"
  v1 "pion/api/dip-bff/v1"
  "pion/app/dip-bff/internal/biz"
	"pion/pkg/util"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type DoctorService struct {
  v1.UnimplementedDoctorServiceServer
  uc *biz.DoctorServiceUsecase
  log *log.Helper
}

func NewDoctorService(uc *biz.DoctorService, logger log.Logger) *DoctorService {
  return &DoctorService{
    log: log.NewHelper(log.With(logger, "module", "service/DoctorService")),
    uc: uc,
  }
}


func (s *DoctorService) DoctorGlobal(ctx context.Context, req *v1.DoctorGlobalRequest) (*v1.DoctorGlobalReply, error) {
  reply, err := s.uc.DoctorGlobal(ctx, &biz.DoctorGlobalRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Page: req.Page,
    PageSize: req.PageSize,
    Desc: req.Desc,
  })
  if err != nil {
    return nil, err
  }
  return &v1.DoctorGlobalReply{
    List: util.BulkItem2Proto[*biz.DoctorGlobalReplyDoctors, *v1.DoctorGlobalReply_Doctors](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *DoctorService) DoctorDept(ctx context.Context, req *v1.DoctorDeptRequest) (*v1.DoctorDeptReply, error) {
  reply, err := s.uc.DoctorDept(ctx, &biz.DoctorDeptRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Page: req.Page,
    PageSize: req.PageSize,
    Desc: req.Desc,
    DeptId: req.DeptId,
  })
  if err != nil {
    return nil, err
  }
  return &v1.DoctorDeptReply{
    List: util.BulkItem2Proto[*biz.DoctorDeptReplyDoctors, *v1.DoctorDeptReply_Doctors](reply.List),
    Count: reply.Count,
  }, nil
}

