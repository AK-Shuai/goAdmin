package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerBusinessRouter)
}

func registerBusinessRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.QingBo{}
	r := v1.Group("/sys-business").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.POST("", api.Insert)
		r.POST("/insert-list", api.InsertList)
		r.DELETE("", api.Delete)
		r.PUT("", api.Update)
		r.POST("/upload-excel", api.UploadExcel)
	}

	r1 := v1.Group("/sys-company-list").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("", api.GetPageCompany)
		r1.POST("", api.InsertCompanyName)
		r1.DELETE("", api.DeleteCompanyName)
	}

	r2 := v1.Group("/sys-service-content").Use(authMiddleware.MiddlewareFunc())
	{
		r2.GET("", api.GetPageServiceContent)
		r2.POST("", api.InsertServiceContent)
		r2.DELETE("", api.DeleteServiceContent)
	}
}
