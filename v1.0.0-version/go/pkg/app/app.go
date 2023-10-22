package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iizyd/xigua-blog/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total"`
}

type ResponseBody struct {
	List any `json:"list"`
	Pager
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	res_body := gin.H{"code": errcode.Success.Code(), "msg": errcode.Success.Msg(), "data": data}
	r.Ctx.JSON(http.StatusOK, res_body)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code(),
		"msg":  errcode.Success.Msg(),
		"data": ResponseBody{
			List: list,
			Pager: Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
