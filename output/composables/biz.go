// Biz Layer for IndicatorValueService
package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
  "google.golang.org/protobuf/types/known/structpb"

)

// GlobalDetailSingleDiseaseMainRequestParam struct represents the request parameters for GlobalDetailSingleDiseaseMain API
type GlobalDetailSingleDiseaseMainRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailDeptSurgeryNumSortRequestParam struct represents the request parameters for GlobalDetailDeptSurgeryNumSort API
type GlobalDetailDeptSurgeryNumSortRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Desc *bool
  Page *int64
  Pagesize *int64
}

// GlobalDetailDoctorSurgeryNumSortRequestParam struct represents the request parameters for GlobalDetailDoctorSurgeryNumSort API
type GlobalDetailDoctorSurgeryNumSortRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Desc *bool
  Page *int64
  Pagesize *int64
}

// GlobalDetailTumourDoctorGroupsRequestParam struct represents the request parameters for GlobalDetailTumourDoctorGroups API
type GlobalDetailTumourDoctorGroupsRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
  Indicator *string
  Sort *string
  Page *int32
  Pagesize *int32
}

// GlobalDetailDoctorsRequestParam struct represents the request parameters for GlobalDetailDoctors API
type GlobalDetailDoctorsRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Page *int32
  Pagesize *int32
  Desc *bool
}

// GlobalDetailCategoryIncomeDetailRequestParam struct represents the request parameters for GlobalDetailCategoryIncomeDetail API
type GlobalDetailCategoryIncomeDetailRequestParam struct {
  Category int32
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailTumourMainRequestParam struct represents the request parameters for GlobalDetailTumourMain API
type GlobalDetailTumourMainRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailDoctorsSortRequestParam struct represents the request parameters for GlobalDetailDoctorsSort API
type GlobalDetailDoctorsSortRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Page *int32
  Pagesize *int32
  Desc *bool
}

// GlobalDetailGroupsRequestParam struct represents the request parameters for GlobalDetailGroups API
type GlobalDetailGroupsRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Page *int32
  Pagesize *int32
  Desc *bool
}

// GlobalDetailCategoryIncomesRequestParam struct represents the request parameters for GlobalDetailCategoryIncomes API
type GlobalDetailCategoryIncomesRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailSurgeryRequestParam struct represents the request parameters for GlobalDetailSurgery API
type GlobalDetailSurgeryRequestParam struct {
  Id int32
  Page *int64
  Pagesize *int64
  Sort *string
  Begindate *string
  Enddate *string
}

// GlobalDetailDeptsRequestParam struct represents the request parameters for GlobalDetailDepts API
type GlobalDetailDeptsRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Page *int32
  Pagesize *int32
  Desc *bool
}

// GlobalListRequestParam struct represents the request parameters for GlobalList API
type GlobalListRequestParam struct {
  Begindate *string
  Enddate *string
  Orgcode *string
}

// GlobalDetailChildrenIndicatorRequestParam struct represents the request parameters for GlobalDetailChildrenIndicator API
type GlobalDetailChildrenIndicatorRequestParam struct {
  Id int32
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailCategoryIncomeDeptDetailRequestParam struct represents the request parameters for GlobalDetailCategoryIncomeDeptDetail API
type GlobalDetailCategoryIncomeDeptDetailRequestParam struct {
  Category int32
  Orgid *int32
  Begindate *string
  Enddate *string
  Sort *string
  Page *int32
  Pagesize *int32
}

// GlobalDetailMainRequestParam struct represents the request parameters for GlobalDetailMain API
type GlobalDetailMainRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailDoctorGroupRequestParam struct represents the request parameters for GlobalDetailDoctorGroup API
type GlobalDetailDoctorGroupRequestParam struct {
  Id int32
  Page *int64
  Pagesize *int64
  Sort *string
  Begindate *string
  Enddate *string
}

// GlobalDetailNumberatorDeptsRequestParam struct represents the request parameters for GlobalDetailNumberatorDepts API
type GlobalDetailNumberatorDeptsRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Page *int32
  Pagesize *int32
  Desc *bool
  Onlyunqualified *bool
}

// GlobalDetailDoctorMainRequestParam struct represents the request parameters for GlobalDetailDoctorMain API
type GlobalDetailDoctorMainRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailTumourDeptsRequestParam struct represents the request parameters for GlobalDetailTumourDepts API
type GlobalDetailTumourDeptsRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
  Indicator *string
  Sort *string
  Page *int32
  Pagesize *int32
}

// GlobalDetailNumberatorDgsRequestParam struct represents the request parameters for GlobalDetailNumberatorDgs API
type GlobalDetailNumberatorDgsRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
  Page *int32
  Pagesize *int32
  Desc *bool
  Onlyunqualified *bool
}

// GlobalDetailSingleDiseaseDoctorGroupsRequestParam struct represents the request parameters for GlobalDetailSingleDiseaseDoctorGroups API
type GlobalDetailSingleDiseaseDoctorGroupsRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
  Indicator *string
  Sort *string
  Page *int32
  Pagesize *int32
}

// GlobalDetailMonthSummaryRequestParam struct represents the request parameters for GlobalDetailMonthSummary API
type GlobalDetailMonthSummaryRequestParam struct {
  Id string
  Orgid *int32
  Begindate *string
  Enddate *string
}

// GlobalDetailSingleDiseaseDeptsRequestParam struct represents the request parameters for GlobalDetailSingleDiseaseDepts API
type GlobalDetailSingleDiseaseDeptsRequestParam struct {
  Orgid *int32
  Begindate *string
  Enddate *string
  Indicator *string
  Sort *string
  Page *int32
  Pagesize *int32
}


// GlobalDetailSingleDiseaseMainReplyParam struct represents the reply parameters for GlobalDetailSingleDiseaseMain API
type GlobalDetailSingleDiseaseMainReplyParam struct {
  List []GlobalDetailSingleDiseaseMainReplyItem
}

// GlobalDetailDeptSurgeryNumSortReplyParam struct represents the reply parameters for GlobalDetailDeptSurgeryNumSort API
type GlobalDetailDeptSurgeryNumSortReplyParam struct {
  List []GlobalDetailDeptSurgeryNumSortReplyItem
  Count int64
}

// GlobalDetailDoctorSurgeryNumSortReplyParam struct represents the reply parameters for GlobalDetailDoctorSurgeryNumSort API
type GlobalDetailDoctorSurgeryNumSortReplyParam struct {
  Reportlist []GlobalDetailDoctorSurgeryNumSortReplySecondItem
  Reportzero int64
  Count int64
  List []GlobalDetailDoctorSurgeryNumSortReplyItem
}

// GlobalDetailTumourDoctorGroupsReplyParam struct represents the reply parameters for GlobalDetailTumourDoctorGroups API
type GlobalDetailTumourDoctorGroupsReplyParam struct {
  Count int32
  List []GlobalDetailTumourDoctorGroupsReplyItem
}

// GlobalDetailDoctorsReplyParam struct represents the reply parameters for GlobalDetailDoctors API
type GlobalDetailDoctorsReplyParam struct {
  Count int32
  List []GlobalDetailDoctorsReplyDoctors
}

// GlobalDetailCategoryIncomeDetailReplyParam struct represents the reply parameters for GlobalDetailCategoryIncomeDetail API
type GlobalDetailCategoryIncomeDetailReplyParam struct {
  List []GlobalDetailCategoryIncomeDetailReplyItem
}

// GlobalDetailTumourMainReplyParam struct represents the reply parameters for GlobalDetailTumourMain API
type GlobalDetailTumourMainReplyParam struct {
  List []GlobalDetailTumourMainReplyItem
}

// GlobalDetailDoctorsSortReplyParam struct represents the reply parameters for GlobalDetailDoctorsSort API
type GlobalDetailDoctorsSortReplyParam struct {
  List []GlobalDetailDoctorsSortReplyItem
  Count int32
}

// GlobalDetailGroupsReplyParam struct represents the reply parameters for GlobalDetailGroups API
type GlobalDetailGroupsReplyParam struct {
  List []GlobalDetailGroupsReplyGroups
  Qualified int32
  Unqualified int32
  Count int32
}

// GlobalDetailCategoryIncomesReplyParam struct represents the reply parameters for GlobalDetailCategoryIncomes API
type GlobalDetailCategoryIncomesReplyParam struct {
  Types []GlobalDetailCategoryIncomesReplyItem
  Categories []GlobalDetailCategoryIncomesReplyItem
}

// GlobalDetailSurgeryReplyParam struct represents the reply parameters for GlobalDetailSurgery API
type GlobalDetailSurgeryReplyParam struct {
  Count int32
  List []GlobalDetailSurgeryReplyItem
}

// GlobalDetailDeptsReplyParam struct represents the reply parameters for GlobalDetailDepts API
type GlobalDetailDeptsReplyParam struct {
  List []GlobalDetailDeptsReplyDepts
  Unqualified int32
  Count int32
  Qualified int32
}

// GlobalListReplyParam struct represents the reply parameters for GlobalList API
type GlobalListReplyParam struct {
  Badnames []string
  List []GlobalListReplyGlobal
  Count int32
  Singlediseases []GlobalListReplyBadIndicatorStat
  Badcount int32
  Nicecount int32
  Tumours []GlobalListReplyBadIndicatorStat
}





// GlobalDetailMainReplyParam struct represents the reply parameters for GlobalDetailMain API
type GlobalDetailMainReplyParam struct {
  Targetnumberator float64
  Target float64
  Pastdenominator float64
  Curnumberator float64
  Curdenominator float64
  Diff float64
  Pastnumberator float64
  Diffnumberator float64
  Targetdenominator float64
  Past float64
  Cur float64
}

// GlobalDetailDoctorGroupReplyParam struct represents the reply parameters for GlobalDetailDoctorGroup API
type GlobalDetailDoctorGroupReplyParam struct {
  List []GlobalDetailDoctorGroupReplyItem
  Count int32
}





// GlobalDetailTumourDeptsReplyParam struct represents the reply parameters for GlobalDetailTumourDepts API
type GlobalDetailTumourDeptsReplyParam struct {
  List []GlobalDetailTumourDeptsReplyItem
  Count int32
}



// GlobalDetailSingleDiseaseDoctorGroupsReplyParam struct represents the reply parameters for GlobalDetailSingleDiseaseDoctorGroups API
type GlobalDetailSingleDiseaseDoctorGroupsReplyParam struct {
  List []GlobalDetailSingleDiseaseDoctorGroupsReplyItem
  Count int32
}

// GlobalDetailMonthSummaryReplyParam struct represents the reply parameters for GlobalDetailMonthSummary API
type GlobalDetailMonthSummaryReplyParam struct {
  List []GlobalDetailMonthSummaryReplyTuple
}

// GlobalDetailSingleDiseaseDeptsReplyParam struct represents the reply parameters for GlobalDetailSingleDiseaseDepts API
type GlobalDetailSingleDiseaseDeptsReplyParam struct {
  Count int32
  List []GlobalDetailSingleDiseaseDeptsReplyItem
}



type GlobalDetailCategoryIncomeDetailReplyItem struct {
  Indicatorname  string
  Value  float64
  Indicatorid  int32
  Target  float64
  Diff  float64
  Guide  string
}

type GlobalDetailCategoryIncomesReplyItem struct {
  Name  string
  Value  float64
}

type GlobalDetailDeptSurgeryNumSortReplyItem struct {
  Deptcode  string
  Deptname  string
  Deptid  int32
  Surgerynum  int64
  Lvfoursurgerynum  int64
  Miniopsurgerynum  int64
}

type GlobalDetailDeptsReplyDepts struct {
  Diff  float64
  Deptid  int32
  Actualvalue  float64
  Targetvalue  float64
  Deptname  string
}

type GlobalDetailDoctorGroupReplyItem struct {
  Surgerynum  int64
  Miniopsurgerynum  int64
  Dgid  int32
  Dgcode  string
  Dgname  string
  Lvfoursurgerynum  int64
}

type GlobalDetailDoctorSurgeryNumSortReplyItem struct {
  Lastyearsurgerynum  int64
  Surgerynum  int64
  Miniopsurgerynum  int64
  Doctorid  int32
  Deptname  string
  Doctorname  string
  Lvfoursurgerynum  int64
  Doctorcode  string
}

type GlobalDetailDoctorSurgeryNumSortReplySecondItem struct {
  Doctorcode  string
  Doctorname  string
  Doctorid  int32
  Percentage  float64
  Deptname  string
}

type GlobalDetailDoctorsReplyDoctors struct {
  Doctorid  int32
  Doctorname  string
  Targetvalue  float64
  Actualvalue  float64
  Diff  float64
}

type GlobalDetailDoctorsSortReplyItem struct {
  Value  float64
  Code  string
  Deptname  string
  Name  string
}

type GlobalDetailGroupsReplyGroups struct {
  Diff  float64
  Dgname  string
  Actualvalue  float64
  Dgid  int32
  Targetvalue  float64
}

type GlobalDetailMonthSummaryReplyTuple struct {
  Ratiodiff  float64
  Sizediff  float64
  Sizevalue  float64
  Ratiotarget  float64
  Sizetarget  float64
  Ratiovalue  float64
  Date  string
}

type GlobalDetailSingleDiseaseDeptsReplyItem struct {
  Target  float64
  Deptname  string
  Singlediseasename  string
  Value  float64
  Diff  float64
  Deptid  int32
}

type GlobalDetailSingleDiseaseDoctorGroupsReplyItem struct {
  Diff  float64
  Target  float64
  Singlediseasename  string
  Value  float64
  Dgname  string
  Dgid  int32
}

type GlobalDetailSingleDiseaseMainReplyItem struct {
  Singlediseasename  string
  Sizevalue  float64
  Avgstaydiff  float64
  Sizediff  float64
  Avgstaytarget  float64
  Avgstayvalue  float64
  Sizetarget  float64
  Deadvalue  float64
  Deadtarget  float64
  Deaddiff  float64
  Avgpricevalue  float64
  Avgpricetarget  float64
  Avgpricediff  float64
}

type GlobalDetailSurgeryReplyItem struct {
  Lastyearvalue  float64
  Value  float64
  Name  string
  Code  string
}

type GlobalDetailTumourDeptsReplyItem struct {
  Diff  float64
  Name  string
  Deptname  string
  Target  float64
  Value  float64
  Deptid  int32
}

type GlobalDetailTumourDoctorGroupsReplyItem struct {
  Target  float64
  Diff  float64
  Value  float64
  Dgid  int32
  Dgname  string
  Name  string
}

type GlobalDetailTumourMainReplyItem struct {
  Diff5  float64
  Name  string
  Diff2  float64
  Target2  float64
  Target4  float64
  Value5  float64
  Target3  float64
  Value3  float64
  Value6  float64
  Target5  float64
  Diff6  float64
  Target6  float64
  Value1  float64
  Diff1  float64
  Value2  float64
  Target1  float64
  Diff3  float64
  Diff4  float64
  Value4  float64
}

type GlobalListReplyBadIndicatorStat struct {
  Diseasenames  []string
  Name  string
  Size  int32
}

type GlobalListReplyGlobal struct {
  Showfactor  int32
  Category3  string
  Mastershowflag  int32
  Submastername  string
  Id  int32
  Showdetails  int32
  Targetvalue  float64
  Numberatordiff  float64
  Denominatorvalue  float64
  Isimport  bool
  Numberatorvalue  float64
  Masterseq  int32
  Category1  string
  Numberatorname  string
  Type  int32
  Numberatorsql  string
  Mastername  string
  Diff  float64
  Value  float64
  Listtype  int32
  Guide  string
  Seq  float64
  Showflag  int32
  Masterid  int32
  Denominatorname  string
  Mastermergeflag  int32
  Denominatorsql  string
  Submasterid  int32
  Attr  string
  Name  string
  Valuetype  int32
  Numberatortarget  float64
  Category2  string
  Masterscore  int32
  Pagedetail  *structpb.Struct
  Description  string
  Morecontents  string
}


type IndicatorValueServiceRepo interface {
  GlobalDetailSingleDiseaseMain(ctx context.Context, param *GlobalDetailSingleDiseaseMainRequestParam) (*GlobalDetailSingleDiseaseMainReplyParam, error)
  GlobalDetailDeptSurgeryNumSort(ctx context.Context, param *GlobalDetailDeptSurgeryNumSortRequestParam) (*GlobalDetailDeptSurgeryNumSortReplyParam, error)
  GlobalDetailDoctorSurgeryNumSort(ctx context.Context, param *GlobalDetailDoctorSurgeryNumSortRequestParam) (*GlobalDetailDoctorSurgeryNumSortReplyParam, error)
  GlobalDetailTumourDoctorGroups(ctx context.Context, param *GlobalDetailTumourDoctorGroupsRequestParam) (*GlobalDetailTumourDoctorGroupsReplyParam, error)
  GlobalDetailDoctors(ctx context.Context, param *GlobalDetailDoctorsRequestParam) (*GlobalDetailDoctorsReplyParam, error)
  GlobalDetailCategoryIncomeDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDetailRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error)
  GlobalDetailTumourMain(ctx context.Context, param *GlobalDetailTumourMainRequestParam) (*GlobalDetailTumourMainReplyParam, error)
  GlobalDetailDoctorsSort(ctx context.Context, param *GlobalDetailDoctorsSortRequestParam) (*GlobalDetailDoctorsSortReplyParam, error)
  GlobalDetailGroups(ctx context.Context, param *GlobalDetailGroupsRequestParam) (*GlobalDetailGroupsReplyParam, error)
  GlobalDetailCategoryIncomes(ctx context.Context, param *GlobalDetailCategoryIncomesRequestParam) (*GlobalDetailCategoryIncomesReplyParam, error)
  GlobalDetailSurgery(ctx context.Context, param *GlobalDetailSurgeryRequestParam) (*GlobalDetailSurgeryReplyParam, error)
  GlobalDetailDepts(ctx context.Context, param *GlobalDetailDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error)
  GlobalList(ctx context.Context, param *GlobalListRequestParam) (*GlobalListReplyParam, error)
  GlobalDetailChildrenIndicator(ctx context.Context, param *GlobalDetailChildrenIndicatorRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error)
  GlobalDetailCategoryIncomeDeptDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDeptDetailRequestParam) (*GlobalDetailDeptsReplyParam, error)
  GlobalDetailMain(ctx context.Context, param *GlobalDetailMainRequestParam) (*GlobalDetailMainReplyParam, error)
  GlobalDetailDoctorGroup(ctx context.Context, param *GlobalDetailDoctorGroupRequestParam) (*GlobalDetailDoctorGroupReplyParam, error)
  GlobalDetailNumberatorDepts(ctx context.Context, param *GlobalDetailNumberatorDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error)
  GlobalDetailDoctorMain(ctx context.Context, param *GlobalDetailDoctorMainRequestParam) (*GlobalListReplyParam, error)
  GlobalDetailTumourDepts(ctx context.Context, param *GlobalDetailTumourDeptsRequestParam) (*GlobalDetailTumourDeptsReplyParam, error)
  GlobalDetailNumberatorDgs(ctx context.Context, param *GlobalDetailNumberatorDgsRequestParam) (*GlobalDetailGroupsReplyParam, error)
  GlobalDetailSingleDiseaseDoctorGroups(ctx context.Context, param *GlobalDetailSingleDiseaseDoctorGroupsRequestParam) (*GlobalDetailSingleDiseaseDoctorGroupsReplyParam, error)
  GlobalDetailMonthSummary(ctx context.Context, param *GlobalDetailMonthSummaryRequestParam) (*GlobalDetailMonthSummaryReplyParam, error)
  GlobalDetailSingleDiseaseDepts(ctx context.Context, param *GlobalDetailSingleDiseaseDeptsRequestParam) (*GlobalDetailSingleDiseaseDeptsReplyParam, error)
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


func (uc *IndicatorValueServiceUsecase) GlobalDetailSingleDiseaseMain(ctx context.Context, param *GlobalDetailSingleDiseaseMainRequestParam) (*GlobalDetailSingleDiseaseMainReplyParam, error) {
	return uc.repo.GlobalDetailSingleDiseaseMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDeptSurgeryNumSort(ctx context.Context, param *GlobalDetailDeptSurgeryNumSortRequestParam) (*GlobalDetailDeptSurgeryNumSortReplyParam, error) {
	return uc.repo.GlobalDetailDeptSurgeryNumSort(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorSurgeryNumSort(ctx context.Context, param *GlobalDetailDoctorSurgeryNumSortRequestParam) (*GlobalDetailDoctorSurgeryNumSortReplyParam, error) {
	return uc.repo.GlobalDetailDoctorSurgeryNumSort(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailTumourDoctorGroups(ctx context.Context, param *GlobalDetailTumourDoctorGroupsRequestParam) (*GlobalDetailTumourDoctorGroupsReplyParam, error) {
	return uc.repo.GlobalDetailTumourDoctorGroups(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctors(ctx context.Context, param *GlobalDetailDoctorsRequestParam) (*GlobalDetailDoctorsReplyParam, error) {
	return uc.repo.GlobalDetailDoctors(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailCategoryIncomeDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDetailRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error) {
	return uc.repo.GlobalDetailCategoryIncomeDetail(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailTumourMain(ctx context.Context, param *GlobalDetailTumourMainRequestParam) (*GlobalDetailTumourMainReplyParam, error) {
	return uc.repo.GlobalDetailTumourMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorsSort(ctx context.Context, param *GlobalDetailDoctorsSortRequestParam) (*GlobalDetailDoctorsSortReplyParam, error) {
	return uc.repo.GlobalDetailDoctorsSort(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailGroups(ctx context.Context, param *GlobalDetailGroupsRequestParam) (*GlobalDetailGroupsReplyParam, error) {
	return uc.repo.GlobalDetailGroups(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailCategoryIncomes(ctx context.Context, param *GlobalDetailCategoryIncomesRequestParam) (*GlobalDetailCategoryIncomesReplyParam, error) {
	return uc.repo.GlobalDetailCategoryIncomes(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailSurgery(ctx context.Context, param *GlobalDetailSurgeryRequestParam) (*GlobalDetailSurgeryReplyParam, error) {
	return uc.repo.GlobalDetailSurgery(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDepts(ctx context.Context, param *GlobalDetailDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error) {
	return uc.repo.GlobalDetailDepts(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalList(ctx context.Context, param *GlobalListRequestParam) (*GlobalListReplyParam, error) {
	return uc.repo.GlobalList(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailChildrenIndicator(ctx context.Context, param *GlobalDetailChildrenIndicatorRequestParam) (*GlobalDetailCategoryIncomeDetailReplyParam, error) {
	return uc.repo.GlobalDetailChildrenIndicator(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailCategoryIncomeDeptDetail(ctx context.Context, param *GlobalDetailCategoryIncomeDeptDetailRequestParam) (*GlobalDetailDeptsReplyParam, error) {
	return uc.repo.GlobalDetailCategoryIncomeDeptDetail(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailMain(ctx context.Context, param *GlobalDetailMainRequestParam) (*GlobalDetailMainReplyParam, error) {
	return uc.repo.GlobalDetailMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorGroup(ctx context.Context, param *GlobalDetailDoctorGroupRequestParam) (*GlobalDetailDoctorGroupReplyParam, error) {
	return uc.repo.GlobalDetailDoctorGroup(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailNumberatorDepts(ctx context.Context, param *GlobalDetailNumberatorDeptsRequestParam) (*GlobalDetailDeptsReplyParam, error) {
	return uc.repo.GlobalDetailNumberatorDepts(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailDoctorMain(ctx context.Context, param *GlobalDetailDoctorMainRequestParam) (*GlobalListReplyParam, error) {
	return uc.repo.GlobalDetailDoctorMain(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailTumourDepts(ctx context.Context, param *GlobalDetailTumourDeptsRequestParam) (*GlobalDetailTumourDeptsReplyParam, error) {
	return uc.repo.GlobalDetailTumourDepts(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailNumberatorDgs(ctx context.Context, param *GlobalDetailNumberatorDgsRequestParam) (*GlobalDetailGroupsReplyParam, error) {
	return uc.repo.GlobalDetailNumberatorDgs(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailSingleDiseaseDoctorGroups(ctx context.Context, param *GlobalDetailSingleDiseaseDoctorGroupsRequestParam) (*GlobalDetailSingleDiseaseDoctorGroupsReplyParam, error) {
	return uc.repo.GlobalDetailSingleDiseaseDoctorGroups(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailMonthSummary(ctx context.Context, param *GlobalDetailMonthSummaryRequestParam) (*GlobalDetailMonthSummaryReplyParam, error) {
	return uc.repo.GlobalDetailMonthSummary(ctx, param)
}

func (uc *IndicatorValueServiceUsecase) GlobalDetailSingleDiseaseDepts(ctx context.Context, param *GlobalDetailSingleDiseaseDeptsRequestParam) (*GlobalDetailSingleDiseaseDeptsReplyParam, error) {
	return uc.repo.GlobalDetailSingleDiseaseDepts(ctx, param)
}

