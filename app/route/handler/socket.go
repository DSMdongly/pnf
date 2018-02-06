package handler

import (
	"github.com/DSMdongly/pnf/socket"

	"github.com/labstack/echo"
)

func Socket() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		con, err := socket.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)

		if err != nil {
			ctx.Logger().Error(err)
			return err
		}

		cli := socket.NewClient(con)
		defer cli.Close()

		cli.Handle()
		
		return nil
	}
}
