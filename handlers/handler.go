package handlers

import (
	"UacademyGo/Blogpost/api_gateway/clients"
	"UacademyGo/Blogpost/api_gateway/config"
)

type handler struct {
	cfg 	config.Config
	grpcClients *clients.GrpcClients
}

func NewHandler(cfg config.Config, grpcClients *clients.GrpcClients) handler {
	return handler{
		cfg: cfg,
		grpcClients: grpcClients,
	}
}
