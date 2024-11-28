// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: econ.proto

// Cut down version of cstrike15_gcmessages.proto

package econ_pb2

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

type CEconItemPreviewDataBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accountid          *uint32                              `protobuf:"varint,1,opt,name=accountid,proto3,oneof" json:"accountid,omitempty"`
	Itemid             *uint64                              `protobuf:"varint,2,opt,name=itemid,proto3,oneof" json:"itemid,omitempty"`
	Defindex           *uint32                              `protobuf:"varint,3,opt,name=defindex,proto3,oneof" json:"defindex,omitempty"`
	Paintindex         *uint32                              `protobuf:"varint,4,opt,name=paintindex,proto3,oneof" json:"paintindex,omitempty"`
	Rarity             *uint32                              `protobuf:"varint,5,opt,name=rarity,proto3,oneof" json:"rarity,omitempty"`
	Quality            *uint32                              `protobuf:"varint,6,opt,name=quality,proto3,oneof" json:"quality,omitempty"`
	Paintwear          *uint32                              `protobuf:"varint,7,opt,name=paintwear,proto3,oneof" json:"paintwear,omitempty"`
	Paintseed          *uint32                              `protobuf:"varint,8,opt,name=paintseed,proto3,oneof" json:"paintseed,omitempty"`
	Killeaterscoretype *uint32                              `protobuf:"varint,9,opt,name=killeaterscoretype,proto3,oneof" json:"killeaterscoretype,omitempty"`
	Killeatervalue     *uint32                              `protobuf:"varint,10,opt,name=killeatervalue,proto3,oneof" json:"killeatervalue,omitempty"`
	Customname         *string                              `protobuf:"bytes,11,opt,name=customname,proto3,oneof" json:"customname,omitempty"`
	Stickers           []*CEconItemPreviewDataBlock_Sticker `protobuf:"bytes,12,rep,name=stickers,proto3" json:"stickers,omitempty"`
	Inventory          *uint32                              `protobuf:"varint,13,opt,name=inventory,proto3,oneof" json:"inventory,omitempty"`
	Origin             *uint32                              `protobuf:"varint,14,opt,name=origin,proto3,oneof" json:"origin,omitempty"`
	Questid            *uint32                              `protobuf:"varint,15,opt,name=questid,proto3,oneof" json:"questid,omitempty"`
	Dropreason         *uint32                              `protobuf:"varint,16,opt,name=dropreason,proto3,oneof" json:"dropreason,omitempty"`
	Musicindex         *uint32                              `protobuf:"varint,17,opt,name=musicindex,proto3,oneof" json:"musicindex,omitempty"`
	Entindex           *int32                               `protobuf:"varint,18,opt,name=entindex,proto3,oneof" json:"entindex,omitempty"`
}

func (x *CEconItemPreviewDataBlock) Reset() {
	*x = CEconItemPreviewDataBlock{}
	mi := &file_econ_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CEconItemPreviewDataBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CEconItemPreviewDataBlock) ProtoMessage() {}

func (x *CEconItemPreviewDataBlock) ProtoReflect() protoreflect.Message {
	mi := &file_econ_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CEconItemPreviewDataBlock.ProtoReflect.Descriptor instead.
func (*CEconItemPreviewDataBlock) Descriptor() ([]byte, []int) {
	return file_econ_proto_rawDescGZIP(), []int{0}
}

func (x *CEconItemPreviewDataBlock) GetAccountid() uint32 {
	if x != nil && x.Accountid != nil {
		return *x.Accountid
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetItemid() uint64 {
	if x != nil && x.Itemid != nil {
		return *x.Itemid
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetDefindex() uint32 {
	if x != nil && x.Defindex != nil {
		return *x.Defindex
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetPaintindex() uint32 {
	if x != nil && x.Paintindex != nil {
		return *x.Paintindex
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetRarity() uint32 {
	if x != nil && x.Rarity != nil {
		return *x.Rarity
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetQuality() uint32 {
	if x != nil && x.Quality != nil {
		return *x.Quality
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetPaintwear() uint32 {
	if x != nil && x.Paintwear != nil {
		return *x.Paintwear
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetPaintseed() uint32 {
	if x != nil && x.Paintseed != nil {
		return *x.Paintseed
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetKilleaterscoretype() uint32 {
	if x != nil && x.Killeaterscoretype != nil {
		return *x.Killeaterscoretype
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetKilleatervalue() uint32 {
	if x != nil && x.Killeatervalue != nil {
		return *x.Killeatervalue
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetCustomname() string {
	if x != nil && x.Customname != nil {
		return *x.Customname
	}
	return ""
}

func (x *CEconItemPreviewDataBlock) GetStickers() []*CEconItemPreviewDataBlock_Sticker {
	if x != nil {
		return x.Stickers
	}
	return nil
}

func (x *CEconItemPreviewDataBlock) GetInventory() uint32 {
	if x != nil && x.Inventory != nil {
		return *x.Inventory
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetOrigin() uint32 {
	if x != nil && x.Origin != nil {
		return *x.Origin
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetQuestid() uint32 {
	if x != nil && x.Questid != nil {
		return *x.Questid
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetDropreason() uint32 {
	if x != nil && x.Dropreason != nil {
		return *x.Dropreason
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetMusicindex() uint32 {
	if x != nil && x.Musicindex != nil {
		return *x.Musicindex
	}
	return 0
}

func (x *CEconItemPreviewDataBlock) GetEntindex() int32 {
	if x != nil && x.Entindex != nil {
		return *x.Entindex
	}
	return 0
}

type CEconItemPreviewDataBlock_Sticker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot      *uint32  `protobuf:"varint,1,opt,name=slot,proto3,oneof" json:"slot,omitempty"`
	StickerId *uint32  `protobuf:"varint,2,opt,name=sticker_id,json=stickerId,proto3,oneof" json:"sticker_id,omitempty"`
	Wear      *float32 `protobuf:"fixed32,3,opt,name=wear,proto3,oneof" json:"wear,omitempty"`
	Scale     *float32 `protobuf:"fixed32,4,opt,name=scale,proto3,oneof" json:"scale,omitempty"`
	Rotation  *float32 `protobuf:"fixed32,5,opt,name=rotation,proto3,oneof" json:"rotation,omitempty"`
	TintId    *uint32  `protobuf:"varint,6,opt,name=tint_id,json=tintId,proto3,oneof" json:"tint_id,omitempty"`
}

func (x *CEconItemPreviewDataBlock_Sticker) Reset() {
	*x = CEconItemPreviewDataBlock_Sticker{}
	mi := &file_econ_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CEconItemPreviewDataBlock_Sticker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CEconItemPreviewDataBlock_Sticker) ProtoMessage() {}

func (x *CEconItemPreviewDataBlock_Sticker) ProtoReflect() protoreflect.Message {
	mi := &file_econ_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CEconItemPreviewDataBlock_Sticker.ProtoReflect.Descriptor instead.
func (*CEconItemPreviewDataBlock_Sticker) Descriptor() ([]byte, []int) {
	return file_econ_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CEconItemPreviewDataBlock_Sticker) GetSlot() uint32 {
	if x != nil && x.Slot != nil {
		return *x.Slot
	}
	return 0
}

func (x *CEconItemPreviewDataBlock_Sticker) GetStickerId() uint32 {
	if x != nil && x.StickerId != nil {
		return *x.StickerId
	}
	return 0
}

func (x *CEconItemPreviewDataBlock_Sticker) GetWear() float32 {
	if x != nil && x.Wear != nil {
		return *x.Wear
	}
	return 0
}

func (x *CEconItemPreviewDataBlock_Sticker) GetScale() float32 {
	if x != nil && x.Scale != nil {
		return *x.Scale
	}
	return 0
}

func (x *CEconItemPreviewDataBlock_Sticker) GetRotation() float32 {
	if x != nil && x.Rotation != nil {
		return *x.Rotation
	}
	return 0
}

func (x *CEconItemPreviewDataBlock_Sticker) GetTintId() uint32 {
	if x != nil && x.TintId != nil {
		return *x.TintId
	}
	return 0
}

var File_econ_proto protoreflect.FileDescriptor

var file_econ_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x63,
	0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x32, 0x22, 0xae, 0x09, 0x0a, 0x19, 0x43, 0x45, 0x63, 0x6f, 0x6e,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x08, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x69,
	0x6e, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x06, 0x72, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x77,
	0x65, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x06, 0x52, 0x09, 0x70, 0x61, 0x69,
	0x6e, 0x74, 0x77, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x65, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x09,
	0x70, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x12,
	0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x61, 0x74, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x08, 0x52, 0x12, 0x6b, 0x69, 0x6c, 0x6c,
	0x65, 0x61, 0x74, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x61, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x09, 0x52, 0x0e, 0x6b, 0x69, 0x6c,
	0x6c, 0x65, 0x61, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0a, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x08, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x32,
	0x2e, 0x43, 0x45, 0x63, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x52, 0x08, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x09,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x0b, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x0c, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x0d, 0x52,
	0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x64,
	0x72, 0x6f, 0x70, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x0e, 0x52, 0x0a, 0x64, 0x72, 0x6f, 0x70, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x0f, 0x52, 0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x48, 0x10, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x1a, 0xfd, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x01, 0x52, 0x09, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x77, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x02, 0x52,
	0x04, 0x77, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x48, 0x03, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x48, 0x04, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x74, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x06, 0x74, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x77,
	0x65, 0x61, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74,
	0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x69, 0x64, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x77, 0x65, 0x61, 0x72,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x65, 0x65, 0x64, 0x42, 0x15,
	0x0a, 0x13, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x61, 0x74, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x64, 0x72, 0x6f, 0x70, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x61, 0x6e, 0x2d, 0x79, 0x61, 0x6e, 0x67, 0x2f,
	0x63, 0x73, 0x32, 0x2d, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2d, 0x67, 0x65, 0x6e, 0x2f,
	0x65, 0x63, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x32, 0x3b, 0x65, 0x63, 0x6f, 0x6e, 0x5f, 0x70, 0x62,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_econ_proto_rawDescOnce sync.Once
	file_econ_proto_rawDescData = file_econ_proto_rawDesc
)

func file_econ_proto_rawDescGZIP() []byte {
	file_econ_proto_rawDescOnce.Do(func() {
		file_econ_proto_rawDescData = protoimpl.X.CompressGZIP(file_econ_proto_rawDescData)
	})
	return file_econ_proto_rawDescData
}

var file_econ_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_econ_proto_goTypes = []any{
	(*CEconItemPreviewDataBlock)(nil),         // 0: econ_pb2.CEconItemPreviewDataBlock
	(*CEconItemPreviewDataBlock_Sticker)(nil), // 1: econ_pb2.CEconItemPreviewDataBlock.Sticker
}
var file_econ_proto_depIdxs = []int32{
	1, // 0: econ_pb2.CEconItemPreviewDataBlock.stickers:type_name -> econ_pb2.CEconItemPreviewDataBlock.Sticker
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_econ_proto_init() }
func file_econ_proto_init() {
	if File_econ_proto != nil {
		return
	}
	file_econ_proto_msgTypes[0].OneofWrappers = []any{}
	file_econ_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_econ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_econ_proto_goTypes,
		DependencyIndexes: file_econ_proto_depIdxs,
		MessageInfos:      file_econ_proto_msgTypes,
	}.Build()
	File_econ_proto = out.File
	file_econ_proto_rawDesc = nil
	file_econ_proto_goTypes = nil
	file_econ_proto_depIdxs = nil
}