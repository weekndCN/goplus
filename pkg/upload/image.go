package upload

import (
    "fmt"
    "log"
    "path"
    "mime/multipart"
    "strings"
    "os"

    "rwplus-backend/pkg/file"
    "rwplus-backend/pkg/setting"
    "rwplus-backend/pkg/util"
)

// 获取图片url绝对路径
func GetImageFullUrl(name string) string {
    return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

// 获取图片名称
func GetImageName(name string) string {
    ext := path.Ext(name)
    fileName := strings.TrimSuffix(name, ext)
    fileName = util.EncodeMD5(fileName)

    return fileName + ext
}

//获取图片保存路径
func GetImagePath() string {
    return setting.AppSetting.ImageSavePath
}

// 获取图片服务绝对路径
func GetImageFullPath() string {
    return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// 检查图片Extension
func CheckImageExt(fileName string) bool {
    ext := file.GetExt(fileName)
    for _, allowExt := range setting.AppSetting.ImageAllowExts {
        if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
            return true
        }
    }

    return false
}

// 检查图片大小
func CheckImageSize(f multipart.File) bool {
    size, err := file.GetSize(f)
    if err != nil {
        log.Println(err)
        return false
    }

    return size <= setting.AppSetting.ImageMaxSize
}

// 检查图片
func CheckImage(src string) error {
    dir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("os.Getwd error: %v", err)
    }

    err = file.IsNotExistMkDir(dir + "/" + src)

    if err != nil {
        return fmt.Errorf("file.IsNotExistMKDir err: %v", err)
    }

    perm := file.CheckPermission(src)
    if perm == true {
        return fmt.Errorf("file.CheckPermission Permission denied src:%s", src)
    }

    return nil
}

