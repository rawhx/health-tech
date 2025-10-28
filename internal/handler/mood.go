package handler

import (
	"health-tech/internal/dto"
	// "health-tech/models"
	"health-tech/pkg/pagination"
	"health-tech/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateMood(c *gin.Context) {
	req := dto.CreateMood{}
	// user := c.MustGet("user").(*models.User)
	err := c.ShouldBind(&req)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "gagal melakukan input", err)
		return
	}

	err = r.services.MoodService.CreateMood(req.UserID, &req)
	
	if err != nil {
		if err.Error() == "tanggal tidak boleh di masa depan" {
			utils.ResponseError(c, http.StatusBadRequest, "tanggal tidak boleh di masa depan", err)
			return
		} else if err.Error() == "mood score harus antara 1 dan 5" {
			utils.ResponseError(c, http.StatusBadRequest, "mood score harus antara 1 dan 5", err)
			return
		}
		utils.ResponseError(c, http.StatusInternalServerError, "internal server error", err)
		return	
	}
	utils.ResponseSuccess(c, http.StatusCreated, "berhasil menambahkan mood", nil)
	return	
}

func (r *Rest) GetUserMoods(c *gin.Context) {
	userID := c.Param("user_id")
	params := pagination.GetPaginationParams(c)

	data, total, err := r.services.MoodService.GetUserMoods(userID, params)
	if err != nil {
		if err.Error() == "data tidak ditemukan" {
			utils.ResponseError(c, http.StatusNotFound, "data user tidak ditemukan", err)
			return
		}
		utils.ResponseError(c, http.StatusInternalServerError, "internal server error", err)
		return	
	}

	meta := pagination.NewMeta(total, params.Page, params.Limit)

	response := dto.DataMoodPagination{
		Meta: meta,
		Moods: data,
	}

	utils.ResponseSuccess(c, http.StatusOK, "berhasil mendapatkan data moods", response)
	return
}

func (r *Rest) GetMoodSummary(c *gin.Context) {
	userID := c.Param("user_id")
	period := c.Query("period") 

	dt, err := r.services.MoodService.GetMoodSummary(userID, period)
	if err != nil {
		if err.Error() == "data tidak ditemukan" {
			utils.ResponseError(c, http.StatusNotFound, "data user tidak ditemukan", err)
			return
		}
		utils.ResponseError(c, http.StatusInternalServerError, "internal server error", err)
		return	
	}

	utils.ResponseSuccess(c, http.StatusOK, "berhasil mendapatkan mood summary", dt)
	return
}
