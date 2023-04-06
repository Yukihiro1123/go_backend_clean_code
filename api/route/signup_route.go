package route

import (
	"go_backend_clean_code/api/controller"
	"go_backend_clean_code/bootstrap"
	"go_backend_clean_code/domain"
	"go_backend_clean_code/mongo"
	"go_backend_clean_code/repository"
	"go_backend_clean_code/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	//user repository
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	//signup controller
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env: env,
	}
	group.POST("/signup", sc.Signup)
}