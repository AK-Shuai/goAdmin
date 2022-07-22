package models

import (
	"go-admin/common/models"
)

type SysQingbo struct {
	models.Model
	Date            int    `json:"Date" gorm:"index;comment:日期"`                       //
	Name            string `json:"Name" gorm:"size:255;comment:名字"`                    //
	ServiceQuality  int    `json:"ServiceQuality" gorm:"index;comment:服务性质 1-除虫，2-除鼠"` //
	ServiceContent  int    `json:"ServiceContent" gorm:"index;comment:服务内容 1-首次，2-售后"`
	ServiceTime     string `json:"ServiceTime" gorm:"size:255;comment:执行时间"`         //
	Address         string `json:"Address" gorm:"size:255;comment:地址"`               //
	Money           int    `json:"Money" gorm:"index;comment:金额"`                    //
	Telephone       string `json:"Telephone" gorm:"size:255;comment:联系人手机号"`         //
	CompanyNameType int    `json:"CompanyNameType" gorm:"index;comment:公司名类型 0-非公司"` //
	CompanyType     int    `json:"CompanyType" gorm:"index;comment:是否公司 1-公司 2-个人"`  //
	Remark          string `json:"Remark" gorm:"size:255;comment:自定义内容"`
	models.ControlBy
	models.ModelTime
}

func (SysQingbo) TableName() string {
	return "sys_qingbo"
}

func (e *SysQingbo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysQingbo) GetId() interface{} {
	return e.Id
}

// SysQingboCompany /*  公司列表
type SysQingboCompany struct {
	models.Model
	CompanyName string `json:"CompanyName" gorm:"comment:公司名"` //
	models.ControlBy
	models.ModelTime
}

func (SysQingboCompany) TableName() string {
	return "sys_qingbo_company_list"
}

func (e *SysQingboCompany) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysQingboCompany) GetId() interface{} {
	return e.Id
}
