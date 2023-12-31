// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: gra.proto

package proto

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

type KolorZolwia int32

const (
	KolorZolwia_XXX    KolorZolwia = 0
	KolorZolwia_RED    KolorZolwia = 1
	KolorZolwia_GREEN  KolorZolwia = 2
	KolorZolwia_BLUE   KolorZolwia = 3
	KolorZolwia_YELLOW KolorZolwia = 4
	KolorZolwia_PURPLE KolorZolwia = 5
)

// Enum value maps for KolorZolwia.
var (
	KolorZolwia_name = map[int32]string{
		0: "XXX",
		1: "RED",
		2: "GREEN",
		3: "BLUE",
		4: "YELLOW",
		5: "PURPLE",
	}
	KolorZolwia_value = map[string]int32{
		"XXX":    0,
		"RED":    1,
		"GREEN":  2,
		"BLUE":   3,
		"YELLOW": 4,
		"PURPLE": 5,
	}
)

func (x KolorZolwia) Enum() *KolorZolwia {
	p := new(KolorZolwia)
	*p = x
	return p
}

func (x KolorZolwia) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KolorZolwia) Descriptor() protoreflect.EnumDescriptor {
	return file_gra_proto_enumTypes[0].Descriptor()
}

func (KolorZolwia) Type() protoreflect.EnumType {
	return &file_gra_proto_enumTypes[0]
}

func (x KolorZolwia) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KolorZolwia.Descriptor instead.
func (KolorZolwia) EnumDescriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{0}
}

type Karta int32

const (
	Karta_XX  Karta = 0 // czerwony jeden do przodu
	Karta_R1  Karta = 1 // czerwony jeden do przodu
	Karta_R2  Karta = 2 // czerwony dwa do przodu
	Karta_R1B Karta = 3 //czerwony jeden do tyłu
	Karta_G1  Karta = 4 // zielony
	Karta_G2  Karta = 5
	Karta_G1B Karta = 6
	Karta_B1  Karta = 7 // niebieski
	Karta_B2  Karta = 8
	Karta_B1B Karta = 9
	Karta_Y1  Karta = 10 // żółty
	Karta_Y2  Karta = 11
	Karta_Y1B Karta = 12
	Karta_P1  Karta = 13 // fioletowy
	Karta_P2  Karta = 14
	Karta_P1B Karta = 15
	Karta_L1  Karta = 16 // ostatni
	Karta_L2  Karta = 17
	Karta_A1  Karta = 18 // dowolny
	Karta_A1B Karta = 19
)

// Enum value maps for Karta.
var (
	Karta_name = map[int32]string{
		0:  "XX",
		1:  "R1",
		2:  "R2",
		3:  "R1B",
		4:  "G1",
		5:  "G2",
		6:  "G1B",
		7:  "B1",
		8:  "B2",
		9:  "B1B",
		10: "Y1",
		11: "Y2",
		12: "Y1B",
		13: "P1",
		14: "P2",
		15: "P1B",
		16: "L1",
		17: "L2",
		18: "A1",
		19: "A1B",
	}
	Karta_value = map[string]int32{
		"XX":  0,
		"R1":  1,
		"R2":  2,
		"R1B": 3,
		"G1":  4,
		"G2":  5,
		"G1B": 6,
		"B1":  7,
		"B2":  8,
		"B1B": 9,
		"Y1":  10,
		"Y2":  11,
		"Y1B": 12,
		"P1":  13,
		"P2":  14,
		"P1B": 15,
		"L1":  16,
		"L2":  17,
		"A1":  18,
		"A1B": 19,
	}
)

func (x Karta) Enum() *Karta {
	p := new(Karta)
	*p = x
	return p
}

func (x Karta) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Karta) Descriptor() protoreflect.EnumDescriptor {
	return file_gra_proto_enumTypes[1].Descriptor()
}

func (Karta) Type() protoreflect.EnumType {
	return &file_gra_proto_enumTypes[1]
}

func (x Karta) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Karta.Descriptor instead.
func (Karta) EnumDescriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{1}
}

type KonfiguracjaGry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiczbaGraczy int32 `protobuf:"varint,1,opt,name=liczbaGraczy,proto3" json:"liczbaGraczy,omitempty"`
}

func (x *KonfiguracjaGry) Reset() {
	*x = KonfiguracjaGry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KonfiguracjaGry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KonfiguracjaGry) ProtoMessage() {}

func (x *KonfiguracjaGry) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KonfiguracjaGry.ProtoReflect.Descriptor instead.
func (*KonfiguracjaGry) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{0}
}

func (x *KonfiguracjaGry) GetLiczbaGraczy() int32 {
	if x != nil {
		return x.LiczbaGraczy
	}
	return 0
}

type NowaGraInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID string `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
}

func (x *NowaGraInfo) Reset() {
	*x = NowaGraInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowaGraInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowaGraInfo) ProtoMessage() {}

func (x *NowaGraInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowaGraInfo.ProtoReflect.Descriptor instead.
func (*NowaGraInfo) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{1}
}

func (x *NowaGraInfo) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

type Dolaczanie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID       string `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
	NazwaGracza string `protobuf:"bytes,2,opt,name=nazwaGracza,proto3" json:"nazwaGracza,omitempty"`
}

func (x *Dolaczanie) Reset() {
	*x = Dolaczanie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dolaczanie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dolaczanie) ProtoMessage() {}

func (x *Dolaczanie) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dolaczanie.ProtoReflect.Descriptor instead.
func (*Dolaczanie) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{2}
}

func (x *Dolaczanie) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

func (x *Dolaczanie) GetNazwaGracza() string {
	if x != nil {
		return x.NazwaGracza
	}
	return ""
}

type StanGry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID      string      `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
	GraczID    string      `protobuf:"bytes,2,opt,name=graczID,proto3" json:"graczID,omitempty"`
	TwojKolor  KolorZolwia `protobuf:"varint,3,opt,name=twojKolor,proto3,enum=KolorZolwia" json:"twojKolor,omitempty"`
	TwojeKarty []Karta     `protobuf:"varint,4,rep,packed,name=twojeKarty,proto3,enum=Karta" json:"twojeKarty,omitempty"`
	Plansza    []*Pole     `protobuf:"bytes,5,rep,name=plansza,proto3" json:"plansza,omitempty"`
	CzyKoniec  bool        `protobuf:"varint,6,opt,name=CzyKoniec,proto3" json:"CzyKoniec,omitempty"`
	KtoWygral  int32       `protobuf:"varint,7,opt,name=KtoWygral,proto3" json:"KtoWygral,omitempty"`
}

func (x *StanGry) Reset() {
	*x = StanGry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StanGry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StanGry) ProtoMessage() {}

func (x *StanGry) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StanGry.ProtoReflect.Descriptor instead.
func (*StanGry) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{3}
}

func (x *StanGry) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

func (x *StanGry) GetGraczID() string {
	if x != nil {
		return x.GraczID
	}
	return ""
}

func (x *StanGry) GetTwojKolor() KolorZolwia {
	if x != nil {
		return x.TwojKolor
	}
	return KolorZolwia_XXX
}

func (x *StanGry) GetTwojeKarty() []Karta {
	if x != nil {
		return x.TwojeKarty
	}
	return nil
}

func (x *StanGry) GetPlansza() []*Pole {
	if x != nil {
		return x.Plansza
	}
	return nil
}

func (x *StanGry) GetCzyKoniec() bool {
	if x != nil {
		return x.CzyKoniec
	}
	return false
}

func (x *StanGry) GetKtoWygral() int32 {
	if x != nil {
		return x.KtoWygral
	}
	return 0
}

type Pole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zolwie []KolorZolwia `protobuf:"varint,2,rep,packed,name=zolwie,proto3,enum=KolorZolwia" json:"zolwie,omitempty"`
}

func (x *Pole) Reset() {
	*x = Pole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pole) ProtoMessage() {}

func (x *Pole) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pole.ProtoReflect.Descriptor instead.
func (*Pole) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{4}
}

func (x *Pole) GetZolwie() []KolorZolwia {
	if x != nil {
		return x.Zolwie
	}
	return nil
}

type RuchGracza struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID        string      `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
	GraczID      string      `protobuf:"bytes,2,opt,name=graczID,proto3" json:"graczID,omitempty"`
	ZagranaKarta Karta       `protobuf:"varint,3,opt,name=zagranaKarta,proto3,enum=Karta" json:"zagranaKarta,omitempty"`
	KolorWybrany KolorZolwia `protobuf:"varint,4,opt,name=kolorWybrany,proto3,enum=KolorZolwia" json:"kolorWybrany,omitempty"`
}

func (x *RuchGracza) Reset() {
	*x = RuchGracza{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuchGracza) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuchGracza) ProtoMessage() {}

func (x *RuchGracza) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuchGracza.ProtoReflect.Descriptor instead.
func (*RuchGracza) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{5}
}

func (x *RuchGracza) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

func (x *RuchGracza) GetGraczID() string {
	if x != nil {
		return x.GraczID
	}
	return ""
}

func (x *RuchGracza) GetZagranaKarta() Karta {
	if x != nil {
		return x.ZagranaKarta
	}
	return Karta_XX
}

func (x *RuchGracza) GetKolorWybrany() KolorZolwia {
	if x != nil {
		return x.KolorWybrany
	}
	return KolorZolwia_XXX
}

var File_gra_proto protoreflect.FileDescriptor

var file_gra_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0f, 0x4b,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x63, 0x6a, 0x61, 0x47, 0x72, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x69, 0x63, 0x7a, 0x62, 0x61, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x69, 0x63, 0x7a, 0x62, 0x61, 0x47, 0x72, 0x61, 0x63,
	0x7a, 0x79, 0x22, 0x23, 0x0a, 0x0b, 0x4e, 0x6f, 0x77, 0x61, 0x47, 0x72, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x22, 0x44, 0x0a, 0x0a, 0x44, 0x6f, 0x6c, 0x61, 0x63,
	0x7a, 0x61, 0x6e, 0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6e,
	0x61, 0x7a, 0x77, 0x61, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x61, 0x7a, 0x77, 0x61, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61, 0x22, 0xea, 0x01,
	0x0a, 0x07, 0x53, 0x74, 0x61, 0x6e, 0x47, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x63, 0x7a, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x61, 0x63, 0x7a, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x09, 0x74, 0x77, 0x6f,
	0x6a, 0x4b, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4b,
	0x6f, 0x6c, 0x6f, 0x72, 0x5a, 0x6f, 0x6c, 0x77, 0x69, 0x61, 0x52, 0x09, 0x74, 0x77, 0x6f, 0x6a,
	0x4b, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0a, 0x74, 0x77, 0x6f, 0x6a, 0x65, 0x4b, 0x61,
	0x72, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x4b, 0x61, 0x72, 0x74,
	0x61, 0x52, 0x0a, 0x74, 0x77, 0x6f, 0x6a, 0x65, 0x4b, 0x61, 0x72, 0x74, 0x79, 0x12, 0x1f, 0x0a,
	0x07, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x7a, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x50, 0x6f, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x7a, 0x61, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x7a, 0x79, 0x4b, 0x6f, 0x6e, 0x69, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x43, 0x7a, 0x79, 0x4b, 0x6f, 0x6e, 0x69, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09,
	0x4b, 0x74, 0x6f, 0x57, 0x79, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x4b, 0x74, 0x6f, 0x57, 0x79, 0x67, 0x72, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x04, 0x50, 0x6f,
	0x6c, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x7a, 0x6f, 0x6c, 0x77, 0x69, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4b, 0x6f, 0x6c, 0x6f, 0x72, 0x5a, 0x6f, 0x6c, 0x77, 0x69, 0x61,
	0x52, 0x06, 0x7a, 0x6f, 0x6c, 0x77, 0x69, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x63,
	0x68, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x72, 0x61, 0x63, 0x7a, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x61, 0x63, 0x7a, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x0c, 0x7a, 0x61, 0x67, 0x72, 0x61,
	0x6e, 0x61, 0x4b, 0x61, 0x72, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x06, 0x2e,
	0x4b, 0x61, 0x72, 0x74, 0x61, 0x52, 0x0c, 0x7a, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x61, 0x4b, 0x61,
	0x72, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x0c, 0x6b, 0x6f, 0x6c, 0x6f, 0x72, 0x57, 0x79, 0x62, 0x72,
	0x61, 0x6e, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4b, 0x6f, 0x6c, 0x6f,
	0x72, 0x5a, 0x6f, 0x6c, 0x77, 0x69, 0x61, 0x52, 0x0c, 0x6b, 0x6f, 0x6c, 0x6f, 0x72, 0x57, 0x79,
	0x62, 0x72, 0x61, 0x6e, 0x79, 0x2a, 0x4c, 0x0a, 0x0b, 0x4b, 0x6f, 0x6c, 0x6f, 0x72, 0x5a, 0x6f,
	0x6c, 0x77, 0x69, 0x61, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x58, 0x58, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x59,
	0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x52, 0x50, 0x4c,
	0x45, 0x10, 0x05, 0x2a, 0xad, 0x01, 0x0a, 0x05, 0x4b, 0x61, 0x72, 0x74, 0x61, 0x12, 0x06, 0x0a,
	0x02, 0x58, 0x58, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a,
	0x02, 0x52, 0x32, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x31, 0x42, 0x10, 0x03, 0x12, 0x06,
	0x0a, 0x02, 0x47, 0x31, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x32, 0x10, 0x05, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x31, 0x42, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x31, 0x10, 0x07, 0x12,
	0x06, 0x0a, 0x02, 0x42, 0x32, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x31, 0x42, 0x10, 0x09,
	0x12, 0x06, 0x0a, 0x02, 0x59, 0x31, 0x10, 0x0a, 0x12, 0x06, 0x0a, 0x02, 0x59, 0x32, 0x10, 0x0b,
	0x12, 0x07, 0x0a, 0x03, 0x59, 0x31, 0x42, 0x10, 0x0c, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x31, 0x10,
	0x0d, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x32, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x31, 0x42,
	0x10, 0x0f, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x31, 0x10, 0x10, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x32,
	0x10, 0x11, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x31, 0x10, 0x12, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x31,
	0x42, 0x10, 0x13, 0x32, 0x7f, 0x0a, 0x03, 0x47, 0x72, 0x61, 0x12, 0x2c, 0x0a, 0x08, 0x4e, 0x6f,
	0x77, 0x79, 0x4d, 0x65, 0x63, 0x7a, 0x12, 0x10, 0x2e, 0x4b, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x63, 0x6a, 0x61, 0x47, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x4e, 0x6f, 0x77, 0x61, 0x47,
	0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0b, 0x44, 0x6f, 0x6c, 0x61,
	0x63, 0x7a, 0x44, 0x6f, 0x47, 0x72, 0x79, 0x12, 0x0b, 0x2e, 0x44, 0x6f, 0x6c, 0x61, 0x63, 0x7a,
	0x61, 0x6e, 0x69, 0x65, 0x1a, 0x08, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x47, 0x72, 0x79, 0x22, 0x00,
	0x12, 0x22, 0x0a, 0x07, 0x4d, 0x6f, 0x6a, 0x52, 0x75, 0x63, 0x68, 0x12, 0x0b, 0x2e, 0x52, 0x75,
	0x63, 0x68, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61, 0x1a, 0x08, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x47,
	0x72, 0x79, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x61, 0x72, 0x61, 0x7a, 0x2f, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x65,
	0x6a, 0x2f, 0x67, 0x72, 0x61, 0x5f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gra_proto_rawDescOnce sync.Once
	file_gra_proto_rawDescData = file_gra_proto_rawDesc
)

func file_gra_proto_rawDescGZIP() []byte {
	file_gra_proto_rawDescOnce.Do(func() {
		file_gra_proto_rawDescData = protoimpl.X.CompressGZIP(file_gra_proto_rawDescData)
	})
	return file_gra_proto_rawDescData
}

var file_gra_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gra_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gra_proto_goTypes = []interface{}{
	(KolorZolwia)(0),        // 0: KolorZolwia
	(Karta)(0),              // 1: Karta
	(*KonfiguracjaGry)(nil), // 2: KonfiguracjaGry
	(*NowaGraInfo)(nil),     // 3: NowaGraInfo
	(*Dolaczanie)(nil),      // 4: Dolaczanie
	(*StanGry)(nil),         // 5: StanGry
	(*Pole)(nil),            // 6: Pole
	(*RuchGracza)(nil),      // 7: RuchGracza
}
var file_gra_proto_depIdxs = []int32{
	0, // 0: StanGry.twojKolor:type_name -> KolorZolwia
	1, // 1: StanGry.twojeKarty:type_name -> Karta
	6, // 2: StanGry.plansza:type_name -> Pole
	0, // 3: Pole.zolwie:type_name -> KolorZolwia
	1, // 4: RuchGracza.zagranaKarta:type_name -> Karta
	0, // 5: RuchGracza.kolorWybrany:type_name -> KolorZolwia
	2, // 6: Gra.NowyMecz:input_type -> KonfiguracjaGry
	4, // 7: Gra.DolaczDoGry:input_type -> Dolaczanie
	7, // 8: Gra.MojRuch:input_type -> RuchGracza
	3, // 9: Gra.NowyMecz:output_type -> NowaGraInfo
	5, // 10: Gra.DolaczDoGry:output_type -> StanGry
	5, // 11: Gra.MojRuch:output_type -> StanGry
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gra_proto_init() }
func file_gra_proto_init() {
	if File_gra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KonfiguracjaGry); i {
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
		file_gra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowaGraInfo); i {
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
		file_gra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dolaczanie); i {
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
		file_gra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StanGry); i {
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
		file_gra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pole); i {
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
		file_gra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuchGracza); i {
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
			RawDescriptor: file_gra_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gra_proto_goTypes,
		DependencyIndexes: file_gra_proto_depIdxs,
		EnumInfos:         file_gra_proto_enumTypes,
		MessageInfos:      file_gra_proto_msgTypes,
	}.Build()
	File_gra_proto = out.File
	file_gra_proto_rawDesc = nil
	file_gra_proto_goTypes = nil
	file_gra_proto_depIdxs = nil
}
