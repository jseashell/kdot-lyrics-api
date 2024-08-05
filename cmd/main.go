package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

var DB_NAME = "kdot-songs-main"

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	data := randomSong()
	body, err := json.Marshal(data)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

func randomSong() Data {
	dbClient := newDbClient()

	av, _ := attributevalue.MarshalMap(map[string]interface{}{
		"ID": uuid.NewString(),
	})

	limit := int32(1)
	limitPtr := &limit

	res, err := dbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		ExclusiveStartKey: av,
		TableName:         aws.String(DB_NAME),
		Limit:             limitPtr,
	})

	if err != nil {
		panic(err)
	}

	randomSong := &RandomSong{}
	attributevalue.UnmarshalMap(res.Items[0], randomSong)

	data := Data{
		ID:     randomSong.ID,
		Lyrics: randomSong.Lyrics,
		Album: Album{
			APIPath:               randomSong.Song.Album.APIPath,
			CoverArtURL:           randomSong.Song.Album.CoverArtURL,
			FullTitle:             randomSong.Song.Album.FullTitle,
			ID:                    randomSong.Song.Album.ID,
			Name:                  randomSong.Song.Album.Name,
			ReleaseDateForDisplay: randomSong.Song.Album.ReleaseDateForDisplay,
			URL:                   randomSong.Song.Album.URL,
		},
		Song: Song{
			ArtistNames:              randomSong.Song.ArtistNames,
			FullTitle:                randomSong.Song.FullTitle,
			HeaderImageThumbnailURL:  randomSong.Song.HeaderImageThumbnailURL,
			HeaderImageURL:           randomSong.Song.HeaderImageURL,
			SongID:                   randomSong.Song.SongID,
			ID:                       randomSong.Song.ID,
			Path:                     randomSong.Song.Path,
			ReleaseDateForDisplay:    randomSong.Song.ReleaseDateForDisplay,
			SongArtImageThumbnailURL: randomSong.Song.SongArtImageThumbnailURL,
			SongArtImageURL:          randomSong.Song.SongArtImageURL,
			Title:                    randomSong.Song.Title,
			URL:                      randomSong.Song.URL,
		},
	}

	return data
}

func newDbClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		slog.Error("Unable to load AWS SDK config.")
		panic(err)
	}
	return dynamodb.NewFromConfig(cfg)
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type Data struct {
	ID     string   `json:"id"`
	Song   Song     `json:"song"`
	Album  Album    `json:"album"`
	Lyrics []string `json:"lyrics"`
}

type Song struct {
	ArtistNames              string `json:"artist_names"`
	FullTitle                string `json:"full_title"`
	HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url"`
	HeaderImageURL           string `json:"header_image_url"`
	SongID                   int    `json:"id"`
	ID                       string `json:"uuid"`
	Path                     string `json:"path"`
	ReleaseDateForDisplay    string `json:"release_date_for_display"`
	SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url"`
	SongArtImageURL          string `json:"song_art_image_url"`
	Title                    string `json:"title"`
	URL                      string `json:"url"`
}

type Album struct {
	APIPath               string `json:"api_path"`
	CoverArtURL           string `json:"cover_art_url"`
	FullTitle             string `json:"full_title"`
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	ReleaseDateForDisplay string `json:"release_date_for_display"`
	URL                   string `json:"url"`
}

type RandomSong struct {
	ID   string `dynamodbav:"ID"`
	Song struct {
		ArtistNames              string `dynamodbav:"ArtistNames"`
		FullTitle                string `dynamodbav:"FullTitle"`
		HeaderImageThumbnailURL  string `dynamodbav:"HeaderImageThumbnailURL"`
		HeaderImageURL           string `dynamodbav:"HeaderImageURL"`
		SongID                   int    `dynamodbav:"SongID"`
		ID                       string `dynamodbav:"ID"`
		Path                     string `dynamodbav:"Path"`
		ReleaseDateForDisplay    string `dynamodbav:"ReleaseDateForDisplay"`
		SongArtImageThumbnailURL string `dynamodbav:"SongArtImageThumbnailURL"`
		SongArtImageURL          string `dynamodbav:"SongArtImageURL"`
		Title                    string `dynamodbav:"Title"`
		URL                      string `dynamodbav:"URL"`
		Album                    struct {
			APIPath               string `dynamodbav:"APIPath"`
			CoverArtURL           string `dynamodbav:"CoverArtURL"`
			FullTitle             string `dynamodbav:"FullTItle"`
			ID                    int    `dynamodbav:"ID"`
			Name                  string `dynamodbav:"Name"`
			ReleaseDateForDisplay string `dynamodbav:"ReleaseDateForDisplay"`
			URL                   string `dynamodbav:"URL"`
		} `dynamodbav:"Album"`
	} `dynamodbav:"Song"`
	Lyrics []string `dynamodbav:"Lyrics"`
}
