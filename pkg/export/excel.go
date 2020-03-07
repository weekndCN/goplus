package export

import "rwplus-backend/pkg/setting"

func GetExcelFullUrl(name string) string {
    return setting.AppSetting.PrefxUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
    return setting.AppSetting.ExportSavePath
}

func GetExcelFullPath() string {
    return setting.AppSetting.RuntimeRootPath() + GetExcelPath()
}