// Code generated by protoc-gen-gogo.
// source: hsm.proto
// DO NOT EDIT!

package caffe2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Each node in the hierarchy contains links to either leaf nodes or more
// non-terminal nodes
type NodeProto struct {
	// Links to non-terminal children nodes
	Children []*NodeProto `protobuf:"bytes,1,rep,name=children" json:"children,omitempty"`
	// Links to terminal (leaf) nodes
	WordIds []int32   `protobuf:"varint,2,rep,name=word_ids,json=wordIds" json:"word_ids,omitempty"`
	Offset  int32     `protobuf:"varint,3,opt,name=offset" json:"offset"`
	Name    string    `protobuf:"bytes,4,opt,name=name" json:"name"`
	Scores  []float32 `protobuf:"fixed32,5,rep,name=scores" json:"scores,omitempty"`
}

func (m *NodeProto) Reset()                    { *m = NodeProto{} }
func (m *NodeProto) String() string            { return proto.CompactTextString(m) }
func (*NodeProto) ProtoMessage()               {}
func (*NodeProto) Descriptor() ([]byte, []int) { return fileDescriptorHsm, []int{0} }

func (m *NodeProto) GetChildren() []*NodeProto {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *NodeProto) GetWordIds() []int32 {
	if m != nil {
		return m.WordIds
	}
	return nil
}

func (m *NodeProto) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *NodeProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeProto) GetScores() []float32 {
	if m != nil {
		return m.Scores
	}
	return nil
}

// Protobuf format to accept hierarchy for hierarchical softmax operator.
// TreeProto points to the root node.
type TreeProto struct {
	RootNode *NodeProto `protobuf:"bytes,1,opt,name=root_node,json=rootNode" json:"root_node,omitempty"`
}

func (m *TreeProto) Reset()                    { *m = TreeProto{} }
func (m *TreeProto) String() string            { return proto.CompactTextString(m) }
func (*TreeProto) ProtoMessage()               {}
func (*TreeProto) Descriptor() ([]byte, []int) { return fileDescriptorHsm, []int{1} }

func (m *TreeProto) GetRootNode() *NodeProto {
	if m != nil {
		return m.RootNode
	}
	return nil
}

// Internal Protobuf format which represents the path in the tree hierarchy for
// each word in the vocabulary.
type HierarchyProto struct {
	Size_ int32        `protobuf:"varint,1,opt,name=size" json:"size"`
	Paths []*PathProto `protobuf:"bytes,2,rep,name=paths" json:"paths,omitempty"`
}

func (m *HierarchyProto) Reset()                    { *m = HierarchyProto{} }
func (m *HierarchyProto) String() string            { return proto.CompactTextString(m) }
func (*HierarchyProto) ProtoMessage()               {}
func (*HierarchyProto) Descriptor() ([]byte, []int) { return fileDescriptorHsm, []int{2} }

func (m *HierarchyProto) GetSize_() int32 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *HierarchyProto) GetPaths() []*PathProto {
	if m != nil {
		return m.Paths
	}
	return nil
}

// Each PathProto belongs to a word and is an array of nodes in the
// path from the root to the leaf (which is the word itself) in the tree.
type PathProto struct {
	WordId    int32            `protobuf:"varint,1,opt,name=word_id,json=wordId" json:"word_id"`
	PathNodes []*PathNodeProto `protobuf:"bytes,2,rep,name=path_nodes,json=pathNodes" json:"path_nodes,omitempty"`
}

func (m *PathProto) Reset()                    { *m = PathProto{} }
func (m *PathProto) String() string            { return proto.CompactTextString(m) }
func (*PathProto) ProtoMessage()               {}
func (*PathProto) Descriptor() ([]byte, []int) { return fileDescriptorHsm, []int{3} }

func (m *PathProto) GetWordId() int32 {
	if m != nil {
		return m.WordId
	}
	return 0
}

func (m *PathProto) GetPathNodes() []*PathNodeProto {
	if m != nil {
		return m.PathNodes
	}
	return nil
}

// Represents a node in the path from the root node all the way down to the
// word (leaf).
type PathNodeProto struct {
	// Parameter matrix offset for this node
	Index int32 `protobuf:"varint,1,opt,name=index" json:"index"`
	// Number of children
	Length int32 `protobuf:"varint,2,opt,name=length" json:"length"`
	// Index of the next node in the path
	Target int32 `protobuf:"varint,3,opt,name=target" json:"target"`
}

func (m *PathNodeProto) Reset()                    { *m = PathNodeProto{} }
func (m *PathNodeProto) String() string            { return proto.CompactTextString(m) }
func (*PathNodeProto) ProtoMessage()               {}
func (*PathNodeProto) Descriptor() ([]byte, []int) { return fileDescriptorHsm, []int{4} }

func (m *PathNodeProto) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PathNodeProto) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *PathNodeProto) GetTarget() int32 {
	if m != nil {
		return m.Target
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeProto)(nil), "caffe2.NodeProto")
	proto.RegisterType((*TreeProto)(nil), "caffe2.TreeProto")
	proto.RegisterType((*HierarchyProto)(nil), "caffe2.HierarchyProto")
	proto.RegisterType((*PathProto)(nil), "caffe2.PathProto")
	proto.RegisterType((*PathNodeProto)(nil), "caffe2.PathNodeProto")
}
func (m *NodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, msg := range m.Children {
			dAtA[i] = 0xa
			i++
			i = encodeVarintHsm(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.WordIds) > 0 {
		for _, num := range m.WordIds {
			dAtA[i] = 0x10
			i++
			i = encodeVarintHsm(dAtA, i, uint64(num))
		}
	}
	dAtA[i] = 0x18
	i++
	i = encodeVarintHsm(dAtA, i, uint64(m.Offset))
	dAtA[i] = 0x22
	i++
	i = encodeVarintHsm(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	if len(m.Scores) > 0 {
		for _, num := range m.Scores {
			dAtA[i] = 0x2d
			i++
			f1 := math.Float32bits(float32(num))
			dAtA[i] = uint8(f1)
			i++
			dAtA[i] = uint8(f1 >> 8)
			i++
			dAtA[i] = uint8(f1 >> 16)
			i++
			dAtA[i] = uint8(f1 >> 24)
			i++
		}
	}
	return i, nil
}

func (m *TreeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TreeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RootNode != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHsm(dAtA, i, uint64(m.RootNode.Size()))
		n2, err := m.RootNode.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *HierarchyProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HierarchyProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintHsm(dAtA, i, uint64(m.Size_))
	if len(m.Paths) > 0 {
		for _, msg := range m.Paths {
			dAtA[i] = 0x12
			i++
			i = encodeVarintHsm(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PathProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PathProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintHsm(dAtA, i, uint64(m.WordId))
	if len(m.PathNodes) > 0 {
		for _, msg := range m.PathNodes {
			dAtA[i] = 0x12
			i++
			i = encodeVarintHsm(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PathNodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PathNodeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintHsm(dAtA, i, uint64(m.Index))
	dAtA[i] = 0x10
	i++
	i = encodeVarintHsm(dAtA, i, uint64(m.Length))
	dAtA[i] = 0x18
	i++
	i = encodeVarintHsm(dAtA, i, uint64(m.Target))
	return i, nil
}

func encodeFixed64Hsm(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Hsm(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHsm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NodeProto) Size() (n int) {
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovHsm(uint64(l))
		}
	}
	if len(m.WordIds) > 0 {
		for _, e := range m.WordIds {
			n += 1 + sovHsm(uint64(e))
		}
	}
	n += 1 + sovHsm(uint64(m.Offset))
	l = len(m.Name)
	n += 1 + l + sovHsm(uint64(l))
	if len(m.Scores) > 0 {
		n += 5 * len(m.Scores)
	}
	return n
}

func (m *TreeProto) Size() (n int) {
	var l int
	_ = l
	if m.RootNode != nil {
		l = m.RootNode.Size()
		n += 1 + l + sovHsm(uint64(l))
	}
	return n
}

func (m *HierarchyProto) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHsm(uint64(m.Size_))
	if len(m.Paths) > 0 {
		for _, e := range m.Paths {
			l = e.Size()
			n += 1 + l + sovHsm(uint64(l))
		}
	}
	return n
}

func (m *PathProto) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHsm(uint64(m.WordId))
	if len(m.PathNodes) > 0 {
		for _, e := range m.PathNodes {
			l = e.Size()
			n += 1 + l + sovHsm(uint64(l))
		}
	}
	return n
}

func (m *PathNodeProto) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHsm(uint64(m.Index))
	n += 1 + sovHsm(uint64(m.Length))
	n += 1 + sovHsm(uint64(m.Target))
	return n
}

func sovHsm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHsm(x uint64) (n int) {
	return sovHsm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHsm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHsm
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &NodeProto{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHsm
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.WordIds = append(m.WordIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHsm
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthHsm
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHsm
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.WordIds = append(m.WordIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field WordIds", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHsm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType == 5 {
				var v uint32
				if (iNdEx + 4) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 4
				v = uint32(dAtA[iNdEx-4])
				v |= uint32(dAtA[iNdEx-3]) << 8
				v |= uint32(dAtA[iNdEx-2]) << 16
				v |= uint32(dAtA[iNdEx-1]) << 24
				v2 := float32(math.Float32frombits(v))
				m.Scores = append(m.Scores, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHsm
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthHsm
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					if (iNdEx + 4) > l {
						return io.ErrUnexpectedEOF
					}
					iNdEx += 4
					v = uint32(dAtA[iNdEx-4])
					v |= uint32(dAtA[iNdEx-3]) << 8
					v |= uint32(dAtA[iNdEx-2]) << 16
					v |= uint32(dAtA[iNdEx-1]) << 24
					v2 := float32(math.Float32frombits(v))
					m.Scores = append(m.Scores, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Scores", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHsm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHsm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TreeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHsm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TreeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TreeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootNode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHsm
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RootNode == nil {
				m.RootNode = &NodeProto{}
			}
			if err := m.RootNode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHsm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHsm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HierarchyProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHsm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HierarchyProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HierarchyProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHsm
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, &PathProto{})
			if err := m.Paths[len(m.Paths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHsm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHsm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PathProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHsm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PathProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WordId", wireType)
			}
			m.WordId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WordId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PathNodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHsm
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PathNodes = append(m.PathNodes, &PathNodeProto{})
			if err := m.PathNodes[len(m.PathNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHsm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHsm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PathNodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHsm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PathNodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathNodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			m.Target = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Target |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHsm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHsm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHsm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHsm
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHsm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHsm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHsm
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHsm(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHsm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHsm   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("hsm.proto", fileDescriptorHsm) }

var fileDescriptorHsm = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xeb, 0xe6, 0xa7, 0xf5, 0x45, 0x20, 0xb0, 0x04, 0x32, 0x08, 0x42, 0x94, 0x85, 0x2c,
	0x64, 0xa8, 0xd8, 0xd8, 0x3a, 0xc1, 0x82, 0xaa, 0xc0, 0x5e, 0xa2, 0xd8, 0x69, 0x22, 0xb5, 0x71,
	0x65, 0x5b, 0xe2, 0xe7, 0x29, 0x78, 0x09, 0xde, 0xa5, 0x23, 0x4f, 0x80, 0x50, 0x79, 0x11, 0x94,
	0xd8, 0x2d, 0x2d, 0x12, 0xe3, 0xfd, 0xce, 0xb9, 0x27, 0x27, 0xd7, 0x80, 0x4b, 0x35, 0x4b, 0xe6,
	0x52, 0x68, 0x41, 0xfc, 0x3c, 0x2b, 0x0a, 0x3e, 0x88, 0xde, 0x11, 0xe0, 0x3b, 0xc1, 0xf8, 0xa8,
	0xa5, 0x97, 0xd0, 0xcf, 0xcb, 0x6a, 0xca, 0x24, 0xaf, 0x29, 0x0a, 0x9d, 0x78, 0x67, 0x70, 0x90,
	0x18, 0x63, 0xb2, 0x36, 0xa5, 0x6b, 0x0b, 0x39, 0x86, 0xfe, 0x93, 0x90, 0x6c, 0x5c, 0x31, 0x45,
	0xbb, 0xa1, 0x13, 0x7b, 0x69, 0xaf, 0x99, 0x6f, 0x99, 0x22, 0xa7, 0xe0, 0x8b, 0xa2, 0x50, 0x5c,
	0x53, 0x27, 0x44, 0xb1, 0x37, 0x74, 0x17, 0x9f, 0xe7, 0x9d, 0xd4, 0x32, 0x42, 0xc1, 0xad, 0xb3,
	0x19, 0xa7, 0x6e, 0x88, 0x62, 0x6c, 0xb5, 0x96, 0x90, 0x23, 0xf0, 0x55, 0x2e, 0x24, 0x57, 0xd4,
	0x0b, 0x9d, 0xb8, 0x9b, 0xda, 0x29, 0xba, 0x06, 0xfc, 0x20, 0xb9, 0xad, 0x99, 0x00, 0x96, 0x42,
	0xe8, 0x71, 0x2d, 0x18, 0xa7, 0x28, 0x44, 0xff, 0xf4, 0x6c, 0x3c, 0xcd, 0x18, 0xdd, 0xc3, 0xde,
	0x4d, 0xc5, 0x65, 0x26, 0xf3, 0xf2, 0xc5, 0x24, 0x50, 0x70, 0x55, 0xf5, 0x6a, 0x96, 0x57, 0xe5,
	0x5a, 0x42, 0x2e, 0xc0, 0x9b, 0x67, 0xba, 0x34, 0x3f, 0xb4, 0x91, 0x3b, 0xca, 0x74, 0x69, 0x72,
	0x8d, 0x1e, 0x3d, 0x02, 0x5e, 0x33, 0x72, 0x06, 0x3d, 0x7b, 0x89, 0xad, 0x48, 0xdf, 0x9c, 0x83,
	0x5c, 0x01, 0x34, 0x4b, 0x6d, 0xe1, 0x55, 0xf2, 0xe1, 0x66, 0xf2, 0x6f, 0x6b, 0x3c, 0xb7, 0xa3,
	0x8a, 0x26, 0xb0, 0xbb, 0xa5, 0x91, 0x13, 0xf0, 0xaa, 0x9a, 0xf1, 0xe7, 0xad, 0x6f, 0x18, 0xd4,
	0x1c, 0x7c, 0xca, 0xeb, 0x89, 0x2e, 0x69, 0x77, 0xb3, 0x80, 0x61, 0x8d, 0xaa, 0x33, 0x39, 0xf9,
	0xfb, 0x1c, 0x86, 0x0d, 0xf7, 0x17, 0xcb, 0x00, 0x7d, 0x2c, 0x03, 0xf4, 0xb5, 0x0c, 0xd0, 0xdb,
	0x77, 0xd0, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x86, 0xcd, 0xc3, 0x80, 0x2a, 0x02, 0x00, 0x00,
}
