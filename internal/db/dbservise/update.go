package dbservise

// Update increment total, action and country counters.
// Return <nil>, if all OK.
// Return error, if MongoDB document updating finish with error.
func (s *service) Update(action, country string) error {
	s.data.Lock()
	s.data.Total++
	s.data.Actions[action].Total++
	s.data.Actions[action].Countries[country].Total++
	s.data.Countries[country].Total++
	s.data.Countries[country].Actions[action].Total++
	s.data.Unlock()
	return nil
}