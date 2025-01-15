package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"

	"github.com/vhall1/expense-tracker/common/grpc/snowflake"
	"github.com/vhall1/expense-tracker/service.snowflake/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SnowflakeGrpcHandler struct {
	snowflakeService types.SnowflakeService
	snowflake.UnimplementedSnowflakeServiceServer
}

func NewGrpcSnowflakeService(grpc *grpc.Server, snowflakeService types.SnowflakeService) {
	h := &SnowflakeGrpcHandler{snowflakeService: snowflakeService}
	snowflake.RegisterSnowflakeServiceServer(grpc, h)
}

func (h *SnowflakeGrpcHandler) GenerateID(ctx context.Context, _ *emptypb.Empty) (*snowflake.GenerateIDResponse, error) {
	id := h.snowflakeService.GenerateID(ctx)

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, id)
	if err != nil {
		return nil, err
	}

	res := &snowflake.GenerateIDResponse{Id: base64.StdEncoding.EncodeToString(buf.Bytes())}

	return res, nil
}
