package handlers

import (
	"backend-go/internal/pkg/config"
	"backend-go/internal/pkg/file_handle"
	"backend-go/internal/pkg/logger"
	"backend-go/internal/pkg/resp"
	"backend-go/utils"
	"errors"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

type UploadResp struct {
	Url string
}

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		resp.Resp(c, 400, "字段格式不正确", nil, -1)
		return
	}

	fileType := utils.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		resp.Resp(c, 400, "文件类型不支持,请修正type字段", nil, -1)
		return
	}

	fileInfo, err := UploadFile(file_handle.FileType(fileType), file, fileHeader)
	if err != nil {
		logger.Errorf("上传失败", err.Error())
		resp.Resp(c, 500, "上传失败", nil, -1)
		return
	}

	resp.Resp(c, 200, "", &UploadResp{
		Url: fileInfo.AccessUrl,
	}, 0)
}

type FileInfo struct {
	Name      string
	AccessUrl string
}

func UploadFile(fileType file_handle.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := file_handle.GetFileName(fileHeader.Filename)
	if !file_handle.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if file_handle.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := file_handle.GetSavePath()
	if file_handle.CheckSavePath(uploadSavePath) {
		if err := file_handle.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if file_handle.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := uploadSavePath + "/" + fileName
	if err := file_handle.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := config.Config.App.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
