// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: flow/entities/collection.proto

package entities

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

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TransactionIds [][]byte `protobuf:"bytes,2,rep,name=transaction_ids,json=transactionIds,proto3" json:"transaction_ids,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_entities_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_flow_entities_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_flow_entities_collection_proto_rawDescGZIP(), []int{0}
}

func (x *Collection) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Collection) GetTransactionIds() [][]byte {
	if x != nil {
		return x.TransactionIds
	}
	return nil
}

type CollectionGuarantee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId     []byte   `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Signatures       [][]byte `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	ReferenceBlockId []byte   `protobuf:"bytes,3,opt,name=reference_block_id,json=referenceBlockId,proto3" json:"reference_block_id,omitempty"`
	Signature        []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	SignerIds        [][]byte `protobuf:"bytes,5,rep,name=signer_ids,json=signerIds,proto3" json:"signer_ids,omitempty"`
}

func (x *CollectionGuarantee) Reset() {
	*x = CollectionGuarantee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_entities_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionGuarantee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionGuarantee) ProtoMessage() {}

func (x *CollectionGuarantee) ProtoReflect() protoreflect.Message {
	mi := &file_flow_entities_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionGuarantee.ProtoReflect.Descriptor instead.
func (*CollectionGuarantee) Descriptor() ([]byte, []int) {
	return file_flow_entities_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionGuarantee) GetCollectionId() []byte {
	if x != nil {
		return x.CollectionId
	}
	return nil
}

func (x *CollectionGuarantee) GetSignatures() [][]byte {
	if x != nil {
		return x.Signatures
	}
	return nil
}

func (x *CollectionGuarantee) GetReferenceBlockId() []byte {
	if x != nil {
		return x.ReferenceBlockId
	}
	return nil
}

func (x *CollectionGuarantee) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *CollectionGuarantee) GetSignerIds() [][]byte {
	if x != nil {
		return x.SignerIds
	}
	return nil
}

var File_flow_entities_collection_proto protoreflect.FileDescriptor

var file_flow_entities_collection_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x45, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x10, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x73, 0x42, 0x50,
	0x0a, 0x1c, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_entities_collection_proto_rawDescOnce sync.Once
	file_flow_entities_collection_proto_rawDescData = file_flow_entities_collection_proto_rawDesc
)

func file_flow_entities_collection_proto_rawDescGZIP() []byte {
	file_flow_entities_collection_proto_rawDescOnce.Do(func() {
		file_flow_entities_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_entities_collection_proto_rawDescData)
	})
	return file_flow_entities_collection_proto_rawDescData
}

var file_flow_entities_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flow_entities_collection_proto_goTypes = []interface{}{
	(*Collection)(nil),          // 0: flow.entities.Collection
	(*CollectionGuarantee)(nil), // 1: flow.entities.CollectionGuarantee
}
var file_flow_entities_collection_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flow_entities_collection_proto_init() }
func file_flow_entities_collection_proto_init() {
	if File_flow_entities_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_entities_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
		file_flow_entities_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionGuarantee); i {
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
			RawDescriptor: file_flow_entities_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_entities_collection_proto_goTypes,
		DependencyIndexes: file_flow_entities_collection_proto_depIdxs,
		MessageInfos:      file_flow_entities_collection_proto_msgTypes,
	}.Build()
	File_flow_entities_collection_proto = out.File
	file_flow_entities_collection_proto_rawDesc = nil
	file_flow_entities_collection_proto_goTypes = nil
	file_flow_entities_collection_proto_depIdxs = nil
}
