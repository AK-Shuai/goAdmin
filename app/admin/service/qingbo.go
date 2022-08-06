package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
	"strconv"
)

type QingBo struct {
	service.Service
}

// GetPage 获取任务列表
func (e *QingBo) GetPage(c *dto.SysQingboGetPageReq, list *[]models.SysQingbo, count *int64) error {
	StartDate, _ := strconv.Atoi(c.StartDate)
	EndDate, _ := strconv.Atoi(c.EndDate)

	err := e.Orm.
		Scopes(
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			BigSmallDate(StartDate, EndDate),
			IsCompany(c.CompanyType),
			OrderType(c.CreatedAtOrder),
			CompanyNameType(c.CompanyNameType),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	return nil
}

// Insert 新增任务
func (e *QingBo) Insert(c *dto.SysQingboControl) error {
	var err error
	var data models.SysQingbo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("Service InsertSysConfig error:%s", err)
		return err
	}
	return nil
}

// InsertList 新增批量任务
func (e *QingBo) InsertList(c *dto.SysQingboControl) error {
	var err error
	var data models.SysQingbo
	for _, date := range c.Dates {
		c.Date = date
		c.Generate(&data)
		err = e.Orm.Create(&data).Error
		if err != nil {
			e.Log.Errorf("Service InsertSysConfig error:%s", err)
			return err
		}
	}
	return nil
}

// UploadExcel 新增excel
func (e *QingBo) UploadExcel(c *dto.SysQingboControlExcel) error {
	var data models.SysQingbo
	var ssq dto.SysQingboExcelList
	Companydata := make([]models.SysQingboCompany, 0)
	Company := make(map[string]int)
	for _, excelList := range c.Datas {

		e.Orm.Find(&Companydata)
		for _, CompanydataList := range Companydata {
			Company[CompanydataList.CompanyName] = CompanydataList.Id
		}
		ssq.Generate(&data, Company, excelList)
		e.Orm.Create(&data)

		//o.JSON(200, gin.H{"code": 200, "data": Companydata, "msg": "msg"})
	}
	return nil
}

// BigSmallDate 日期筛选
func BigSmallDate(startDate int, endDate int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if startDate > endDate {
			return db.Where("date = ?", startDate)
		}
		if startDate != 0 && endDate != 0 {
			return db.Where("date >= ?", startDate).Where("date <= ?", endDate)
		} else {
			return db
		}
	}
}

// IsCompany 服务性质筛选
func IsCompany(CompanyType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if CompanyType != 0 {
			return db.Where("company_type = ?", CompanyType)
		} else {
			return db
		}
	}
}

// CompanyNameType 公司筛选
func CompanyNameType(CompanyNameType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if CompanyNameType != 0 {
			return db.Where("company_name_type = ?", CompanyNameType)
		} else {
			return db
		}
	}
}

// OrderType 排序类型
func OrderType(createdAtOrder string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("id " + createdAtOrder)
	}
}

// GetPageCompany 获取公司列表
func (e *QingBo) GetPageCompany(c *dto.SysQingboGetPageCompanyReq, list *[]models.SysQingboCompany, count *int64) error {

	// 获取全部公司名
	if c.IsAll == true {
		err := e.Orm.
			Scopes(
				OrderType(c.CreatedAtOrder),
			).
			Find(list).Limit(-1).Offset(-1).
			Count(count).Error
		if err != nil {
			e.Log.Errorf("Service GetSysConfigPage error:%s", err)
			return err
		}
	}
	// 分页获取公司名
	err := e.Orm.
		Scopes(
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			OrderType(c.CreatedAtOrder),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	return nil
}

// InsertCompanyName 新增公司名
func (e *QingBo) InsertCompanyName(c *dto.SysQingboCompanyControl) error {
	var err error
	var data models.SysQingboCompany
	c.Generate(&data)
	err = e.Orm.Where(" company_name = ?", c.CompanyName).Find(&data).Error

	if &data.Id != nil && data.Id >= 1 {
		return errors.New("已存在")
	} else if err != nil {
		return errors.New("查询失败")
	}
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("Service InsertSysConfig error:%s", err)
		return err
	}
	return nil
}

// RemoveQingBoList 删除QingBoList
func (e *QingBo) RemoveQingBoList(d *dto.SysQingDeleteReq, p *actions.DataPermission) error {
	var data models.SysQingbo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysApi error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *QingBo) UpdateQingBoList(c *dto.SysQingUpdateReq, p *actions.DataPermission) error {
	var err error
	var model = models.SysQingbo{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.Orm.Save(&model)
	err = db.Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysConfig error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// RemoveQingBoCompanyName 删除 CompanyName
func (e *QingBo) RemoveQingBoCompanyName(d *dto.SysQingCompanyDeleteReq, p *actions.DataPermission) error {
	var data models.SysQingboCompany

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysApi error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
