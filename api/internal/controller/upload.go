package controller

import (
	"api/api"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

// 获取签名(web端直传用)
func (controllerThis *Upload) Sign(ctx context.Context, req *api.UploadSignReq) (res *api.UploadSignRes, err error) {
	request := g.RequestFromCtx(ctx)
	option := utils.AliyunOssSignOption{
		ExpireTime: 15 * 60,
		Dir:        fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`)),
		MinSize:    0,
		MaxSize:    100 * 1024 * 1024,
	}

	if g.Cfg().MustGet(ctx, `uploadCallbackEnable`).Bool() {
		option.Callback = utils.AliyunOssCallback{
			CallbackUrl:      gstr.Replace(request.GetUrl(), request.URL.Path, `/upload/notify`, 1),
			CallbackBody:     `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
			CallbackBodyType: `application/x-www-form-urlencoded`,
		}
	}

	config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`})
	upload := utils.NewAliyunOss(ctx, config)
	signInfo, _ := upload.CreateSign(option)
	res = &api.UploadSignRes{}
	gconv.Struct(signInfo, &res)
	// utils.HttpSuccessJson(g.RequestFromCtx(ctx), signInfo, 0)
	return
}

// 获取Sts Token(App端直传用)
func (controllerThis *Upload) Sts(ctx context.Context, req *api.UploadStsReq) (res *api.UploadStsRes, err error) {
	request := g.RequestFromCtx(ctx)
	config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssRoleArn`})
	dir := fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`))
	option := utils.AliyunOssStsOption{
		SessionName: `oss_app_sts_token`,
		ExpireTime:  15 * 60,
		Policy:      `{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:` + gconv.String(config[`aliyunOssBucket`]) + `/` + dir + `*"]}],"Version": "1"}`,
	}
	if g.Cfg().MustGet(ctx, `uploadCallbackEnable`).Bool() {
		option.Callback = utils.AliyunOssCallback{
			CallbackUrl:      gstr.Replace(request.GetUrl(), request.URL.Path, `/upload/notify`, 1),
			CallbackBody:     `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
			CallbackBodyType: `application/x-www-form-urlencoded`,
		}
	}

	upload := utils.NewAliyunOss(ctx, config)
	stsInfo, _ := upload.GetStsToken(option)
	stsInfo[`endpoint`] = config[`aliyunOssHost`]
	stsInfo[`bucket`] = config[`aliyunOssBucket`]
	stsInfo[`dir`] = dir
	stsInfo[`callbackUrl`] = option.Callback.CallbackUrl
	stsInfo[`callbackBody`] = option.Callback.CallbackBody
	stsInfo[`callbackBodyType`] = option.Callback.CallbackBodyType
	request.Response.WriteJsonExit(stsInfo) //必须按阿里云官方文档要求的格式返回。App端的SDK才能用
	return
}

// 回调
func (controllerThis *Upload) Notify(ctx context.Context, req *api.UploadNotifyReq) (res *api.UploadNotifyRes, err error) {
	r := g.RequestFromCtx(ctx)
	filename := r.Get(`filename`).String()
	width := r.Get(`width`).String()
	height := r.Get(`height`).String()

	config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssHost`, `aliyunOssBucket`})
	upload := utils.NewAliyunOss(r.GetCtx(), config)
	err = upload.Notify(r)
	if err != nil {
		return
	}

	res = &api.UploadNotifyRes{
		Url: upload.GetBucketHost() + `/` + filename + `?w=` + width + `&h=` + height, //需要记录宽高，ios显示瀑布流必须知道宽高。直接存在query内
	}
	return
}
