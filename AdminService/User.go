// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package AdminService

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type User struct {
	_tab flatbuffers.Table
}

func GetRootAsUser(buf []byte, offset flatbuffers.UOffsetT) *User {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &User{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUser(buf []byte, offset flatbuffers.UOffsetT) *User {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &User{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *User) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *User) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *User) Id() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *User) MutateId(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *User) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *User) Email() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func UserStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func UserAddId(builder *flatbuffers.Builder, id int32) {
	builder.PrependInt32Slot(0, id, 0)
}
func UserAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func UserAddEmail(builder *flatbuffers.Builder, email flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(email), 0)
}
func UserEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
