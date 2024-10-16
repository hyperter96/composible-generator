package data

import (
	"context"
  dipV1 "pion/api/dip/v1"
  "pion/app/dip-bff/internal/biz"
	"pion/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.DoctorServiceRepo = (*DoctorServiceRepo)(nil)

type DoctorServiceRepo struct {
  dip dipV1.DoctorServiceClient
  log *log.Helper
}

func NewDoctorServiceRepo(logger log.Logger, dip dipV1.DoctorServiceClient) biz.DoctorServiceRepo {
  return &DoctorServiceRepo{
    dip: dip,
    log: log.NewHelper(log.With(logger, "module", "data/dip")),
  }
}


func (r DoctorServiceRepo) DoctorGlobal(ctx context.Context, param *biz.DoctorGlobalRequestParam) (*biz.DoctorGlobalReplyParam, error) {
  reply, err := r.dip.DoctorGlobal(ctx, &dipV1.DoctorGlobalRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Page: param.Page,
    PageSize: param.PageSize,
    Desc: param.Desc,
  })
  if err != nil {
    return nil, err
  }
  return &biz.DoctorGlobalReplyParam{
    List: util.BulkProto2Biz[*v1.DoctorGlobalReply_Doctors, *biz.DoctorGlobalReplyDoctors](reply.List),
    Count: reply.Count,
  }, nil
}

func (r DoctorServiceRepo) DoctorDept(ctx context.Context, param *biz.DoctorDeptRequestParam) (*biz.DoctorDeptReplyParam, error) {
  reply, err := r.dip.DoctorDept(ctx, &dipV1.DoctorDeptRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Page: param.Page,
    PageSize: param.PageSize,
    Desc: param.Desc,
    DeptId: param.DeptId,
  })
  if err != nil {
    return nil, err
  }
  return &biz.DoctorDeptReplyParam{
    List: util.BulkProto2Biz[*v1.DoctorDeptReply_Doctors, *biz.DoctorDeptReplyDoctors](reply.List),
    Count: reply.Count,
  }, nil
}

