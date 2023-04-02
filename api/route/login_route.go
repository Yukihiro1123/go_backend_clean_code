package route

import (
	"go_backend_clean_code/api/controller"
	"go_backend_clean_code/bootstrap"
	"go_backend_clean_code/domain"
	"go_backend_clean_code/repository"
	"go_backend_clean_code/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env: env,
	}
	group.POST("/login", lc.Login)
}