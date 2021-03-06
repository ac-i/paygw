//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: paygw/proto/app/msg.proto

package app

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

// App runtime mode
type AppMode int32

const (
	//
	AppMode_APP_MODE_UNSPECIFIED AppMode = 0
	//
	AppMode_APP_MODE_PROD AppMode = 1
	//
	AppMode_APP_MODE_STAGE AppMode = 2
	//
	AppMode_APP_MODE_TEST AppMode = 3
	//
	AppMode_APP_MODE_DEV AppMode = 4
	//
	AppMode_APP_MODE_FREE AppMode = 5
)

// Enum value maps for AppMode.
var (
	AppMode_name = map[int32]string{
		0: "APP_MODE_UNSPECIFIED",
		1: "APP_MODE_PROD",
		2: "APP_MODE_STAGE",
		3: "APP_MODE_TEST",
		4: "APP_MODE_DEV",
		5: "APP_MODE_FREE",
	}
	AppMode_value = map[string]int32{
		"APP_MODE_UNSPECIFIED": 0,
		"APP_MODE_PROD":        1,
		"APP_MODE_STAGE":       2,
		"APP_MODE_TEST":        3,
		"APP_MODE_DEV":         4,
		"APP_MODE_FREE":        5,
	}
)

func (x AppMode) Enum() *AppMode {
	p := new(AppMode)
	*p = x
	return p
}

func (x AppMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppMode) Descriptor() protoreflect.EnumDescriptor {
	return file_paygw_proto_app_msg_proto_enumTypes[0].Descriptor()
}

func (AppMode) Type() protoreflect.EnumType {
	return &file_paygw_proto_app_msg_proto_enumTypes[0]
}

func (x AppMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppMode.Descriptor instead.
func (AppMode) EnumDescriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{0}
}

// App domain
type AppDom int32

const (
	//
	AppDom_APP_DOM_UNSPECIFIED AppDom = 0
	//
	AppDom_APP_DOM_ALL AppDom = 1
	//
	AppDom_APP_DOM_PAYCARD AppDom = 2
	//
	AppDom_APP_DOM_PORT AppDom = 3
	//
	AppDom_APP_DOM_STORE AppDom = 4
)

// Enum value maps for AppDom.
var (
	AppDom_name = map[int32]string{
		0: "APP_DOM_UNSPECIFIED",
		1: "APP_DOM_ALL",
		2: "APP_DOM_PAYCARD",
		3: "APP_DOM_PORT",
		4: "APP_DOM_STORE",
	}
	AppDom_value = map[string]int32{
		"APP_DOM_UNSPECIFIED": 0,
		"APP_DOM_ALL":         1,
		"APP_DOM_PAYCARD":     2,
		"APP_DOM_PORT":        3,
		"APP_DOM_STORE":       4,
	}
)

func (x AppDom) Enum() *AppDom {
	p := new(AppDom)
	*p = x
	return p
}

func (x AppDom) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppDom) Descriptor() protoreflect.EnumDescriptor {
	return file_paygw_proto_app_msg_proto_enumTypes[1].Descriptor()
}

func (AppDom) Type() protoreflect.EnumType {
	return &file_paygw_proto_app_msg_proto_enumTypes[1]
}

func (x AppDom) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppDom.Descriptor instead.
func (AppDom) EnumDescriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{1}
}

// Config configures app instance
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Environment *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	//
	Active int32 `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	//
	Server *Server `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetEnvironment() *Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *Config) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *Config) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

// App runtime environment context
type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	AppName string `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	//
	AppMode AppMode `protobuf:"varint,3,opt,name=app_mode,json=appMode,proto3,enum=proto.app.AppMode" json:"app_mode,omitempty"`
	//
	AppHost string `protobuf:"bytes,4,opt,name=app_host,json=appHost,proto3" json:"app_host,omitempty"`
	//
	AppPid int32 `protobuf:"varint,5,opt,name=app_pid,json=appPid,proto3" json:"app_pid,omitempty"`
	//
	AppStart int64 `protobuf:"varint,6,opt,name=app_start,json=appStart,proto3" json:"app_start,omitempty"`
	// Activate Servers & Storages [grpc restgw rest]
	Servers []string `protobuf:"bytes,7,rep,name=servers,proto3" json:"servers,omitempty"`
	// Activate Storages [mem grpc aql cql]
	Storages []string `protobuf:"bytes,8,rep,name=storages,proto3" json:"storages,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Environment) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *Environment) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Environment) GetAppMode() AppMode {
	if x != nil {
		return x.AppMode
	}
	return AppMode_APP_MODE_UNSPECIFIED
}

func (x *Environment) GetAppHost() string {
	if x != nil {
		return x.AppHost
	}
	return ""
}

func (x *Environment) GetAppPid() int32 {
	if x != nil {
		return x.AppPid
	}
	return 0
}

func (x *Environment) GetAppStart() int64 {
	if x != nil {
		return x.AppStart
	}
	return 0
}

func (x *Environment) GetServers() []string {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *Environment) GetStorages() []string {
	if x != nil {
		return x.Storages
	}
	return nil
}

// Domain Server
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	ServGrpc *ServGRPC `protobuf:"bytes,2,opt,name=serv_grpc,json=servGrpc,proto3" json:"serv_grpc,omitempty"`
	//
	ServRestgw *ServRESTGW `protobuf:"bytes,3,opt,name=serv_restgw,json=servRestgw,proto3" json:"serv_restgw,omitempty"`
	//
	ServRest *ServREST `protobuf:"bytes,4,opt,name=serv_rest,json=servRest,proto3" json:"serv_rest,omitempty"` // unimplemented
	//
	CliGrpc *CliGRPC `protobuf:"bytes,5,opt,name=cli_grpc,json=cliGrpc,proto3" json:"cli_grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *Server) GetServGrpc() *ServGRPC {
	if x != nil {
		return x.ServGrpc
	}
	return nil
}

func (x *Server) GetServRestgw() *ServRESTGW {
	if x != nil {
		return x.ServRestgw
	}
	return nil
}

func (x *Server) GetServRest() *ServREST {
	if x != nil {
		return x.ServRest
	}
	return nil
}

func (x *Server) GetCliGrpc() *CliGRPC {
	if x != nil {
		return x.CliGrpc
	}
	return nil
}

//
type CliGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	// grpc.DialContext(ctx context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error)
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *CliGRPC) Reset() {
	*x = CliGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CliGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CliGRPC) ProtoMessage() {}

func (x *CliGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CliGRPC.ProtoReflect.Descriptor instead.
func (*CliGRPC) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{3}
}

func (x *CliGRPC) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *CliGRPC) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *CliGRPC) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

//
type ServGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	// Listen(network string, address string)
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	// Listen(network string, address string)
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// serv storage in ServGRPC only
	// map<string, StorageDriver> storage_drivers = 5;
	// serv storage in ServGRPC only
	Storage *Storage `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *ServGRPC) Reset() {
	*x = ServGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServGRPC) ProtoMessage() {}

func (x *ServGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServGRPC.ProtoReflect.Descriptor instead.
func (*ServGRPC) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{4}
}

func (x *ServGRPC) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *ServGRPC) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *ServGRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *ServGRPC) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ServGRPC) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

//
type ServRESTGW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	// RegisterPortServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption)
	GrpcEndpoint string `protobuf:"bytes,3,opt,name=grpc_endpoint,json=grpcEndpoint,proto3" json:"grpc_endpoint,omitempty"`
	// http.Server struct { Addr string
	RestgwAddr string `protobuf:"bytes,4,opt,name=restgw_addr,json=restgwAddr,proto3" json:"restgw_addr,omitempty"`
}

func (x *ServRESTGW) Reset() {
	*x = ServRESTGW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServRESTGW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServRESTGW) ProtoMessage() {}

func (x *ServRESTGW) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServRESTGW.ProtoReflect.Descriptor instead.
func (*ServRESTGW) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{5}
}

func (x *ServRESTGW) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *ServRESTGW) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *ServRESTGW) GetGrpcEndpoint() string {
	if x != nil {
		return x.GrpcEndpoint
	}
	return ""
}

func (x *ServRESTGW) GetRestgwAddr() string {
	if x != nil {
		return x.RestgwAddr
	}
	return ""
}

//
type ServREST struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	// http.Server struct { Addr string
	RestAddr string `protobuf:"bytes,3,opt,name=rest_addr,json=restAddr,proto3" json:"rest_addr,omitempty"`
}

func (x *ServREST) Reset() {
	*x = ServREST{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServREST) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServREST) ProtoMessage() {}

func (x *ServREST) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServREST.ProtoReflect.Descriptor instead.
func (*ServREST) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{6}
}

func (x *ServREST) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *ServREST) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *ServREST) GetRestAddr() string {
	if x != nil {
		return x.RestAddr
	}
	return ""
}

// Storage
type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	// serv storage in ServGRPC only
	Drivers map[string]*StorageDriver `protobuf:"bytes,2,rep,name=drivers,proto3" json:"drivers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{7}
}

func (x *Storage) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *Storage) GetDrivers() map[string]*StorageDriver {
	if x != nil {
		return x.Drivers
	}
	return nil
}

//
type StorageDriver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	//
	Hosts []string `protobuf:"bytes,3,rep,name=hosts,proto3" json:"hosts,omitempty"`
	//
	Auth *BasicAuth `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *StorageDriver) Reset() {
	*x = StorageDriver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageDriver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageDriver) ProtoMessage() {}

func (x *StorageDriver) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageDriver.ProtoReflect.Descriptor instead.
func (*StorageDriver) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{8}
}

func (x *StorageDriver) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *StorageDriver) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *StorageDriver) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *StorageDriver) GetAuth() *BasicAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

//
type BasicAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Active int32 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	//
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	//
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	//
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *BasicAuth) Reset() {
	*x = BasicAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_app_msg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAuth) ProtoMessage() {}

func (x *BasicAuth) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_app_msg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAuth.ProtoReflect.Descriptor instead.
func (*BasicAuth) Descriptor() ([]byte, []int) {
	return file_paygw_proto_app_msg_proto_rawDescGZIP(), []int{9}
}

func (x *BasicAuth) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *BasicAuth) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *BasicAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BasicAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_paygw_proto_app_msg_proto protoreflect.FileDescriptor

var file_paygw_proto_app_msg_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x61, 0x79, 0x67, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x22, 0x85, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x38, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xf6,
	0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x41, 0x70, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61,
	0x70, 0x70, 0x5f, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x70,
	0x70, 0x50, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x70, 0x70, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x22, 0xeb, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x47, 0x52,
	0x50, 0x43, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x47, 0x72, 0x70, 0x63, 0x12, 0x36, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x67, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x52, 0x45, 0x53, 0x54, 0x47, 0x57, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x52, 0x65,
	0x73, 0x74, 0x67, 0x77, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x5f, 0x72, 0x65, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x52, 0x45, 0x53, 0x54, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x52, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6c, 0x69, 0x47, 0x52, 0x50, 0x43, 0x52, 0x07, 0x63, 0x6c,
	0x69, 0x47, 0x72, 0x70, 0x63, 0x22, 0x51, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x47, 0x52, 0x50, 0x43,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x72,
	0x76, 0x47, 0x52, 0x50, 0x43, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x52, 0x45, 0x53, 0x54, 0x47, 0x57, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67,
	0x72, 0x70, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x74, 0x67, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x67, 0x77, 0x41, 0x64, 0x64, 0x72, 0x22, 0x57, 0x0a, 0x08,
	0x53, 0x65, 0x72, 0x76, 0x52, 0x45, 0x53, 0x54, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x22, 0xb2, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x1a, 0x54, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7f, 0x0a, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x28, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x73, 0x0a, 0x09, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x2a, 0x82, 0x01, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x41, 0x50, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x50, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x50, 0x50,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x41, 0x50, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x03,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x50, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x56,
	0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x46,
	0x52, 0x45, 0x45, 0x10, 0x05, 0x2a, 0x6c, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x44, 0x6f, 0x6d, 0x12,
	0x17, 0x0a, 0x13, 0x41, 0x50, 0x50, 0x5f, 0x44, 0x4f, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x5f,
	0x44, 0x4f, 0x4d, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x50, 0x50,
	0x5f, 0x44, 0x4f, 0x4d, 0x5f, 0x50, 0x41, 0x59, 0x43, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x41, 0x50, 0x50, 0x5f, 0x44, 0x4f, 0x4d, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x03,
	0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x50, 0x5f, 0x44, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x4f, 0x52,
	0x45, 0x10, 0x04, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x61, 0x79, 0x67, 0x77, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paygw_proto_app_msg_proto_rawDescOnce sync.Once
	file_paygw_proto_app_msg_proto_rawDescData = file_paygw_proto_app_msg_proto_rawDesc
)

func file_paygw_proto_app_msg_proto_rawDescGZIP() []byte {
	file_paygw_proto_app_msg_proto_rawDescOnce.Do(func() {
		file_paygw_proto_app_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_paygw_proto_app_msg_proto_rawDescData)
	})
	return file_paygw_proto_app_msg_proto_rawDescData
}

var file_paygw_proto_app_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_paygw_proto_app_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_paygw_proto_app_msg_proto_goTypes = []interface{}{
	(AppMode)(0),          // 0: proto.app.AppMode
	(AppDom)(0),           // 1: proto.app.AppDom
	(*Config)(nil),        // 2: proto.app.Config
	(*Environment)(nil),   // 3: proto.app.Environment
	(*Server)(nil),        // 4: proto.app.Server
	(*CliGRPC)(nil),       // 5: proto.app.CliGRPC
	(*ServGRPC)(nil),      // 6: proto.app.ServGRPC
	(*ServRESTGW)(nil),    // 7: proto.app.ServRESTGW
	(*ServREST)(nil),      // 8: proto.app.ServREST
	(*Storage)(nil),       // 9: proto.app.Storage
	(*StorageDriver)(nil), // 10: proto.app.StorageDriver
	(*BasicAuth)(nil),     // 11: proto.app.BasicAuth
	nil,                   // 12: proto.app.Storage.DriversEntry
}
var file_paygw_proto_app_msg_proto_depIdxs = []int32{
	3,  // 0: proto.app.Config.environment:type_name -> proto.app.Environment
	4,  // 1: proto.app.Config.server:type_name -> proto.app.Server
	0,  // 2: proto.app.Environment.app_mode:type_name -> proto.app.AppMode
	6,  // 3: proto.app.Server.serv_grpc:type_name -> proto.app.ServGRPC
	7,  // 4: proto.app.Server.serv_restgw:type_name -> proto.app.ServRESTGW
	8,  // 5: proto.app.Server.serv_rest:type_name -> proto.app.ServREST
	5,  // 6: proto.app.Server.cli_grpc:type_name -> proto.app.CliGRPC
	9,  // 7: proto.app.ServGRPC.storage:type_name -> proto.app.Storage
	12, // 8: proto.app.Storage.drivers:type_name -> proto.app.Storage.DriversEntry
	11, // 9: proto.app.StorageDriver.auth:type_name -> proto.app.BasicAuth
	10, // 10: proto.app.Storage.DriversEntry.value:type_name -> proto.app.StorageDriver
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_paygw_proto_app_msg_proto_init() }
func file_paygw_proto_app_msg_proto_init() {
	if File_paygw_proto_app_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paygw_proto_app_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CliGRPC); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServGRPC); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServRESTGW); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServREST); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageDriver); i {
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
		file_paygw_proto_app_msg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAuth); i {
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
			RawDescriptor: file_paygw_proto_app_msg_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_paygw_proto_app_msg_proto_goTypes,
		DependencyIndexes: file_paygw_proto_app_msg_proto_depIdxs,
		EnumInfos:         file_paygw_proto_app_msg_proto_enumTypes,
		MessageInfos:      file_paygw_proto_app_msg_proto_msgTypes,
	}.Build()
	File_paygw_proto_app_msg_proto = out.File
	file_paygw_proto_app_msg_proto_rawDesc = nil
	file_paygw_proto_app_msg_proto_goTypes = nil
	file_paygw_proto_app_msg_proto_depIdxs = nil
}
