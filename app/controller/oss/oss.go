package oss

import (
	"aliyunoss/pkg/oss"
	"aliyunoss/pkg/response"
	"github.com/gin-gonic/gin"
)

type OssReq struct {
	OssObjectName string `json:"oss_object_name,omitempty"`
	LocalFile     string `json:"local_file,omitempty"`
}

func Show(c *gin.Context) {
	files, err := oss.ListFile()
	if err != nil {
		response.ErrorJSON(c, "获取失败")
		return
	}
	response.SuccessJSONWithField(c, "获取成功", response.Field{
		"files": files,
	})
}

func Upload(c *gin.Context) {
	ossReq := OssReq{}
	if err := c.ShouldBindJSON(&ossReq); err != nil {
		response.ErrorJSON(c, "参数错误")
		return
	}

	if err := oss.UploadFile(ossReq.OssObjectName, ossReq.LocalFile); err != nil {
		response.ErrorJSON(c, "上传失败")
		return
	}
	response.SuccessJSON(c, "上传成功")
}

func Download(c *gin.Context) {
	ossReq := &OssReq{}

	if err := c.ShouldBindJSON(ossReq); err != nil {
		response.ErrorJSON(c, "参数错误")
		return
	}
	if err := oss.DownloadFile(ossReq.OssObjectName, ossReq.LocalFile); err != nil {
		response.ErrorJSON(c, "下载失败")
		return
	}
	response.SuccessJSON(c, "下载成功")
}

func Delete(c *gin.Context) {
	ossReq := &OssReq{}

	if err := c.ShouldBindJSON(ossReq); err != nil {
		response.ErrorJSON(c, "参数错误")
		return
	}
	if err := oss.DeleteFile(ossReq.OssObjectName); err != nil {
		response.ErrorJSON(c, "删除失败")
		return
	}
	response.SuccessJSON(c, "删除成功")
}
