package reply

type PongRelply struct {
}

var pongBytes = []byte("+PONG\r\n")

func (r PongRelply) ToBytes() []byte {
	return pongBytes
}

func MakePongReply() *PongRelply {
	return &PongRelply{}
}

type OkReply struct{}

var okBytes = []byte("+OK\r\n")

func (r *OkReply) ToBytes() []byte {
	return okBytes
}

var theOkReply = new(OkReply)

func MakeOkReply() *OkReply {
	return theOkReply
}

type NullBulkReply struct{}

var nullBulkBytes = []byte("$-1\r\n")

func (n NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

type EmptyMultiBulkReply struct{}

var emptyMultiBulkBytes = []byte("*0\r\n")

func (r *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

type NoReply struct{}

var noBytes = []byte("")

func (n NoReply) ToBytes() []byte {
	return noBytes
}
