package main

import (
	"anime-grpc/api"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	port = ":8080"
)

func main() {
	ctx := context.Background()
	connectionStartTime := time.Now()
	connection, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", port, err)
	}
	log.Printf("Connected in %d microsec", time.Since(connectionStartTime))

	animeRepositoryClient := api.NewAnimeServiceClient(connection)
	animeList, err := animeRepositoryClient.GetAll(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("Could not get anime: %v", err)
	}
	log.Printf("Got list of anime: %v", animeList.Anime)

	validId, invalidId := 1, 5
	anime, err := animeRepositoryClient.Get(ctx, &api.AnimeRequestId{Id: int64(validId)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got anime with id %d: %v", validId, anime)

	_, err = animeRepositoryClient.Get(ctx, &api.AnimeRequestId{Id: int64(invalidId)})
	if err != nil {
		log.Printf("Got error: %v", err)
	}

	newAnime := &api.Anime{
		Id:            2,
		Title:         "Classroom of Elite",
		TitleJapanese: "ようこそ実力至上主義の教室へ",
		Source:        "Light Novel",
		Episodes:      12,
		Kind:          "TV",
		Score:         8.4,
		Status:        "Released",
		Synopsis: "Kodo Ikusei High School is a prestigious educational institution with excellent " +
			"performance, where 100% of students go to universities or find work. Students are allowed to " +
			"wear any hairstyle and bring any personal belongings. " +
			"Kodo Ikusei is a paradise-like school, but the truth is that this approach " +
			"only applies to the brightest students.",
	}

	postNewAnime, err := animeRepositoryClient.Post(ctx, newAnime)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Anime was successfully created: %v", postNewAnime)

	updatedAnime := &api.Anime{
		Id:            1,
		Title:         "K-on",
		TitleJapanese: "けいおん!",
		Source:        "Manga",
		Episodes:      13,
		Kind:          "TV",
		Score:         9.4,
		Status:        "Released",
		Synopsis: "Yui Hirasawa is a young and carefree girl who has just entered high school. On " +
			"the very first day of school, she notices an advertising poster for the light music " +
			"club and, burning with the desire to join there, goes to enroll in the participants.",
	}
	updatedNewAnime, err := animeRepositoryClient.Put(ctx, updatedAnime)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Anime was successfully updated: %v", updatedNewAnime)

	animeList, err = animeRepositoryClient.GetAll(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("Could not get anime: %v", err)
	}
	log.Printf("Got list of anime: %v", animeList.Anime)
	anime, err = animeRepositoryClient.Get(ctx, &api.AnimeRequestId{Id: int64(2)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got anime with id %d: %v", 2, anime)
	_, err = animeRepositoryClient.Delete(ctx, &api.AnimeRequestId{Id: 2})
	if err != nil {
		log.Fatal(err)
	}
	animeList, err = animeRepositoryClient.GetAll(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("Could not get anime: %v", err)
	}
	log.Printf("Got list of anime: %v", animeList.Anime)
}
