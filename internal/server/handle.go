/*
 * Copyright 2022 by Mel2oo <https://github.com/saferun/owl>
 *
 * Licensed under the GNU General Public License version 3 (GPLv3)
 *
 * If you distribute GPL-licensed software the license requires
 * that you also distribute the complete, corresponding source
 * code (as defined by GPL) to that GPL-licensed software.
 *
 * You should have received a copy of the GNU General Public License
 * with this program. If not, see <https://www.gnu.org/licenses/>
 */

package server

import (
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) register() error {
	r := s.Group("/owl/v1")

	{
		r.POST("/submit", s.submit)
		r.POST("/restlt", s.result)
	}

	return nil
}

func (s *server) submit(c echo.Context) error {
	type request struct {
		File *multipart.FileHeader `form:"file"`
	}

	type resppnse struct {
		Code int `json:"code"`
	}

	var req request

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusOK, &resppnse{ParamErrorCode})
	}

	return c.JSON(http.StatusOK, &resppnse{RequestOk})
}

func (s *server) result(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
