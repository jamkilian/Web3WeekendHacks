package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//go struct for GetEndPoints
type EndPoints struct {
	Endpoints []string `json:"data"`
}

type Tracks []struct {
	Artwork struct {
		One50X150   string `json:"150x150"`
		Four80X480  string `json:"480x480"`
		One000X1000 string `json:"1000x1000"`
	} `json:"artwork"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	ID          string `json:"id"`
	Mood        string `json:"mood"`
	ReleaseDate string `json:"release_date"`
	RemixOf     struct {
		Tracks interface{} `json:"tracks"`
	} `json:"remix_of"`
	RepostCount   int    `json:"repost_count"`
	FavoriteCount int    `json:"favorite_count"`
	Tags          string `json:"tags"`
	Title         string `json:"title"`
	User          struct {
		AlbumCount int    `json:"album_count"`
		Bio        string `json:"bio"`
		CoverPhoto struct {
			Six40X  string `json:"640x"`
			Two000X string `json:"2000x"`
		} `json:"cover_photo"`
		FolloweeCount  int    `json:"followee_count"`
		FollowerCount  int    `json:"follower_count"`
		Handle         string `json:"handle"`
		ID             string `json:"id"`
		IsVerified     bool   `json:"is_verified"`
		Location       string `json:"location"`
		Name           string `json:"name"`
		PlaylistCount  int    `json:"playlist_count"`
		ProfilePicture struct {
			One50X150   string `json:"150x150"`
			Four80X480  string `json:"480x480"`
			One000X1000 string `json:"1000x1000"`
		} `json:"profile_picture"`
		RepostCount int `json:"repost_count"`
		TrackCount  int `json:"track_count"`
	} `json:"user"`
	Duration     int  `json:"duration"`
	Downloadable bool `json:"downloadable"`
	PlayCount    int  `json:"play_count"`
}

func main() {
	var endPoint string = ""
	var trackId string = ""
	endPoint = GetEndPoints()
	fmt.Println(endPoint)
	trackId = GetTracks(endPoint)
	fmt.Println(trackId)

}

func GetEndPoints() string {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://api.audius.co")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to EndPoints struct
	var EndpointsStruct EndPoints
	json.Unmarshal(bodyBytes, &EndpointsStruct)
	fmt.Printf("API Response as struct %+v\n", EndpointsStruct)
	if len(EndpointsStruct.Endpoints) > 0 {
		return EndpointsStruct.Endpoints[0]
	} else {
		return "error"
	}
}

func GetTracks(endPoint string) string {
	fmt.Println("2. Performing Http Get...")
	resp, err := http.Get(endPoint + "/v1/tracks/search?query=baauer b2b&app_name=testapp")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var tracksStruct Tracks
	json.Unmarshal(bodyBytes, &tracksStruct)
	fmt.Printf("API Response as struct %+v\n", tracksStruct)
	if len(tracksStruct) > 0 {
		return tracksStruct[0].ID

	} else {
		return "error"
	}
}

/*
// find the play_count, favorite_count and the repost_count
func GetSingleTrack() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://api.audius.co")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
}
*/
