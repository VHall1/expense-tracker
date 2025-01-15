package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"net/http"

	"github.com/vhall1/expense-tracker/common/grpc/snowflake"
	"github.com/vhall1/expense-tracker/common/util"
	"github.com/vhall1/expense-tracker/service.snowflake/types"
)

type SnowflakeHttpHandler struct {
	snowflakeService types.SnowflakeService
}

func NewHttpSnowflakeService(snowflakeService types.SnowflakeService) *SnowflakeHttpHandler {
	return &SnowflakeHttpHandler{snowflakeService: snowflakeService}
}

func (h *SnowflakeHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", h.GenerateID)
}

func (h *SnowflakeHttpHandler) GenerateID(w http.ResponseWriter, r *http.Request) {
	id := h.snowflakeService.GenerateID(r.Context())

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, id)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &snowflake.GenerateIDResponse{Id: base64.StdEncoding.EncodeToString(buf.Bytes())}
	util.WriteJSON(w, http.StatusOK, res)
}
