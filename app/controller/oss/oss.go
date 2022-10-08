package oss

import (
	"aliyunoss/pkg/oss"
	"github.com/gin-gonic/gin"
)

type OssReq struct {
	OssObjectName string `json:"oss_object_name,omitempty"`
	LocalFile     string `json:"local_file,omitempty"`
}

func Show(c *gin.Context) {
	files, err := oss.ListFile()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "获取失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    files,
	})
}

func Upload(c *gin.Context) {
	ossReq := OssReq{}
	if err := c.ShouldBindJSON(&ossReq); err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}

	if err := oss.UploadFile(ossReq.OssObjectName, ossReq.LocalFile); err != nil {
		c.JSON(400, gin.H{
			"message": "上传失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "上传成功",
	})
}

func Download(c *gin.Context) {
	ossReq := &OssReq{}

	if err := c.ShouldBindJSON(ossReq); err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}
	if err := oss.DownloadFile(ossReq.OssObjectName, ossReq.LocalFile); err != nil {
		c.JSON(400, gin.H{
			"message": "下载失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "下载成功",
	})
}

func Delete(c *gin.Context) {
	ossReq := &OssReq{}

	if err := c.ShouldBindJSON(ossReq); err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}
	if err := oss.DeleteFile(ossReq.OssObjectName); err != nil {
		c.JSON(400, gin.H{
			"message": "删除失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}
