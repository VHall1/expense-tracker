package config

import (
	"net/http"

	"github.com/vhall1/expense-tracker/service.snowflake/handler"
	"github.com/vhall1/expense-tracker/service.snowflake/lib"
)

func RegisterRoutes(mux *http.ServeMux) error {
	// TODO: pull machine ID from config
	snowflakeService, err := lib.NewSnowflake(int64(0))
	snowflakeHandler := handler.NewSnowflakeHandler(snowflakeService)
	if err != nil {
		return err
	}

	mux.Handle("GET /snowflake", snowflakeHandler.HandleGenSnowflake())

	return nil
}
