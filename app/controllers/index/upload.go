package index

import (
	"gofiber/app/module/context"
	"gofiber/utils"
)

// Upload 上传文件
func Upload(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	//解析文件数据
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	//获取所有文件
	files := form.File["file"]
	var saveFilePath string
	var filePaths []string

	//上传文件
	for _, file := range files {
		saveFilePath, err = utils.Upload(file)
		if err != nil {
			return err
		}
		filePaths = append(filePaths, saveFilePath)
	}

	return ctx.SuccessJson(filePaths)
}
