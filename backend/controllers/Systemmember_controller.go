package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jirayuSai/app/ent"
	"github.com/jirayuSai/app/ent/systemmember"
)

// SystemmemberController defines the struct for the systemmember controller
type SystemmemberController struct {
	client *ent.Client
	router gin.IRouter
}
type Systemmember struct {
}

// CreateSystemmember handles POST requests for adding systemmember entities
// @Summary Create systemmember
// @Description Create ussystemmemberer
// @ID create-systemmember
// @Accept   json
// @Produce  json
// @Param systemmember body ent.Systemmember true "Systemmember entity"
// @Success 200 {object} ent.Systemmember
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemmembers [post]
func (ctl *SystemmemberController) CreateSystemmember(c *gin.Context) {
	obj := ent.Systemmember{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "systemmember binding failed",
		})
		return
	}

	sm, err := ctl.client.Systemmember.
		Create().
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sm)
}

// GetSystemmember handles GET requests to retrieve a systemmember entity
// @Summary Get a systemmember entity by ID
// @Description get systemmember by ID
// @ID get-systemmember
// @Produce  json
// @Param id path int true "Systemmember ID"
// @Success 200 {object} ent.Systemmember
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemmembers/{id} [get]
func (ctl *SystemmemberController) GetSystemmember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	sm, err := ctl.client.Systemmember.
		Query().
		Where(systemmember.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sm)
}

// ListSystemmember handles request to get a list of systemmember entities
// @Summary List systemmember entities
// @Description list systemmember entities
// @ID list-systemmember
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Systemmember
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemmembers [get]
func (ctl *SystemmemberController) ListSystemmember(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	systemmembers, err := ctl.client.Systemmember.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, systemmembers)
}

// DeleteSystemmember handles DELETE requests to delete a systemmember entity
// @Summary Delete a systemmember entity by ID
// @Description get systemmember by ID
// @ID delete-systemmember
// @Produce  json
// @Param id path int true "Systemmember ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemmembers/{id} [delete]
func (ctl *SystemmemberController) DeleteSystemmember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Systemmember.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateSystemmember handles PUT requests to update a systemmember entity
// @Summary Update a systemmember entity by ID
// @Description update systemmember by ID
// @ID update-systemmember
// @Accept   json
// @Produce  json
// @Param id path int true "Systemmember ID"
// @Param systemmember body ent.Systemmember true "Systemmember entity"
// @Success 200 {object} ent.Systemmember
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemmembers/{id} [put]
func (ctl *SystemmemberController) UpdateSystemmember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Systemmember{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "systemmember binding failed",
		})
		return
	}
	obj.ID = int(id)
	sm, err := ctl.client.Systemmember.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, sm)
}

// NewSystemmemberController creates and registers handles for the systemmember controller
func NewSystemmemberController(router gin.IRouter, client *ent.Client) *SystemmemberController {
	smc := &SystemmemberController{
		client: client,
		router: router,
	}
	smc.register()
	return smc
}

// InitSystemmemberController registers routes to the main engine
func (ctl *SystemmemberController) register() {
	systemmembers := ctl.router.Group("/systemmembers")

	systemmembers.GET("", ctl.ListSystemmember)

	// CRUD
	systemmembers.POST("", ctl.CreateSystemmember)
	systemmembers.GET(":id", ctl.GetSystemmember)
	systemmembers.PUT(":id", ctl.UpdateSystemmember)
	systemmembers.DELETE(":id", ctl.DeleteSystemmember)
}
