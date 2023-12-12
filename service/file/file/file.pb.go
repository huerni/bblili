// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: file.proto

package file

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type UpLoadFileBySlicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileMD5     string `protobuf:"bytes,1,opt,name=fileMD5,proto3" json:"fileMD5,omitempty"`
	FileType    string `protobuf:"bytes,2,opt,name=fileType,proto3" json:"fileType,omitempty"`
	SliceNumber int32  `protobuf:"varint,3,opt,name=sliceNumber,proto3" json:"sliceNumber,omitempty"`
	SliceSize   int32  `protobuf:"varint,4,opt,name=sliceSize,proto3" json:"sliceSize,omitempty"`
}

func (x *UpLoadFileBySlicesRequest) Reset() {
	*x = UpLoadFileBySlicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpLoadFileBySlicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpLoadFileBySlicesRequest) ProtoMessage() {}

func (x *UpLoadFileBySlicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpLoadFileBySlicesRequest.ProtoReflect.Descriptor instead.
func (*UpLoadFileBySlicesRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *UpLoadFileBySlicesRequest) GetFileMD5() string {
	if x != nil {
		return x.FileMD5
	}
	return ""
}

func (x *UpLoadFileBySlicesRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *UpLoadFileBySlicesRequest) GetSliceNumber() int32 {
	if x != nil {
		return x.SliceNumber
	}
	return 0
}

func (x *UpLoadFileBySlicesRequest) GetSliceSize() int32 {
	if x != nil {
		return x.SliceSize
	}
	return 0
}

type UploadFileBySliceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadFileBySliceResponse) Reset() {
	*x = UploadFileBySliceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileBySliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileBySliceResponse) ProtoMessage() {}

func (x *UploadFileBySliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileBySliceResponse.ProtoReflect.Descriptor instead.
func (*UploadFileBySliceResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileBySliceResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFileRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{5}
}

type GetFileMD5Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *GetFileMD5Request) Reset() {
	*x = GetFileMD5Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileMD5Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileMD5Request) ProtoMessage() {}

func (x *GetFileMD5Request) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileMD5Request.ProtoReflect.Descriptor instead.
func (*GetFileMD5Request) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{6}
}

func (x *GetFileMD5Request) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type GetFileMD5Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetFileMD5Response) Reset() {
	*x = GetFileMD5Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileMD5Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileMD5Response) ProtoMessage() {}

func (x *GetFileMD5Response) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileMD5Response.ProtoReflect.Descriptor instead.
func (*GetFileMD5Response) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{7}
}

func (x *GetFileMD5Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GetSuccessChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileMD5         string `protobuf:"bytes,1,opt,name=fileMD5,proto3" json:"fileMD5,omitempty"`
	Filetype        string `protobuf:"bytes,2,opt,name=filetype,proto3" json:"filetype,omitempty"`
	TotalChunkCount int32  `protobuf:"varint,3,opt,name=totalChunkCount,proto3" json:"totalChunkCount,omitempty"`
}

func (x *GetSuccessChunkRequest) Reset() {
	*x = GetSuccessChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuccessChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuccessChunkRequest) ProtoMessage() {}

func (x *GetSuccessChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuccessChunkRequest.ProtoReflect.Descriptor instead.
func (*GetSuccessChunkRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{8}
}

func (x *GetSuccessChunkRequest) GetFileMD5() string {
	if x != nil {
		return x.FileMD5
	}
	return ""
}

func (x *GetSuccessChunkRequest) GetFiletype() string {
	if x != nil {
		return x.Filetype
	}
	return ""
}

func (x *GetSuccessChunkRequest) GetTotalChunkCount() int32 {
	if x != nil {
		return x.TotalChunkCount
	}
	return 0
}

type GetSuccessChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadID string `protobuf:"bytes,1,opt,name=uploadID,proto3" json:"uploadID,omitempty"`
	UuId     string `protobuf:"bytes,2,opt,name=uuId,proto3" json:"uuId,omitempty"`
	Chunks   string `protobuf:"bytes,3,opt,name=chunks,proto3" json:"chunks,omitempty"`
	Uploaded bool   `protobuf:"varint,4,opt,name=uploaded,proto3" json:"uploaded,omitempty"`
}

func (x *GetSuccessChunkResponse) Reset() {
	*x = GetSuccessChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuccessChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuccessChunkResponse) ProtoMessage() {}

func (x *GetSuccessChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuccessChunkResponse.ProtoReflect.Descriptor instead.
func (*GetSuccessChunkResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{9}
}

func (x *GetSuccessChunkResponse) GetUploadID() string {
	if x != nil {
		return x.UploadID
	}
	return ""
}

func (x *GetSuccessChunkResponse) GetUuId() string {
	if x != nil {
		return x.UuId
	}
	return ""
}

func (x *GetSuccessChunkResponse) GetChunks() string {
	if x != nil {
		return x.Chunks
	}
	return ""
}

func (x *GetSuccessChunkResponse) GetUploaded() bool {
	if x != nil {
		return x.Uploaded
	}
	return false
}

type GetFileUrlByMD5Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md5 string `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
}

func (x *GetFileUrlByMD5Request) Reset() {
	*x = GetFileUrlByMD5Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileUrlByMD5Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileUrlByMD5Request) ProtoMessage() {}

func (x *GetFileUrlByMD5Request) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileUrlByMD5Request.ProtoReflect.Descriptor instead.
func (*GetFileUrlByMD5Request) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{10}
}

func (x *GetFileUrlByMD5Request) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

type GetFileUrlByMD5Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetFileUrlByMD5Response) Reset() {
	*x = GetFileUrlByMD5Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileUrlByMD5Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileUrlByMD5Response) ProtoMessage() {}

func (x *GetFileUrlByMD5Response) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileUrlByMD5Response.ProtoReflect.Descriptor instead.
func (*GetFileUrlByMD5Response) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{11}
}

func (x *GetFileUrlByMD5Response) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x22, 0x91, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x42, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6c, 0x69, 0x63, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x2d, 0x0a, 0x19, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x42, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d,
	0x44, 0x35, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x78, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x22, 0x2a, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x79, 0x4d, 0x44, 0x35, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x22, 0x2b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x79, 0x4d, 0x44, 0x35, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x32, 0xa8, 0x03, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x55, 0x70, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x12, 0x17, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x12, 0x17, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x1c, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x79, 0x4d,
	0x44, 0x35, 0x12, 0x1c, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x72, 0x6c, 0x42, 0x79, 0x4d, 0x44, 0x35, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x42, 0x79, 0x4d, 0x44, 0x35, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_file_proto_goTypes = []interface{}{
	(*Request)(nil),                   // 0: file.Request
	(*Response)(nil),                  // 1: file.Response
	(*UpLoadFileBySlicesRequest)(nil), // 2: file.UpLoadFileBySlicesRequest
	(*UploadFileBySliceResponse)(nil), // 3: file.UploadFileBySliceResponse
	(*DeleteFileRequest)(nil),         // 4: file.DeleteFileRequest
	(*DeleteFileResponse)(nil),        // 5: file.DeleteFileResponse
	(*GetFileMD5Request)(nil),         // 6: file.GetFileMD5Request
	(*GetFileMD5Response)(nil),        // 7: file.GetFileMD5Response
	(*GetSuccessChunkRequest)(nil),    // 8: file.GetSuccessChunkRequest
	(*GetSuccessChunkResponse)(nil),   // 9: file.GetSuccessChunkResponse
	(*GetFileUrlByMD5Request)(nil),    // 10: file.GetFileUrlByMD5Request
	(*GetFileUrlByMD5Response)(nil),   // 11: file.GetFileUrlByMD5Response
}
var file_file_proto_depIdxs = []int32{
	0,  // 0: file.File.Ping:input_type -> file.Request
	2,  // 1: file.File.UploadFileSlices:input_type -> file.UpLoadFileBySlicesRequest
	6,  // 2: file.File.DeleteFileMD5:input_type -> file.GetFileMD5Request
	6,  // 3: file.File.GetFileMD5:input_type -> file.GetFileMD5Request
	8,  // 4: file.File.GetSuccessChunk:input_type -> file.GetSuccessChunkRequest
	10, // 5: file.File.GetFileUrlByMD5:input_type -> file.GetFileUrlByMD5Request
	1,  // 6: file.File.Ping:output_type -> file.Response
	3,  // 7: file.File.UploadFileSlices:output_type -> file.UploadFileBySliceResponse
	7,  // 8: file.File.DeleteFileMD5:output_type -> file.GetFileMD5Response
	7,  // 9: file.File.GetFileMD5:output_type -> file.GetFileMD5Response
	9,  // 10: file.File.GetSuccessChunk:output_type -> file.GetSuccessChunkResponse
	11, // 11: file.File.GetFileUrlByMD5:output_type -> file.GetFileUrlByMD5Response
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpLoadFileBySlicesRequest); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileBySliceResponse); i {
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
		file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
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
		file_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileResponse); i {
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
		file_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileMD5Request); i {
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
		file_file_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileMD5Response); i {
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
		file_file_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuccessChunkRequest); i {
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
		file_file_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuccessChunkResponse); i {
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
		file_file_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileUrlByMD5Request); i {
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
		file_file_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileUrlByMD5Response); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
