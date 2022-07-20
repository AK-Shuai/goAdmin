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
	}
	r1 := v1.Group("/sys-company-list").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("", api.GetPageCompany)
		r1.POST("", api.InsertCompanyName)
	}
}
