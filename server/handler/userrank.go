package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//----------------------------------------------------------------

// 全データベースを読み取るAPI ただし制御機構付
func (h *Handler) GetMessageCounts(c echo.Context) error {
	//クエリパラメタでページ制御 50*page
	pagestr := c.QueryParam("page")
	if pagestr == "" {
		return c.JSON(http.StatusOK, h.nowhavingdata)
	}
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, h.nowhavingdata[page*50:(page+1)*50])
}

// グループ単位で読み取り グループuuidをパスパラメータに
func (h *Handler) GetMessageCountsWithGroup(c echo.Context) error {
	//グループ探索
	groupid := c.Param("groupid")
	res, httpres, err := h.GetGroupMembers(groupid)
	if err != nil {
		return c.String(httpres.StatusCode, err.Error())
	}
	//クエリパラメタでページ制御 50*page
	pagestr := c.QueryParam("page")
	if pagestr == "" {
		return c.JSON(httpres.StatusCode, res)
	}
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(httpres.StatusCode, res[page*50:(page+1)*50])
}
