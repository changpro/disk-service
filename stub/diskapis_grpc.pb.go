// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: diskapis.proto

package stub

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*UploadFileRsp, error)
	GetDirsAndFiles(ctx context.Context, in *GetDirsAndFilesReq, opts ...grpc.CallOption) (*GetDirsAndFilesRsp, error)
	GetFileDetail(ctx context.Context, in *GetFileDetailReq, opts ...grpc.CallOption) (*GetFileDetailRsp, error)
	MakeNewFolder(ctx context.Context, in *MakeNewFolderReq, opts ...grpc.CallOption) (*MakeNewFolderRsp, error)
	SetHiddenDoc(ctx context.Context, in *SetHiddenDocReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MoveToRecycleBin(ctx context.Context, in *MoveToRecycleBinReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RecoverDocs(ctx context.Context, in *RecoverDocsReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRecycleBinList(ctx context.Context, in *GetRecycleBinListReq, opts ...grpc.CallOption) (*GetRecycleBinListRsp, error)
	SoftDelete(ctx context.Context, in *SoftDeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HardDelete(ctx context.Context, in *HardDeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Rename(ctx context.Context, in *RenameReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MoveToPath(ctx context.Context, in *MoveToPathReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CopyToPath(ctx context.Context, in *CopyToPathReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateShare(ctx context.Context, in *CreateShareReq, opts ...grpc.CallOption) (*CreateShareRsp, error)
	RetrieveShareToPath(ctx context.Context, in *RetrieveShareToPathReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetShareRecords(ctx context.Context, in *GetShareRecordsReq, opts ...grpc.CallOption) (*GetShareRecordsRsp, error)
	GetShareDetail(ctx context.Context, in *GetShareDetailReq, opts ...grpc.CallOption) (*GetShareDetailRsp, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*UploadFileRsp, error) {
	out := new(UploadFileRsp)
	err := c.cc.Invoke(ctx, "/FileService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetDirsAndFiles(ctx context.Context, in *GetDirsAndFilesReq, opts ...grpc.CallOption) (*GetDirsAndFilesRsp, error) {
	out := new(GetDirsAndFilesRsp)
	err := c.cc.Invoke(ctx, "/FileService/GetDirsAndFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileDetail(ctx context.Context, in *GetFileDetailReq, opts ...grpc.CallOption) (*GetFileDetailRsp, error) {
	out := new(GetFileDetailRsp)
	err := c.cc.Invoke(ctx, "/FileService/GetFileDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) MakeNewFolder(ctx context.Context, in *MakeNewFolderReq, opts ...grpc.CallOption) (*MakeNewFolderRsp, error) {
	out := new(MakeNewFolderRsp)
	err := c.cc.Invoke(ctx, "/FileService/MakeNewFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) SetHiddenDoc(ctx context.Context, in *SetHiddenDocReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/SetHiddenDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) MoveToRecycleBin(ctx context.Context, in *MoveToRecycleBinReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/MoveToRecycleBin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) RecoverDocs(ctx context.Context, in *RecoverDocsReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/RecoverDocs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetRecycleBinList(ctx context.Context, in *GetRecycleBinListReq, opts ...grpc.CallOption) (*GetRecycleBinListRsp, error) {
	out := new(GetRecycleBinListRsp)
	err := c.cc.Invoke(ctx, "/FileService/GetRecycleBinList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) SoftDelete(ctx context.Context, in *SoftDeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/SoftDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) HardDelete(ctx context.Context, in *HardDeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/HardDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) Rename(ctx context.Context, in *RenameReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) MoveToPath(ctx context.Context, in *MoveToPathReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/MoveToPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CopyToPath(ctx context.Context, in *CopyToPathReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/CopyToPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CreateShare(ctx context.Context, in *CreateShareReq, opts ...grpc.CallOption) (*CreateShareRsp, error) {
	out := new(CreateShareRsp)
	err := c.cc.Invoke(ctx, "/FileService/CreateShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) RetrieveShareToPath(ctx context.Context, in *RetrieveShareToPathReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FileService/RetrieveShareToPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetShareRecords(ctx context.Context, in *GetShareRecordsReq, opts ...grpc.CallOption) (*GetShareRecordsRsp, error) {
	out := new(GetShareRecordsRsp)
	err := c.cc.Invoke(ctx, "/FileService/GetShareRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetShareDetail(ctx context.Context, in *GetShareDetailReq, opts ...grpc.CallOption) (*GetShareDetailRsp, error) {
	out := new(GetShareDetailRsp)
	err := c.cc.Invoke(ctx, "/FileService/GetShareDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations should embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	UploadFile(context.Context, *UploadFileReq) (*UploadFileRsp, error)
	GetDirsAndFiles(context.Context, *GetDirsAndFilesReq) (*GetDirsAndFilesRsp, error)
	GetFileDetail(context.Context, *GetFileDetailReq) (*GetFileDetailRsp, error)
	MakeNewFolder(context.Context, *MakeNewFolderReq) (*MakeNewFolderRsp, error)
	SetHiddenDoc(context.Context, *SetHiddenDocReq) (*emptypb.Empty, error)
	MoveToRecycleBin(context.Context, *MoveToRecycleBinReq) (*emptypb.Empty, error)
	RecoverDocs(context.Context, *RecoverDocsReq) (*emptypb.Empty, error)
	GetRecycleBinList(context.Context, *GetRecycleBinListReq) (*GetRecycleBinListRsp, error)
	SoftDelete(context.Context, *SoftDeleteReq) (*emptypb.Empty, error)
	HardDelete(context.Context, *HardDeleteReq) (*emptypb.Empty, error)
	Rename(context.Context, *RenameReq) (*emptypb.Empty, error)
	MoveToPath(context.Context, *MoveToPathReq) (*emptypb.Empty, error)
	CopyToPath(context.Context, *CopyToPathReq) (*emptypb.Empty, error)
	CreateShare(context.Context, *CreateShareReq) (*CreateShareRsp, error)
	RetrieveShareToPath(context.Context, *RetrieveShareToPathReq) (*emptypb.Empty, error)
	GetShareRecords(context.Context, *GetShareRecordsReq) (*GetShareRecordsRsp, error)
	GetShareDetail(context.Context, *GetShareDetailReq) (*GetShareDetailRsp, error)
}

// UnimplementedFileServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) UploadFile(context.Context, *UploadFileReq) (*UploadFileRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileServiceServer) GetDirsAndFiles(context.Context, *GetDirsAndFilesReq) (*GetDirsAndFilesRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDirsAndFiles not implemented")
}
func (UnimplementedFileServiceServer) GetFileDetail(context.Context, *GetFileDetailReq) (*GetFileDetailRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileDetail not implemented")
}
func (UnimplementedFileServiceServer) MakeNewFolder(context.Context, *MakeNewFolderReq) (*MakeNewFolderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeNewFolder not implemented")
}
func (UnimplementedFileServiceServer) SetHiddenDoc(context.Context, *SetHiddenDocReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHiddenDoc not implemented")
}
func (UnimplementedFileServiceServer) MoveToRecycleBin(context.Context, *MoveToRecycleBinReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveToRecycleBin not implemented")
}
func (UnimplementedFileServiceServer) RecoverDocs(context.Context, *RecoverDocsReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverDocs not implemented")
}
func (UnimplementedFileServiceServer) GetRecycleBinList(context.Context, *GetRecycleBinListReq) (*GetRecycleBinListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecycleBinList not implemented")
}
func (UnimplementedFileServiceServer) SoftDelete(context.Context, *SoftDeleteReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDelete not implemented")
}
func (UnimplementedFileServiceServer) HardDelete(context.Context, *HardDeleteReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HardDelete not implemented")
}
func (UnimplementedFileServiceServer) Rename(context.Context, *RenameReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedFileServiceServer) MoveToPath(context.Context, *MoveToPathReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveToPath not implemented")
}
func (UnimplementedFileServiceServer) CopyToPath(context.Context, *CopyToPathReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyToPath not implemented")
}
func (UnimplementedFileServiceServer) CreateShare(context.Context, *CreateShareReq) (*CreateShareRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShare not implemented")
}
func (UnimplementedFileServiceServer) RetrieveShareToPath(context.Context, *RetrieveShareToPathReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveShareToPath not implemented")
}
func (UnimplementedFileServiceServer) GetShareRecords(context.Context, *GetShareRecordsReq) (*GetShareRecordsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShareRecords not implemented")
}
func (UnimplementedFileServiceServer) GetShareDetail(context.Context, *GetShareDetailReq) (*GetShareDetailRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShareDetail not implemented")
}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadFile(ctx, req.(*UploadFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetDirsAndFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDirsAndFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetDirsAndFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/GetDirsAndFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetDirsAndFiles(ctx, req.(*GetDirsAndFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/GetFileDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileDetail(ctx, req.(*GetFileDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_MakeNewFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeNewFolderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).MakeNewFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/MakeNewFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).MakeNewFolder(ctx, req.(*MakeNewFolderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_SetHiddenDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHiddenDocReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SetHiddenDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/SetHiddenDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SetHiddenDoc(ctx, req.(*SetHiddenDocReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_MoveToRecycleBin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveToRecycleBinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).MoveToRecycleBin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/MoveToRecycleBin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).MoveToRecycleBin(ctx, req.(*MoveToRecycleBinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_RecoverDocs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverDocsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).RecoverDocs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/RecoverDocs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).RecoverDocs(ctx, req.(*RecoverDocsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetRecycleBinList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecycleBinListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetRecycleBinList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/GetRecycleBinList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetRecycleBinList(ctx, req.(*GetRecycleBinListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_SoftDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SoftDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SoftDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/SoftDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SoftDelete(ctx, req.(*SoftDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_HardDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HardDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).HardDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/HardDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).HardDelete(ctx, req.(*HardDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Rename(ctx, req.(*RenameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_MoveToPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveToPathReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).MoveToPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/MoveToPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).MoveToPath(ctx, req.(*MoveToPathReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CopyToPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyToPathReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CopyToPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/CopyToPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CopyToPath(ctx, req.(*CopyToPathReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CreateShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/CreateShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateShare(ctx, req.(*CreateShareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_RetrieveShareToPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveShareToPathReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).RetrieveShareToPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/RetrieveShareToPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).RetrieveShareToPath(ctx, req.(*RetrieveShareToPathReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetShareRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShareRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetShareRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/GetShareRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetShareRecords(ctx, req.(*GetShareRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetShareDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShareDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetShareDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/GetShareDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetShareDetail(ctx, req.(*GetShareDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _FileService_UploadFile_Handler,
		},
		{
			MethodName: "GetDirsAndFiles",
			Handler:    _FileService_GetDirsAndFiles_Handler,
		},
		{
			MethodName: "GetFileDetail",
			Handler:    _FileService_GetFileDetail_Handler,
		},
		{
			MethodName: "MakeNewFolder",
			Handler:    _FileService_MakeNewFolder_Handler,
		},
		{
			MethodName: "SetHiddenDoc",
			Handler:    _FileService_SetHiddenDoc_Handler,
		},
		{
			MethodName: "MoveToRecycleBin",
			Handler:    _FileService_MoveToRecycleBin_Handler,
		},
		{
			MethodName: "RecoverDocs",
			Handler:    _FileService_RecoverDocs_Handler,
		},
		{
			MethodName: "GetRecycleBinList",
			Handler:    _FileService_GetRecycleBinList_Handler,
		},
		{
			MethodName: "SoftDelete",
			Handler:    _FileService_SoftDelete_Handler,
		},
		{
			MethodName: "HardDelete",
			Handler:    _FileService_HardDelete_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _FileService_Rename_Handler,
		},
		{
			MethodName: "MoveToPath",
			Handler:    _FileService_MoveToPath_Handler,
		},
		{
			MethodName: "CopyToPath",
			Handler:    _FileService_CopyToPath_Handler,
		},
		{
			MethodName: "CreateShare",
			Handler:    _FileService_CreateShare_Handler,
		},
		{
			MethodName: "RetrieveShareToPath",
			Handler:    _FileService_RetrieveShareToPath_Handler,
		},
		{
			MethodName: "GetShareRecords",
			Handler:    _FileService_GetShareRecords_Handler,
		},
		{
			MethodName: "GetShareDetail",
			Handler:    _FileService_GetShareDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diskapis.proto",
}

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	RegisterNewUser(ctx context.Context, in *RegisterNewUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ModifyPassword(ctx context.Context, in *ModifyPasswordReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileRsp, error)
	ModifyUserProfile(ctx context.Context, in *ModifyUserProfileReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserStorage(ctx context.Context, in *UpdateUserStorageReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RegisterNewUser(ctx context.Context, in *RegisterNewUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/RegisterNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ModifyPassword(ctx context.Context, in *ModifyPasswordReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/ModifyPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileRsp, error) {
	out := new(GetUserProfileRsp)
	err := c.cc.Invoke(ctx, "/AuthService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ModifyUserProfile(ctx context.Context, in *ModifyUserProfileReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/ModifyUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUserStorage(ctx context.Context, in *UpdateUserStorageReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/UpdateUserStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations should embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	RegisterNewUser(context.Context, *RegisterNewUserReq) (*emptypb.Empty, error)
	SignIn(context.Context, *SignInReq) (*emptypb.Empty, error)
	ModifyPassword(context.Context, *ModifyPasswordReq) (*emptypb.Empty, error)
	GetUserProfile(context.Context, *GetUserProfileReq) (*GetUserProfileRsp, error)
	ModifyUserProfile(context.Context, *ModifyUserProfileReq) (*emptypb.Empty, error)
	UpdateUserStorage(context.Context, *UpdateUserStorageReq) (*emptypb.Empty, error)
}

// UnimplementedAuthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) RegisterNewUser(context.Context, *RegisterNewUserReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNewUser not implemented")
}
func (UnimplementedAuthServiceServer) SignIn(context.Context, *SignInReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServiceServer) ModifyPassword(context.Context, *ModifyPasswordReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPassword not implemented")
}
func (UnimplementedAuthServiceServer) GetUserProfile(context.Context, *GetUserProfileReq) (*GetUserProfileRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedAuthServiceServer) ModifyUserProfile(context.Context, *ModifyUserProfileReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserProfile not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUserStorage(context.Context, *UpdateUserStorageReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserStorage not implemented")
}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_RegisterNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNewUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/RegisterNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterNewUser(ctx, req.(*RegisterNewUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignIn(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ModifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ModifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/ModifyPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ModifyPassword(ctx, req.(*ModifyPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserProfile(ctx, req.(*GetUserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ModifyUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ModifyUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/ModifyUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ModifyUserProfile(ctx, req.(*ModifyUserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUserStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserStorageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUserStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/UpdateUserStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUserStorage(ctx, req.(*UpdateUserStorageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewUser",
			Handler:    _AuthService_RegisterNewUser_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthService_SignIn_Handler,
		},
		{
			MethodName: "ModifyPassword",
			Handler:    _AuthService_ModifyPassword_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _AuthService_GetUserProfile_Handler,
		},
		{
			MethodName: "ModifyUserProfile",
			Handler:    _AuthService_ModifyUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserStorage",
			Handler:    _AuthService_UpdateUserStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diskapis.proto",
}
