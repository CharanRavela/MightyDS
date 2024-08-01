package stack

type StackStats struct {
	// TODO: Need to add stats for other operations of this package

	tOps  int // total operations performed pop + push + peek + clear includes failed attempts
	tPops int // total pop operations performed includes failed attempts

	nPops   int // no of pop operations performed only positive attempts
	nPushes int // no of push operations performed only postive attempts

	// there are no failed attempts in push operations so excluded
	nFailedPops int // no of failed pop operations

	currStackSize    int // current size of the stack
	avgStackSize     int // average size of the stack till now
	highestStackSize int // max size stack has growed to till now

	nEmptyStack int // no of times stacks became empty positively

	// Perfromance
	maxPushTime int // max time to push an element
	maxPopTime  int // max time to pop an element

	minPushTime int // min time to push an element
	minPopTime  int // min time to pop an element

	avgPushTime int // average time to push an element
	avgPopTime  int // average time to pop an element
}

func (ss *StackStats) NewStackStats() *StackStats {
	return &StackStats{
		tOps:  0,
		tPops: 0,

		nPops:   0,
		nPushes: 0,

		nFailedPops: 0,

		currStackSize:    0,
		avgStackSize:     0,
		highestStackSize: 0,

		nEmptyStack: 0,

		maxPushTime: 0,
		maxPopTime:  0,

		minPushTime: 0,
		minPopTime:  0,

		avgPushTime: 0,
		avgPopTime:  0,
	}
}

func (ss *StackStats) UpdateStats() {
	ss.tOps = +1
}
