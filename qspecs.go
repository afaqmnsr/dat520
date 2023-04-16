package gorumspaxos

import (
	pb "dat520/lab5/gorumspaxos/proto"
)

// PaxosQSpec is a quorum specification object for Paxos.
// It only holds the quorum size.
// TODO(student): If necessary add additional fields
// DO NOT remove the existing fields in the structure
type PaxosQSpec struct {
	qSize int
}

// NewPaxosQSpec returns a quorum specification object for Paxos
// for the given quorum size.
func NewPaxosQSpec(quorumSize int) pb.QuorumSpec {
	return &PaxosQSpec{
		qSize: quorumSize,
	}
}

// PrepareQF is the quorum function to process the replies of the Prepare RPC call.
// Proposer handle PromiseMsgs returned by the Acceptors, and any promiseslots added
// to the PromiseMsg should be in the increasing order.
func (qs PaxosQSpec) PrepareQF(prepare *pb.PrepareMsg,
	replies map[uint32]*pb.PromiseMsg) (*pb.PromiseMsg, bool) {
	// TODO(student) complete the PrepareQF
	return nil, false
}

// AcceptQF is the quorum function for the Accept quorum RPC call
// This is where the Proposer handle LearnMsgs to determine if a
// value has been decided by the Acceptors.
// The quorum function returns true if a value has been decided,
// and the corresponding LearnMsg holds the round number and value
// that was decided. If false is returned, no value was decided.
func (qs PaxosQSpec) AcceptQF(accMsg *pb.AcceptMsg, replies map[uint32]*pb.LearnMsg) (*pb.LearnMsg, bool) {
	// TODO (student) complete the AcceptQF
	return nil, false
}

// CommitQF is the quorum function for the Commit quorum RPC call.
// This function just waits for a quorum of empty replies,
// indicating that at least a quorum of Learners have committed
// the value decided by the Acceptors.
func (qs PaxosQSpec) CommitQF(_ *pb.LearnMsg, replies map[uint32]*pb.Empty) (*pb.Empty, bool) {
	// TODO (student) complete the CommitQF
	return nil, false
}

// ClientHandleQF is the quorum function  for the ClientHandle quorum RPC call.
// This functions waits for replies from the majority of replicas. Received replies
// should be validated before returning the response
func (qs PaxosQSpec) ClientHandleQF(in *pb.Value, replies map[uint32]*pb.Response) (*pb.Response, bool) {
	// TODO (student) complete the ClientHandleQF
	return nil, false
}
