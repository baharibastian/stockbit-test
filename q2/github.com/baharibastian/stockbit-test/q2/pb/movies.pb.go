// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: proto/movies.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MovieData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *MovieData) Reset() {
	*x = MovieData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieData) ProtoMessage() {}

func (x *MovieData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieData.ProtoReflect.Descriptor instead.
func (*MovieData) Descriptor() ([]byte, []int) {
	return file_proto_movies_proto_rawDescGZIP(), []int{0}
}

func (x *MovieData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieData) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieData) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieData) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type SearchMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Searchword string `protobuf:"bytes,1,opt,name=searchword,proto3" json:"searchword,omitempty"`
	Pagination int32  `protobuf:"varint,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *SearchMoviesRequest) Reset() {
	*x = SearchMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMoviesRequest) ProtoMessage() {}

func (x *SearchMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMoviesRequest.ProtoReflect.Descriptor instead.
func (*SearchMoviesRequest) Descriptor() ([]byte, []int) {
	return file_proto_movies_proto_rawDescGZIP(), []int{1}
}

func (x *SearchMoviesRequest) GetSearchword() string {
	if x != nil {
		return x.Searchword
	}
	return ""
}

func (x *SearchMoviesRequest) GetPagination() int32 {
	if x != nil {
		return x.Pagination
	}
	return 0
}

type SearchMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search       []*MovieData `protobuf:"bytes,1,rep,name=Search,proto3" json:"Search,omitempty"`
	TotalResults string       `protobuf:"bytes,2,opt,name=totalResults,proto3" json:"totalResults,omitempty"`
	Response     string       `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *SearchMoviesResponse) Reset() {
	*x = SearchMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMoviesResponse) ProtoMessage() {}

func (x *SearchMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMoviesResponse.ProtoReflect.Descriptor instead.
func (*SearchMoviesResponse) Descriptor() ([]byte, []int) {
	return file_proto_movies_proto_rawDescGZIP(), []int{2}
}

func (x *SearchMoviesResponse) GetSearch() []*MovieData {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *SearchMoviesResponse) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *SearchMoviesResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIDRequest) Reset() {
	*x = GetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRequest) ProtoMessage() {}

func (x *GetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRequest.ProtoReflect.Descriptor instead.
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_movies_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string    `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year       string    `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string    `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released   string    `protobuf:"bytes,4,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime    string    `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string    `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string    `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer     string    `protobuf:"bytes,8,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors     string    `protobuf:"bytes,9,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot       string    `protobuf:"bytes,10,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language   string    `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string    `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string    `protobuf:"bytes,13,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string    `protobuf:"bytes,14,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Ratings    []*Rating `protobuf:"bytes,15,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
	Metascore  string    `protobuf:"bytes,16,opt,name=Metascore,proto3" json:"Metascore,omitempty"`
	ImdbRating string    `protobuf:"bytes,17,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string    `protobuf:"bytes,18,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	ImdbID     string    `protobuf:"bytes,19,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type       string    `protobuf:"bytes,20,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD        string    `protobuf:"bytes,21,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string    `protobuf:"bytes,22,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production string    `protobuf:"bytes,23,opt,name=Production,proto3" json:"Production,omitempty"`
	Website    string    `protobuf:"bytes,24,opt,name=Website,proto3" json:"Website,omitempty"`
	Response   string    `protobuf:"bytes,25,opt,name=Response,proto3" json:"Response,omitempty"`
	Error      string    `protobuf:"bytes,26,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *GetByIDResponse) Reset() {
	*x = GetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDResponse) ProtoMessage() {}

func (x *GetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDResponse.ProtoReflect.Descriptor instead.
func (*GetByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_movies_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIDResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetByIDResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetByIDResponse) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *GetByIDResponse) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *GetByIDResponse) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *GetByIDResponse) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *GetByIDResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *GetByIDResponse) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *GetByIDResponse) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *GetByIDResponse) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *GetByIDResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetByIDResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetByIDResponse) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *GetByIDResponse) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *GetByIDResponse) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetByIDResponse) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *GetByIDResponse) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *GetByIDResponse) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *GetByIDResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *GetByIDResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetByIDResponse) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *GetByIDResponse) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *GetByIDResponse) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *GetByIDResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetByIDResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *GetByIDResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_proto_movies_proto_rawDescGZIP(), []int{5}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_movies_proto protoreflect.FileDescriptor

var file_proto_movies_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x55, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7a, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xaa, 0x05, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x50, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x07, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x44, 0x56, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x74, 0x0a, 0x0d, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61,
	0x68, 0x61, 0x72, 0x69, 0x62, 0x61, 0x73, 0x74, 0x69, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x62, 0x69, 0x74, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x71, 0x32, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_movies_proto_rawDescOnce sync.Once
	file_proto_movies_proto_rawDescData = file_proto_movies_proto_rawDesc
)

func file_proto_movies_proto_rawDescGZIP() []byte {
	file_proto_movies_proto_rawDescOnce.Do(func() {
		file_proto_movies_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_movies_proto_rawDescData)
	})
	return file_proto_movies_proto_rawDescData
}

var file_proto_movies_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_movies_proto_goTypes = []interface{}{
	(*MovieData)(nil),            // 0: MovieData
	(*SearchMoviesRequest)(nil),  // 1: SearchMoviesRequest
	(*SearchMoviesResponse)(nil), // 2: SearchMoviesResponse
	(*GetByIDRequest)(nil),       // 3: GetByIDRequest
	(*GetByIDResponse)(nil),      // 4: GetByIDResponse
	(*Rating)(nil),               // 5: Rating
}
var file_proto_movies_proto_depIdxs = []int32{
	0, // 0: SearchMoviesResponse.Search:type_name -> MovieData
	5, // 1: GetByIDResponse.Ratings:type_name -> Rating
	1, // 2: MoviesService.Search:input_type -> SearchMoviesRequest
	3, // 3: MoviesService.GetByID:input_type -> GetByIDRequest
	2, // 4: MoviesService.Search:output_type -> SearchMoviesResponse
	4, // 5: MoviesService.GetByID:output_type -> GetByIDResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_movies_proto_init() }
func file_proto_movies_proto_init() {
	if File_proto_movies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_movies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMoviesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMoviesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_movies_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_movies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_movies_proto_goTypes,
		DependencyIndexes: file_proto_movies_proto_depIdxs,
		MessageInfos:      file_proto_movies_proto_msgTypes,
	}.Build()
	File_proto_movies_proto = out.File
	file_proto_movies_proto_rawDesc = nil
	file_proto_movies_proto_goTypes = nil
	file_proto_movies_proto_depIdxs = nil
}
