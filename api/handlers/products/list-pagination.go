package products

import (
	"FinalProject/dtos"
	"FinalProject/models"
	"FinalProject/utils"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ListProductRes struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	UserId uint   `json:"user_id"`
}

func (h *ProductHandler) Pagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := http.StatusOK
		currentUserId, ok := ctx.Get("currentUserId")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You need to login to get detail"})
			return
		}
		pagination := utils.GeneratePaginationRequest(ctx)
		response := Pagination(h, ctx, pagination, currentUserId.(uint))
		if !response.Success {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, response)
	}
}

func Pagination(h *ProductHandler, ctx *gin.Context, pagination *dtos.Pagination, userId uint) dtos.Response {
	operationResult, totalPages := h.GetPagination(pagination, userId)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dtos.Pagination)

	// get current url path
	urlPath := ctx.Request.URL.Path

	// search query params
	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return dtos.Response{Success: true, Data: data}
}

func (h *ProductHandler) GetPagination(pagination *dtos.Pagination, userId uint) (RepositoryResult, int) {
	var products []models.Product

	var totalRows int64 = 0
	totalPages, fromRow, toRow := 0, 0, 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := h.Db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	for _, value := range searchs {
		column := value.Column
		action := value.Action
		query := value.Query

	L: // fix warning
		switch action {
		case "equals":
			whereQuery := fmt.Sprintf("%s = ?", column)
			find = find.Where(whereQuery, query)
			break L
		case "contains":
			whereQuery := fmt.Sprintf("%s LIKE ?", column)
			find = find.Where(whereQuery, "%"+query+"%")
			break L
		case "in":
			whereQuery := fmt.Sprintf("%s IN (?)", column)
			queryArray := strings.Split(query, ",")
			find = find.Where(whereQuery, queryArray)
			break L
		}
	}

	find = find.Where("user_id = ?", userId).Find(&products)

	// clean data to return
	datas := []ListProductRes{}
	for _, product := range products {
		datas = append(datas, ListProductRes{
			ID:     product.ID,
			Name:   product.Name,
			Price:  product.Price,
			UserId: product.UserID,
		})
	}

	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = datas

	// count all data
	errCount := h.Db.Model(&models.Product{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = int(totalRows)

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
