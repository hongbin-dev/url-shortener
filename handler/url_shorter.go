package handler

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"myapp/model"
	"net/http"
)

type SaveRequest struct {
	OriginUrl string
}

func (h *Handler) SaveUrl(c echo.Context) (err error) {
	saveRequest := new(SaveRequest)
	if err = c.Bind(saveRequest); err != nil {
		return err
	}

	targetUrl := Algorithm(saveRequest.OriginUrl)

	urls := &model.Urls{
		Id:         bson.NewObjectId(),
		OriginLink: saveRequest.OriginUrl,
		ShortLink:  targetUrl,
	}

	db := h.DB.Clone()
	defer db.Close()

	if err := db.DB("url-shortener").C("urls").Insert(urls); err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(200, urls)
}

func (h *Handler) FindUrl(c echo.Context) (err error) {
	sourceUrl := c.QueryParam("url")
	urls := new(model.Urls)

	db := h.DB.Clone()
	defer db.Close()

	if err := db.DB("url-shortener").C("urls").
		Find(bson.M{"origin_link": sourceUrl}).One(urls); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return err
	}
	return c.JSON(200, urls)
}

func Algorithm(originUrl string) string {
	return originUrl
}
