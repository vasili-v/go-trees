package iptree32

// !!!DON'T EDIT!!! Generated by infobloxopen/go-trees/etc from <name>tree{{.bits}} with etc -s uint32 -d uintX.yaml -t ./<name>tree\{\{.bits\}\}

import (
	"fmt"
	"math/bits"
)

// key32BitSize is an alias for bitsize of 32-bit radix tree's key.
const key32BitSize = 32

var (
	masks32  = []uint32{
		0x00000000, 0x80000000, 0xc0000000, 0xe0000000,
		0xf0000000, 0xf8000000, 0xfc000000, 0xfe000000,
		0xff000000, 0xff800000, 0xffc00000, 0xffe00000,
		0xfff00000, 0xfff80000, 0xfffc0000, 0xfffe0000,
		0xffff0000, 0xffff8000, 0xffffc000, 0xffffe000,
		0xfffff000, 0xfffff800, 0xfffffc00, 0xfffffe00,
		0xffffff00, 0xffffff80, 0xffffffc0, 0xffffffe0,
		0xfffffff0, 0xfffffff8, 0xfffffffc, 0xfffffffe,
		0xffffffff}
)

// node32 is an element of radix tree with 32-bit unsigned integer as a key.
type node32 struct {
	// key stores key for current node.
	key uint32
	// bits is a number of significant bits in key.
	bits uint8
	// leaf indicates if the node is leaf node and contains any data in value.
	leaf bool
	// value contains data associated with key.
	value uint32

	chld [2]*node32
}

// Dot dumps tree to Graphviz .dot format
func (n *node32) Dot() string {
	body := ""

	// Iterate all nodes using breadth-first search algorithm.
	i := 0
	queue := []*node32{n}
	for len(queue) > 0 {
		c := queue[0]
		body += fmt.Sprintf("N%d %s\n", i, c.dotString())
		if c != nil && (c.chld[0] != nil || c.chld[1] != nil) {
			// Children for current node if any always go to the end of the queue
			// so we can know their indices using current queue length.
			body += fmt.Sprintf("N%d -> { N%d N%d }\n", i, i+len(queue), i+len(queue)+1)
			queue = append(append(queue, c.chld[0]), c.chld[1])
		}

		queue = queue[1:]
		i++
	}

	return "digraph d {\n" + body + "}\n"
}

// Insert puts new leaf to radix tree and returns pointer to new root. The method uses copy on write strategy so old root doesn't see the change.
func (n *node32) Insert(key uint32, bits int, value uint32) *node32 {
	// Adjust bits.
	if bits < 0 {
		bits = 0
	} else if bits > key32BitSize {
		bits = key32BitSize
	}

	return n.insert(newNode32(key, uint8(bits), true, value))
}

// InplaceInsert puts new leaf to radix tree (or replaces value in existing one). The method inserts data directly to current tree so make sure you have exclusive access to it.
func (n *node32) InplaceInsert(key uint32, bits int, value uint32) *node32 {
	// Adjust bits.
	if bits < 0 {
		bits = 0
	} else if bits > key32BitSize {
		bits = key32BitSize
	}

	return n.inplaceInsert(key, uint8(bits), value)
}

// Enumerate returns channel which is populated by nodes with data in order of their keys.
func (n *node32) Enumerate() chan *node32 {
	ch := make(chan *node32)

	go func() {
		defer close(ch)

		// If tree is empty -
		if n == nil {
			// return nothing.
			return
		}

		n.enumerate(ch)
	}()

	return ch
}

// Match locates node which key is equal to or "contains" the key passed as argument.
func (n *node32) Match(key uint32, bits int) (uint32, bool) {
	// If tree is empty -
	if n == nil {
		// report nothing.
		return 0, false
	}

	// Adjust bits.
	if bits < 0 {
		bits = 0
	} else if bits > key32BitSize {
		bits = key32BitSize
	}

	r := n.match(key, uint8(bits))
	if r == nil {
		return 0, false
	}

	return r.value, true
}

// ExactMatch locates node which exactly matches given key.
func (n *node32) ExactMatch(key uint32, bits int) (uint32, bool) {
	// If tree is empty -
	if n == nil {
		// report nothing.
		return 0, false
	}

	// Adjust bits.
	if bits < 0 {
		bits = 0
	} else if bits > key32BitSize {
		bits = key32BitSize
	}

	r := n.exactMatch(key, uint8(bits))
	if r == nil {
		return 0, false
	}

	return r.value, true
}

// Delete removes subtree which is contained by given key. The method uses copy on write strategy.
func (n *node32) Delete(key uint32, bits int) (*node32, bool) {
	// If tree is empty -
	if n == nil {
		// report nothing.
		return n, false
	}

	// Adjust bits.
	if bits < 0 {
		bits = 0
	} else if bits > key32BitSize {
		bits = key32BitSize
	}

	return n.del(key, uint8(bits))
}

func (n *node32) dotString() string {
	if n == nil {
		return "[label=\"nil\"]"
	}

	if n.leaf {
		return fmt.Sprintf("[label=\"k: %08x, b: %d, v: \\\"%d\\\"\"]", n.key, n.bits, n.value)
	}

	return fmt.Sprintf("[label=\"k: %08x, b: %d\"]", n.key, n.bits)
}

func (n *node32) insert(c *node32) *node32 {
	if n == nil {
		return c
	}

	// Find number of common most significant bits (NCSB):
	// 1. xor operation puts zeroes at common bits;
	// 2. or masks put ones so that zeroes can't go after smaller number of significant bits (NSB)
	// 3. count of leading zeroes gives number of common bits
	bits := uint8(bits.LeadingZeros32((n.key ^ c.key) | ^masks32[n.bits] | ^masks32[c.bits]))

	// There are three cases possible:
	// - NCSB less than number of significant bits (NSB) of current tree node:
	if bits < n.bits {
		// (branch for current tree node is determined by a bit right after the last common bit)
		branch := (n.key >> (key32BitSize - 1 - bits)) & 1

		// - NCSB equals to NSB of candidate node:
		if bits == c.bits {
			// make new root from the candidate and put current node to one of its branch;
			c.chld[branch] = n
			return c
		}

		// - NCSB less than NSB of candidate node (it can't be greater because bits after NSB don't count):
		// make new root (non-leaf node)
		m := newNode32(c.key&masks32[bits], bits, false, 0)
		// with current tree node at one of branches
		m.chld[branch] = n
		// and the candidate at the other.
		m.chld[1-branch] = c

		return m
	}

	// - keys are equal (NCSB not less than NSB of current tree node and both numbers are equal):
	if c.bits == n.bits {
		// replace current node with the candidate.
		c.chld = n.chld
		return c
	}

	// - current tree node contains candidate node:
	// make new root as a copy of current tree node;
	m := newNode32(n.key, n.bits, n.leaf, n.value)
	m.chld = n.chld

	// (branch for the candidate is determined by a bit right after the last common bit)
	branch := (c.key >> (key32BitSize - 1 - bits)) & 1
	// insert it to correct branch.
	m.chld[branch] = m.chld[branch].insert(c)

	return m
}

func (n *node32) inplaceInsert(key uint32, sbits uint8, value uint32) *node32 {
	var (
		p      *node32
		branch uint32
	)

	r := n

	for n != nil {
		cbits := uint8(bits.LeadingZeros32((n.key ^ key) | ^masks32[n.bits] | ^masks32[sbits]))
		if cbits < n.bits {
			pBranch := branch
			branch = (n.key >> (key32BitSize - 1 - cbits)) & 1

			var m *node32

			if cbits == sbits {
				m = newNode32(key, sbits, true, value)
				m.chld[branch] = n
			} else {
				m = newNode32(key&masks32[cbits], cbits, false, 0)
				m.chld[1-branch] = newNode32(key, sbits, true, value)
			}

			m.chld[branch] = n
			if p == nil {
				r = m
			} else {
				p.chld[pBranch] = m
			}

			return r
		}

		if sbits == n.bits {
			n.key = key
			n.leaf = true
			n.value = value
			return r
		}

		p = n
		branch = (key >> (key32BitSize - 1 - cbits)) & 1
		n = n.chld[branch]
	}

	n = newNode32(key, sbits, true, value)
	if p == nil {
		return n
	}

	p.chld[branch] = n
	return r
}

func (n *node32) enumerate(ch chan *node32) {
	// Implemented by depth-first search.
	if n.leaf {
		ch <- n
	}

	if n.chld[0] != nil {
		n.chld[0].enumerate(ch)
	}

	if n.chld[1] != nil {
		n.chld[1].enumerate(ch)
	}
}

func (n *node32) match(key uint32, bits uint8) *node32 {
	// If can't be contained in current root node -
	if n.bits > bits {
		// report nothing.
		return nil
	}

	// If NSB of current tree node is the same as key has -
	if n.bits == bits {
		// return current node only if it contains data (leaf node) and masked keys are equal.
		if n.leaf && (n.key^key)&masks32[n.bits] == 0 {
			return n
		}

		return nil
	}

	// If key can be contained by current tree node -
	if (n.key^key)&masks32[n.bits] != 0 {
		// but it isn't report nothing.
		return nil
	}

	// Otherwise jump to branch by key bit right after NSB of current tree node
	c := n.chld[(key>>(key32BitSize-1-n.bits))&1]
	if c != nil {
		// and check if child on the branch has anything.
		r := c.match(key, bits)
		if r != nil {
			return r
		}
	}

	// If nothing matches check if current node contains any data.
	if n.leaf {
		return n
	}

	return nil
}

func (n *node32) exactMatch(key uint32, bits uint8) *node32 {
	// If can't be contained in current root node -
	if n.bits > bits {
		// report nothing.
		return nil
	}

	// If NSB of current tree node is the same as key has -
	if n.bits == bits {
		// return current node only if it contains data (leaf node) and masked keys are equal.
		if n.leaf && (n.key^key)&masks32[n.bits] == 0 {
			return n
		}

		return nil
	}

	// If key can be contained by current tree node -
	if (n.key^key)&masks32[n.bits] != 0 {
		// but it isn't report nothing.
		return nil
	}

	// Otherwise jump to branch by key bit right after NSB of current tree node
	c := n.chld[(key>>(key32BitSize-1-n.bits))&1]
	if c != nil {
		// and check if child on the branch has anything.
		r := c.exactMatch(key, bits)
		if r != nil {
			return r
		}
	}

	return nil
}

func (n *node32) del(key uint32, bits uint8) (*node32, bool) {
	// If key can contain current tree node -
	if bits <= n.bits {
		// report empty new tree and put deletion mark if it contains indeed.
		if (n.key^key)&masks32[bits] == 0 {
			return nil, true
		}

		return n, false
	}

	// If key can be contained by current tree node -
	if (n.key^key)&masks32[n.bits] != 0 {
		// but it isn't report nothing.
		return n, false
	}

	// Otherwise jump to branch by key bit right after NSB of current tree node
	branch := (key >> (key32BitSize - 1 - n.bits)) & 1
	c := n.chld[branch]
	if c == nil {
		// report nothing if the branch is empty.
		return n, false
	}

	// Try to remove from subtree
	c, ok := c.del(key, bits)
	if !ok {
		// and report nothing if nothing has been deleted.
		return n, false
	}

	// If child of non-leaf node has been completely deleted -
	if c == nil && !n.leaf {
		// drop the node.
		return n.chld[1-branch], true
	}

	// If deletion happens inside the branch then copy current node.
	m := newNode32(n.key, n.bits, n.leaf, n.value)
	m.chld = n.chld

	// Replace changed child with new one and return new root with deletion mark set.
	m.chld[branch] = c
	return m, true
}

func newNode32(key uint32, bits uint8, leaf bool, value uint32) *node32 {
	return &node32{
		key:   key,
		bits:  bits,
		leaf:  leaf,
		value: value}
}
