package server

type stub1Type struct {
	featureFlags uint32
}

type stub2Type struct {
	featureFlags uint32
}

func (s1 *stub1Type) stub1Method() {
	s1.featureFlags = 0
}

func (s2 *stub2Type) stub2Method() {
	s2.featureFlags = 0
}

type i1 interface {
	stub1Method()
	stub2Method()
}

func aggle(test i1) {
	test.stub1Method()
	//test.stub2Method()
}
