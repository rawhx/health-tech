package handler

import (
	"health-tech/internal/dto"
	"health-tech/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(c *gin.Context) {
	req := dto.UserCreateRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "gagal melakukan input", err)
		return
	}

	err = r.services.UserService.CreateUser(req)
	if err != nil {
		if err.Error() == "email sudah terdaftar" {
			utils.ResponseError(c, http.StatusBadRequest, "gagal membuat user baru", err)
			return
		} else if err.Error() == "gagal melakukan hash password" {
			utils.ResponseError(c, http.StatusInternalServerError, "gagal melakukan hash", err)
			return
		}
		utils.ResponseError(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "berhasil membuat user baru", nil)
}

func (r *Rest) Login(c *gin.Context) {
	req := dto.UserLoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "gagal melakukan input", err)
		return
	}

	data, err := r.services.UserService.Login(req)
	if err != nil {
		if err.Error() == "gagal cek password" {
			utils.ResponseError(c, http.StatusBadRequest, "email atau password salah", err)
			return
		} else if err.Error() == "data tidak ditemukan" {
			utils.ResponseError(c, http.StatusNotFound, "data user tidak ditemukan", err)
			return
		} else if err.Error() == "gagal membuat token" {
			utils.ResponseError(c, http.StatusInternalServerError, "gagal membuat token", err)
			return
		}
		utils.ResponseError(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, "berhasil melakukan login", data)
}