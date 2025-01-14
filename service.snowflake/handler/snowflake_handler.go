package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/http"

	"github.com/vhall1/expense-tracker/service.snowflake/lib"
)

type SnowflakeHandler struct {
	SnowflakeService *lib.Snowflake
}

func NewSnowflakeHandler(snowflakeService *lib.Snowflake) *SnowflakeHandler {
	return &SnowflakeHandler{
		SnowflakeService: snowflakeService,
	}
}

func (h *SnowflakeHandler) HandleGenSnowflake() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := h.SnowflakeService.GenerateID()

		// Create a byte buffer to store the binary representation
		var buf bytes.Buffer
		// Write the int64 into the buffer in big-endian byte order
		err := binary.Write(&buf, binary.BigEndian, id)
		if err != nil {
			fmt.Println("Error writing int64 to buffer:", err)
		}

		fmt.Fprint(w, base64.StdEncoding.EncodeToString(buf.Bytes()))
	})
}
