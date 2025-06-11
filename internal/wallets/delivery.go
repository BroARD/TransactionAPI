package wallets

import "github.com/labstack/echo/v4"

type Hanlders interface {
	GetByID() echo.HandlerFunc
}
