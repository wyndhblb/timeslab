package timeslab

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import "github.com/tinylib/msgp/msgp"

// DecodeMsg implements msgp.Decodable
func (z *Resolution) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zxvk int32
		zxvk, err = dc.ReadInt32()
		(*z) = Resolution(zxvk)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Resolution) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Resolution) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Resolution) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zbzg int32
		zbzg, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Resolution(zbzg)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Resolution) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}
