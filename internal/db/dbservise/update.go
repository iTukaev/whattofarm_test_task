package dbservise

// Update increment total, action and country counters.
// Return <nil>, if all OK.
// Return error, if MongoDB document updating finish with error.
func (s *Service) Update(action, country string) {
	s.Lock()
	s.data.Total++
	if _, ok := s.data.Actions[action]; !ok {
		s.data.Actions[action] = new(SubCountries)
		s.data.Actions[action].Countries = make(map[string]*TotalCounter)
	}
	if _, ok := s.data.Actions[action].Countries[country]; !ok {
		s.data.Actions[action].Countries[country] = new(TotalCounter)
	}
	s.data.Actions[action].Countries[country].Total++
	s.data.Actions[action].Total++

	if _, ok := s.data.Countries[country]; !ok {
		s.data.Countries[country] = new(SubActions)
		s.data.Countries[country].Actions = make(map[string]*TotalCounter)
	}
	if _, ok := s.data.Countries[country].Actions[action]; !ok {
		s.data.Countries[country].Actions[action] = new(TotalCounter)
	}
	s.data.Countries[country].Actions[action].Total++
	s.data.Countries[country].Total++
	s.Unlock()
}