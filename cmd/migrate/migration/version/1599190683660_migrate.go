package version

import (
	"go-admin/app/admin/models"
	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"
	"gorm.io/gorm"
	"path/filepath"
	"runtime"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	fileName = filepath.Base(fileName)
	fileName = fileName[:len(fileName)-3]
	migration.Migrate.SetVersion(fileName, _1599190683660Test)
}

func _1599190683660Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var err error

		list := []models.CasbinRule{
			{"p", "admin", "/api/v1/menulist", "GET", "", "", ""},
			{"p", "admin", "/api/v1/menu", "POST", "", "", ""},
			{"p", "admin", "/api/v1/dict/databytype/", "GET", "", "", ""},
			{"p", "admin", "/api/v1/menu", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/menu/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/sysUserList", "GET", "", "", ""},
			{"p", "admin", "/api/v1/sysUser/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/sysUser/", "GET", "", "", ""},
			{"p", "admin", "/api/v1/sysUser", "POST", "", "", ""},
			{"p", "admin", "/api/v1/sysUser", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/sysUser/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/user/profile", "GET", "", "", ""},
			{"p", "admin", "/api/v1/rolelist", "GET", "", "", ""},
			{"p", "admin", "/api/v1/role/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/role", "POST", "", "", ""},
			{"p", "admin", "/api/v1/role", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/role/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/configList", "GET", "", "", ""},
			{"p", "admin", "/api/v1/config/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/config", "POST", "", "", ""},
			{"p", "admin", "/api/v1/config", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/config/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/menurole", "GET", "", "", ""},
			{"p", "admin", "/api/v1/roleMenuTreeselect/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/menuTreeselect", "GET", "", "", ""},
			{"p", "admin", "/api/v1/rolemenu", "GET", "", "", ""},
			{"p", "admin", "/api/v1/rolemenu", "POST", "", "", ""},
			{"p", "admin", "/api/v1/rolemenu/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/deptList", "GET", "", "", ""},
			{"p", "admin", "/api/v1/dept/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/dept", "POST", "", "", ""},
			{"p", "admin", "/api/v1/dept", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/dept/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/dict/datalist", "GET", "", "", ""},
			{"p", "admin", "/api/v1/dict/data/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/dict/databytype/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/dict/data", "POST", "", "", ""},
			{"p", "admin", "/api/v1/dict/data/", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/dict/data/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/dict/typelist", "GET", "", "", ""},
			{"p", "admin", "/api/v1/dict/type/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/dict/type", "POST", "", "", ""},
			{"p", "admin", "/api/v1/dict/type", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/dict/type/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/postlist", "GET", "", "", ""},
			{"p", "admin", "/api/v1/post/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/post", "POST", "", "", ""},
			{"p", "admin", "/api/v1/post", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/post/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/menu/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/menuids", "GET", "", "", ""},
			{"p", "admin", "/api/v1/loginloglist", "GET", "", "", ""},
			{"p", "admin", "/api/v1/loginlog/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/operloglist", "GET", "", "", ""},
			{"p", "admin", "/api/v1/getinfo", "GET", "", "", ""},
			{"p", "admin", "/api/v1/roledatascope", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/roleDeptTreeselect/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/deptTree", "GET", "", "", ""},
			{"p", "admin", "/api/v1/configKey/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/logout", "POST", "", "", ""},
			{"p", "admin", "/api/v1/user/avatar", "POST", "", "", ""},
			{"p", "admin", "/api/v1/user/pwd", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/dict/typeoptionselect", "GET", "", "", ""},
			{"p", "admin", "/api/v1/sysjobList", "GET", "", "", ""},
			{"p", "admin", "/api/v1/sysjob/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/sysjob", "POST", "", "", ""},
			{"p", "admin", "/api/v1/sysjob", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/sysjob/:id", "DELETE", "", "", ""},
			{"p", "admin", "/api/v1/syssettingList", "GET", "", "", ""},
			{"p", "admin", "/api/v1/syssetting/:id", "GET", "", "", ""},
			{"p", "admin", "/api/v1/syssetting", "POST", "", "", ""},
			{"p", "admin", "/api/v1/syssetting", "PUT", "", "", ""},
			{"p", "admin", "/api/v1/syssetting/:id", "DELETE", "", "", ""},
		}

		err = db.Create(list).Error
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
