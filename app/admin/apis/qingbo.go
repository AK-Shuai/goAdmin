package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type QingBo struct {
	api.Api
}

// GetPage 获取任务列表
func (e QingBo) GetPage(c *gin.Context) {
	s := service.QingBo{}
	req := dto.SysQingboGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	SysQingbo := models.SysQingbo{}
	e.Orm.AutoMigrate(&SysQingbo)
	/*
		SysQingbo := models.SysQingbo{}
		e.Orm.AutoMigrate(&SysQingbo)
		c.JSON(200, gin.H{"code": 200, "data": 1, "msg": "建表成功"})
	*/
	list := make([]models.SysQingbo, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	// c.JSON(200, gin.H{"code": 200, "data":  numbers, "msg": "msg"})
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")

}

// Insert 新增个人任务
func (e QingBo) Insert(c *gin.Context) {
	s := service.QingBo{}
	req := dto.SysQingboControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// InsertList 新增公司任务
func (e QingBo) InsertList(c *gin.Context) {
	s := service.QingBo{}
	req := dto.SysQingboControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))

	err = s.InsertList(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// UploadExcel 上传任务excel
func (e QingBo) UploadExcel(c *gin.Context) {
	s := service.QingBo{}
	req := dto.SysQingboControlExcel{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetCreateBy(user.GetUserId(c))
	err = s.UploadExcel(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(1, "创建成功")
}

func (e QingBo) Delete(c *gin.Context) {
	req := dto.SysQingDeleteReq{}
	s := service.QingBo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	p := actions.GetPermissionFromContext(c)
	err = s.RemoveQingBoList(&req, p)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e QingBo) Update(c *gin.Context) {
	req := dto.SysQingUpdateReq{}
	s := service.QingBo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)
	err = s.UpdateQingBoList(&req, p)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// GetPageCompany 获取公司名列表
func (e QingBo) GetPageCompany(c *gin.Context) {
	s := service.QingBo{}
	req := dto.SysQingboGetPageCompanyReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	/*
		SysQingboCompany := models.SysQingboCompany{}
		e.Orm.AutoMigrate(&SysQingboCompany)
		c.JSON(200, gin.H{"code": 200, "data": 1, "msg": "建表成功"})
	*/
	list := make([]models.SysQingboCompany, 0)
	var count int64
	err = s.GetPageCompany(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	// c.JSON(200, gin.H{"code": 200, "data":  "", "msg": "查询成功"})
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")

}

// InsertCompanyName 新增公司名
func (e QingBo) InsertCompanyName(c *gin.Context) {
	s := service.QingBo{}
	req := dto.SysQingboCompanyControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))

	err = s.InsertCompanyName(&req)
	if err != nil {
		if err.Error() == "已存在" {
			e.Error(500, err, "已存在名字")
			return
		}
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

func (e QingBo) DeleteCompanyName(c *gin.Context) {
	req := dto.SysQingCompanyDeleteReq{}
	s := service.QingBo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	p := actions.GetPermissionFromContext(c)
	err = s.RemoveQingBoCompanyName(&req, p)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}
