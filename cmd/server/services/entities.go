package services

import (
	"context"
	"log"

	"github.com/yeyus/dmr-entities/internal/pkg/db"
	"github.com/yeyus/dmr-entities/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EntitiesServer struct {
	DB *db.DBManager
}

func (s *EntitiesServer) GetEntity(context context.Context, request *api.GetEntityRequest) (*api.Entity, error) {
	requestID := "unknown"
	entity, err := s.DB.GetEntity(request.GetId())
	if err != nil {
		log.Printf("[GetEntity] {%} error while fetching from db: %s", requestID, err.Error())
	}

	return entity, err
}

func (s *EntitiesServer) SearchEntity(context.Context, *api.SearchEntitiesRequest) (*api.EntitiesResponse, error) {
	return nil, status.New(codes.Unimplemented, "method not implemented").Err()
}

func (s *EntitiesServer) GetEntitiesBySystemID(context.Context, *api.GetEntitiesBySystemIDRequest) (*api.EntitiesResponse, error) {
	return nil, status.New(codes.Unimplemented, "method not implemented").Err()
}

func (s *EntitiesServer) GetEntitiesByCallsign(context.Context, *api.GetEntitiesByCallsignRequest) (*api.EntitiesResponse, error) {
	return nil, status.New(codes.Unimplemented, "method not implemented").Err()
}
