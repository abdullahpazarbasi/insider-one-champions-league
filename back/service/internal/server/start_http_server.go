package server

func (s HTTPServer) Start(addr string) error {
	return s.engine.Start(addr)
}
