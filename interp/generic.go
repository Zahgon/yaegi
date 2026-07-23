package interp

func genAST(sc *scope, root *node, types []*itype) (*node, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func copyNode(n, anc *node, recursive bool) *node { _ = "STUB: not implemented"; return nil }

func inferTypesFromCall(sc *scope, fun *node, args []*node) ([]*itype, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkConstraint(it, ct *itype) error { _ = "STUB: not implemented"; return nil }
