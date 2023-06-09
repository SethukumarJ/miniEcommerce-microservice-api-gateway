package auth

import (
    "github.com/gin-gonic/gin"
    "github.com/SethukumarJ/go-grpc-api-gateway/pkg/auth/routes"
    "github.com/SethukumarJ/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
    svc := &ServiceClient{
        Client: InitServiceClient(c),
    }

    routes := r.Group("/auth")
    routes.DELETE("/delete/:id",svc.Delete)
    routes.POST("/register", svc.Register)
    routes.POST("/login", svc.Login)

    return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
    routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
    routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
    routes.Delete(ctx, svc.Client)
}