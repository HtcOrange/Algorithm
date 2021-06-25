func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{-1, nil}
    last := head
    for l1 != nil && l2 != nil {
        var next *ListNode
        if l1.Val < l2.Val {
            next = l1.Next
            last.Next = l1
            l1 = next
        }else {
            next = l2.Next
            last.Next = l2
            l2 = next
        }
        last = last.Next
    }
    if l1 != nil {
        last.Next = l1
    }
    if l2 != nil {
        last.Next = l2
    }
    return head.Next
}