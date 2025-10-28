package pagination

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalData  int `json:"total_data"`
	TotalPages int `json:"total_pages"`
}

type Params struct {
	Page  int
	Limit int
	Skip  int
}

func GetPaginationParams(c *gin.Context) Params {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	skip := (page - 1) * limit
	return Params{
		Page:  page,
		Limit: limit,
		Skip:  skip,
	}
}

func NewMeta(totalData, page, limit int) Meta {
	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	return Meta{
		Page:       page,
		Limit:      limit,
		TotalData:  totalData,
		TotalPages: totalPages,
	}
}
