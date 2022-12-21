package clients

import (
	"UacademyGo/Blogpost/api_gateway/config"
	"UacademyGo/Blogpost/api_gateway/protogen/blogpost"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author blogpost.AuthorServiceClient
	Article blogpost.ArticleServiceClient
	Auth blogpost.AuthServiceClient

	conns []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connectAuthor, err := grpc.Dial(cfg.AuthorServiceGrpcHost+cfg.AuthorServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	author := blogpost.NewAuthorServiceClient(connectAuthor)

	connectArticle, err := grpc.Dial(cfg.ArticleServiceGrpcHost+cfg.ArticleServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	article := blogpost.NewArticleServiceClient(connectArticle)

	connectAuth, err := grpc.Dial(cfg.AuthServiceGrpcHost+cfg.AuthServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	auth := blogpost.NewAuthServiceClient(connectAuth)

	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Author: author,
		Article: article,
		Auth: auth,
		conns: append(conns, connectAuthor, connectArticle, connectAuth),
	}, nil
}

func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}