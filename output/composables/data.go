package data

import (
	"context"
  performanceV1 "pion/api/performance/v1"
  "pion/app/bff/internal/biz"
	"pion/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.IndicatorValueServiceRepo = (*IndicatorValueServiceRepo)(nil)

type IndicatorValueServiceRepo struct {
  performance performanceV1.IndicatorValueServiceClient
  log *log.Helper
}

func NewIndicatorValueServiceRepo(logger Logger, performance performanceV1.IndicatorValueServiceClient) biz.IndicatorValueServiceRepo {
  return &IndicatorValueService{
    performance: performance,
    log: log.NewHelper(log.With(logger, "module", "data/performance")),
  }
}


func (r IndicatorValueServiceRepo) GlobalDetailCategoryIncomeDeptDetail(ctx context.Context, param *biz.GlobalDetailCategoryIncomeDeptDetailRequestParam) (*biz.GlobalDetailDeptsReplyParam, error) {
  reply, err := r.performance.GlobalDetailCategoryIncomeDeptDetail(ctx, &performanceV1.GlobalDetailCategoryIncomeDeptDetailRequest{
    Category: param.Category,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Sort: param.Sort,
    Page: param.Page,
    PageSize: param.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailDeptsReplyParam{
    Unqualified: reply.Unqualified,
    List: util.BulkProto2Biz[*v1.GlobalDetailDeptsReply_Depts, *biz.GlobalDetailDeptsReplyDepts](reply.List),
    Count: reply.Count,
    Qualified: reply.Qualified,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailChildrenIndicator(ctx context.Context, param *biz.GlobalDetailChildrenIndicatorRequestParam) (*biz.GlobalDetailCategoryIncomeDetailReplyParam, error) {
  reply, err := r.performance.GlobalDetailChildrenIndicator(ctx, &performanceV1.GlobalDetailChildrenIndicatorRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailCategoryIncomeDetailReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailCategoryIncomeDetailReply_Item, *biz.GlobalDetailCategoryIncomeDetailReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailDoctorMain(ctx context.Context, param *biz.GlobalDetailDoctorMainRequestParam) (*biz.GlobalListReplyParam, error) {
  reply, err := r.performance.GlobalDetailDoctorMain(ctx, &performanceV1.GlobalDetailDoctorMainRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalListReplyParam{
    Count: reply.Count,
    SingleDiseases: util.BulkProto2Biz[*v1.GlobalListReply_BadIndicatorStat, *biz.GlobalListReplyBadIndicatorStat](reply.SingleDiseases),
    BadCount: reply.BadCount,
    Tumours: util.BulkProto2Biz[*v1.GlobalListReply_BadIndicatorStat, *biz.GlobalListReplyBadIndicatorStat](reply.Tumours),
    BadNames: reply.BadNames,
    NiceCount: reply.NiceCount,
    List: util.BulkProto2Biz[*v1.GlobalListReply_Global, *biz.GlobalListReplyGlobal](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailSurgery(ctx context.Context, param *biz.GlobalDetailSurgeryRequestParam) (*biz.GlobalDetailSurgeryReplyParam, error) {
  reply, err := r.performance.GlobalDetailSurgery(ctx, &performanceV1.GlobalDetailSurgeryRequest{
    Id: param.Id,
    Page: param.Page,
    PageSize: param.PageSize,
    Sort: param.Sort,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailSurgeryReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailSurgeryReply_Item, *biz.GlobalDetailSurgeryReplyItem](reply.List),
    Count: reply.Count,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailSingleDiseaseMain(ctx context.Context, param *biz.GlobalDetailSingleDiseaseMainRequestParam) (*biz.GlobalDetailSingleDiseaseMainReplyParam, error) {
  reply, err := r.performance.GlobalDetailSingleDiseaseMain(ctx, &performanceV1.GlobalDetailSingleDiseaseMainRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailSingleDiseaseMainReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailSingleDiseaseMainReply_Item, *biz.GlobalDetailSingleDiseaseMainReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailDoctors(ctx context.Context, param *biz.GlobalDetailDoctorsRequestParam) (*biz.GlobalDetailDoctorsReplyParam, error) {
  reply, err := r.performance.GlobalDetailDoctors(ctx, &performanceV1.GlobalDetailDoctorsRequest{
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
  return &biz.GlobalDetailDoctorsReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailDoctorsReply_Doctors, *biz.GlobalDetailDoctorsReplyDoctors](reply.List),
    Count: reply.Count,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailGroups(ctx context.Context, param *biz.GlobalDetailGroupsRequestParam) (*biz.GlobalDetailGroupsReplyParam, error) {
  reply, err := r.performance.GlobalDetailGroups(ctx, &performanceV1.GlobalDetailGroupsRequest{
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
  return &biz.GlobalDetailGroupsReplyParam{
    Unqualified: reply.Unqualified,
    Qualified: reply.Qualified,
    List: util.BulkProto2Biz[*v1.GlobalDetailGroupsReply_Groups, *biz.GlobalDetailGroupsReplyGroups](reply.List),
    Count: reply.Count,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailDoctorsSort(ctx context.Context, param *biz.GlobalDetailDoctorsSortRequestParam) (*biz.GlobalDetailDoctorsSortReplyParam, error) {
  reply, err := r.performance.GlobalDetailDoctorsSort(ctx, &performanceV1.GlobalDetailDoctorsSortRequest{
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
  return &biz.GlobalDetailDoctorsSortReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailDoctorsSortReply_Item, *biz.GlobalDetailDoctorsSortReplyItem](reply.List),
    Count: reply.Count,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailNumberatorDgs(ctx context.Context, param *biz.GlobalDetailNumberatorDgsRequestParam) (*biz.GlobalDetailGroupsReplyParam, error) {
  reply, err := r.performance.GlobalDetailNumberatorDgs(ctx, &performanceV1.GlobalDetailNumberatorDgsRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Page: param.Page,
    PageSize: param.PageSize,
    Desc: param.Desc,
    OnlyUnqualified: param.OnlyUnqualified,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailGroupsReplyParam{
    Unqualified: reply.Unqualified,
    Qualified: reply.Qualified,
    List: util.BulkProto2Biz[*v1.GlobalDetailGroupsReply_Groups, *biz.GlobalDetailGroupsReplyGroups](reply.List),
    Count: reply.Count,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailMonthSummary(ctx context.Context, param *biz.GlobalDetailMonthSummaryRequestParam) (*biz.GlobalDetailMonthSummaryReplyParam, error) {
  reply, err := r.performance.GlobalDetailMonthSummary(ctx, &performanceV1.GlobalDetailMonthSummaryRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailMonthSummaryReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailMonthSummaryReply_Tuple, *biz.GlobalDetailMonthSummaryReplyTuple](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailCategoryIncomeDetail(ctx context.Context, param *biz.GlobalDetailCategoryIncomeDetailRequestParam) (*biz.GlobalDetailCategoryIncomeDetailReplyParam, error) {
  reply, err := r.performance.GlobalDetailCategoryIncomeDetail(ctx, &performanceV1.GlobalDetailCategoryIncomeDetailRequest{
    Category: param.Category,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailCategoryIncomeDetailReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailCategoryIncomeDetailReply_Item, *biz.GlobalDetailCategoryIncomeDetailReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailDeptSurgeryNumSort(ctx context.Context, param *biz.GlobalDetailDeptSurgeryNumSortRequestParam) (*biz.GlobalDetailDeptSurgeryNumSortReplyParam, error) {
  reply, err := r.performance.GlobalDetailDeptSurgeryNumSort(ctx, &performanceV1.GlobalDetailDeptSurgeryNumSortRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Desc: param.Desc,
    Page: param.Page,
    PageSize: param.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailDeptSurgeryNumSortReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailDeptSurgeryNumSortReply_Item, *biz.GlobalDetailDeptSurgeryNumSortReplyItem](reply.List),
    Count: reply.Count,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailSingleDiseaseDoctorGroups(ctx context.Context, param *biz.GlobalDetailSingleDiseaseDoctorGroupsRequestParam) (*biz.GlobalDetailSingleDiseaseDoctorGroupsReplyParam, error) {
  reply, err := r.performance.GlobalDetailSingleDiseaseDoctorGroups(ctx, &performanceV1.GlobalDetailSingleDiseaseDoctorGroupsRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Indicator: param.Indicator,
    Sort: param.Sort,
    Page: param.Page,
    PageSize: param.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailSingleDiseaseDoctorGroupsReplyParam{
    Count: reply.Count,
    List: util.BulkProto2Biz[*v1.GlobalDetailSingleDiseaseDoctorGroupsReply_Item, *biz.GlobalDetailSingleDiseaseDoctorGroupsReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailTumourDepts(ctx context.Context, param *biz.GlobalDetailTumourDeptsRequestParam) (*biz.GlobalDetailTumourDeptsReplyParam, error) {
  reply, err := r.performance.GlobalDetailTumourDepts(ctx, &performanceV1.GlobalDetailTumourDeptsRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Indicator: param.Indicator,
    Sort: param.Sort,
    Page: param.Page,
    PageSize: param.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailTumourDeptsReplyParam{
    Count: reply.Count,
    List: util.BulkProto2Biz[*v1.GlobalDetailTumourDeptsReply_Item, *biz.GlobalDetailTumourDeptsReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailMain(ctx context.Context, param *biz.GlobalDetailMainRequestParam) (*biz.GlobalDetailMainReplyParam, error) {
  reply, err := r.performance.GlobalDetailMain(ctx, &performanceV1.GlobalDetailMainRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailMainReplyParam{
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

func (r IndicatorValueServiceRepo) GlobalDetailNumberatorDepts(ctx context.Context, param *biz.GlobalDetailNumberatorDeptsRequestParam) (*biz.GlobalDetailDeptsReplyParam, error) {
  reply, err := r.performance.GlobalDetailNumberatorDepts(ctx, &performanceV1.GlobalDetailNumberatorDeptsRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Page: param.Page,
    PageSize: param.PageSize,
    Desc: param.Desc,
    OnlyUnqualified: param.OnlyUnqualified,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailDeptsReplyParam{
    Unqualified: reply.Unqualified,
    List: util.BulkProto2Biz[*v1.GlobalDetailDeptsReply_Depts, *biz.GlobalDetailDeptsReplyDepts](reply.List),
    Count: reply.Count,
    Qualified: reply.Qualified,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailCategoryIncomes(ctx context.Context, param *biz.GlobalDetailCategoryIncomesRequestParam) (*biz.GlobalDetailCategoryIncomesReplyParam, error) {
  reply, err := r.performance.GlobalDetailCategoryIncomes(ctx, &performanceV1.GlobalDetailCategoryIncomesRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailCategoryIncomesReplyParam{
    Categories: util.BulkProto2Biz[*v1.GlobalDetailCategoryIncomesReply_Item, *biz.GlobalDetailCategoryIncomesReplyItem](reply.Categories),
    Types: util.BulkProto2Biz[*v1.GlobalDetailCategoryIncomesReply_Item, *biz.GlobalDetailCategoryIncomesReplyItem](reply.Types),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailTumourDoctorGroups(ctx context.Context, param *biz.GlobalDetailTumourDoctorGroupsRequestParam) (*biz.GlobalDetailTumourDoctorGroupsReplyParam, error) {
  reply, err := r.performance.GlobalDetailTumourDoctorGroups(ctx, &performanceV1.GlobalDetailTumourDoctorGroupsRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Indicator: param.Indicator,
    Sort: param.Sort,
    Page: param.Page,
    PageSize: param.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailTumourDoctorGroupsReplyParam{
    Count: reply.Count,
    List: util.BulkProto2Biz[*v1.GlobalDetailTumourDoctorGroupsReply_Item, *biz.GlobalDetailTumourDoctorGroupsReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailDoctorSurgeryNumSort(ctx context.Context, param *biz.GlobalDetailDoctorSurgeryNumSortRequestParam) (*biz.GlobalDetailDoctorSurgeryNumSortReplyParam, error) {
  reply, err := r.performance.GlobalDetailDoctorSurgeryNumSort(ctx, &performanceV1.GlobalDetailDoctorSurgeryNumSortRequest{
    Id: param.Id,
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Desc: param.Desc,
    Page: param.Page,
    PageSize: param.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailDoctorSurgeryNumSortReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailDoctorSurgeryNumSortReply_Item, *biz.GlobalDetailDoctorSurgeryNumSortReplyItem](reply.List),
    ReportZero: reply.ReportZero,
    Count: reply.Count,
    ReportList: util.BulkProto2Biz[*v1.GlobalDetailDoctorSurgeryNumSortReply_SecondItem, *biz.GlobalDetailDoctorSurgeryNumSortReplySecondItem](reply.ReportList),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailDepts(ctx context.Context, param *biz.GlobalDetailDeptsRequestParam) (*biz.GlobalDetailDeptsReplyParam, error) {
  reply, err := r.performance.GlobalDetailDepts(ctx, &performanceV1.GlobalDetailDeptsRequest{
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
  return &biz.GlobalDetailDeptsReplyParam{
    Unqualified: reply.Unqualified,
    List: util.BulkProto2Biz[*v1.GlobalDetailDeptsReply_Depts, *biz.GlobalDetailDeptsReplyDepts](reply.List),
    Count: reply.Count,
    Qualified: reply.Qualified,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailSingleDiseaseDepts(ctx context.Context, param *biz.GlobalDetailSingleDiseaseDeptsRequestParam) (*biz.GlobalDetailSingleDiseaseDeptsReplyParam, error) {
  reply, err := r.performance.GlobalDetailSingleDiseaseDepts(ctx, &performanceV1.GlobalDetailSingleDiseaseDeptsRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    Indicator: param.Indicator,
    Sort: param.Sort,
    Page: param.Page,
    PageSize: param.PageSize,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailSingleDiseaseDeptsReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailSingleDiseaseDeptsReply_Item, *biz.GlobalDetailSingleDiseaseDeptsReplyItem](reply.List),
    Count: reply.Count,
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailTumourMain(ctx context.Context, param *biz.GlobalDetailTumourMainRequestParam) (*biz.GlobalDetailTumourMainReplyParam, error) {
  reply, err := r.performance.GlobalDetailTumourMain(ctx, &performanceV1.GlobalDetailTumourMainRequest{
    OrgId: param.OrgId,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailTumourMainReplyParam{
    List: util.BulkProto2Biz[*v1.GlobalDetailTumourMainReply_Item, *biz.GlobalDetailTumourMainReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalDetailDoctorGroup(ctx context.Context, param *biz.GlobalDetailDoctorGroupRequestParam) (*biz.GlobalDetailDoctorGroupReplyParam, error) {
  reply, err := r.performance.GlobalDetailDoctorGroup(ctx, &performanceV1.GlobalDetailDoctorGroupRequest{
    Id: param.Id,
    Page: param.Page,
    PageSize: param.PageSize,
    Sort: param.Sort,
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalDetailDoctorGroupReplyParam{
    Count: reply.Count,
    List: util.BulkProto2Biz[*v1.GlobalDetailDoctorGroupReply_Item, *biz.GlobalDetailDoctorGroupReplyItem](reply.List),
  }, nil
}

func (r IndicatorValueServiceRepo) GlobalList(ctx context.Context, param *biz.GlobalListRequestParam) (*biz.GlobalListReplyParam, error) {
  reply, err := r.performance.GlobalList(ctx, &performanceV1.GlobalListRequest{
    BeginDate: param.BeginDate,
    EndDate: param.EndDate,
    OrgCode: param.OrgCode,
  })
  if err != nil {
    return nil, err
  }
  return &biz.GlobalListReplyParam{
    Count: reply.Count,
    SingleDiseases: util.BulkProto2Biz[*v1.GlobalListReply_BadIndicatorStat, *biz.GlobalListReplyBadIndicatorStat](reply.SingleDiseases),
    BadCount: reply.BadCount,
    Tumours: util.BulkProto2Biz[*v1.GlobalListReply_BadIndicatorStat, *biz.GlobalListReplyBadIndicatorStat](reply.Tumours),
    BadNames: reply.BadNames,
    NiceCount: reply.NiceCount,
    List: util.BulkProto2Biz[*v1.GlobalListReply_Global, *biz.GlobalListReplyGlobal](reply.List),
  }, nil
}

