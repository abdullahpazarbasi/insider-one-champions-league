package server

func (s *Server) Start(address string) error {
	return s.router.Start(address)
}
