package game

import (
	"fmt"
	"net"
	"sync"
)

var GlobalServerListener *ServerListener

type ServerListener struct {
	listener    net.Listener
	Connections map[int32]*Connection
	mutex       sync.Mutex
	nextId      int32
}

func NewAndServe(addr string) (*ServerListener, error) {
	server := &ServerListener{
		Connections: make(map[int32]*Connection),
		mutex:       sync.Mutex{},
		nextId:      0,
	}

	if err := server.Start(addr); err != nil {
		return nil, err
	}

	return server, nil
}

func (server *ServerListener) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server.listener = listener
	go server.acceptConnections()
	return nil
}

func (server *ServerListener) GetConnection(id int32) *Connection {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	conn, ok := server.Connections[id]
	if !ok {
		return nil
	}
	return conn
}

func (server *ServerListener) ProcessConnectionMessages() {

	server.mutex.Lock()
	defer server.mutex.Unlock()

	for id, conn := range server.Connections {
		if !conn.HandleMessages() {
			delete(server.Connections, id)
		}
	}
}

func (server *ServerListener) acceptConnections() {
	defer server.listener.Close()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		fmt.Printf("Acceped new Connection: %s\n", conn.RemoteAddr())

		connection := NewConnection(conn)

		server.mutex.Lock()
		server.Connections[server.nextId] = connection
		server.nextId++
		server.mutex.Unlock()

		go connection.Start()
	}
}

func (s *ServerListener) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, conn := range s.Connections {
		conn.Close("Server Stopped")
	}

	s.listener.Close()
}
