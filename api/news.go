package api

import (
	"encoding/json"
	db "github.com/Reoptima/GoLearnPrject/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createNewsRequest struct {
	Text json.RawMessage `json:"text"`
}

type News struct {
	ID   int64  `json:"ID"`
	Text string `json:"text"`
}

func (server *Server) createNews(ctx *gin.Context) {
	var req createNewsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createdNews, err := server.store.CreateNews(ctx, string(req.Text))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createdNews)
}

type getNewsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getNews(ctx *gin.Context) {
	var req getNewsRequest
	newsList, err := server.store.GetNews(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newsList)
}

type listNewsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listNews(ctx *gin.Context) {
	var req listNewsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListNewsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	newsList, err := server.store.ListNews(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newsList)
}
