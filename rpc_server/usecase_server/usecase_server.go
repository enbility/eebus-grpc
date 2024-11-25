package usecase_server

import (
	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-grpc/rpc_server"
	"github.com/enbility/eebus-grpc/rpc_services/control_service"
	log "github.com/enbility/eebus-grpc/utils/logging"
	"github.com/enbility/eebus-grpc/utils/type_conversions"
	spineapi "github.com/enbility/spine-go/api"
)

type UseCaseServer interface {
	OnUseCaseEvent(ski string, device spineapi.DeviceRemoteInterface, entity spineapi.EntityRemoteInterface, event api.EventType)
	GetUseCase() api.UseCaseInterface
	GetEvents() chan *control_service.SubscribeUseCaseEventsResponse
	rpc_server.RPCServer
}

type BaseUseCaseServer struct {
	events chan *control_service.SubscribeUseCaseEventsResponse

	log.Logger
}

func NewBaseUseCaseServer() *BaseUseCaseServer {
	return &BaseUseCaseServer{
		events: make(chan *control_service.SubscribeUseCaseEventsResponse),
	}
}

func (h *BaseUseCaseServer) OnUseCaseEvent(ski string, device spineapi.DeviceRemoteInterface, entity spineapi.EntityRemoteInterface, event api.EventType) {
	remoteEntityAddress := type_conversions.ConvertEEBUSEntityAddress(entity.Address().Entity)
	useCaseEvent := type_conversions.ConvertEEBUSEventType(event)
	useCaseEventResponse := control_service.SubscribeUseCaseEventsResponse{
		RemoteSki:           ski,
		RemoteEntityAddress: remoteEntityAddress,
		UseCaseEvent:        &useCaseEvent,
	}
	h.events <- &useCaseEventResponse
}

func (h *BaseUseCaseServer) GetEvents() chan *control_service.SubscribeUseCaseEventsResponse {
	return h.events
}
