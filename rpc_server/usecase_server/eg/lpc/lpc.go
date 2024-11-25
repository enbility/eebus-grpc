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
	uc_eg_lpc "github.com/enbility/eebus-go/usecases/eg/lpc"
	"github.com/enbility/eebus-grpc/rpc_server/usecase_server"
	"github.com/enbility/eebus-grpc/rpc_services/usecases/eg/lpc"
	"github.com/enbility/eebus-grpc/utils/type_conversions"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"google.golang.org/grpc"
)

type Server struct {
	eebusService *service.Service
	uclpc        *uc_eg_lpc.LPC

	lpc.UnimplementedEnergyGuardLPCControlServer
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
	server.uclpc = uc_eg_lpc.NewLPC(localEntity, server.OnUseCaseEvent)
	eebusService.AddUseCase(server.uclpc)
	return server
}

func (h *Server) GetUseCase() api.UseCaseInterface {
	return h.uclpc
}

func (h *Server) ConsumptionLimit(_ context.Context, req *lpc.ConsumptionLimitRequest) (*lpc.ConsumptionLimitResponse, error) {
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))
	eebus_limit, resultErr := h.uclpc.ConsumptionLimit(remoteEntity)
	rpc_limit := type_conversions.ConvertEEBUSLoadLimit(&eebus_limit)
	return &lpc.ConsumptionLimitResponse{
		Limit: &rpc_limit,
	}, resultErr
}

func (h *Server) WriteConsumptionLimit(_ context.Context, req *lpc.WriteConsumptionLimitRequest) (*lpc.WriteConsumptionLimitResponse, error) {
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))

	resultCh := make(chan model.ResultDataType)
	resultCB := func(result model.ResultDataType) {
		resultCh <- result
	}

	msgCounter, resultErr := h.uclpc.WriteConsumptionLimit(
		remoteEntity,
		type_conversions.ConvertRPCLoadLimit(req.GetLimit()),
		resultCB,
	)

	result := <-resultCh
	errorNumber := uint64(*result.ErrorNumber)

	return &lpc.WriteConsumptionLimitResponse{
		ErrorNumber: &errorNumber,
		Description: (*string)(result.Description),
		MsgCounter:  uint64(*msgCounter),
	}, resultErr

}
func (h *Server) FailsafeConsumptionActivePowerLimit(_ context.Context, req *lpc.FailsafeConsumptionActivePowerLimitRequest) (*lpc.FailsafeConsumptionActivePowerLimitResponse, error) {
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))

	limit, resultErr := h.uclpc.FailsafeConsumptionActivePowerLimit(remoteEntity)

	return &lpc.FailsafeConsumptionActivePowerLimitResponse{
		Limit: limit,
	}, resultErr
}

func (h *Server) WriteFailsafeConsumptionActivePowerLimit(_ context.Context, req *lpc.WriteFailsafeConsumptionActivePowerLimitRequest) (*lpc.WriteFailsafeConsumptionActivePowerLimitResponse, error) {
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))

	msgCounter, resultErr := h.uclpc.WriteFailsafeConsumptionActivePowerLimit(
		remoteEntity,
		req.GetLimit(),
	)

	return &lpc.WriteFailsafeConsumptionActivePowerLimitResponse{
		MsgCounter: uint64(*msgCounter),
	}, resultErr
}

func (h *Server) FailsafeDurationMinimum(_ context.Context, req *lpc.FailsafeDurationMinimumRequest) (*lpc.FailsafeDurationMinimumResponse, error) {
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))

	duration, resultErr := h.uclpc.FailsafeDurationMinimum(remoteEntity)

	return &lpc.FailsafeDurationMinimumResponse{
		DurationNanoseconds: uint64(duration.Nanoseconds()),
	}, resultErr
}

func (h *Server) WriteFailsafeDurationMinimum(_ context.Context, req *lpc.WriteFailsafeDurationMinimumRequest) (*lpc.WriteFailsafeDurationMinimumResponse, error) {
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))

	msgCounter, resultErr := h.uclpc.WriteFailsafeDurationMinimum(
		remoteEntity,
		time.Duration(req.GetDurationNanoseconds())*time.Nanosecond,
	)

	return &lpc.WriteFailsafeDurationMinimumResponse{
		MsgCounter: uint64(*msgCounter),
	}, resultErr
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
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))

	withinDuration := h.uclpc.IsHeartbeatWithinDuration(remoteEntity)

	return &lpc.IsHeartbeatWithinDurationResponse{
		WithinDuration: withinDuration,
	}, nil
}

func (h *Server) ConsumptionNominalMax(_ context.Context, req *lpc.ConsumptionNominalMaxRequest) (*lpc.ConsumptionNominalMaxResponse, error) {
	remoteDevice := h.eebusService.LocalDevice().RemoteDeviceForSki(req.GetRemoteSki())
	remoteEntity := remoteDevice.Entity(type_conversions.ConvertRPCEntityAddress(req.GetRemoteEntityAddress()))

	nominalMax, resultErr := h.uclpc.ConsumptionNominalMax(remoteEntity)

	return &lpc.ConsumptionNominalMaxResponse{
		ConsumptionNominalMax: nominalMax,
	}, resultErr
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
	lpc.RegisterEnergyGuardLPCControlServer(h.grpcServer, h)
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
