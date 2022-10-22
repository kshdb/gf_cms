package response

import (
	"context"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"

	"github.com/gogf/gf/v2/frame/g"
)

type sResponse struct {
}

var (
	insResponse           = sResponse{}
	codeSuccessDefault    = gcode.CodeOK.Code()
	MessageSuccessDefault = "操作成功"
)

func init() {
	service.RegisterResponse(New())
}

func New() *sResponse {
	return &sResponse{}
}

func Response() *sResponse {
	return &insResponse
}

// SuccessJson 返回成功json
func (s *sResponse) SuccessJson(ctx context.Context, code int, message string, data interface{}) {
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// SuccessJsonDefault 返回默认成功json
func (s *sResponse) SuccessJsonDefault(ctx context.Context) {
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), service.Response().SuccessMessageDefault(), g.Map{})
}

// SuccessCodeDefault 获取默认成功code码
func (s *sResponse) SuccessCodeDefault() int {
	return codeSuccessDefault
}

// SuccessMessageDefault 获取默认成功提示语
func (s *sResponse) SuccessMessageDefault() string {
	return MessageSuccessDefault
}

// View 模板渲染
func (s *sResponse) View(ctx context.Context, template string, data g.Map) (err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl(template, data)
	if err != nil {
		return err
	}
	return
}

func (s *sResponse) ErrorTpl(ctx context.Context, code int, message string) (err error) {
	template := "tpl/error.html"
	data := g.Map{
		"code":    code,
		"message": message,
	}
	err = g.RequestFromCtx(ctx).Response.WriteTpl(template, data)
	if err != nil {
		return err
	}
	return
}
