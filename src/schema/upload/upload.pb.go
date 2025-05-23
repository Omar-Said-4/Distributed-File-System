// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: upload.proto

package upload

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MasterUploadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MasterUploadRequest) Reset() {
	*x = MasterUploadRequest{}
	mi := &file_upload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MasterUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterUploadRequest) ProtoMessage() {}

func (x *MasterUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterUploadRequest.ProtoReflect.Descriptor instead.
func (*MasterUploadRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{0}
}

type MasterUploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodeIp        string                 `protobuf:"bytes,1,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	NodePort      string                 `protobuf:"bytes,2,opt,name=node_port,json=nodePort,proto3" json:"node_port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MasterUploadResponse) Reset() {
	*x = MasterUploadResponse{}
	mi := &file_upload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MasterUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterUploadResponse) ProtoMessage() {}

func (x *MasterUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterUploadResponse.ProtoReflect.Descriptor instead.
func (*MasterUploadResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *MasterUploadResponse) GetNodeIp() string {
	if x != nil {
		return x.NodeIp
	}
	return ""
}

func (x *MasterUploadResponse) GetNodePort() string {
	if x != nil {
		return x.NodePort
	}
	return ""
}

type FileInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath      string                 `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileSize      uint64                 `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	mi := &file_upload_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *FileInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfo) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileInfo) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type UploadFileRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*UploadFileRequest_FileInfo
	//	*UploadFileRequest_Chunks
	Data          isUploadFileRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	mi := &file_upload_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadFileRequest) GetFileInfo() *FileInfo {
	if x != nil {
		if x, ok := x.Data.(*UploadFileRequest_FileInfo); ok {
			return x.FileInfo
		}
	}
	return nil
}

func (x *UploadFileRequest) GetChunks() []byte {
	if x != nil {
		if x, ok := x.Data.(*UploadFileRequest_Chunks); ok {
			return x.Chunks
		}
	}
	return nil
}

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_FileInfo struct {
	FileInfo *FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3,oneof"`
}

type UploadFileRequest_Chunks struct {
	Chunks []byte `protobuf:"bytes,2,opt,name=chunks,proto3,oneof"`
}

func (*UploadFileRequest_FileInfo) isUploadFileRequest_Data() {}

func (*UploadFileRequest_Chunks) isUploadFileRequest_Data() {}

type UploadFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	mi := &file_upload_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{4}
}

type NotifyMasterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodeId        uint32                 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	FileInfo      *FileInfo              `protobuf:"bytes,2,opt,name=file_info,json=fileInfo,proto3" json:"file_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotifyMasterRequest) Reset() {
	*x = NotifyMasterRequest{}
	mi := &file_upload_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyMasterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMasterRequest) ProtoMessage() {}

func (x *NotifyMasterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMasterRequest.ProtoReflect.Descriptor instead.
func (*NotifyMasterRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{5}
}

func (x *NotifyMasterRequest) GetNodeId() uint32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NotifyMasterRequest) GetFileInfo() *FileInfo {
	if x != nil {
		return x.FileInfo
	}
	return nil
}

type NotifyMasterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotifyMasterResponse) Reset() {
	*x = NotifyMasterResponse{}
	mi := &file_upload_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyMasterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMasterResponse) ProtoMessage() {}

func (x *NotifyMasterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMasterResponse.ProtoReflect.Descriptor instead.
func (*NotifyMasterResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{6}
}

type ConfirmUploadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileInfo      *FileInfo              `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3" json:"file_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmUploadRequest) Reset() {
	*x = ConfirmUploadRequest{}
	mi := &file_upload_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmUploadRequest) ProtoMessage() {}

func (x *ConfirmUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmUploadRequest.ProtoReflect.Descriptor instead.
func (*ConfirmUploadRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{7}
}

func (x *ConfirmUploadRequest) GetFileInfo() *FileInfo {
	if x != nil {
		return x.FileInfo
	}
	return nil
}

type ConfirmUploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmUploadResponse) Reset() {
	*x = ConfirmUploadResponse{}
	mi := &file_upload_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmUploadResponse) ProtoMessage() {}

func (x *ConfirmUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmUploadResponse.ProtoReflect.Descriptor instead.
func (*ConfirmUploadResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{8}
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a,
	0x14, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x70, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x61, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x66,
	0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x0a, 0x13,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xc1, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x13, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x49,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x2e, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData []byte
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upload_proto_rawDesc), len(file_upload_proto_rawDesc)))
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_upload_proto_goTypes = []any{
	(*MasterUploadRequest)(nil),   // 0: upload.MasterUploadRequest
	(*MasterUploadResponse)(nil),  // 1: upload.MasterUploadResponse
	(*FileInfo)(nil),              // 2: upload.FileInfo
	(*UploadFileRequest)(nil),     // 3: upload.UploadFileRequest
	(*UploadFileResponse)(nil),    // 4: upload.UploadFileResponse
	(*NotifyMasterRequest)(nil),   // 5: upload.NotifyMasterRequest
	(*NotifyMasterResponse)(nil),  // 6: upload.NotifyMasterResponse
	(*ConfirmUploadRequest)(nil),  // 7: upload.ConfirmUploadRequest
	(*ConfirmUploadResponse)(nil), // 8: upload.ConfirmUploadResponse
}
var file_upload_proto_depIdxs = []int32{
	2, // 0: upload.UploadFileRequest.file_info:type_name -> upload.FileInfo
	2, // 1: upload.NotifyMasterRequest.file_info:type_name -> upload.FileInfo
	2, // 2: upload.ConfirmUploadRequest.file_info:type_name -> upload.FileInfo
	0, // 3: upload.UploadService.MasterRequestUpload:input_type -> upload.MasterUploadRequest
	3, // 4: upload.UploadService.UploadFile:input_type -> upload.UploadFileRequest
	5, // 5: upload.UploadService.NotifyMaster:input_type -> upload.NotifyMasterRequest
	7, // 6: upload.UploadService.ConfirmUpload:input_type -> upload.ConfirmUploadRequest
	1, // 7: upload.UploadService.MasterRequestUpload:output_type -> upload.MasterUploadResponse
	4, // 8: upload.UploadService.UploadFile:output_type -> upload.UploadFileResponse
	6, // 9: upload.UploadService.NotifyMaster:output_type -> upload.NotifyMasterResponse
	8, // 10: upload.UploadService.ConfirmUpload:output_type -> upload.ConfirmUploadResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	file_upload_proto_msgTypes[3].OneofWrappers = []any{
		(*UploadFileRequest_FileInfo)(nil),
		(*UploadFileRequest_Chunks)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upload_proto_rawDesc), len(file_upload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}
