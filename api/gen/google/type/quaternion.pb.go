// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: google/type/quaternion.proto

package _type

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

// A quaternion is defined as the quotient of two directed lines in a
// three-dimensional space or equivalently as the quotient of two Euclidean
// vectors (https://en.wikipedia.org/wiki/Quaternion).
//
// Quaternions are often used in calculations involving three-dimensional
// rotations (https://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation),
// as they provide greater mathematical robustness by avoiding the gimbal lock
// problems that can be encountered when using Euler angles
// (https://en.wikipedia.org/wiki/Gimbal_lock).
//
// Quaternions are generally represented in this form:
//
//	w + xi + yj + zk
//
// where x, y, z, and w are real numbers, and i, j, and k are three imaginary
// numbers.
//
// Our naming choice `(x, y, z, w)` comes from the desire to avoid confusion for
// those interested in the geometric properties of the quaternion in the 3D
// Cartesian space. Other texts often use alternative names or subscripts, such
// as `(a, b, c, d)`, `(1, i, j, k)`, or `(0, 1, 2, 3)`, which are perhaps
// better suited for mathematical interpretations.
//
// To avoid any confusion, as well as to maintain compatibility with a large
// number of software libraries, the quaternions represented using the protocol
// buffer below *must* follow the Hamilton convention, which defines `ij = k`
// (i.e. a right-handed algebra), and therefore:
//
//	i^2 = j^2 = k^2 = ijk = −1
//	ij = −ji = k
//	jk = −kj = i
//	ki = −ik = j
//
// Please DO NOT use this to represent quaternions that follow the JPL
// convention, or any of the other quaternion flavors out there.
//
// Definitions:
//
//   - Quaternion norm (or magnitude): `sqrt(x^2 + y^2 + z^2 + w^2)`.
//   - Unit (or normalized) quaternion: a quaternion whose norm is 1.
//   - Pure quaternion: a quaternion whose scalar component (`w`) is 0.
//   - Rotation quaternion: a unit quaternion used to represent rotation.
//   - Orientation quaternion: a unit quaternion used to represent orientation.
//
// A quaternion can be normalized by dividing it by its norm. The resulting
// quaternion maintains the same direction, but has a norm of 1, i.e. it moves
// on the unit sphere. This is generally necessary for rotation and orientation
// quaternions, to avoid rounding errors:
// https://en.wikipedia.org/wiki/Rotation_formalisms_in_three_dimensions
//
// Note that `(x, y, z, w)` and `(-x, -y, -z, -w)` represent the same rotation,
// but normalization would be even more useful, e.g. for comparison purposes, if
// it would produce a unique representation. It is thus recommended that `w` be
// kept positive, which can be achieved by changing all the signs when `w` is
// negative.
type Quaternion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The x component.
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	// The y component.
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	// The z component.
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	// The scalar component.
	W float64 `protobuf:"fixed64,4,opt,name=w,proto3" json:"w,omitempty"`
}

func (x *Quaternion) Reset() {
	*x = Quaternion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_type_quaternion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quaternion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quaternion) ProtoMessage() {}

func (x *Quaternion) ProtoReflect() protoreflect.Message {
	mi := &file_google_type_quaternion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quaternion.ProtoReflect.Descriptor instead.
func (*Quaternion) Descriptor() ([]byte, []int) {
	return file_google_type_quaternion_proto_rawDescGZIP(), []int{0}
}

func (x *Quaternion) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Quaternion) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Quaternion) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Quaternion) GetW() float64 {
	if x != nil {
		return x.W
	}
	return 0
}

var File_google_type_quaternion_proto protoreflect.FileDescriptor

var file_google_type_quaternion_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x71, 0x75,
	0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x0a, 0x51,
	0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x7a, 0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x77, 0x42, 0x97, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0f, 0x51, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2d, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x47, 0x54, 0x58, 0xaa, 0x02, 0x0b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0xca, 0x02, 0x0b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x54, 0x79,
	0x70, 0x65, 0xe2, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x54, 0x79, 0x70, 0x65,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_type_quaternion_proto_rawDescOnce sync.Once
	file_google_type_quaternion_proto_rawDescData = file_google_type_quaternion_proto_rawDesc
)

func file_google_type_quaternion_proto_rawDescGZIP() []byte {
	file_google_type_quaternion_proto_rawDescOnce.Do(func() {
		file_google_type_quaternion_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_type_quaternion_proto_rawDescData)
	})
	return file_google_type_quaternion_proto_rawDescData
}

var file_google_type_quaternion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_type_quaternion_proto_goTypes = []interface{}{
	(*Quaternion)(nil), // 0: google.type.Quaternion
}
var file_google_type_quaternion_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_type_quaternion_proto_init() }
func file_google_type_quaternion_proto_init() {
	if File_google_type_quaternion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_type_quaternion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quaternion); i {
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
			RawDescriptor: file_google_type_quaternion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_type_quaternion_proto_goTypes,
		DependencyIndexes: file_google_type_quaternion_proto_depIdxs,
		MessageInfos:      file_google_type_quaternion_proto_msgTypes,
	}.Build()
	File_google_type_quaternion_proto = out.File
	file_google_type_quaternion_proto_rawDesc = nil
	file_google_type_quaternion_proto_goTypes = nil
	file_google_type_quaternion_proto_depIdxs = nil
}
