package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jirayuSai/app/ent"
	"github.com/jirayuSai/app/ent/mmedicine"
)

// MmedicineController defines the struct for the mmedicine controller
type MmedicineController struct {
	client *ent.Client
	router gin.IRouter
}

//CreateMmedicine handles POST requests for adding mmedicine entities
// @Summary Create mmedicine
// @Description Create mmedicine
// @ID create-mmedicine
// @Accept   json
// @Produce  json
// @Param mmedicine body ent.Mmedicine true "Mmedicine entity"
// @Success 200 {object} ent.Mmedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mmedicines [post]
func (ctl *MmedicineController) CreateMmedicine(c *gin.Context) {
	obj := ent.Mmedicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "mmedicine binding failed",
		})
		return
	}

	me, err := ctl.client.Mmedicine.
		Create().
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, me)
}

// GetMmedicine handles GET requests to retrieve a mmedicine entity
// @Summary Get a mmedicine entity by ID
// @Description get mmedicine by ID
// @ID get-mmedicine
// @Produce  json
// @Param id path int true "Mmedicine ID"
// @Success 200 {object} ent.Mmedicine
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mmedicines/{id} [get]
func (ctl *MmedicineController) GetMmedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	me, err := ctl.client.Mmedicine.
		Query().
		Where(mmedicine.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, me)
}

// ListMmedicine handles request to get a list of mmedicine entities
// @Summary List mmedicine entities
// @Description list mmedicine entities
// @ID list-mmedicine
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Mmedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mmedicines [get]
func (ctl *MmedicineController) ListMmedicine(c *gin.Context) {
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

	mmedicines, err := ctl.client.Mmedicine.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, mmedicines)
}

// DeleteMmedicine handles DELETE requests to delete a mmedicine entity
// @Summary Delete a mmedicine entity by ID
// @Description get mmedicine by ID
// @ID delete-mmedicine
// @Produce  json
// @Param id path int true "Mmedicine ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mmedicines/{id} [delete]
func (ctl *MmedicineController) DeleteMmedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Mmedicine.
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

// UpdateMmedicine handles PUT requests to update a mmedicine entity
// @Summary Update a mmedicine entity by ID
// @Description update mmedicine by ID
// @ID update-mmedicine
// @Accept   json
// @Produce  json
// @Param id path int true "Mmedicine ID"
// @Param mmedicine body ent.Mmedicine true "Mmedicine entity"
// @Success 200 {object} ent.Mmedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mmedicines/{id} [put]
func (ctl *MmedicineController) UpdateMmedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Mmedicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "mmedicine binding failed",
		})
		return
	}
	obj.ID = int(id)
	pt, err := ctl.client.Mmedicine.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, pt)
}

// NewMmedicineController creates and registers handles for the mmedicine controller
func NewMmedicineController(router gin.IRouter, client *ent.Client) *MmedicineController {
	ptc := &MmedicineController{
		client: client,
		router: router,
	}
	ptc.register()
	return ptc
}

// InitMmedicineController registers routes to the main engine
func (ctl *MmedicineController) register() {
	mmedicines := ctl.router.Group("/mmedicines")

	mmedicines.GET("", ctl.ListMmedicine)

	// CRUD
	mmedicines.POST("", ctl.CreateMmedicine)
	mmedicines.GET(":id", ctl.GetMmedicine)
	mmedicines.PUT(":id", ctl.UpdateMmedicine)
	mmedicines.DELETE(":id", ctl.DeleteMmedicine)
}
