package httpfiber

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"errors"
	"homework8/internal/models/ads"
	"homework8/internal/ports/dto"
	adsvc "homework8/internal/services/ad"
	usersvc "homework8/internal/services/user"
	"homework8/internal/usecase"
)

func handleAdError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, usecase.Err4002_001):
		c.Status(http.StatusBadRequest)
		return c.JSON(AdErrorResponse(err))
	case errors.Is(err, usecase.Err4002_002):
		c.Status(http.StatusBadRequest)
		return c.JSON(AdErrorResponse(err))
	case errors.Is(err, usecase.Err4002_003):
		c.Status(http.StatusBadRequest)
		return c.JSON(AdErrorResponse(err))
	case errors.Is(err, usecase.Err4001_002):
		c.Status(http.StatusForbidden)
		return c.JSON(AdErrorResponse(err))
	default:
		c.Status(http.StatusInternalServerError)
		return c.JSON(AdErrorResponse(err))
	}
}

// Метод для создания объявления (ad)
func createAd(adSvc adsvc.Service, userSvc usersvc.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody dto.CreateAdRequest
		err := c.BodyParser(&reqBody)
		if err != nil {
			return handleAdError(c, err)
		}
		uc := usecase.CreateAd(adSvc, userSvc)
		ad, err := uc.Execute(reqBody.ToAd(), reqBody.UserID)
		if err != nil {
			return handleAdError(c, err)
		}
		return c.JSON(AdSuccessResponse(&ad))
	}
}

// Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
func changeAdStatus(adSvc adsvc.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody dto.ChangeAdStatusRequest
		if err := c.BodyParser(&reqBody); err != nil {
			return handleAdError(c, err)
		}

		adID, err := c.ParamsInt("ad_id")
		if err != nil {
			return handleAdError(c, err)
		}

		uc := usecase.ChangeAdStatus(adSvc)
		ad, err := uc.Execute(int64(adID), reqBody.Published, int64(reqBody.UserID))
		if err != nil {
			return handleAdError(c, err)
		}

		return c.JSON(AdSuccessResponse(&ad))
	}
}

// Метод для обновления текста(Text) или заголовка(Title) объявления
func updateAd(adSvc adsvc.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody dto.UpdateAdRequest
		if err := c.BodyParser(&reqBody); err != nil {
			return handleAdError(c, err)
		}

		adID, err := c.ParamsInt("ad_id")
		if err != nil {
			return handleAdError(c, err)
		}

		uc := usecase.UpdateAd(adSvc)
		ad, err := uc.Execute(int64(adID), reqBody.Title, reqBody.Text, reqBody.UserID)
		if err != nil {
			return handleAdError(c, err)
		}

		return c.JSON(AdSuccessResponse(&ad))
	}
}

func listAd(adSvc adsvc.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
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
				return handleAdError(c, err)
			}
		}

		if publishedStr := c.Query("published"); publishedStr != "" {
			if p, err := strconv.ParseBool(publishedStr); err == nil {
				filter.SetPublished(p)
			} else {
				return handleAdError(c, err)
			}
		} else {
			// По дефолту выводим опубликованные
			filter.SetPublished(true)
		}

		if createdDate := c.Query("created_date"); createdDate != "" {
			// Проверка формата YYYY-MM-DD
			if _, err := time.Parse(time.DateOnly, createdDate); err != nil {
				return handleAdError(c, err)
			}
			filter.SetCreatedDate(createdDate)
		}

		uc := usecase.ListAd(adSvc)
		adList, err := uc.Execute(filter)
		if err != nil {
			return handleAdError(c, err)
		}
		return c.JSON(ListAdResponse(adList))
	}
}
