package mycode

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type Iterator struct{}

func (i *Iterator) hasNext() bool { return false }

func (i *Iterator) next() int { return 0 }

type PeekingIterator struct {
	it           *Iterator
	nextCache    int
	hasNextCache bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	pi := &PeekingIterator{
		it:           iter,
		hasNextCache: iter.hasNext(),
	}

	if pi.hasNextCache {
		pi.nextCache = iter.next()
	}

	return pi
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasNextCache
}

func (this *PeekingIterator) next() int {
	rst := this.nextCache

	if this.hasNextCache = this.it.hasNext(); this.hasNextCache {
		this.nextCache = this.it.next()
	}

	return rst
}

func (this *PeekingIterator) peek() int {
	return this.nextCache
}
