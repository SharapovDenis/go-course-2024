package httpgin

import (
	"errors"
	"homework8/internal/models/ads"
	"homework8/internal/ports/dto"
	adsvc "homework8/internal/services/ad"
	usersvc "homework8/internal/services/user"
	"homework8/internal/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func handleAdError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, usecase.Err4002_001):
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
	case errors.Is(err, usecase.Err4002_002):
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
	case errors.Is(err, usecase.Err4002_003):
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
	case errors.Is(err, usecase.Err4001_002):
		c.JSON(http.StatusForbidden, AdErrorResponse(err))
	default:
		c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
	}
}

func createAd(adSvc adsvc.Service, userSvc usersvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody dto.CreateAdRequest
		if err := c.BindJSON(&reqBody); err != nil {
			handleAdError(c, err)
			return
		}
		uc := usecase.CreateAd(adSvc, userSvc)
		ad, err := uc.Execute(reqBody.ToAd(), reqBody.UserID)
		if err != nil {
			handleAdError(c, err)
			return
		}
		c.JSON(http.StatusOK, AdSuccessResponse(&ad))
	}
}

func changeAdStatus(adSvc adsvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody dto.ChangeAdStatusRequest
		if err := c.BindJSON(&reqBody); err != nil {
			handleAdError(c, err)
			return
		}

		strAdID := c.Param("ad_id")
		if strAdID == "" {
			handleAdError(c, errors.New("ai_id is empty"))
			return
		}

		intAdID, err := strconv.ParseInt(strAdID, 10, 64)
		if err != nil {
			handleAdError(c, err)
			return
		}

		uc := usecase.ChangeAdStatus(adSvc)
		ad, err := uc.Execute(intAdID, reqBody.Published, int64(reqBody.UserID))
		if err != nil {
			handleAdError(c, err)
			return
		}

		c.JSON(http.StatusOK, AdSuccessResponse(&ad))
	}
}

func updateAd(adSvc adsvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody dto.UpdateAdRequest
		if err := c.BindJSON(&reqBody); err != nil {
			handleAdError(c, err)
			return
		}

		strAdID := c.Param("ad_id")
		if strAdID == "" {
			handleAdError(c, errors.New("empty ai_id"))
			return
		}

		intAdID, err := strconv.ParseInt(strAdID, 10, 64)
		if err != nil {
			handleAdError(c, err)
			return
		}

		uc := usecase.UpdateAd(adSvc)
		ad, err := uc.Execute(intAdID, reqBody.Title, reqBody.Text, reqBody.UserID)
		if err != nil {
			handleAdError(c, err)
			return
		}
		c.JSON(http.StatusOK, AdSuccessResponse(&ad))
	}
}

func listAd(adSvc adsvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := ads.NewFilter()

		if title := c.Query("title"); title != "" {
			filter.SetTitle(title)
		}

		if text := c.Query("title"); text != "" {
			filter.SetText(text)
		}

		if authorIDStr := c.Query("author_id"); authorIDStr != "" {
			if id, err := strconv.ParseInt(authorIDStr, 10, 64); err == nil {
				filter.SetAuthorID(id)
			} else {
				handleAdError(c, err)
				return
			}
		}

		if publishedStr := c.Query("published"); publishedStr != "" {
			if p, err := strconv.ParseBool(publishedStr); err == nil {
				filter.SetPublished(p)
			} else {
				handleAdError(c, err)
				return
			}
		} else {
			// По дефолту выводим опубликованные
			filter.SetPublished(true)
		}

		if createdDate := c.Query("created_date"); createdDate != "" {
			// Проверка формата YYYY-MM-DD
			if _, err := time.Parse(time.DateOnly, createdDate); err != nil {
				handleAdError(c, err)
				return
			}
			filter.SetCreatedDate(createdDate)
		}

		uc := usecase.ListAd(adSvc)
		adList, err := uc.Execute(filter)
		if err != nil {
			handleAdError(c, err)
			return
		}
		c.JSON(http.StatusOK, ListAdResponse(adList))
	}
}
