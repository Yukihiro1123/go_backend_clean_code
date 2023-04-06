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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}