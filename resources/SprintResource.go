package resources

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample_golang_application/common"
	"sample_golang_application/models"
	"sample_golang_application/service"
	"strconv"
)

type SprintResource struct {
	//sprintService service.ISprintService
}

var Sprint SprintResource

var Service service.ISprintService

func init() {
	Service = &service.SprintServiceTxn{SprintService:&service.SprintService{}}
}

func (resource *SprintResource) CreateSprint(c *gin.Context) {
	var sprint models.Sprint
	c.BindJSON(&sprint)
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	sprintId, err := Service.CreateSprint(&sprint)
	if err != nil {
		common.Error.Println("Sprint could not be created", err)
		//return err
	}
	common.Info.Println("Created sprint with id ", sprintId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": sprint.Id})
	return

}

func (resource *SprintResource) GetSprint(c *gin.Context) {
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	sprint, _ := Service.GetSprint(sprintId)
	common.Info.Println("sprint found %d", sprintId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": sprint})
	return
}

func (resource *SprintResource) UpdateSprint(c *gin.Context) {
	var sprint models.Sprint
	c.BindJSON(&sprint)
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	sprint.Id = sprintId
	err := Service.UpdateSprint(&sprint)

	if err != nil {
		common.Error.Println("Sprint could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sprint updated successfully!"})
	return
}

func (resource *SprintResource) DeleteSprint(c *gin.Context) {
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	err := Service.DeleteSprint(sprintId)
	if err != nil {
		common.Error.Println("Sprint could not be deleted ", sprintId)
		//return err
	}
	common.Info.Println("Sprint deleted ", sprintId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}

func (resource *SprintResource) CreateSprintStories(c *gin.Context) {
	var story models.Story
	c.BindJSON(&story)
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	storyId, err := Service.CreateSprintStories(sprintId, &story)
	if err != nil {
		common.Error.Println("Story could not be created", err)
		//return err
	}
	common.Info.Println("Created story with id ", storyId)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": story.Id})
	return

}

func (resource *SprintResource) UpdateSprintStories(c *gin.Context) {
	var story models.Story
	c.BindJSON(&story)
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	story.Id = storyId
	err := Service.UpdateSprintStories(sprintId, &story)

	if err != nil {
		common.Error.Println("Story could not be updated", err)
		//return err
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story updated successfully!"})
	return
}

func (resource *SprintResource) DeleteSprintStories(c *gin.Context) {
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	err := Service.DeleteSprintStories(sprintId, storyId)
	if err != nil {
		common.Error.Println("Story could not be deleted ", storyId)
		//return err
	}
	common.Info.Println("Story deleted ", storyId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully! "})
	return
}

func (resource *SprintResource) GetSprintStories(c *gin.Context) {
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	story, _ := Service.GetSprintStories(sprintId, storyId)
	common.Info.Println("story found %d", storyId)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": story})
	return
}

func (resource *SprintResource) UnlinkSprintStories(c *gin.Context) {
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	Service.UnlinkSprintStories(sprintId, storyId)
	return
}

func (resource *SprintResource) LinkSprintStories(c *gin.Context) {
	sprintId, _ := strconv.Atoi(c.Param("sprintId"))
	storyId, _ := strconv.Atoi(c.Param("storyId"))
	Service.LinkSprintStories(sprintId, storyId)
	return
}
