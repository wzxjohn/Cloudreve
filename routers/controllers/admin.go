package controllers

import (
	"github.com/HFO4/cloudreve/pkg/aria2"
	"github.com/HFO4/cloudreve/pkg/email"
	"github.com/HFO4/cloudreve/pkg/request"
	"github.com/HFO4/cloudreve/pkg/serializer"
	"github.com/HFO4/cloudreve/service/admin"
	"github.com/gin-gonic/gin"
	"io"
)

// AdminSummary 获取管理站点概况
func AdminSummary(c *gin.Context) {
	var service admin.NoParamService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Summary()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminNews 获取社区新闻
func AdminNews(c *gin.Context) {
	r := request.HTTPClient{}
	res := r.Request("GET", "https://forum.cloudreve.org/api/discussions?include=startUser%2ClastUser%2CstartPost%2Ctags&filter%5Bq%5D=%20tag%3Anotice&sort=-startTime&", nil)
	io.Copy(c.Writer, res.Response.Body)
}

// AdminChangeSetting 获取站点设定项
func AdminChangeSetting(c *gin.Context) {
	var service admin.BatchSettingChangeService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Change()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminGetSetting 获取站点设置
func AdminGetSetting(c *gin.Context) {
	var service admin.BatchSettingGet
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminGetGroups 获取用户组列表
func AdminGetGroups(c *gin.Context) {
	var service admin.NoParamService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GroupList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminReloadService 重新加载子服务
func AdminReloadService(c *gin.Context) {
	service := c.Param("service")
	switch service {
	case "email":
		email.Init()
	case "aria2":
		aria2.Init(true)
	}

	c.JSON(200, serializer.Response{})
}

// AdminSendTestMail 发送测试邮件
func AdminSendTestMail(c *gin.Context) {
	var service admin.MailTestService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Send()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminListRedeems 列出激活码
func AdminListRedeems(c *gin.Context) {
	var service admin.AdminListService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Redeems()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminGenerateRedeems 生成激活码
func AdminGenerateRedeems(c *gin.Context) {
	var service admin.GenerateRedeemsService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Generate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminDeleteRedeem 删除激活码
func AdminDeleteRedeem(c *gin.Context) {
	var service admin.SingleIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.DeleteRedeem()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminTestAria2 测试aria2连接
func AdminTestAria2(c *gin.Context) {
	var service admin.Aria2TestService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Test()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminListPolicy 列出存储策略
func AdminListPolicy(c *gin.Context) {
	var service admin.AdminListService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Policies()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminTestPath 测试本地路径可用性
func AdminTestPath(c *gin.Context) {
	var service admin.PathTestService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Test()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminTestSlave 测试从机可用性
func AdminTestSlave(c *gin.Context) {
	var service admin.SlaveTestService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Test()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminAddPolicy 新建存储策略
func AdminAddPolicy(c *gin.Context) {
	var service admin.AddPolicyService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminAddCORS 创建跨域策略
func AdminAddCORS(c *gin.Context) {
	var service admin.PolicyService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.AddCORS()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminAddSCF 创建回调函数
func AdminAddSCF(c *gin.Context) {
	var service admin.PolicyService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.AddSCF()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminOneDriveOAuth 获取 OneDrive OAuth URL
func AdminOneDriveOAuth(c *gin.Context) {
	var service admin.PolicyService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetOAuth(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminGetPolicy 获取存储策略详情
func AdminGetPolicy(c *gin.Context) {
	var service admin.PolicyService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminDeletePolicy 删除存储策略
func AdminDeletePolicy(c *gin.Context) {
	var service admin.PolicyService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminListGroup 列出用户组
func AdminListGroup(c *gin.Context) {
	var service admin.AdminListService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Groups()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminAddGroup 新建用户组
func AdminAddGroup(c *gin.Context) {
	var service admin.AddGroupService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Add()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminDeleteGroup 删除用户组
func AdminDeleteGroup(c *gin.Context) {
	var service admin.GroupService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminGetGroup 获取用户组详情
func AdminGetGroup(c *gin.Context) {
	var service admin.GroupService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminListUser 列出用户
func AdminListUser(c *gin.Context) {
	var service admin.AdminListService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Users()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}