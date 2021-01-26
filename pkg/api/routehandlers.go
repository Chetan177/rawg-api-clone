package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"rawg/pkg/model"
)

func (s *Server) getGames(c echo.Context) error {
	req := new(model.GameRequest)
	if err := c.Bind(req); err != nil {
		return s.logAndReturnResponse(c, err)
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
