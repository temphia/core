package fencer

// pool is pool of same type of policy with is kept at
// some const amount all the time
type pool struct {
	pool chan policy
}

func NewPool(max int) *pool {
	return &pool{
		pool: make(chan policy, max),
	}
}

func (p *pool) Borrow() policy {
	var c policy
	select {
	case c = <-p.pool:
	default:

	}
	return c
}

func (p *pool) Return(c policy) {
	select {
	case p.pool <- c:
	default:
		// let it go, let it go...
	}
}
