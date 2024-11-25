// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package lpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/service"
	uc_cs_lpc "github.com/enbility/eebus-go/usecases/cs/lpc"
	"github.com/enbility/eebus-grpc/rpc_server/usecase_server"
	"github.com/enbility/eebus-grpc/rpc_services/common_types"
	"github.com/enbility/eebus-grpc/rpc_services/usecases/cs/lpc"
	"github.com/enbility/eebus-grpc/utils/type_conversions"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"google.golang.org/grpc"
)

type Server struct {
	eebusService *service.Service
	uclpc        *uc_cs_lpc.LPC

	lpc.UnimplementedControllableSystemLPCControlServer
	usecase_server.BaseUseCaseServer

	grpcServer *grpc.Server
}

func NewServer(eebusService *service.Service, localEntity spineapi.EntityLocalInterface) *Server {
	server := &Server{
		eebusService:      eebusService,
		uclpc:             nil,
		BaseUseCaseServer: *usecase_server.NewBaseUseCaseServer(),
		grpcServer:        nil,
	}
	server.uclpc = uc_cs_lpc.NewLPC(localEntity, server.OnUseCaseEvent)
	eebusService.AddUseCase(server.uclpc)
	return server
}

func (h *Server) GetUseCase() api.UseCaseInterface {
	return h.uclpc
}

func (h *Server) ConsumptionLimit(_ context.Context, req *lpc.ConsumptionLimitRequest) (*lpc.ConsumptionLimitResponse, error) {
	eebus_limit, resultErr := h.uclpc.ConsumptionLimit()
	rpc_limit := type_conversions.ConvertEEBUSLoadLimit(&eebus_limit)

	res := &lpc.ConsumptionLimitResponse{
		LoadLimit: &rpc_limit,
	}
	return res, resultErr
}

func (h *Server) SetConsumptionLimit(_ context.Context, req *lpc.SetConsumptionLimitRequest) (*lpc.SetConsumptionLimitResponse, error) {
	eebus_limit := type_conversions.ConvertRPCLoadLimit(req.GetLoadLimit())
	h.uclpc.IsHeartbeatWithinDuration()
	resultErr := h.uclpc.SetConsumptionLimit(eebus_limit)

	return &lpc.SetConsumptionLimitResponse{}, resultErr
}

func (h *Server) PendingConsumptionLimit(_ context.Context, req *lpc.PendingConsumptionLimitRequest) (*lpc.PendingConsumptionLimitResponse, error) {
	eebus_load_limits := h.uclpc.PendingConsumptionLimits()
	rpc_load_limits := make(map[uint64]*common_types.LoadLimit)
	for eebus_id, eebus_limit := range eebus_load_limits {
		rpc_limit := type_conversions.ConvertEEBUSLoadLimit(&eebus_limit)
		rpc_id := uint64(eebus_id)
		rpc_load_limits[rpc_id] = &rpc_limit
	}

	res := &lpc.PendingConsumptionLimitResponse{
		LoadLimits: rpc_load_limits,
	}
	return res, nil
}

func (h *Server) ApproveOrDenyConsumptionLimit(_ context.Context, req *lpc.ApproveOrDenyConsumptionLimitRequest) (*lpc.ApproveOrDenyConsumptionLimitResponse, error) {
	h.uclpc.ApproveOrDenyConsumptionLimit(
		model.MsgCounterType(req.GetMsgCounter()),
		req.GetApprove(),
		req.GetReason(),
	)
	return &lpc.ApproveOrDenyConsumptionLimitResponse{}, nil
}

func (h *Server) FailsafeConsumptionActivePowerLimit(_ context.Context, req *lpc.FailsafeConsumptionActivePowerLimitRequest) (*lpc.FailsafeConsumptionActivePowerLimitResponse, error) {
	limit, isChangeable, resultErr := h.uclpc.FailsafeConsumptionActivePowerLimit()

	res := &lpc.FailsafeConsumptionActivePowerLimitResponse{
		Limit:        limit,
		IsChangeable: isChangeable,
	}

	return res, resultErr
}

func (h *Server) SetFailsafeConsumptionActivePowerLimit(_ context.Context, req *lpc.SetFailsafeConsumptionActivePowerLimitRequest) (*lpc.SetFailsafeConsumptionActivePowerLimitResponse, error) {
	resultErr := h.uclpc.SetFailsafeConsumptionActivePowerLimit(
		req.GetValue(),
		req.GetIsChangeable(),
	)
	return &lpc.SetFailsafeConsumptionActivePowerLimitResponse{}, resultErr
}

func (h *Server) FailsafeDurationMinimum(_ context.Context, req *lpc.FailsafeDurationMinimumRequest) (*lpc.FailsafeDurationMinimumResponse, error) {
	duration, isChangeable, resultErr := h.uclpc.FailsafeDurationMinimum()

	res := &lpc.FailsafeDurationMinimumResponse{
		DurationNanoseconds: duration.Nanoseconds(),
		IsChangeable:        isChangeable,
	}

	return res, resultErr
}

func (h *Server) SetFailsafeDurationMinimum(_ context.Context, req *lpc.SetFailsafeDurationMinimumRequest) (*lpc.SetFailsafeDurationMinimumResponse, error) {
	resultErr := h.uclpc.SetFailsafeDurationMinimum(
		time.Duration(req.GetDurationNanoseconds())*time.Nanosecond,
		req.GetIsChangeable(),
	)

	return &lpc.SetFailsafeDurationMinimumResponse{}, resultErr
}

func (h *Server) StartHeartbeat(_ context.Context, req *lpc.StartHeartbeatRequest) (*lpc.StartHeartbeatResponse, error) {
	h.uclpc.StartHeartbeat()

	return &lpc.StartHeartbeatResponse{}, nil
}

func (h *Server) StopHeartbeat(_ context.Context, req *lpc.StopHeartbeatRequest) (*lpc.StopHeartbeatResponse, error) {
	h.uclpc.StopHeartbeat()
	return &lpc.StopHeartbeatResponse{}, nil
}

func (h *Server) IsHeartbeatWithinDuration(_ context.Context, req *lpc.IsHeartbeatWithinDurationRequest) (*lpc.IsHeartbeatWithinDurationResponse, error) {
	is_within_duration := h.uclpc.IsHeartbeatWithinDuration()
	res := &lpc.IsHeartbeatWithinDurationResponse{
		IsWithinDuration: is_within_duration,
	}
	return res, nil
}

func (h *Server) ConsumptionNominalMax(_ context.Context, req *lpc.ConsumptionNominalMaxRequest) (*lpc.ConsumptionNominalMaxResponse, error) {
	value, resultErr := h.uclpc.ConsumptionNominalMax()

	res := &lpc.ConsumptionNominalMaxResponse{
		Value: value,
	}

	return res, resultErr
}

func (h *Server) SetConsumptionNominalMax(_ context.Context, req *lpc.SetConsumptionNominalMaxRequest) (*lpc.SetConsumptionNominalMaxResponse, error) {
	resultErr := h.uclpc.SetConsumptionNominalMax(req.GetValue())

	return &lpc.SetConsumptionNominalMaxResponse{}, resultErr
}

func (h *Server) Start(port *int) (int, error) {
	var res int
	if port != nil {
		return -1, fmt.Errorf("port is specifiedm but not allowed")
	}
	if h.grpcServer != nil {
		return -1, fmt.Errorf("grpc server is already running")
	}
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	res = lis.Addr().(*net.TCPAddr).Port
	fmt.Printf("server running at port %d\n", res)

	if err != nil {
		return -1, fmt.Errorf("failed to listen: %v", err)
	}
	h.grpcServer = grpc.NewServer()
	lpc.RegisterControllableSystemLPCControlServer(h.grpcServer, h)
	go func() {
		if err := h.grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()

	return res, nil
}

func (h *Server) Stop() error {
	if h.grpcServer == nil {
		return fmt.Errorf("grpc server is not running")
	}
	h.grpcServer.Stop()
	h.grpcServer = nil

	return nil
}
