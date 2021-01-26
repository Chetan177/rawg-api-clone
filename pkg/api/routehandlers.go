package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"rawg/pkg/model"
)

func (s *Server) getGames(c echo.Context) error {
	req := &model.GameRequest{
		Page:     1,
		PageSize: 5,
	}
	if err := c.Bind(req); err != nil {
		return s.logAndReturnResponse(c, err)
	}
	log.Debug("getGame req %+v", req)
	fmt.Printf("%+v", req)
	// serach based on slog only
	data := new(model.Game)
	if req.Search != "" {
		filter := bson.D{{"slug", req.Search}}
		if err := s.db.Get(filter, data); err != nil {
			return s.logAndReturnResponse(c, err)
		}
		return c.JSON(200, data)
	}

	return s.logAndReturnResponse(c, nil)
}

func (s *Server) logAndReturnResponse(c echo.Context, err error) error {
	var res model.Response
	switch err.(type) {
	case nil:
		res = StatusMap[StatusSuccess]
	case *echo.HTTPError:
		res = StatusMap[ValidationError]
	}

	if err != nil {
		log.Error("%s error %+v", res.Msg, err)
	}
	return c.JSON(res.Status, res)
}
