package route

import (
	"go_backend_clean_code/api/controller"
	"go_backend_clean_code/bootstrap"
	"go_backend_clean_code/domain"
	"go_backend_clean_code/repository"
	"go_backend_clean_code/usecase"
	"time"

	"go_backend_clean_code/mongo"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env: env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}