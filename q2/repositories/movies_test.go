package repositories

import (
	"errors"
	"github.com/baharibastian/stockbit-test/q2/entity"
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func Test_moviesRepo_Search(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	type args struct {
		keyword string
		page    string
	}
	tests := []struct {
		id      int
		name    string
		args    args
		want    entity.SearchMoviesResult
		wantErr bool
	}{
		{
			id:   1,
			name: "happy flow",
			args: args{
				keyword: "batman",
				page:    "1",
			},
			want: entity.SearchMoviesResult{
				Movies: []entity.SearchMovieResultData{
					{
						Title:  "batman",
						Year:   "2016",
						IMDBId: "tt1231",
						Poster: "https",
					},
				},
				TotalResults: "22",
				Response:     "True",
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "invalid base url",
			args: args{
				keyword: "batman",
				page:    "1",
			},
			want:    entity.SearchMoviesResult{},
			wantErr: true,
		},
		{
			id:   3,
			name: "error when hit the API",
			args: args{
				keyword: "batman",
				page:    "1",
			},
			want:    entity.SearchMoviesResult{},
			wantErr: true,
		},
		{
			id:   4,
			name: "error decode response json",
			args: args{
				keyword: "batman",
				page:    "1",
			},
			want: entity.SearchMoviesResult{
				Movies: []entity.SearchMovieResultData{
					{
						Title:  "batman",
						Year:   "2016",
						IMDBId: "tt1231",
						Poster: "https",
					},
				},
				TotalResults: "22",
			},
			wantErr: true,
		},
		{
			id:   5,
			name: "movie not found",
			args: args{
				keyword: "batman",
				page:    "1",
			},
			want: entity.SearchMoviesResult{
				Response: "False",
				Error:    "Movie not found!",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := moviesRepo{
				baseURL: "https://someurl.com/",
				apiKey:  "adas",
			}

			switch tt.id {
			case 1:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&page=1&s=batman", func(request *http.Request) (*http.Response, error) {
					resp, _ := httpmock.NewJsonResponse(200, map[string]interface{}{
						"Search": []map[string]interface{}{
							{
								"Title":  "batman",
								"Year":   "2016",
								"imdbID": "tt1231",
								"Poster": "https",
							},
						},
						"totalResults": "22",
						"Response":     "True",
					})
					return resp, nil
				})
			case 2:
				mr.baseURL = "'http:/'"
				httpmock.RegisterResponder("GET", "'http:/'?apikey=adas&page=1&s=batman", func(request *http.Request) (*http.Response, error) {
					return nil, nil
				})
			case 3:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&page=1&s=batman", func(request *http.Request) (*http.Response, error) {
					return nil, errors.New("error http")
				})
			case 4:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&page=1&s=batman", func(request *http.Request) (*http.Response, error) {
					resp, _ := httpmock.NewJsonResponse(200, map[string]interface{}{
						"Search": []map[string]interface{}{
							{
								"Title":  "batman",
								"Year":   "2016",
								"imdbID": "tt1231",
								"Poster": "https",
							},
						},
						"totalResults": "22",
						"Response":     true,
					})
					return resp, nil
				})
			case 5:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&page=1&s=batman", func(request *http.Request) (*http.Response, error) {
					resp, _ := httpmock.NewJsonResponse(200, map[string]interface{}{
						"Response": "False",
						"Error":    "Movie not found!",
					})
					return resp, nil
				})
			}
			got, err := mr.Search(tt.args.keyword, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moviesRepo_GetByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	type args struct {
		id string
	}
	tests := []struct {
		id      int
		name    string
		args    args
		want    entity.SingleMovieResult
		wantErr bool
	}{
		{
			id:   1,
			name: "happy flow",
			args: args{
				id: "tt121",
			},
			want: entity.SingleMovieResult{
				Title:    "Batman",
				Year:     "2018",
				Response: "True",
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "invalid base url",
			args: args{
				id: "t121",
			},
			want:    entity.SingleMovieResult{},
			wantErr: true,
		},
		{
			id:   3,
			name: "error when hit the API",
			args: args{
				id: "t12121",
			},
			want:    entity.SingleMovieResult{},
			wantErr: true,
		},
		{
			id:   4,
			name: "error decode response json",
			args: args{
				id: "t1212",
			},
			want: entity.SingleMovieResult{
				Title: "Batman",
				Year:  "2017",
			},
			wantErr: true,
		},
		{
			id:   5,
			name: "movie not found",
			args: args{
				id: "t99182",
			},
			want: entity.SingleMovieResult{
				Response: "False",
				Error:    "incorrect ID",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := moviesRepo{
				baseURL: "https://someurl.com/",
				apiKey:  "adas",
			}
			switch tt.id {
			case 1:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&i=tt121", func(request *http.Request) (*http.Response, error) {
					resp, _ := httpmock.NewJsonResponse(200, map[string]interface{}{
						"Title":    "Batman",
						"Year":     "2018",
						"Response": "True",
					})
					return resp, nil
				})
			case 2:
				mr.baseURL = "'http:/'"
				httpmock.RegisterResponder("GET", "'http:/'?apikey=adas&i=t121", func(request *http.Request) (*http.Response, error) {
					return nil, nil
				})
			case 3:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&i=t12121", func(request *http.Request) (*http.Response, error) {
					return nil, errors.New("error http")
				})
			case 4:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&i=t1212", func(request *http.Request) (*http.Response, error) {
					resp, _ := httpmock.NewJsonResponse(200, map[string]interface{}{
						"Title":    "Batman",
						"Year":     "2017",
						"Response": true,
					})
					return resp, nil
				})
			case 5:
				httpmock.RegisterResponder("GET", "https://someurl.com/?apikey=adas&i=t99182", func(request *http.Request) (*http.Response, error) {
					resp, _ := httpmock.NewJsonResponse(200, map[string]interface{}{
						"Response": "False",
						"Error":    "incorrect ID",
					})
					return resp, nil
				})
			}
			got, err := mr.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMovies(t *testing.T) {
	type args struct {
		baseURL string
		apiKey  string
	}
	tests := []struct {
		name string
		args args
		want *moviesRepo
	}{
		{
			name: "happy flow",
			args: args{
				baseURL: "https://someurl.com/",
				apiKey:  "adas",
			},
			want: &moviesRepo{
				baseURL: "https://someurl.com/",
				apiKey:  "adas",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMovies(tt.args.baseURL, tt.args.apiKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMovies() = %v, want %v", got, tt.want)
			}
		})
	}
}
