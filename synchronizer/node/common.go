package node

import "errors"

var (
	ErrHeaderTraversalAheadOfProvider            = errors.New("the HeaderTraversal's internal state is ahead of the provider")
	ErrHeaderTraversalAndProviderMismatchedState = errors.New("the HeaderTraversal and provider have diverged in state")
	ErrHeaderTraversalCheckHeaderByHashDelDbData = errors.New("the HeaderTraversal headerList[0].ParentHash != dbLatestHeader.Hash()")
)
