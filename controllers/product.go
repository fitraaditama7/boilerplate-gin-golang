package controllers

import (
	"database/sql"
	"golang-boilerplate/helper"
	"golang-boilerplate/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetProducts(c *gin.Context) {
	var (
		res      http.ResponseWriter = c.Writer
		products []structs.Product
	)
	err := idb.DB.Find(&products).Where("deleted_at", nil).Error
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
	} else {
		helper.Responses(res, http.StatusOK, products)
	}
}

func (idb *InDB) GetProduct(c *gin.Context) {
	var (
		res     http.ResponseWriter = c.Writer
		product []structs.Product
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&product).Where("deleted_at", nil).Error
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
	} else {
		helper.Responses(res, http.StatusOK, product)
	}
}

func (idb *InDB) CreateProduct(c *gin.Context) {
	var (
		product structs.Product
		res     http.ResponseWriter = c.Writer
	)

	c.BindJSON(product)

	err := idb.DB.Create(&product).Error
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
	} else {
		helper.Responses(res, http.StatusOK, product)
	}
}

func (idb *InDB) UpdateProduct(c *gin.Context) {
	var (
		res     http.ResponseWriter = c.Writer
		product structs.Product
	)

	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).Where("deleted_at", nil).First(&product).Error

	if err == sql.ErrNoRows {
		helper.ErrorCustomStatus(res, http.StatusNotFound, "Data Tidak Ditemukann")
	}
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&product)
		idb.DB.Save(&product)
		helper.Responses(res, http.StatusOK, product)
	}
}

func (idb *InDB) DeleteProduct(c *gin.Context) {
	var (
		res     http.ResponseWriter = c.Writer
		product structs.Product
	)

	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).Where("deleted_at", nil).First(&product).Error

	if err == sql.ErrNoRows {
		helper.ErrorCustomStatus(res, http.StatusNotFound, "Data tidak ditemukan")
	}
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusNotFound, err.Error())
	} else {
		product.DeletedAt = time.Now()
		idb.DB.Save(&product)
		helper.Responses(res, http.StatusOK, product)
	}
}
