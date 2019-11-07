package services

import (
	"context"
	"log"

	"github.com/yeyus/dmr-entities/internal/pkg/db"
	"github.com/yeyus/dmr-entities/pkg/api"
	"github.com/yeyus/go-grpc-interceptor/xrequestid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EntitiesServer struct {
	DB *db.DBManager
}

// GetEntity - Handler for GetEntity rpc call
func (s *EntitiesServer) GetEntity(context context.Context, request *api.GetEntityRequest) (*api.Entity, error) {
	requestID := xrequestid.FromContext(context)
	entity, err := s.DB.GetEntity(request.GetId())
	if err != nil {
		log.Printf("[GetEntity] {%s} error while fetching from db: %s", requestID, err.Error())
	}

	log.Printf("[GetEntity] {%s} for ID %d", requestID, request.GetId())
	return entity, err
}

// SearchEntity - Handler for SearchEntity rpc call
func (s *EntitiesServer) SearchEntity(context.Context, *api.SearchEntitiesRequest) (*api.EntitiesResponse, error) {
	return nil, status.New(codes.Unimplemented, "method not implemented").Err()
}

// GetEntitiesBySystemID - Handler for GetEntitiesBySystemID rpc call
func (s *EntitiesServer) GetEntitiesBySystemID(context.Context, *api.GetEntitiesBySystemIDRequest) (*api.EntitiesResponse, error) {
	return nil, status.New(codes.Unimplemented, "method not implemented").Err()
}

// GetEntitiesByCallsign - Handler for GetEntitiesByCallsign rpc call
func (s *EntitiesServer) GetEntitiesByCallsign(context.Context, *api.GetEntitiesByCallsignRequest) (*api.EntitiesResponse, error) {
	return nil, status.New(codes.Unimplemented, "method not implemented").Err()
}
