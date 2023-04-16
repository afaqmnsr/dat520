package gorumspaxos

import (
	"container/list"
	"dat520/lab3/leaderdetector"
	pb "dat520/lab5/gorumspaxos/proto"
	"log"
	"sync"
	"time"

	"github.com/relab/gorums"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Operation int

const (
	NoAction Operation = iota
	Prepare
	Accept
	Commit
)

// Proposer represents a proposer as defined by the Multi-Paxos algorithm.
// TODO(student) Add additional fields if necessary
// DO NOT remove the existing fields in the structure
type Proposer struct {
	sync.RWMutex
	id                 int                           // id: id of the replica
	crndID             int32                         // crndID: holds the current round of the replica, initialized to 0
	aduSlotID          uint32                        // aduSlotID: is the slot id for which the commit is completed.
	nextSlotID         uint32                        // nextSlotID: is used in processing the requests, initialized to 0
	config             *pb.Configuration             // config: configuration used for multipaxos, assigned in createConfiguration
	phaseOneDone       bool                          // phaseOneDone: indicates if the phase1 is done, initialized to false
	ld                 leaderdetector.LeaderDetector // ld: leader detector implementation module
	leader             int                           // leader: initial leader of the consensus
	nextOperation      Operation                     // nextOperation: holds the next operation to be performed
	phaseOneWaitTime   time.Duration                 // phaseOneWaitTime: duration until phase1 operation timeout
	phaseTwoWaitTime   time.Duration                 // phaseTwoWaitTime: duration until phase2 operation timeout
	acceptMsgQueue     *list.List                    // acceptMsgQueue: queue to hold the pending accept requests as part of prepare operation
	learnMsg           *pb.LearnMsg                  // learnMsg: next learnMsg to be processed. (may become a queue if alpha increased)
	qspec              pb.QuorumSpec                 // qspec: QuorumSpec object to be used in the creation of configuration
	nodeMap            map[string]uint32             // nodeMap: is the map of the address to the id
	clientRequestQueue *list.List                    // clientRequestQueue: queue used to store the pending client requests
	statusChannel      chan bool                     // statusChannel: used to pass the status of the operation to move to next operation
	stop               chan struct{}                 // stop: channel used to end the replica functionality
	manager            *pb.Manager                   // manager holds the gorums manager to create. modify and close the configuration
}

// NewProposer returns a new Multi-Paxos proposer.
func NewProposer(args NewProposerArgs) *Proposer {
	mgr := pb.NewManager(gorums.WithDialTimeout(args.phaseOneWaitTime),
		gorums.WithGrpcDialOptions(
			grpc.WithBlock(), // block until connections are made
			grpc.WithTransportCredentials(insecure.NewCredentials()), // disable TLS
		),
	)
	return &Proposer{
		id:                 args.id,
		aduSlotID:          args.aduSlotID,
		phaseOneWaitTime:   args.phaseOneWaitTime,
		phaseTwoWaitTime:   args.phaseTwoWaitTime,
		ld:                 args.leaderDetector,
		leader:             args.leaderDetector.Leader(),
		qspec:              args.qspec,
		nodeMap:            args.nodeMap,
		stop:               make(chan struct{}),
		crndID:             0,
		nextSlotID:         0,
		nextOperation:      NoAction,
		acceptMsgQueue:     list.New(),
		clientRequestQueue: list.New(),
		statusChannel:      make(chan bool, 1),
		manager:            mgr,
	}
}

// Internal: checks if the current replica is the leader
func (p *Proposer) isLeader() bool {
	p.RLock()
	defer p.RUnlock()
	return p.leader == p.id
}

// runPhaseOne runs the phase one (prepare->promise) of the protocol.
// This function should be called only if the replica is the leader and
// phase1 has not already completed.
// 1. Create the configuration with all replicas, call the createConfiguration method
// 2. Create a PrepareMsg with the current round ID and incremented aduSlotID
// 3. Send the prepare message to all the replicas (quorum call)
// 4. If succeeded set phaseOneDone to true and nextOperation to "accept"
// 5. If the response Promise Message contains the slots, prepare an accept
// message for each of the slot and add it to the accept queue
func (p *Proposer) runPhaseOne() error {
	// TODO(student): complete the runPhaseOne
	return nil
}

// Runs Phase 1, if not already completed.
// Performs Accept after phase1 successfully completes.
// Performs Commit after accept successfully completes.
func (p *Proposer) runMultiPaxos() {
	if p.isLeader() {
		switch {
		case !p.isPhaseOneDone():
			p.increaseCrnd()
			p.sendStatus(p.runPhaseOne(), Accept)
		case p.nextOperation == Accept:
			p.sendStatus(p.performAccept(), Commit)
		case p.nextOperation == Commit:
			p.sendStatus(p.performCommit(), Accept)
		}
	}
}

// sendStatus: internal function to mark the next operation and send status on the
// statusChannel channel
func (p *Proposer) sendStatus(err error, nextOperation Operation) {
	if err != nil {
		log.Printf("Node id %d\t operation %d failed %v", p.id, p.nextOperation, err)
		p.statusChannel <- true
	} else {
		log.Printf("Node id %d\t operation %d succeeded next operation %d", p.id, p.nextOperation, nextOperation)
		p.Lock()
		p.nextOperation = nextOperation
		p.Unlock()
		p.statusChannel <- false
	}
}

// Start starts proposer's main run loop as a separate goroutine.
// The separate goroutine should start an infinite loop and
// use a select mechanism to wait for the following events
// 1. on the status channel to conduct the paxos phases and operations
// 2. on the channel returned by leader detector subscribe method
// 3. on stop channel to stop the relpica
// default case is to process any pending client requests by adding to the
// clientRequestQueue and call runMultiPaxos
// If no requests are available sleep for RequestWaitTime
// If a signal is received on the status channel,
// move to the next phase by calling runMultiPaxos
// If the replica is the new leader, then reset the phaseOneDone,
// acceptReqQueue, clientRequestQueue, nextOperation and call runMultiPaxos
func (p *Proposer) Start() {
	go func() {
		// TODO(student): Complete the function
	}()
}

// Stop stops the proposer's run loop.
func (p *Proposer) Stop() {
	p.stop <- struct{}{}
}

// IncrementAllDecidedUpTo increments the Proposer's adu variable by one.
func (p *Proposer) IncrementAllDecidedUpTo() {
	p.Lock()
	defer p.Unlock()
	p.aduSlotID++
}

// increaseCrnd increases the proposer's current round (crnd field)
// with the size of the Paxos configuration.
func (p *Proposer) increaseCrnd() {
	p.Lock()
	defer p.Unlock()
	p.crndID = p.crndID + int32(len(p.nodeMap))
}

// Perform accept on the servers and return error if required
// 1. Check if any pending accept requests in the acceptReqQueue to process
// 2. Check if any pending client requests in the clientRequestQueue to process
// 3. Increment the nextSlotID and prepare an accept message with the pending request,
//    crndID and nextSlotID.
// 4. Perform accept quorum call on the configuration and set the learnMsg with the
//    return value of the quorum call.
func (p *Proposer) performAccept() error {
	//TODO(student) Complete the performAccept function
	return nil
}

// performCommit performs the commit operation of the protocol.
// Call commit quorum call with the learnMsg returned in the accept call.
func (p *Proposer) performCommit() error {
	// TODO(student): Complete the performCommit function
	return nil
}

// returns phaseOneDone
func (p *Proposer) isPhaseOneDone() bool {
	p.RLock()
	defer p.RUnlock()
	return p.phaseOneDone
}

// createConfiguration creates a configuration with the addresses and quorum spec.
func (p *Proposer) createConfiguration() error {
	// TODO(student): Complete the createConfiguration function
	return nil
}

// ProcessRequest: processes the request from the client by putting it into the
// clientRequestQueue and calling the getResponse to get the response matching
// the request.
func (p *Proposer) AddRequestToQ(request *pb.Value) {
	if p.isLeader() {
		log.Printf("Node id %d\t is adding the request to Queue %v", p.id, request)
		p.Lock()
		p.clientRequestQueue.PushBack(request)
		p.Unlock()
	}
}
