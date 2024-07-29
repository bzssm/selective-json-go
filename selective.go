package selective_json_go

type SelectiveMarshaller struct {
	Scenario string
}

func NewSelectiveMarshaller(scenario string) *SelectiveMarshaller {
	return &SelectiveMarshaller{
		Scenario: scenario,
	}
}

func (s *SelectiveMarshaller) Marshal(v any) ([]byte, error) {
	e := newEncodeState()
	e.setScenario(s.Scenario)
	defer encodeStatePool.Put(e)

	err := e.marshal(v, encOpts{escapeHTML: true})
	if err != nil {
		return nil, err
	}
	buf := append([]byte(nil), e.Bytes()...)

	return buf, nil
}
