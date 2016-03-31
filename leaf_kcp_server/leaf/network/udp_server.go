package network

import (
	"leaf/kcp"
	"leaf/log"
	"sync"
)

type UDPServer struct {
	Addr       string
	MaxConnNum int
	NewAgent   func(*kcp.UDPSession) Agent
	ln         *kcp.Listener

	wgLn sync.WaitGroup
}

func (server *UDPServer) Start() {
	server.init()
	go server.run()
}

func (server *UDPServer) init() {
	ln, err := kcp.Listen(kcp.MODE_NORMAL, server.Addr)
	if err != nil {
		log.Fatal("%v", err)
	}

	if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		log.Release("invalid MaxConnNum, reset to %v", server.MaxConnNum)
	}
	if server.NewAgent == nil {
		log.Fatal("NewAgent must not be nil")
	}

	server.ln = ln

}

func (server *UDPServer) run() {
	server.wgLn.Add(1)
	defer server.wgLn.Done()

	for {
		conn, err := server.ln.Accept()
		if err != nil {
			log.Release("accept error: %v", err)
			return
		}
		if server.ln.SessionCount() >= server.MaxConnNum {
			conn.Close()
			log.Debug("too many connections")
			continue
		}

		agent := server.NewAgent(conn)

		go func() {
			agent.Run()
			agent.OnClose()

		}()
	}
}

func (server *UDPServer) Close() {
	server.ln.Close()
	server.wgLn.Wait()
}
