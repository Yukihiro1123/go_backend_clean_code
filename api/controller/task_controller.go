package controller

import (
	"go_backend_clean_code/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

//gin.Context を使うことで、URL に付随したパラメータの取得や POST で送信されたデータの取得などを行うことが可能
func (tc *TaskController) Create(c *gin.Context) {
	var task domain.Task
	//jsonをtask structにエンコード
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	//c.Set("x-user-id", userID)
	//We can get the UserID from the HTTP Web Framework Context
	userID := c.GetString("x-user-id")
	task.ID = primitive.NewObjectID()
	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")
	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}