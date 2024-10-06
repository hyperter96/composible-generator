// api.ts

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { ref } from 'vue';

// 创建 Axios 实例
const apiClient: AxiosInstance = axios.create({
  baseURL: 'https://api.example.com', // 替换为你的API基础URL
  timeout: 10000, // 请求超时
});






export function globalDetailDoctorMain(params: IndicatorValueServiceGlobalDetailDoctorMainRequest) {
  const globalDetailDoctorMainData = ref<V1GlobalListReply | null>(null);
  const globalDetailDoctorMainError = ref<string | null>(null);
  const globalDetailDoctorMainLoading = ref<boolean>(false);

  async function globalDetailDoctorMainExecute() {
    globalDetailDoctorMainLoading.value = true;
    globalDetailDoctorMainError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailDoctorMain({
        ...params
      });
      globalDetailDoctorMainData.value = response.data;
    } catch (err: any) {
      globalDetailDoctorMainError.value = err.message || 'An error occurred';
    } finally {
      globalDetailDoctorMainLoading.value = false;
    }
  }

  return {
    globalDetailDoctorMainData,
    globalDetailDoctorMainError,
    globalDetailDoctorMainLoading,
    globalDetailDoctorMainExecute,
  };
}

export function globalDetailCategoryIncomes(params: IndicatorValueServiceGlobalDetailCategoryIncomesRequest) {
  const globalDetailCategoryIncomesData = ref<V1GlobalDetailCategoryIncomesReply | null>(null);
  const globalDetailCategoryIncomesError = ref<string | null>(null);
  const globalDetailCategoryIncomesLoading = ref<boolean>(false);

  async function globalDetailCategoryIncomesExecute() {
    globalDetailCategoryIncomesLoading.value = true;
    globalDetailCategoryIncomesError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailCategoryIncomes({
        ...params
      });
      globalDetailCategoryIncomesData.value = response.data;
    } catch (err: any) {
      globalDetailCategoryIncomesError.value = err.message || 'An error occurred';
    } finally {
      globalDetailCategoryIncomesLoading.value = false;
    }
  }

  return {
    globalDetailCategoryIncomesData,
    globalDetailCategoryIncomesError,
    globalDetailCategoryIncomesLoading,
    globalDetailCategoryIncomesExecute,
  };
}

export function globalDetailTumourMain(params: IndicatorValueServiceGlobalDetailTumourMainRequest) {
  const globalDetailTumourMainData = ref<V1GlobalDetailTumourMainReply | null>(null);
  const globalDetailTumourMainError = ref<string | null>(null);
  const globalDetailTumourMainLoading = ref<boolean>(false);

  async function globalDetailTumourMainExecute() {
    globalDetailTumourMainLoading.value = true;
    globalDetailTumourMainError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailTumourMain({
        ...params
      });
      globalDetailTumourMainData.value = response.data;
    } catch (err: any) {
      globalDetailTumourMainError.value = err.message || 'An error occurred';
    } finally {
      globalDetailTumourMainLoading.value = false;
    }
  }

  return {
    globalDetailTumourMainData,
    globalDetailTumourMainError,
    globalDetailTumourMainLoading,
    globalDetailTumourMainExecute,
  };
}

export function globalDetailSingleDiseaseDepts(params: IndicatorValueServiceGlobalDetailSingleDiseaseDeptsRequest) {
  const globalDetailSingleDiseaseDeptsData = ref<V1GlobalDetailSingleDiseaseDeptsReply | null>(null);
  const globalDetailSingleDiseaseDeptsError = ref<string | null>(null);
  const globalDetailSingleDiseaseDeptsLoading = ref<boolean>(false);

  async function globalDetailSingleDiseaseDeptsExecute() {
    globalDetailSingleDiseaseDeptsLoading.value = true;
    globalDetailSingleDiseaseDeptsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailSingleDiseaseDepts({
        ...params
      });
      globalDetailSingleDiseaseDeptsData.value = response.data;
    } catch (err: any) {
      globalDetailSingleDiseaseDeptsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailSingleDiseaseDeptsLoading.value = false;
    }
  }

  return {
    globalDetailSingleDiseaseDeptsData,
    globalDetailSingleDiseaseDeptsError,
    globalDetailSingleDiseaseDeptsLoading,
    globalDetailSingleDiseaseDeptsExecute,
  };
}

export function globalDetailNumberatorDepts(params: IndicatorValueServiceGlobalDetailNumberatorDeptsRequest) {
  const globalDetailNumberatorDeptsData = ref<V1GlobalDetailDeptsReply | null>(null);
  const globalDetailNumberatorDeptsError = ref<string | null>(null);
  const globalDetailNumberatorDeptsLoading = ref<boolean>(false);

  async function globalDetailNumberatorDeptsExecute() {
    globalDetailNumberatorDeptsLoading.value = true;
    globalDetailNumberatorDeptsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailNumberatorDepts({
        ...params
      });
      globalDetailNumberatorDeptsData.value = response.data;
    } catch (err: any) {
      globalDetailNumberatorDeptsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailNumberatorDeptsLoading.value = false;
    }
  }

  return {
    globalDetailNumberatorDeptsData,
    globalDetailNumberatorDeptsError,
    globalDetailNumberatorDeptsLoading,
    globalDetailNumberatorDeptsExecute,
  };
}

export function globalDetailDeptSurgeryNumSort(params: IndicatorValueServiceGlobalDetailDeptSurgeryNumSortRequest) {
  const globalDetailDeptSurgeryNumSortData = ref<V1GlobalDetailDeptSurgeryNumSortReply | null>(null);
  const globalDetailDeptSurgeryNumSortError = ref<string | null>(null);
  const globalDetailDeptSurgeryNumSortLoading = ref<boolean>(false);

  async function globalDetailDeptSurgeryNumSortExecute() {
    globalDetailDeptSurgeryNumSortLoading.value = true;
    globalDetailDeptSurgeryNumSortError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailDeptSurgeryNumSort({
        ...params
      });
      globalDetailDeptSurgeryNumSortData.value = response.data;
    } catch (err: any) {
      globalDetailDeptSurgeryNumSortError.value = err.message || 'An error occurred';
    } finally {
      globalDetailDeptSurgeryNumSortLoading.value = false;
    }
  }

  return {
    globalDetailDeptSurgeryNumSortData,
    globalDetailDeptSurgeryNumSortError,
    globalDetailDeptSurgeryNumSortLoading,
    globalDetailDeptSurgeryNumSortExecute,
  };
}

export function globalDetailCategoryIncomeDetail(params: IndicatorValueServiceGlobalDetailCategoryIncomeDetailRequest) {
  const globalDetailCategoryIncomeDetailData = ref<V1GlobalDetailCategoryIncomeDetailReply | null>(null);
  const globalDetailCategoryIncomeDetailError = ref<string | null>(null);
  const globalDetailCategoryIncomeDetailLoading = ref<boolean>(false);

  async function globalDetailCategoryIncomeDetailExecute() {
    globalDetailCategoryIncomeDetailLoading.value = true;
    globalDetailCategoryIncomeDetailError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailCategoryIncomeDetail({
        ...params
      });
      globalDetailCategoryIncomeDetailData.value = response.data;
    } catch (err: any) {
      globalDetailCategoryIncomeDetailError.value = err.message || 'An error occurred';
    } finally {
      globalDetailCategoryIncomeDetailLoading.value = false;
    }
  }

  return {
    globalDetailCategoryIncomeDetailData,
    globalDetailCategoryIncomeDetailError,
    globalDetailCategoryIncomeDetailLoading,
    globalDetailCategoryIncomeDetailExecute,
  };
}

export function globalList(params: IndicatorValueServiceGlobalListRequest) {
  const globalListData = ref<V1GlobalListReply | null>(null);
  const globalListError = ref<string | null>(null);
  const globalListLoading = ref<boolean>(false);

  async function globalListExecute() {
    globalListLoading.value = true;
    globalListError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalList({
        ...params
      });
      globalListData.value = response.data;
    } catch (err: any) {
      globalListError.value = err.message || 'An error occurred';
    } finally {
      globalListLoading.value = false;
    }
  }

  return {
    globalListData,
    globalListError,
    globalListLoading,
    globalListExecute,
  };
}

export function globalDetailTumourDoctorGroups(params: IndicatorValueServiceGlobalDetailTumourDoctorGroupsRequest) {
  const globalDetailTumourDoctorGroupsData = ref<V1GlobalDetailTumourDoctorGroupsReply | null>(null);
  const globalDetailTumourDoctorGroupsError = ref<string | null>(null);
  const globalDetailTumourDoctorGroupsLoading = ref<boolean>(false);

  async function globalDetailTumourDoctorGroupsExecute() {
    globalDetailTumourDoctorGroupsLoading.value = true;
    globalDetailTumourDoctorGroupsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailTumourDoctorGroups({
        ...params
      });
      globalDetailTumourDoctorGroupsData.value = response.data;
    } catch (err: any) {
      globalDetailTumourDoctorGroupsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailTumourDoctorGroupsLoading.value = false;
    }
  }

  return {
    globalDetailTumourDoctorGroupsData,
    globalDetailTumourDoctorGroupsError,
    globalDetailTumourDoctorGroupsLoading,
    globalDetailTumourDoctorGroupsExecute,
  };
}

export function globalDetailDoctors(params: IndicatorValueServiceGlobalDetailDoctorsRequest) {
  const globalDetailDoctorsData = ref<V1GlobalDetailDoctorsReply | null>(null);
  const globalDetailDoctorsError = ref<string | null>(null);
  const globalDetailDoctorsLoading = ref<boolean>(false);

  async function globalDetailDoctorsExecute() {
    globalDetailDoctorsLoading.value = true;
    globalDetailDoctorsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailDoctors({
        ...params
      });
      globalDetailDoctorsData.value = response.data;
    } catch (err: any) {
      globalDetailDoctorsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailDoctorsLoading.value = false;
    }
  }

  return {
    globalDetailDoctorsData,
    globalDetailDoctorsError,
    globalDetailDoctorsLoading,
    globalDetailDoctorsExecute,
  };
}

export function globalDetailDoctorSurgeryNumSort(params: IndicatorValueServiceGlobalDetailDoctorSurgeryNumSortRequest) {
  const globalDetailDoctorSurgeryNumSortData = ref<V1GlobalDetailDoctorSurgeryNumSortReply | null>(null);
  const globalDetailDoctorSurgeryNumSortError = ref<string | null>(null);
  const globalDetailDoctorSurgeryNumSortLoading = ref<boolean>(false);

  async function globalDetailDoctorSurgeryNumSortExecute() {
    globalDetailDoctorSurgeryNumSortLoading.value = true;
    globalDetailDoctorSurgeryNumSortError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailDoctorSurgeryNumSort({
        ...params
      });
      globalDetailDoctorSurgeryNumSortData.value = response.data;
    } catch (err: any) {
      globalDetailDoctorSurgeryNumSortError.value = err.message || 'An error occurred';
    } finally {
      globalDetailDoctorSurgeryNumSortLoading.value = false;
    }
  }

  return {
    globalDetailDoctorSurgeryNumSortData,
    globalDetailDoctorSurgeryNumSortError,
    globalDetailDoctorSurgeryNumSortLoading,
    globalDetailDoctorSurgeryNumSortExecute,
  };
}

export function globalDetailDoctorsSort(params: IndicatorValueServiceGlobalDetailDoctorsSortRequest) {
  const globalDetailDoctorsSortData = ref<V1GlobalDetailDoctorsSortReply | null>(null);
  const globalDetailDoctorsSortError = ref<string | null>(null);
  const globalDetailDoctorsSortLoading = ref<boolean>(false);

  async function globalDetailDoctorsSortExecute() {
    globalDetailDoctorsSortLoading.value = true;
    globalDetailDoctorsSortError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailDoctorsSort({
        ...params
      });
      globalDetailDoctorsSortData.value = response.data;
    } catch (err: any) {
      globalDetailDoctorsSortError.value = err.message || 'An error occurred';
    } finally {
      globalDetailDoctorsSortLoading.value = false;
    }
  }

  return {
    globalDetailDoctorsSortData,
    globalDetailDoctorsSortError,
    globalDetailDoctorsSortLoading,
    globalDetailDoctorsSortExecute,
  };
}

export function globalDetailSingleDiseaseMain(params: IndicatorValueServiceGlobalDetailSingleDiseaseMainRequest) {
  const globalDetailSingleDiseaseMainData = ref<V1GlobalDetailSingleDiseaseMainReply | null>(null);
  const globalDetailSingleDiseaseMainError = ref<string | null>(null);
  const globalDetailSingleDiseaseMainLoading = ref<boolean>(false);

  async function globalDetailSingleDiseaseMainExecute() {
    globalDetailSingleDiseaseMainLoading.value = true;
    globalDetailSingleDiseaseMainError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailSingleDiseaseMain({
        ...params
      });
      globalDetailSingleDiseaseMainData.value = response.data;
    } catch (err: any) {
      globalDetailSingleDiseaseMainError.value = err.message || 'An error occurred';
    } finally {
      globalDetailSingleDiseaseMainLoading.value = false;
    }
  }

  return {
    globalDetailSingleDiseaseMainData,
    globalDetailSingleDiseaseMainError,
    globalDetailSingleDiseaseMainLoading,
    globalDetailSingleDiseaseMainExecute,
  };
}

export function globalDetailDepts(params: IndicatorValueServiceGlobalDetailDeptsRequest) {
  const globalDetailDeptsData = ref<V1GlobalDetailDeptsReply | null>(null);
  const globalDetailDeptsError = ref<string | null>(null);
  const globalDetailDeptsLoading = ref<boolean>(false);

  async function globalDetailDeptsExecute() {
    globalDetailDeptsLoading.value = true;
    globalDetailDeptsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailDepts({
        ...params
      });
      globalDetailDeptsData.value = response.data;
    } catch (err: any) {
      globalDetailDeptsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailDeptsLoading.value = false;
    }
  }

  return {
    globalDetailDeptsData,
    globalDetailDeptsError,
    globalDetailDeptsLoading,
    globalDetailDeptsExecute,
  };
}

export function globalDetailChildrenIndicator(params: IndicatorValueServiceGlobalDetailChildrenIndicatorRequest) {
  const globalDetailChildrenIndicatorData = ref<V1GlobalDetailCategoryIncomeDetailReply | null>(null);
  const globalDetailChildrenIndicatorError = ref<string | null>(null);
  const globalDetailChildrenIndicatorLoading = ref<boolean>(false);

  async function globalDetailChildrenIndicatorExecute() {
    globalDetailChildrenIndicatorLoading.value = true;
    globalDetailChildrenIndicatorError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailChildrenIndicator({
        ...params
      });
      globalDetailChildrenIndicatorData.value = response.data;
    } catch (err: any) {
      globalDetailChildrenIndicatorError.value = err.message || 'An error occurred';
    } finally {
      globalDetailChildrenIndicatorLoading.value = false;
    }
  }

  return {
    globalDetailChildrenIndicatorData,
    globalDetailChildrenIndicatorError,
    globalDetailChildrenIndicatorLoading,
    globalDetailChildrenIndicatorExecute,
  };
}

export function globalDetailMonthSummary(params: IndicatorValueServiceGlobalDetailMonthSummaryRequest) {
  const globalDetailMonthSummaryData = ref<V1GlobalDetailMonthSummaryReply | null>(null);
  const globalDetailMonthSummaryError = ref<string | null>(null);
  const globalDetailMonthSummaryLoading = ref<boolean>(false);

  async function globalDetailMonthSummaryExecute() {
    globalDetailMonthSummaryLoading.value = true;
    globalDetailMonthSummaryError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailMonthSummary({
        ...params
      });
      globalDetailMonthSummaryData.value = response.data;
    } catch (err: any) {
      globalDetailMonthSummaryError.value = err.message || 'An error occurred';
    } finally {
      globalDetailMonthSummaryLoading.value = false;
    }
  }

  return {
    globalDetailMonthSummaryData,
    globalDetailMonthSummaryError,
    globalDetailMonthSummaryLoading,
    globalDetailMonthSummaryExecute,
  };
}

export function globalDetailSingleDiseaseDoctorGroups(params: IndicatorValueServiceGlobalDetailSingleDiseaseDoctorGroupsRequest) {
  const globalDetailSingleDiseaseDoctorGroupsData = ref<V1GlobalDetailSingleDiseaseDoctorGroupsReply | null>(null);
  const globalDetailSingleDiseaseDoctorGroupsError = ref<string | null>(null);
  const globalDetailSingleDiseaseDoctorGroupsLoading = ref<boolean>(false);

  async function globalDetailSingleDiseaseDoctorGroupsExecute() {
    globalDetailSingleDiseaseDoctorGroupsLoading.value = true;
    globalDetailSingleDiseaseDoctorGroupsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailSingleDiseaseDoctorGroups({
        ...params
      });
      globalDetailSingleDiseaseDoctorGroupsData.value = response.data;
    } catch (err: any) {
      globalDetailSingleDiseaseDoctorGroupsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailSingleDiseaseDoctorGroupsLoading.value = false;
    }
  }

  return {
    globalDetailSingleDiseaseDoctorGroupsData,
    globalDetailSingleDiseaseDoctorGroupsError,
    globalDetailSingleDiseaseDoctorGroupsLoading,
    globalDetailSingleDiseaseDoctorGroupsExecute,
  };
}

export function globalDetailNumberatorDgs(params: IndicatorValueServiceGlobalDetailNumberatorDgsRequest) {
  const globalDetailNumberatorDgsData = ref<V1GlobalDetailGroupsReply | null>(null);
  const globalDetailNumberatorDgsError = ref<string | null>(null);
  const globalDetailNumberatorDgsLoading = ref<boolean>(false);

  async function globalDetailNumberatorDgsExecute() {
    globalDetailNumberatorDgsLoading.value = true;
    globalDetailNumberatorDgsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailNumberatorDgs({
        ...params
      });
      globalDetailNumberatorDgsData.value = response.data;
    } catch (err: any) {
      globalDetailNumberatorDgsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailNumberatorDgsLoading.value = false;
    }
  }

  return {
    globalDetailNumberatorDgsData,
    globalDetailNumberatorDgsError,
    globalDetailNumberatorDgsLoading,
    globalDetailNumberatorDgsExecute,
  };
}

export function globalDetailTumourDepts(params: IndicatorValueServiceGlobalDetailTumourDeptsRequest) {
  const globalDetailTumourDeptsData = ref<V1GlobalDetailTumourDeptsReply | null>(null);
  const globalDetailTumourDeptsError = ref<string | null>(null);
  const globalDetailTumourDeptsLoading = ref<boolean>(false);

  async function globalDetailTumourDeptsExecute() {
    globalDetailTumourDeptsLoading.value = true;
    globalDetailTumourDeptsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailTumourDepts({
        ...params
      });
      globalDetailTumourDeptsData.value = response.data;
    } catch (err: any) {
      globalDetailTumourDeptsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailTumourDeptsLoading.value = false;
    }
  }

  return {
    globalDetailTumourDeptsData,
    globalDetailTumourDeptsError,
    globalDetailTumourDeptsLoading,
    globalDetailTumourDeptsExecute,
  };
}

export function globalDetailMain(params: IndicatorValueServiceGlobalDetailMainRequest) {
  const globalDetailMainData = ref<V1GlobalDetailMainReply | null>(null);
  const globalDetailMainError = ref<string | null>(null);
  const globalDetailMainLoading = ref<boolean>(false);

  async function globalDetailMainExecute() {
    globalDetailMainLoading.value = true;
    globalDetailMainError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailMain({
        ...params
      });
      globalDetailMainData.value = response.data;
    } catch (err: any) {
      globalDetailMainError.value = err.message || 'An error occurred';
    } finally {
      globalDetailMainLoading.value = false;
    }
  }

  return {
    globalDetailMainData,
    globalDetailMainError,
    globalDetailMainLoading,
    globalDetailMainExecute,
  };
}

export function globalDetailCategoryIncomeDeptDetail(params: IndicatorValueServiceGlobalDetailCategoryIncomeDeptDetailRequest) {
  const globalDetailCategoryIncomeDeptDetailData = ref<V1GlobalDetailDeptsReply | null>(null);
  const globalDetailCategoryIncomeDeptDetailError = ref<string | null>(null);
  const globalDetailCategoryIncomeDeptDetailLoading = ref<boolean>(false);

  async function globalDetailCategoryIncomeDeptDetailExecute() {
    globalDetailCategoryIncomeDeptDetailLoading.value = true;
    globalDetailCategoryIncomeDeptDetailError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailCategoryIncomeDeptDetail({
        ...params
      });
      globalDetailCategoryIncomeDeptDetailData.value = response.data;
    } catch (err: any) {
      globalDetailCategoryIncomeDeptDetailError.value = err.message || 'An error occurred';
    } finally {
      globalDetailCategoryIncomeDeptDetailLoading.value = false;
    }
  }

  return {
    globalDetailCategoryIncomeDeptDetailData,
    globalDetailCategoryIncomeDeptDetailError,
    globalDetailCategoryIncomeDeptDetailLoading,
    globalDetailCategoryIncomeDeptDetailExecute,
  };
}

export function globalDetailSurgery(params: IndicatorValueServiceGlobalDetailSurgeryRequest) {
  const globalDetailSurgeryData = ref<V1GlobalDetailSurgeryReply | null>(null);
  const globalDetailSurgeryError = ref<string | null>(null);
  const globalDetailSurgeryLoading = ref<boolean>(false);

  async function globalDetailSurgeryExecute() {
    globalDetailSurgeryLoading.value = true;
    globalDetailSurgeryError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailSurgery({
        ...params
      });
      globalDetailSurgeryData.value = response.data;
    } catch (err: any) {
      globalDetailSurgeryError.value = err.message || 'An error occurred';
    } finally {
      globalDetailSurgeryLoading.value = false;
    }
  }

  return {
    globalDetailSurgeryData,
    globalDetailSurgeryError,
    globalDetailSurgeryLoading,
    globalDetailSurgeryExecute,
  };
}

export function globalDetailGroups(params: IndicatorValueServiceGlobalDetailGroupsRequest) {
  const globalDetailGroupsData = ref<V1GlobalDetailGroupsReply | null>(null);
  const globalDetailGroupsError = ref<string | null>(null);
  const globalDetailGroupsLoading = ref<boolean>(false);

  async function globalDetailGroupsExecute() {
    globalDetailGroupsLoading.value = true;
    globalDetailGroupsError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailGroups({
        ...params
      });
      globalDetailGroupsData.value = response.data;
    } catch (err: any) {
      globalDetailGroupsError.value = err.message || 'An error occurred';
    } finally {
      globalDetailGroupsLoading.value = false;
    }
  }

  return {
    globalDetailGroupsData,
    globalDetailGroupsError,
    globalDetailGroupsLoading,
    globalDetailGroupsExecute,
  };
}

export function globalDetailDoctorGroup(params: IndicatorValueServiceGlobalDetailDoctorGroupRequest) {
  const globalDetailDoctorGroupData = ref<V1GlobalDetailDoctorGroupReply | null>(null);
  const globalDetailDoctorGroupError = ref<string | null>(null);
  const globalDetailDoctorGroupLoading = ref<boolean>(false);

  async function globalDetailDoctorGroupExecute() {
    globalDetailDoctorGroupLoading.value = true;
    globalDetailDoctorGroupError.value = null;

    try {
      const response = await IndicatorValueServiceApi.indicatorValueServiceGlobalDetailDoctorGroup({
        ...params
      });
      globalDetailDoctorGroupData.value = response.data;
    } catch (err: any) {
      globalDetailDoctorGroupError.value = err.message || 'An error occurred';
    } finally {
      globalDetailDoctorGroupLoading.value = false;
    }
  }

  return {
    globalDetailDoctorGroupData,
    globalDetailDoctorGroupError,
    globalDetailDoctorGroupLoading,
    globalDetailDoctorGroupExecute,
  };
}


