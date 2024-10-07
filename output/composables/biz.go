// Biz Layer for IndicatorValueService
package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
  "google.golang.org/protobuf/types/known/structpb"

)

// GlobalDetailCategoryIncomeDeptDetailRequestParam struct represents the request parameters for GlobalDetailCategoryIncomeDeptDetail API
type GlobalDetailCategoryIncomeDeptDetailRequestParam struct {
  Category int32
  OrgId *int32
  BeginDate *string
  EndDate *string
  Sort *string
  Page *int32
  PageSize *int32
}

// GlobalDetailChildrenIndicatorRequestParam struct represents the request parameters for GlobalDetailChildrenIndicator API
type GlobalDetailChildrenIndicatorRequestParam struct {
  Id int32
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailDoctorMainRequestParam struct represents the request parameters for GlobalDetailDoctorMain API
type GlobalDetailDoctorMainRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailSurgeryRequestParam struct represents the request parameters for GlobalDetailSurgery API
type GlobalDetailSurgeryRequestParam struct {
  Id int32
  Page *int64
  PageSize *int64
  Sort *string
  BeginDate *string
  EndDate *string
}

// GlobalDetailSingleDiseaseMainRequestParam struct represents the request parameters for GlobalDetailSingleDiseaseMain API
type GlobalDetailSingleDiseaseMainRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailDoctorsRequestParam struct represents the request parameters for GlobalDetailDoctors API
type GlobalDetailDoctorsRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Page *int32
  PageSize *int32
  Desc *bool
}

// GlobalDetailGroupsRequestParam struct represents the request parameters for GlobalDetailGroups API
type GlobalDetailGroupsRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Page *int32
  PageSize *int32
  Desc *bool
}

// GlobalDetailDoctorsSortRequestParam struct represents the request parameters for GlobalDetailDoctorsSort API
type GlobalDetailDoctorsSortRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Page *int32
  PageSize *int32
  Desc *bool
}

// GlobalDetailNumberatorDgsRequestParam struct represents the request parameters for GlobalDetailNumberatorDgs API
type GlobalDetailNumberatorDgsRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Page *int32
  PageSize *int32
  Desc *bool
  OnlyUnqualified *bool
}

// GlobalDetailMonthSummaryRequestParam struct represents the request parameters for GlobalDetailMonthSummary API
type GlobalDetailMonthSummaryRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailCategoryIncomeDetailRequestParam struct represents the request parameters for GlobalDetailCategoryIncomeDetail API
type GlobalDetailCategoryIncomeDetailRequestParam struct {
  Category int32
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailDeptSurgeryNumSortRequestParam struct represents the request parameters for GlobalDetailDeptSurgeryNumSort API
type GlobalDetailDeptSurgeryNumSortRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Desc *bool
  Page *int64
  PageSize *int64
}

// GlobalDetailSingleDiseaseDoctorGroupsRequestParam struct represents the request parameters for GlobalDetailSingleDiseaseDoctorGroups API
type GlobalDetailSingleDiseaseDoctorGroupsRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
  Indicator *string
  Sort *string
  Page *int32
  PageSize *int32
}

// GlobalDetailTumourDeptsRequestParam struct represents the request parameters for GlobalDetailTumourDepts API
type GlobalDetailTumourDeptsRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
  Indicator *string
  Sort *string
  Page *int32
  PageSize *int32
}

// GlobalDetailMainRequestParam struct represents the request parameters for GlobalDetailMain API
type GlobalDetailMainRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailNumberatorDeptsRequestParam struct represents the request parameters for GlobalDetailNumberatorDepts API
type GlobalDetailNumberatorDeptsRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Page *int32
  PageSize *int32
  Desc *bool
  OnlyUnqualified *bool
}

// GlobalDetailCategoryIncomesRequestParam struct represents the request parameters for GlobalDetailCategoryIncomes API
type GlobalDetailCategoryIncomesRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailTumourDoctorGroupsRequestParam struct represents the request parameters for GlobalDetailTumourDoctorGroups API
type GlobalDetailTumourDoctorGroupsRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
  Indicator *string
  Sort *string
  Page *int32
  PageSize *int32
}

// GlobalDetailDoctorSurgeryNumSortRequestParam struct represents the request parameters for GlobalDetailDoctorSurgeryNumSort API
type GlobalDetailDoctorSurgeryNumSortRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Desc *bool
  Page *int64
  PageSize *int64
}

// GlobalDetailDeptsRequestParam struct represents the request parameters for GlobalDetailDepts API
type GlobalDetailDeptsRequestParam struct {
  Id string
  OrgId *int32
  BeginDate *string
  EndDate *string
  Page *int32
  PageSize *int32
  Desc *bool
}

// GlobalDetailSingleDiseaseDeptsRequestParam struct represents the request parameters for GlobalDetailSingleDiseaseDepts API
type GlobalDetailSingleDiseaseDeptsRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
  Indicator *string
  Sort *string
  Page *int32
  PageSize *int32
}

// GlobalDetailTumourMainRequestParam struct represents the request parameters for GlobalDetailTumourMain API
type GlobalDetailTumourMainRequestParam struct {
  OrgId *int32
  BeginDate *string
  EndDate *string
}

// GlobalDetailDoctorGroupRequestParam struct represents the request parameters for GlobalDetailDoctorGroup API
type GlobalDetailDoctorGroupRequestParam struct {
  Id int32
  Page *int64
  PageSize *int64
  Sort *string
  BeginDate *string
  EndDate *string
}

// GlobalListRequestParam struct represents the request parameters for GlobalList API
type GlobalListRequestParam struct {
  BeginDate *string
  EndDate *string
  OrgCode *string
}


// GlobalDetailDeptsReplyParam struct represents the reply parameters for GlobalDetailCategoryIncomeDeptDetail API
type GlobalDetailDeptsReplyParam struct {
  Unqualified int32
  List []GlobalDetailDeptsReplyDepts
  Count int32
  Qualified int32
}

// GlobalDetailCategoryIncomeDetailReplyParam struct represents the reply parameters for GlobalDetailChildrenIndicator API
type GlobalDetailCategoryIncomeDetailReplyParam struct {
  List []GlobalDetailCategoryIncomeDetailReplyItem
}

// GlobalListReplyParam struct represents the reply parameters for GlobalDetailDoctorMain API
type GlobalListReplyParam struct {
  Count int32
  SingleDiseases []GlobalListReplyBadIndicatorStat
  BadCount int32
  Tumours []GlobalListReplyBadIndicatorStat
  BadNames []string
  NiceCount int32
  List []GlobalListReplyGlobal
}

// GlobalDetailSurgeryReplyParam struct represents the reply parameters for GlobalDetailSurgery API
type GlobalDetailSurgeryReplyParam struct {
  List []GlobalDetailSurgeryReplyItem
  Count int32
}

// GlobalDetailSingleDiseaseMainReplyParam struct represents the reply parameters for GlobalDetailSingleDiseaseMain API
type GlobalDetailSingleDiseaseMainReplyParam struct {
  List []GlobalDetailSingleDiseaseMainReplyItem
}

// GlobalDetailDoctorsReplyParam struct represents the reply parameters for GlobalDetailDoctors API
type GlobalDetailDoctorsReplyParam struct {
  List []GlobalDetailDoctorsReplyDoctors
  Count int32
}

// GlobalDetailGroupsReplyParam struct represents the reply parameters for GlobalDetailGroups API
type GlobalDetailGroupsReplyParam struct {
  Unqualified int32
  Qualified int32
  List []GlobalDetailGroupsReplyGroups
  Count int32
}

// GlobalDetailDoctorsSortReplyParam struct represents the reply parameters for GlobalDetailDoctorsSort API
type GlobalDetailDoctorsSortReplyParam struct {
  List []GlobalDetailDoctorsSortReplyItem
  Count int32
}



// GlobalDetailMonthSummaryReplyParam struct represents the reply parameters for GlobalDetailMonthSummary API
type GlobalDetailMonthSummaryReplyParam struct {
  List []GlobalDetailMonthSummaryReplyTuple
}



// GlobalDetailDeptSurgeryNumSortReplyParam struct represents the reply parameters for GlobalDetailDeptSurgeryNumSort API
type GlobalDetailDeptSurgeryNumSortReplyParam struct {
  List []GlobalDetailDeptSurgeryNumSortReplyItem
  Count int64
}

// GlobalDetailSingleDiseaseDoctorGroupsReplyParam struct represents the reply parameters for GlobalDetailSingleDiseaseDoctorGroups API
type GlobalDetailSingleDiseaseDoctorGroupsReplyParam struct {
  Count int32
  List []GlobalDetailSingleDiseaseDoctorGroupsReplyItem
}

// GlobalDetailTumourDeptsReplyParam struct represents the reply parameters for GlobalDetailTumourDepts API
type GlobalDetailTumourDeptsReplyParam struct {
  Count int32
  List []GlobalDetailTumourDeptsReplyItem
}

// GlobalDetailMainReplyParam struct represents the reply parameters for GlobalDetailMain API
type GlobalDetailMainReplyParam struct {
  PastNumberator float64
  Cur float64
  TargetDenominator float64
  TargetNumberator float64
  CurNumberator float64
  Diff float64
  Target float64
  CurDenominator float64
  DiffNumberator float64
  PastDenominator float64
  Past float64
}



// GlobalDetailCategoryIncomesReplyParam struct represents the reply parameters for GlobalDetailCategoryIncomes API
type GlobalDetailCategoryIncomesReplyParam struct {
  Categories []GlobalDetailCategoryIncomesReplyItem
  Types []GlobalDetailCategoryIncomesReplyItem
}

// GlobalDetailTumourDoctorGroupsReplyParam struct represents the reply parameters for GlobalDetailTumourDoctorGroups API
type GlobalDetailTumourDoctorGroupsReplyParam struct {
  Count int32
  List []GlobalDetailTumourDoctorGroupsReplyItem
}

// GlobalDetailDoctorSurgeryNumSortReplyParam struct represents the reply parameters for GlobalDetailDoctorSurgeryNumSort API
type GlobalDetailDoctorSurgeryNumSortReplyParam struct {
  List []GlobalDetailDoctorSurgeryNumSortReplyItem
  ReportZero int64
  Count int64
  ReportList []GlobalDetailDoctorSurgeryNumSortReplySecondItem
}



// GlobalDetailSingleDiseaseDeptsReplyParam struct represents the reply parameters for GlobalDetailSingleDiseaseDepts API
type GlobalDetailSingleDiseaseDeptsReplyParam struct {
  List []GlobalDetailSingleDiseaseDeptsReplyItem
  Count int32
}

// GlobalDetailTumourMainReplyParam struct represents the reply parameters for GlobalDetailTumourMain API
type GlobalDetailTumourMainReplyParam struct {
  List []GlobalDetailTumourMainReplyItem
}

// GlobalDetailDoctorGroupReplyParam struct represents the reply parameters for GlobalDetailDoctorGroup API
type GlobalDetailDoctorGroupReplyParam struct {
  Count int32
  List []GlobalDetailDoctorGroupReplyItem
}





type GlobalDetailCategoryIncomeDetailReplyItem struct {
  IndicatorId  int32
  Value  float64
  Diff  float64
  IndicatorName  string
  Guide  string
  Target  float64
}

type GlobalDetailCategoryIncomesReplyItem struct {
  Name  string
  Value  float64
}

type GlobalDetailDeptSurgeryNumSortReplyItem struct {
  DeptId  int32
  DeptName  string
  DeptCode  string
  LvFourSurgeryNum  int64
  SurgeryNum  int64
  MiniOpSurgeryNum  int64
}

type GlobalDetailDeptsReplyDepts struct {
  ActualValue  float64
  TargetValue  float64
  DeptId  int32
  DeptName  string
  Diff  float64
}

type GlobalDetailDoctorGroupReplyItem struct {
  DgCode  string
  LvFourSurgeryNum  int64
  DgId  int32
  DgName  string
  SurgeryNum  int64
  MiniOpSurgeryNum  int64
}

type GlobalDetailDoctorSurgeryNumSortReplyItem struct {
  DeptName  string
  MiniOpSurgeryNum  int64
  DoctorId  int32
  LastYearSurgeryNum  int64
  LvFourSurgeryNum  int64
  DoctorCode  string
  DoctorName  string
  SurgeryNum  int64
}

type GlobalDetailDoctorSurgeryNumSortReplySecondItem struct {
  DoctorId  int32
  Percentage  float64
  DoctorName  string
  DoctorCode  string
  DeptName  string
}

type GlobalDetailDoctorsReplyDoctors struct {
  ActualValue  float64
  DoctorName  string
  DoctorId  int32
  Diff  float64
  TargetValue  float64
}

type GlobalDetailDoctorsSortReplyItem struct {
  Value  float64
  Name  string
  Code  string
  DeptName  string
}

type GlobalDetailGroupsReplyGroups struct {
  DgId  int32
  TargetValue  float64
  DgName  string
  ActualValue  float64
  Diff  float64
}

type GlobalDetailMonthSummaryReplyTuple struct {
  SizeDiff  float64
  SizeTarget  float64
  Date  string
  SizeValue  float64
  RatioTarget  float64
  RatioDiff  float64
  RatioValue  float64
}

type GlobalDetailSingleDiseaseDeptsReplyItem struct {
  DeptId  int32
  Target  float64
  Diff  float64
  SingleDiseaseName  string
  Value  float64
  DeptName  string
}

type GlobalDetailSingleDiseaseDoctorGroupsReplyItem struct {
  Target  float64
  Diff  float64
  DgName  string
  SingleDiseaseName  string
  DgId  int32
  Value  float64
}

type GlobalDetailSingleDiseaseMainReplyItem struct {
  SizeTarget  float64
  SizeValue  float64
  AvgStayTarget  float64
  SingleDiseaseName  string
  DeadValue  float64
  AvgStayValue  float64
  AvgPriceDiff  float64
  DeadTarget  float64
  SizeDiff  float64
  AvgPriceTarget  float64
  AvgPriceValue  float64
  DeadDiff  float64
  AvgStayDiff  float64
}

type GlobalDetailSurgeryReplyItem struct {
  Code  string
  Value  float64
  Name  string
  LastYearValue  float64
}

type GlobalDetailTumourDeptsReplyItem struct {
  Value  float64
  Target  float64
  Name  string
  Diff  float64
  DeptId  int32
  DeptName  string
}

type GlobalDetailTumourDoctorGroupsReplyItem struct {
  DgId  int32
  DgName  string
  Name  string
  Target  float64
  Value  float64
  Diff  float64
}

type GlobalDetailTumourMainReplyItem struct {
  Diff5  float64
  Target6  float64
  Diff2  float64
  Value4  float64
  Target1  float64
  Diff4  float64
  Target3  float64
  Value5  float64
  Diff6  float64
  Value1  float64
  Target5  float64
  Target4  float64
  Name  string
  Target2  float64
  Value3  float64
  Diff1  float64
  Diff3  float64
  Value6  float64
  Value2  float64
}

type GlobalListReplyBadIndicatorStat struct {
  Name  string
  Size  int32
  DiseaseNames  []string
}

type GlobalListReplyGlobal struct {
  ValueType  int32
  DenominatorSql  string
  Description  string
  Id  int32
  Category3  string
  Seq  float64
  NumberatorDiff  float64
  Name  string
  Attr  string
  NumberatorSql  string
  MasterShowFlag  int32
  ShowDetails  int32
  NumberatorName  string
  Category2  string
  TargetValue  float64
  DenominatorName  string
  ShowFlag  int32
  SubMasterId  int32
  MasterSeq  int32
  MoreContents  string
  NumberatorValue  float64
  Category1  string
  IsImport  bool
  MasterScore  int32
  MasterName  string
  MasterMergeFlag  int32
  DenominatorValue  float64
  ListType  int32
  SubMasterName  string
  ShowFactor  int32
  Diff  float64
  Guide  string
  Type  int32
  NumberatorTarget  float64
  Value  float64
  MasterId  int32
  PageDetail  *structpb.Struct
}


type IndicatorValueServiceRepo interface {
  GlobalDetailCategoryIncomeDeptDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDeptDetailRequestParam) (*GlobalDetailDeptsReplyParam, error)
  GlobalDetailChildrenIndicator(ctx context.Context, param *GlobalDetailChildrenIndicatorRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error)
  GlobalDetailDoctorMain(ctx context.Context, param *GlobalDetailDoctorMainRequestParam) (*GlobalListReplyParam, error)
  GlobalDetailSurgery(ctx context.Context, param *GlobalDetailSurgeryRequestParam) (*GlobalDetailSurgeryReplyParam, error)
  GlobalDetailSingleDiseaseMain(ctx context.Context, param *GlobalDetailSingleDiseaseMainRequestParam) (*GlobalDetailSingleDiseaseMainReplyParam, error)
  GlobalDetailDoctors(ctx context.Context, param *GlobalDetailDoctorsRequestParam) (*GlobalDetailDoctorsReplyParam, error)
  GlobalDetailGroups(ctx context.Context, param *GlobalDetailGroupsRequestParam) (*GlobalDetailGroupsReplyParam, error)
  GlobalDetailDoctorsSort(ctx context.Context, param *GlobalDetailDoctorsSortRequestParam) (*GlobalDetailDoctorsSortReplyParam, error)
  GlobalDetailNumberatorDgs(ctx context.Context, param *GlobalDetailNumberatorDgsRequestParam) (*GlobalDetailGroupsReplyParam, error)
  GlobalDetailMonthSummary(ctx context.Context, param *GlobalDetailMonthSummaryRequestParam) (*GlobalDetailMonthSummaryReplyParam, error)
  GlobalDetailCategoryIncomeDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDetailRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error)
  GlobalDetailDeptSurgeryNumSort(ctx context.Context, param *GlobalDetailDeptSurgeryNumSortRequestParam) (*GlobalDetailDeptSurgeryNumSortReplyParam, error)
  GlobalDetailSingleDiseaseDoctorGroups(ctx context.Context, param *GlobalDetailSingleDiseaseDoctorGroupsRequestParam) (*GlobalDetailSingleDiseaseDoctorGroupsReplyParam, error)
  GlobalDetailTumourDepts(ctx context.Context, param *GlobalDetailTumourDeptsRequestParam) (*GlobalDetailTumourDeptsReplyParam, error)
  GlobalDetailMain(ctx context.Context, param *GlobalDetailMainRequestParam) (*GlobalDetailMainReplyParam, error)
  GlobalDetailNumberatorDepts(ctx context.Context, param *GlobalDetailNumberatorDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error)
  GlobalDetailCategoryIncomes(ctx context.Context, param *GlobalDetailCategoryIncomesRequestParam) (*GlobalDetailCategoryIncomesReplyParam, error)
  GlobalDetailTumourDoctorGroups(ctx context.Context, param *GlobalDetailTumourDoctorGroupsRequestParam) (*GlobalDetailTumourDoctorGroupsReplyParam, error)
  GlobalDetailDoctorSurgeryNumSort(ctx context.Context, param *GlobalDetailDoctorSurgeryNumSortRequestParam) (*GlobalDetailDoctorSurgeryNumSortReplyParam, error)
  GlobalDetailDepts(ctx context.Context, param *GlobalDetailDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error)
  GlobalDetailSingleDiseaseDepts(ctx context.Context, param *GlobalDetailSingleDiseaseDeptsRequestParam) (*GlobalDetailSingleDiseaseDeptsReplyParam, error)
  GlobalDetailTumourMain(ctx context.Context, param *GlobalDetailTumourMainRequestParam) (*GlobalDetailTumourMainReplyParam, error)
  GlobalDetailDoctorGroup(ctx context.Context, param *GlobalDetailDoctorGroupRequestParam) (*GlobalDetailDoctorGroupReplyParam, error)
  GlobalList(ctx context.Context, param *GlobalListRequestParam) (*GlobalListReplyParam, error)
}

type IndicatorValueServiceUsecase struct {
	repo IndicatorValueServiceRepo
	log  *log.Helper
}

func NewIndicatorValueServiceUsecase(repo IndicatorValueServiceRepo, logger log.Logger) *IndicatorValueServiceUsecase {
	return &IndicatorValueServiceUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/indicator_value_service")),
	}
}


func (uc *IndicatorValueServiceUsecase) GlobalDetailCategoryIncomeDeptDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDeptDetailRequestParam) (*GlobalDetailDeptsReplyParam, error) {
	return uc.repo.GlobalDetailCategoryIncomeDeptDetail(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailChildrenIndicator(ctx context.Context, param *GlobalDetailChildrenIndicatorRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error) {
	return uc.repo.GlobalDetailChildrenIndicator(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorMain(ctx context.Context, param *GlobalDetailDoctorMainRequestParam) (*GlobalListReplyParam, error) {
	return uc.repo.GlobalDetailDoctorMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailSurgery(ctx context.Context, param *GlobalDetailSurgeryRequestParam) (*GlobalDetailSurgeryReplyParam, error) {
	return uc.repo.GlobalDetailSurgery(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailSingleDiseaseMain(ctx context.Context, param *GlobalDetailSingleDiseaseMainRequestParam) (*GlobalDetailSingleDiseaseMainReplyParam, error) {
	return uc.repo.GlobalDetailSingleDiseaseMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctors(ctx context.Context, param *GlobalDetailDoctorsRequestParam) (*GlobalDetailDoctorsReplyParam, error) {
	return uc.repo.GlobalDetailDoctors(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailGroups(ctx context.Context, param *GlobalDetailGroupsRequestParam) (*GlobalDetailGroupsReplyParam, error) {
	return uc.repo.GlobalDetailGroups(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorsSort(ctx context.Context, param *GlobalDetailDoctorsSortRequestParam) (*GlobalDetailDoctorsSortReplyParam, error) {
	return uc.repo.GlobalDetailDoctorsSort(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailNumberatorDgs(ctx context.Context, param *GlobalDetailNumberatorDgsRequestParam) (*GlobalDetailGroupsReplyParam, error) {
	return uc.repo.GlobalDetailNumberatorDgs(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailMonthSummary(ctx context.Context, param *GlobalDetailMonthSummaryRequestParam) (*GlobalDetailMonthSummaryReplyParam, error) {
	return uc.repo.GlobalDetailMonthSummary(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailCategoryIncomeDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDetailRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error) {
	return uc.repo.GlobalDetailCategoryIncomeDetail(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDeptSurgeryNumSort(ctx context.Context, param *GlobalDetailDeptSurgeryNumSortRequestParam) (*GlobalDetailDeptSurgeryNumSortReplyParam, error) {
	return uc.repo.GlobalDetailDeptSurgeryNumSort(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailSingleDiseaseDoctorGroups(ctx context.Context, param *GlobalDetailSingleDiseaseDoctorGroupsRequestParam) (*GlobalDetailSingleDiseaseDoctorGroupsReplyParam, error) {
	return uc.repo.GlobalDetailSingleDiseaseDoctorGroups(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailTumourDepts(ctx context.Context, param *GlobalDetailTumourDeptsRequestParam) (*GlobalDetailTumourDeptsReplyParam, error) {
	return uc.repo.GlobalDetailTumourDepts(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailMain(ctx context.Context, param *GlobalDetailMainRequestParam) (*GlobalDetailMainReplyParam, error) {
	return uc.repo.GlobalDetailMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailNumberatorDepts(ctx context.Context, param *GlobalDetailNumberatorDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error) {
	return uc.repo.GlobalDetailNumberatorDepts(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailCategoryIncomes(ctx context.Context, param *GlobalDetailCategoryIncomesRequestParam) (*GlobalDetailCategoryIncomesReplyParam, error) {
	return uc.repo.GlobalDetailCategoryIncomes(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailTumourDoctorGroups(ctx context.Context, param *GlobalDetailTumourDoctorGroupsRequestParam) (*GlobalDetailTumourDoctorGroupsReplyParam, error) {
	return uc.repo.GlobalDetailTumourDoctorGroups(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorSurgeryNumSort(ctx context.Context, param *GlobalDetailDoctorSurgeryNumSortRequestParam) (*GlobalDetailDoctorSurgeryNumSortReplyParam, error) {
	return uc.repo.GlobalDetailDoctorSurgeryNumSort(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDepts(ctx context.Context, param *GlobalDetailDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error) {
	return uc.repo.GlobalDetailDepts(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailSingleDiseaseDepts(ctx context.Context, param *GlobalDetailSingleDiseaseDeptsRequestParam) (*GlobalDetailSingleDiseaseDeptsReplyParam, error) {
	return uc.repo.GlobalDetailSingleDiseaseDepts(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailTumourMain(ctx context.Context, param *GlobalDetailTumourMainRequestParam) (*GlobalDetailTumourMainReplyParam, error) {
	return uc.repo.GlobalDetailTumourMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorGroup(ctx context.Context, param *GlobalDetailDoctorGroupRequestParam) (*GlobalDetailDoctorGroupReplyParam, error) {
	return uc.repo.GlobalDetailDoctorGroup(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalList(ctx context.Context, param *GlobalListRequestParam) (*GlobalListReplyParam, error) {
	return uc.repo.GlobalList(ctx, param)
}

