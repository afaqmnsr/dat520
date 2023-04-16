package gorumspaxos

import (
	pb "dat520/lab5/gorumspaxos/proto"
)

// Acceptor represents an acceptor as defined by the Multi-Paxos algorithm.
// For current implementation alpha is assumed to be 1. In future if the alpha is
// increased, then there would be significant changes to the current implementation.
type Acceptor struct {
	rnd           *pb.Round                  // rnd: is the current round in which the acceptor is participating
	slots         map[uint32]*pb.PromiseSlot // slots: is the internal data structure maintained by the acceptor to remember the slots
	maxSeenSlotId uint32                     // maxSeenSlotId: is the highest slot for which the prepare is received
}

// NewAcceptor returns a new Multi-Paxos acceptor
func NewAcceptor() *Acceptor {
	return &Acceptor{
		rnd:   NoRound(),
		slots: make(map[uint32]*pb.PromiseSlot),
	}
}

// handlePrepare takes a prepare message and returns a promise message according to
// the Multi-Paxos algorithm. If the prepare message is invalid, nil should be returned.
func (a *Acceptor) handlePrepare(prp *pb.PrepareMsg) (prm *pb.PromiseMsg) {
	// TODO(student): Complete the handlePrepare
	return &pb.PromiseMsg{}
}

// handleAccept takes an accept message and returns a learn message according to
// the Multi-Paxos algorithm. If the accept message is invalid, nil should be returned.
func (a *Acceptor) handleAccept(acc *pb.AcceptMsg) (lrn *pb.LearnMsg) {
	// TODO(student): Complete the handleAccept
	return &pb.LearnMsg{}
}
