package endpoint

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "log"
    "net/http"
    "simpleHttp/internal/pkg/constants"
    "time"
)

type Service interface {
    DaysLeft() int32
}

type Endpoint struct {
    s Service
}

func NonConst() time.Time{
    localTime := time.Now()
    return localTime
}

func New(s Service) *Endpoint  {
    return &Endpoint{
        s: s,
    }
}

func (e *Endpoint) Status(ctx echo.Context) error {

    d := e.s.DaysLeft()

    s := fmt.Sprint(constants.DaysLeft, constants.Space, d)

    if err := ctx.String(http.StatusOK, s); err != nil {
        return err
    }

    return nil
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
    return func(ctx echo.Context) error {

        val := ctx.Request().Header.Get("User-Role")

        if val == constants.RoleA {
            log.Println(constants.RoleAText)
        } else  {
            log.Println(constants.RoleNoAText)
        }

        if err := next(ctx); err != nil {
            return err
        }

        return nil
    }
}
/*func MustLoad() int{
YearInt := 2025
return YearInt
}*/

/*func parceDate()  {
dt := "2020-06-25"
dt1, _ := time.Parse("2006-01-02", dt)
fmt.Println("mydate:", dt1.Add(1*time.Hour).Format("02.01.2006 3:04pm"))
}*/
