type MyCircularDeque struct {
    Head *LinkNode
    Tail *LinkNode
    Length int
    Capacity int
}

type LinkNode struct {
    Val int
    Last *LinkNode
    Next *LinkNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    myCircularDeque := MyCircularDeque{
        Head: &LinkNode{-1, nil, nil},
        Tail: &LinkNode{-1, nil, nil},
        Length: 0,
        Capacity: k,
    }
    myCircularDeque.Head.Next = myCircularDeque.Tail
    myCircularDeque.Tail.Last = myCircularDeque.Head
    return myCircularDeque
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() == true {
        return false
    }
    data := &LinkNode{value, nil, nil}
    data.Next = this.Head.Next
    this.Head.Next.Last = data
    this.Head.Next = data
    data.Last = this.Head
    this.Length ++
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() == true {
        return false
    }
    data := &LinkNode{value, nil, nil}
    data.Last = this.Tail.Last
    this.Tail.Last.Next = data
    this.Tail.Last = data
    data.Next = this.Tail
    this.Length ++
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() == true {
        return false
    }
    this.Head.Next = this.Head.Next.Next
    this.Head.Next.Last = this.Head
    this.Length --
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() == true {
        return false
    }
    this.Tail.Last = this.Tail.Last.Last
    this.Tail.Last.Next = this.Tail
    this.Length --
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() == true {
        return -1
    }
    return this.Head.Next.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() == true {
        return -1
    }
    return this.Tail.Last.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    if this.Length == 0 {
        return true
    }
    return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    if this.Length == this.Capacity {
        return true
    }
    return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */