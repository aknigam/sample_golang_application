package resources

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample_golang_application/common"
	"sample_golang_application/models"
	"sample_golang_application/repository"
	"sample_golang_application/service"
	"strconv"
)

type StoryResource struct {
	storyService service.StoryService
}

var Story StoryResource

func init() {
	Story := StoryResource{}
	Story.storyService = service.StoryService{StoryRepo: repository.StoryRepository{}}
}

func (resource *StoryResource) CreateStoryTasks(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	taskId, _ := strconv.Atoi(c.Param("taskId"))
	taskId, err := resource.storyService.CreateStoryTasks(storyId, &task)
	if err != nil {
		common.Error.Println("Task could not be created", err)
		//return err
	}
	common.Info.Println("Created task with id ", taskId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": task.Id})
	return

}

func (resource *StoryResource) UpdateStoryTasks(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	taskId, _ := strconv.Atoi(c.Param("taskId"))
	task.Id = taskId
	err := resource.storyService.UpdateStoryTasks(storyId, &task)

	if err != nil {
		common.Error.Println("Task could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Task updated successfully!"})
	return
}

func (resource *StoryResource) DeleteStoryTasks(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	taskId, _ := strconv.Atoi(c.Param("taskId"))
	err := resource.storyService.DeleteStoryTasks(storyId, taskId)
	if err != nil {
		common.Error.Println("Task could not be deleted ", taskId)
		//return err
	}
	common.Info.Println("Task deleted ", taskId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}

func (resource *StoryResource) GetStoryTasks(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	taskId, _ := strconv.Atoi(c.Param("taskId"))
	task, _ := resource.storyService.GetStoryTasks(storyId, taskId)
	common.Info.Println("task found %d", taskId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": task})
	return
}

func (resource *StoryResource) UnlinkStoryTasks(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	taskId, _ := strconv.Atoi(c.Param("taskId"))
	resource.storyService.UnlinkStoryTasks(storyId, taskId)
	return
}

func (resource *StoryResource) LinkStoryTasks(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	taskId, _ := strconv.Atoi(c.Param("taskId"))
	resource.storyService.LinkStoryTasks(storyId, taskId)
	return
}

func (resource *StoryResource) CreateStoryComments(c *gin.Context) {
	var comment models.Comment
	c.BindJSON(&comment)
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	commentId, err := resource.storyService.CreateStoryComments(storyId, &comment)
	if err != nil {
		common.Error.Println("Comment could not be created", err)
		//return err
	}
	common.Info.Println("Created comment with id ", commentId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": comment.Id})
	return

}

func (resource *StoryResource) UpdateStoryComments(c *gin.Context) {
	var comment models.Comment
	c.BindJSON(&comment)
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	comment.Id = commentId
	err := resource.storyService.UpdateStoryComments(storyId, &comment)

	if err != nil {
		common.Error.Println("Comment could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Comment updated successfully!"})
	return
}

func (resource *StoryResource) DeleteStoryComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	err := resource.storyService.DeleteStoryComments(storyId, commentId)
	if err != nil {
		common.Error.Println("Comment could not be deleted ", commentId)
		//return err
	}
	common.Info.Println("Comment deleted ", commentId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}

func (resource *StoryResource) GetStoryComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	comment, _ := resource.storyService.GetStoryComments(storyId, commentId)
	common.Info.Println("comment found %d", commentId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": comment})
	return
}

func (resource *StoryResource) UnlinkStoryComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	resource.storyService.UnlinkStoryComments(storyId, commentId)
	return
}

func (resource *StoryResource) LinkStoryComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	resource.storyService.LinkStoryComments(storyId, commentId)
	return
}

func (resource *StoryResource) CreateStoryPoComments(c *gin.Context) {
	var comment models.Comment
	c.BindJSON(&comment)
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	commentId, err := resource.storyService.CreateStoryPoComments(storyId, &comment)
	if err != nil {
		common.Error.Println("Comment could not be created", err)
		//return err
	}
	common.Info.Println("Created comment with id ", commentId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": comment.Id})
	return

}

func (resource *StoryResource) UpdateStoryPoComments(c *gin.Context) {
	var comment models.Comment
	c.BindJSON(&comment)
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	comment.Id = commentId
	err := resource.storyService.UpdateStoryPoComments(storyId, &comment)

	if err != nil {
		common.Error.Println("Comment could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Comment updated successfully!"})
	return
}

func (resource *StoryResource) DeleteStoryPoComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	err := resource.storyService.DeleteStoryPoComments(storyId, commentId)
	if err != nil {
		common.Error.Println("Comment could not be deleted ", commentId)
		//return err
	}
	common.Info.Println("Comment deleted ", commentId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}

func (resource *StoryResource) GetStoryPoComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	comment, _ := resource.storyService.GetStoryPoComments(storyId, commentId)
	common.Info.Println("comment found %d", commentId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": comment})
	return
}

func (resource *StoryResource) UnlinkStoryPoComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	resource.storyService.UnlinkStoryPoComments(storyId, commentId)
	return
}

func (resource *StoryResource) LinkStoryPoComments(c *gin.Context) {
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	resource.storyService.LinkStoryPoComments(storyId, commentId)
	return
}
