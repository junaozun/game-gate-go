package codec

import (
	"sync"

	"github.com/golang/protobuf/proto"
)

type protoBuffer struct {
	proto.Buffer
	lastMarshaledSize uint32
}

var protoBufferPool = &sync.Pool{
	New: func() interface{} {
		return &protoBuffer{
			Buffer:            proto.Buffer{},
			lastMarshaledSize: 16,
		}
	},
}
