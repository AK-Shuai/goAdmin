package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"
)

// SysQingboGetPageReq 任务列表
type SysQingboGetPageReq struct {
	dto.Pagination  `search:"-"`
	Name            string `form:"name" search:"type:contains;column:name;table:sys_qingbo"`
	Date            int    `form:"Date" search:"type:contains;column:date;table:sys_qingbo"`
	ServiceQuality  int    `form:"ServiceQuality" search:"type:contains;column:service_quality;table:sys_qingbo"`
	ServiceContent  int    `form:"ServiceContent" search:"type:contains;column:service_content;table:sys_qingbo"`
	ServiceTime     string `form:"ServiceTime" search:"type:contains;column:service_time;table:sys_qingbo"`
	Address         string `form:"Address"search:"type:contains;column:address;table:sys_qingbo"`
	Money           int    `form:"Money" search:"type:contains;column:money;table:sys_qingbo"`
	Telephone       string `form:"Telephone" search:"type:contains;column:telephone;table:sys_qingbo"`
	CompanyNameType int    `form:"CompanyNameType" search:"type:contains;column:company_name_type;table:sys_qingbo"`
	CompanyType     int    `form:"CompanyType" search:"type:contains;column:company_type;table:sys_qingbo"`
	StartDate       string `form:"startDate" search:"type:contains;column:date;table:sys_qingbo"`
	EndDate         string `form:"endDate" search:"type:contains;column:date;table:sys_qingbo"`
	CreatedAtOrder  string `form:"createdAtOrder"`
	Remark          string `form:"Remark" search:"type:contains;column:remark;table:sys_qingbo"`
}

type SysQingboControl struct {
	Id              int      `uri:"Id" comment:"编码"`                // 编码
	Date            string   `json:"Date" comment:"日期"`             //
	Name            string   `json:"Name" comment:"名字"`             //
	ServiceQuality  int      `json:"ServiceQuality" comment:"服务性质"` //
	ServiceContent  int      `json:"ServiceContent" comment:"服务内容"`
	ServiceTime     string   `json:"ServiceTime" comment:"执行时间"`       //
	Address         string   `json:"Address" comment:"地址"`             //
	Money           int      `json:"Money" comment:"金额"`               //
	Telephone       string   `json:"Telephone" comment:"联系人手机号"`       //
	CompanyNameType int      `json:"CompanyNameType" comment:"公司名称类型"` //
	CompanyType     int      `json:"CompanyType" comment:"是否公司"`       //
	Dates           []string `json:"Dates"`
	Remark          string   `json:"Remark" comment:"自定义内容"` //
	common.ControlBy
}

func (m *SysQingboGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// GetId 获取数据对应的ID
func (s *SysQingboControl) GetId() interface{} {
	return s.Id
}

func (s *SysQingboControl) Generate(model *models.SysQingbo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}

	intDate, err := strconv.Atoi(s.Date)
	if err != nil {

	}
	model.Date = intDate
	model.Name = s.Name
	model.Address = s.Address
	model.ServiceTime = s.ServiceTime
	model.ServiceContent = s.ServiceContent
	model.ServiceQuality = s.ServiceQuality
	model.Money = s.Money
	model.Telephone = s.Telephone
	model.CompanyType = s.CompanyType
	model.CompanyNameType = s.CompanyNameType
	model.Remark = s.Remark
}

type SysQingboControlExcel struct {
	Datas []SysQingboExcelList `json:datas`
	common.ControlBy
}

type SysQingboExcelList struct {
	Id              int    `uri:"Id" comment:"编码"`                // 编码
	Date            string `json:"Date" comment:"日期"`             //
	Name            string `json:"Name" comment:"名字"`             //
	ServiceQuality  string `json:"ServiceQuality" comment:"服务性质"` //
	ServiceContent  string `json:"ServiceContent" comment:"服务内容"`
	ServiceTime     string `json:"ServiceTime" comment:"执行时间"`       //
	Address         string `json:"Address" comment:"地址"`             //
	Money           int    `json:"Money" comment:"金额"`               //
	Telephone       string `json:"Telephone" comment:"联系人手机号"`       //
	CompanyNameType string `json:"CompanyNameType" comment:"公司名称类型"` //
	CompanyType     string `json:"CompanyType" comment:"是否公司"`       //
	Remark          string `json:"Remark" comment:"自定义内容"`           //
}

// GetId 获取数据对应的ID
func (s *SysQingboExcelList) GetId() interface{} {
	return s.Id
}

func (s *SysQingboExcelList) Generate(model *models.SysQingbo, result map[string]int, excelList SysQingboExcelList) {

	if excelList.Id == 0 {
		model.Model = common.Model{Id: excelList.Id}
	}

	intDate, _ := strconv.Atoi(excelList.Date)

	model.Date = intDate
	model.Name = excelList.Name
	model.Address = excelList.Address
	model.ServiceTime = excelList.ServiceTime
	model.ServiceContent = s.UpServiceContent(excelList)
	model.ServiceQuality = s.UpServiceQuality(excelList)
	model.Money = excelList.Money
	model.Telephone = excelList.Telephone
	model.CompanyType = s.UpCompanyType(excelList)
	model.CompanyNameType = result[excelList.CompanyNameType]
	model.Remark = excelList.Remark
}

func (s *SysQingboExcelList) UpServiceContent(excelList SysQingboExcelList) int {
	var ServiceContent int

	if excelList.ServiceContent == "除虫" {
		ServiceContent = 1
	} else if excelList.ServiceContent == "除鼠" {
		ServiceContent = 2
	}
	return ServiceContent
}

func (s *SysQingboExcelList) UpServiceQuality(excelList SysQingboExcelList) int {
	var ServiceQuality int

	if excelList.ServiceQuality == "首次" {
		ServiceQuality = 1
	} else if excelList.ServiceQuality == "售后" {
		ServiceQuality = 2
	}
	return ServiceQuality
}

func (s *SysQingboExcelList) UpCompanyType(excelList SysQingboExcelList) int {
	var CompanyType int

	if excelList.CompanyType == "是" {
		CompanyType = 1
	} else if excelList.CompanyType == "否" {
		CompanyType = 2
	}
	return CompanyType
}

// SysQingboGetPageCompanyReq 公司列表
type SysQingboGetPageCompanyReq struct {
	dto.Pagination `search:"-"`
	CreatedAtOrder string `form:"createdAtOrder"`
	IsAll          bool   `form:"isAll"`
	// CompanyName    string `form:"CompanyName" search:"type:contains;column:company_name;table:sys_qingbo_company_list"`
}

type SysQingboCompanyControl struct {
	Id          int    `uri:"Id" comment:"编码"` // 编码
	CompanyName string `json:"CompanyName" comment:"公司名"`
	common.ControlBy
}

// GetId 获取数据对应的ID
func (s *SysQingboCompanyControl) GetId() interface{} {
	return s.Id
}

func (s *SysQingboCompanyControl) Generate(model *models.SysQingboCompany) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}

	model.CompanyName = s.CompanyName
}

// SysQingDeleteReq 功能删除请求参数
type SysQingDeleteReq struct {
	Id int `json:"id"`
}

func (s *SysQingDeleteReq) GetId() interface{} {
	return s.Id
}

type SysQingUpdateReq struct {
	Id              int    `uri:"id" comment:"编码"`                // 编码
	Date            int    `json:"Date" comment:"日期"`             //
	Name            string `json:"Name" comment:"名字"`             //
	ServiceQuality  int    `json:"ServiceQuality" comment:"服务性质"` //
	ServiceContent  int    `json:"ServiceContent" comment:"服务内容"`
	ServiceTime     string `json:"ServiceTime" comment:"执行时间"`       //
	Address         string `json:"Address" comment:"地址"`             //
	Money           int    `json:"Money" comment:"金额"`               //
	Telephone       string `json:"Telephone" comment:"联系人手机号"`       //
	CompanyNameType int    `json:"CompanyNameType" comment:"公司名称类型"` //
	CompanyType     int    `json:"CompanyType" comment:"是否公司"`       //
	Remark          string `json:"Remark" comment:"自定义内容"`           //
	common.ControlBy
}

func (s *SysQingUpdateReq) GetId() interface{} {
	return s.Id
}

func (s *SysQingUpdateReq) Generate(model *models.SysQingbo) {
	if s.Id != 0 {
		model.Id = s.Id
	}

	model.Date = s.Date
	model.Name = s.Name
	model.Address = s.Address
	model.ServiceTime = s.ServiceTime
	model.ServiceContent = s.ServiceContent
	model.ServiceQuality = s.ServiceQuality
	model.Money = s.Money
	model.Telephone = s.Telephone
	model.CompanyType = s.CompanyType
	model.CompanyNameType = s.CompanyNameType
	model.Remark = s.Remark
}

// SysQingCompanyDeleteReq 功能删除请求参数
type SysQingCompanyDeleteReq struct {
	Id int `json:"id"`
}

func (s *SysQingCompanyDeleteReq) GetId() interface{} {
	return s.Id
}
