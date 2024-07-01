package game

import (
	"encoding/binary"
	"fmt"
	"main/network"
	"net"
)

type Connection struct {
	conn      net.Conn
	connected bool
	incoming  *network.NetworkQueue
	World     *World
	Player    *Player
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		conn:      conn,
		connected: true,
		incoming:  network.NewNetworkQueue(),
	}
}

func (c *Connection) Start() {
	for c.connected {
		payloadSize, err := c.readPayloadSize()
		if err != nil {
			c.Close(fmt.Sprintf("PayloadSize Read Error: %s", err.Error()))
			break
		}

		payload, err := c.readBytes(payloadSize)
		if err != nil {
			c.Close(fmt.Sprintf("Payload Read Error: %s", err.Error()))
			break
		}

		rdr := network.NewNetworkReader(payload)
		messageId := rdr.ReadByte()

		m := network.NewIncomingMessage(messageId)
		m.Read(rdr)

		remaining := rdr.RemainingBytes()
		if rdr.HasError() || remaining > 0 {
			c.Close(fmt.Sprintf("failed to handle message: %d with %d remaining bytes", messageId, remaining))
			break
		}

		c.incoming.Lock()
		c.incoming.Push(m)
		c.incoming.Unlock()
	}
}

func (c *Connection) HandleMessages() bool {

	if !c.connected {
		return false
	}

	c.incoming.Lock()
	defer c.incoming.Unlock()

	for c.incoming.Size() > 0 {

		message := c.incoming.Pop()
		err := c.HandleMessage(message)
		if err != nil {
			c.Close(err.Error())
			return false
		}
	}
	return true
}

func (c *Connection) readPayloadSize() (int, error) {
	buffer := make([]byte, 4)
	bytesRead := 0
	for bytesRead < 4 {
		m, err := c.conn.Read(buffer[bytesRead:])
		if err != nil {
			return -1, err
		}
		bytesRead += m
	}
	payloadSize := int(binary.BigEndian.Uint32(buffer)) - 4
	return payloadSize, nil
}

func (c *Connection) readBytes(n int) ([]byte, error) {
	buffer := make([]byte, n)
	bytesRead := 0
	for bytesRead < n {
		m, err := c.conn.Read(buffer[bytesRead:])
		if err != nil {
			return nil, err
		}
		bytesRead += m
	}
	return buffer, nil
}

func (c *Connection) Close(reason string) {
	if !c.connected {
		return
	}
	c.connected = false

	fmt.Printf("Connection was closed: %s\n", reason)

	err := c.conn.Close()
	if err != nil {
		fmt.Printf("Failed to close connection")
	}
}

// Handlers for client Messages

func (c *Connection) HandleMessage(m network.IncomingMessage) error {
	switch msg := m.(type) {
	case *network.HelloMessage:
		c.HandleHelloMessage(msg)
	case *network.LoadMessage:
		c.HandleLoadMessage(msg)
	case *network.CreateMessage:
		c.HandleCreateMessage(msg)
	default:
		fmt.Println("Unknown message handler type")
	}
	return nil
}

func (c *Connection) HandleHelloMessage(m *network.HelloMessage) {

	// Anything below 0 is unique
	// For now just reset to 0 (Nexus)
	// -1 is Tutorial
	// -2 is Nexus

	if m.GameId < 0 {
		m.GameId = 0
	}

	c.World = GlobalWorldManager.GetWorld(m.GameId)
	if c.World == nil {
		c.Close("Null World")
		return
	}

	c.SendMessage(network.MapInfoMessage(c.World.Width, c.World.Height, c.World.IdName, c.World.DisplayName, 0, 0, 0, false, false))
}

func (c *Connection) HandleLoadMessage(m *network.LoadMessage) {

	c.Player = c.World.CreatePlayer(c, 32.5, 32.5)

	c.SendMessage(network.CreateSuccessMessage(c.Player.Id, 0))
}

func (c *Connection) HandleCreateMessage(m *network.CreateMessage) {

	c.Player = c.World.CreatePlayer(c, 32.5, 32.5)

	c.SendMessage(network.CreateSuccessMessage(c.Player.Id, 0))
}

func (c *Connection) SendMessage(data []byte) {
	_, err := c.conn.Write(data)
	if err != nil {
		c.Close("Send failed")
	}
}

// Server State

func (c *Connection) NewTick(dt float64) {
	tickTime := int32(dt * 1000.0)

	fmt.Println("TickTime:", tickTime)

	_ = tickTime
}