package main

import (
	"anime-grpc/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	port = ":8080"
)

type AnimeRepository struct {
	anime map[int64]*api.Anime
	api.UnimplementedAnimeServiceServer
}

func NewAnimeRepository() *AnimeRepository {
	return &AnimeRepository{
		anime: map[int64]*api.Anime{
			1: {
				Id:            1,
				Title:         "K-on!",
				TitleJapanese: "けいおん!",
				Source:        "Manga",
				Episodes:      13,
				Kind:          "TV",
				Score:         8.4,
				Status:        "Released",
				Synopsis: "Yui Hirasawa is a young and carefree girl who has just entered high school. On " +
					"the very first day of school, she notices an advertising poster for the light music " +
					"club and, burning with the desire to join there, goes to enroll in the participants.",
			},
		},
	}
}

func (ar AnimeRepository) GetAll(context.Context, *api.Empty) (*api.AnimeList, error) {
	var animeList []*api.Anime
	for _, anime := range ar.anime {
		animeList = append(animeList, anime)
	}
	result := api.AnimeList{Anime: animeList}
	return &result, nil
}

func (ar AnimeRepository) Get(_ context.Context, req *api.AnimeRequestId) (*api.Anime, error) {
	if anime, ok := ar.anime[req.Id]; ok {
		return anime, nil
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Anime with id %v doesn't exist", req.Id))
}

func (ar AnimeRepository) Post(_ context.Context, req *api.Anime) (*api.Anime, error) {
	anime := &api.Anime{
		Id:            req.Id,
		Title:         req.Title,
		TitleJapanese: req.TitleJapanese,
		Source:        req.Source,
		Episodes:      req.Episodes,
		Kind:          req.Kind,
		Score:         req.Score,
		Status:        req.Status,
		Synopsis:      req.Synopsis,
	}
	ar.anime[req.Id] = req
	return anime, nil
}

func (ar AnimeRepository) Put(_ context.Context, req *api.Anime) (*api.Anime, error) {
	anime := &api.Anime{
		Id:            req.Id,
		Title:         req.Title,
		TitleJapanese: req.TitleJapanese,
		Source:        req.Source,
		Episodes:      req.Episodes,
		Kind:          req.Kind,
		Score:         req.Score,
		Status:        req.Status,
		Synopsis:      req.Synopsis,
	}
	ar.anime[req.Id] = req
	return anime, nil
}

func (ar AnimeRepository) Delete(_ context.Context, req *api.AnimeRequestId) (*api.Empty, error) {
	if _, ok := ar.anime[req.Id]; ok {
		delete(ar.anime, req.Id)
		return &api.Empty{}, nil
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Anime with id %v doesn't exist", req.Id))
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot listen to %s: %v", port, err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	animeRepository := NewAnimeRepository()

	api.RegisterAnimeServiceServer(grpcServer, animeRepository)
	log.Printf("Serving on %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve on %v: %v", listener.Addr(), err)
	}
}
