package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iamzhiyudong/xigua-blog/global"
	"github.com/iamzhiyudong/xigua-blog/internal/service"
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
	"github.com/iamzhiyudong/xigua-blog/pkg/convert"
	"github.com/iamzhiyudong/xigua-blog/pkg/errcode"
	"github.com/iamzhiyudong/xigua-blog/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
