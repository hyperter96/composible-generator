package service

import (
	"context"
  v1 "pion/api/bff/v1"
  "pion/app/bff/internal/biz"
	"pion/pkg/util"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type IndicatorValueService struct {
  v1.UnimplementedIndicatorValueServiceServer
  uc *biz.IndicatorValueServiceUsecase
  log *log.Helper
}

func NewIndicatorValueService(uc *biz.IndicatorValueService, logger log.Logger) *IndicatorValueService {
  return &IndicatorValueService{
    log: log.NewHelper(log.With(logger, "module", "service/IndicatorValueService")),
    uc: uc,
  }
}


func (s *IndicatorValueService) GlobalDetailCategoryIncomeDeptDetail(ctx context.Context, req *v1.GlobalDetailCategoryIncomeDeptDetailRequest) (*v1.GlobalDetailDeptsReply, error) {
  reply, err := s.uc.GlobalDetailCategoryIncomeDeptDetail(ctx, &biz.GlobalDetailCategoryIncomeDeptDetailRequestParam{
    Category: req.Category,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Sort: req.Sort,
    Page: req.Page,
    PageSize: req.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailDeptsReply{
    Unqualified: reply.Unqualified,
    List: util.BulkItem2Proto[*biz.GlobalDetailDeptsReplyDepts, *v1.GlobalDetailDeptsReply_Depts](reply.List),
    Count: reply.Count,
    Qualified: reply.Qualified,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailChildrenIndicator(ctx context.Context, req *v1.GlobalDetailChildrenIndicatorRequest) (*v1.GlobalDetailCategoryIncomeDetailReply, error) {
  reply, err := s.uc.GlobalDetailChildrenIndicator(ctx, &biz.GlobalDetailChildrenIndicatorRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailCategoryIncomeDetailReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailCategoryIncomeDetailReplyItem, *v1.GlobalDetailCategoryIncomeDetailReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailDoctorMain(ctx context.Context, req *v1.GlobalDetailDoctorMainRequest) (*v1.GlobalListReply, error) {
  reply, err := s.uc.GlobalDetailDoctorMain(ctx, &biz.GlobalDetailDoctorMainRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalListReply{
    Count: reply.Count,
    SingleDiseases: util.BulkItem2Proto[*biz.GlobalListReplyBadIndicatorStat, *v1.GlobalListReply_BadIndicatorStat](reply.SingleDiseases),
    BadCount: reply.BadCount,
    Tumours: util.BulkItem2Proto[*biz.GlobalListReplyBadIndicatorStat, *v1.GlobalListReply_BadIndicatorStat](reply.Tumours),
    BadNames: reply.BadNames,
    NiceCount: reply.NiceCount,
    List: util.BulkItem2Proto[*biz.GlobalListReplyGlobal, *v1.GlobalListReply_Global](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailSurgery(ctx context.Context, req *v1.GlobalDetailSurgeryRequest) (*v1.GlobalDetailSurgeryReply, error) {
  reply, err := s.uc.GlobalDetailSurgery(ctx, &biz.GlobalDetailSurgeryRequestParam{
    Id: req.Id,
    Page: req.Page,
    PageSize: req.PageSize,
    Sort: req.Sort,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailSurgeryReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailSurgeryReplyItem, *v1.GlobalDetailSurgeryReply_Item](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailSingleDiseaseMain(ctx context.Context, req *v1.GlobalDetailSingleDiseaseMainRequest) (*v1.GlobalDetailSingleDiseaseMainReply, error) {
  reply, err := s.uc.GlobalDetailSingleDiseaseMain(ctx, &biz.GlobalDetailSingleDiseaseMainRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailSingleDiseaseMainReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailSingleDiseaseMainReplyItem, *v1.GlobalDetailSingleDiseaseMainReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailDoctors(ctx context.Context, req *v1.GlobalDetailDoctorsRequest) (*v1.GlobalDetailDoctorsReply, error) {
  reply, err := s.uc.GlobalDetailDoctors(ctx, &biz.GlobalDetailDoctorsRequestParam{
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
  return &v1.GlobalDetailDoctorsReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailDoctorsReplyDoctors, *v1.GlobalDetailDoctorsReply_Doctors](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailGroups(ctx context.Context, req *v1.GlobalDetailGroupsRequest) (*v1.GlobalDetailGroupsReply, error) {
  reply, err := s.uc.GlobalDetailGroups(ctx, &biz.GlobalDetailGroupsRequestParam{
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
  return &v1.GlobalDetailGroupsReply{
    Unqualified: reply.Unqualified,
    Qualified: reply.Qualified,
    List: util.BulkItem2Proto[*biz.GlobalDetailGroupsReplyGroups, *v1.GlobalDetailGroupsReply_Groups](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailDoctorsSort(ctx context.Context, req *v1.GlobalDetailDoctorsSortRequest) (*v1.GlobalDetailDoctorsSortReply, error) {
  reply, err := s.uc.GlobalDetailDoctorsSort(ctx, &biz.GlobalDetailDoctorsSortRequestParam{
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
  return &v1.GlobalDetailDoctorsSortReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailDoctorsSortReplyItem, *v1.GlobalDetailDoctorsSortReply_Item](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailNumberatorDgs(ctx context.Context, req *v1.GlobalDetailNumberatorDgsRequest) (*v1.GlobalDetailGroupsReply, error) {
  reply, err := s.uc.GlobalDetailNumberatorDgs(ctx, &biz.GlobalDetailNumberatorDgsRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Page: req.Page,
    PageSize: req.PageSize,
    Desc: req.Desc,
    OnlyUnqualified: req.OnlyUnqualified,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailGroupsReply{
    Unqualified: reply.Unqualified,
    Qualified: reply.Qualified,
    List: util.BulkItem2Proto[*biz.GlobalDetailGroupsReplyGroups, *v1.GlobalDetailGroupsReply_Groups](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailMonthSummary(ctx context.Context, req *v1.GlobalDetailMonthSummaryRequest) (*v1.GlobalDetailMonthSummaryReply, error) {
  reply, err := s.uc.GlobalDetailMonthSummary(ctx, &biz.GlobalDetailMonthSummaryRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailMonthSummaryReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailMonthSummaryReplyTuple, *v1.GlobalDetailMonthSummaryReply_Tuple](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailCategoryIncomeDetail(ctx context.Context, req *v1.GlobalDetailCategoryIncomeDetailRequest) (*v1.GlobalDetailCategoryIncomeDetailReply, error) {
  reply, err := s.uc.GlobalDetailCategoryIncomeDetail(ctx, &biz.GlobalDetailCategoryIncomeDetailRequestParam{
    Category: req.Category,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailCategoryIncomeDetailReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailCategoryIncomeDetailReplyItem, *v1.GlobalDetailCategoryIncomeDetailReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailDeptSurgeryNumSort(ctx context.Context, req *v1.GlobalDetailDeptSurgeryNumSortRequest) (*v1.GlobalDetailDeptSurgeryNumSortReply, error) {
  reply, err := s.uc.GlobalDetailDeptSurgeryNumSort(ctx, &biz.GlobalDetailDeptSurgeryNumSortRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Desc: req.Desc,
    Page: req.Page,
    PageSize: req.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailDeptSurgeryNumSortReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailDeptSurgeryNumSortReplyItem, *v1.GlobalDetailDeptSurgeryNumSortReply_Item](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailSingleDiseaseDoctorGroups(ctx context.Context, req *v1.GlobalDetailSingleDiseaseDoctorGroupsRequest) (*v1.GlobalDetailSingleDiseaseDoctorGroupsReply, error) {
  reply, err := s.uc.GlobalDetailSingleDiseaseDoctorGroups(ctx, &biz.GlobalDetailSingleDiseaseDoctorGroupsRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Indicator: req.Indicator,
    Sort: req.Sort,
    Page: req.Page,
    PageSize: req.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailSingleDiseaseDoctorGroupsReply{
    Count: reply.Count,
    List: util.BulkItem2Proto[*biz.GlobalDetailSingleDiseaseDoctorGroupsReplyItem, *v1.GlobalDetailSingleDiseaseDoctorGroupsReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailTumourDepts(ctx context.Context, req *v1.GlobalDetailTumourDeptsRequest) (*v1.GlobalDetailTumourDeptsReply, error) {
  reply, err := s.uc.GlobalDetailTumourDepts(ctx, &biz.GlobalDetailTumourDeptsRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Indicator: req.Indicator,
    Sort: req.Sort,
    Page: req.Page,
    PageSize: req.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailTumourDeptsReply{
    Count: reply.Count,
    List: util.BulkItem2Proto[*biz.GlobalDetailTumourDeptsReplyItem, *v1.GlobalDetailTumourDeptsReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailMain(ctx context.Context, req *v1.GlobalDetailMainRequest) (*v1.GlobalDetailMainReply, error) {
  reply, err := s.uc.GlobalDetailMain(ctx, &biz.GlobalDetailMainRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailMainReply{
    PastNumberator: reply.PastNumberator,
    Cur: reply.Cur,
    TargetDenominator: reply.TargetDenominator,
    TargetNumberator: reply.TargetNumberator,
    CurNumberator: reply.CurNumberator,
    Diff: reply.Diff,
    Target: reply.Target,
    CurDenominator: reply.CurDenominator,
    DiffNumberator: reply.DiffNumberator,
    PastDenominator: reply.PastDenominator,
    Past: reply.Past,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailNumberatorDepts(ctx context.Context, req *v1.GlobalDetailNumberatorDeptsRequest) (*v1.GlobalDetailDeptsReply, error) {
  reply, err := s.uc.GlobalDetailNumberatorDepts(ctx, &biz.GlobalDetailNumberatorDeptsRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Page: req.Page,
    PageSize: req.PageSize,
    Desc: req.Desc,
    OnlyUnqualified: req.OnlyUnqualified,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailDeptsReply{
    Unqualified: reply.Unqualified,
    List: util.BulkItem2Proto[*biz.GlobalDetailDeptsReplyDepts, *v1.GlobalDetailDeptsReply_Depts](reply.List),
    Count: reply.Count,
    Qualified: reply.Qualified,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailCategoryIncomes(ctx context.Context, req *v1.GlobalDetailCategoryIncomesRequest) (*v1.GlobalDetailCategoryIncomesReply, error) {
  reply, err := s.uc.GlobalDetailCategoryIncomes(ctx, &biz.GlobalDetailCategoryIncomesRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailCategoryIncomesReply{
    Categories: util.BulkItem2Proto[*biz.GlobalDetailCategoryIncomesReplyItem, *v1.GlobalDetailCategoryIncomesReply_Item](reply.Categories),
    Types: util.BulkItem2Proto[*biz.GlobalDetailCategoryIncomesReplyItem, *v1.GlobalDetailCategoryIncomesReply_Item](reply.Types),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailTumourDoctorGroups(ctx context.Context, req *v1.GlobalDetailTumourDoctorGroupsRequest) (*v1.GlobalDetailTumourDoctorGroupsReply, error) {
  reply, err := s.uc.GlobalDetailTumourDoctorGroups(ctx, &biz.GlobalDetailTumourDoctorGroupsRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Indicator: req.Indicator,
    Sort: req.Sort,
    Page: req.Page,
    PageSize: req.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailTumourDoctorGroupsReply{
    Count: reply.Count,
    List: util.BulkItem2Proto[*biz.GlobalDetailTumourDoctorGroupsReplyItem, *v1.GlobalDetailTumourDoctorGroupsReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailDoctorSurgeryNumSort(ctx context.Context, req *v1.GlobalDetailDoctorSurgeryNumSortRequest) (*v1.GlobalDetailDoctorSurgeryNumSortReply, error) {
  reply, err := s.uc.GlobalDetailDoctorSurgeryNumSort(ctx, &biz.GlobalDetailDoctorSurgeryNumSortRequestParam{
    Id: req.Id,
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Desc: req.Desc,
    Page: req.Page,
    PageSize: req.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailDoctorSurgeryNumSortReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailDoctorSurgeryNumSortReplyItem, *v1.GlobalDetailDoctorSurgeryNumSortReply_Item](reply.List),
    ReportZero: reply.ReportZero,
    Count: reply.Count,
    ReportList: util.BulkItem2Proto[*biz.GlobalDetailDoctorSurgeryNumSortReplySecondItem, *v1.GlobalDetailDoctorSurgeryNumSortReply_SecondItem](reply.ReportList),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailDepts(ctx context.Context, req *v1.GlobalDetailDeptsRequest) (*v1.GlobalDetailDeptsReply, error) {
  reply, err := s.uc.GlobalDetailDepts(ctx, &biz.GlobalDetailDeptsRequestParam{
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
  return &v1.GlobalDetailDeptsReply{
    Unqualified: reply.Unqualified,
    List: util.BulkItem2Proto[*biz.GlobalDetailDeptsReplyDepts, *v1.GlobalDetailDeptsReply_Depts](reply.List),
    Count: reply.Count,
    Qualified: reply.Qualified,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailSingleDiseaseDepts(ctx context.Context, req *v1.GlobalDetailSingleDiseaseDeptsRequest) (*v1.GlobalDetailSingleDiseaseDeptsReply, error) {
  reply, err := s.uc.GlobalDetailSingleDiseaseDepts(ctx, &biz.GlobalDetailSingleDiseaseDeptsRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    Indicator: req.Indicator,
    Sort: req.Sort,
    Page: req.Page,
    PageSize: req.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailSingleDiseaseDeptsReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailSingleDiseaseDeptsReplyItem, *v1.GlobalDetailSingleDiseaseDeptsReply_Item](reply.List),
    Count: reply.Count,
  }, nil
}

func (s *IndicatorValueService) GlobalDetailTumourMain(ctx context.Context, req *v1.GlobalDetailTumourMainRequest) (*v1.GlobalDetailTumourMainReply, error) {
  reply, err := s.uc.GlobalDetailTumourMain(ctx, &biz.GlobalDetailTumourMainRequestParam{
    OrgId: req.OrgId,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailTumourMainReply{
    List: util.BulkItem2Proto[*biz.GlobalDetailTumourMainReplyItem, *v1.GlobalDetailTumourMainReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalDetailDoctorGroup(ctx context.Context, req *v1.GlobalDetailDoctorGroupRequest) (*v1.GlobalDetailDoctorGroupReply, error) {
  reply, err := s.uc.GlobalDetailDoctorGroup(ctx, &biz.GlobalDetailDoctorGroupRequestParam{
    Id: req.Id,
    Page: req.Page,
    PageSize: req.PageSize,
    Sort: req.Sort,
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalDetailDoctorGroupReply{
    Count: reply.Count,
    List: util.BulkItem2Proto[*biz.GlobalDetailDoctorGroupReplyItem, *v1.GlobalDetailDoctorGroupReply_Item](reply.List),
  }, nil
}

func (s *IndicatorValueService) GlobalList(ctx context.Context, req *v1.GlobalListRequest) (*v1.GlobalListReply, error) {
  reply, err := s.uc.GlobalList(ctx, &biz.GlobalListRequestParam{
    BeginDate: req.BeginDate,
    EndDate: req.EndDate,
    OrgCode: req.OrgCode,
  })
  if err != nil {
    return nil, err
  }
  return &v1.GlobalListReply{
    Count: reply.Count,
    SingleDiseases: util.BulkItem2Proto[*biz.GlobalListReplyBadIndicatorStat, *v1.GlobalListReply_BadIndicatorStat](reply.SingleDiseases),
    BadCount: reply.BadCount,
    Tumours: util.BulkItem2Proto[*biz.GlobalListReplyBadIndicatorStat, *v1.GlobalListReply_BadIndicatorStat](reply.Tumours),
    BadNames: reply.BadNames,
    NiceCount: reply.NiceCount,
    List: util.BulkItem2Proto[*biz.GlobalListReplyGlobal, *v1.GlobalListReply_Global](reply.List),
  }, nil
}

