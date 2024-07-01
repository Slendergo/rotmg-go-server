package game

import (
	"encoding/binary"
	"fmt"
	"main/network"
	"net"
)

type Connection struct {
	conn      net.Conn
	Connected bool
	incoming  *network.NetworkQueue

	// State
	tickId int32
	World  *World
	Player *Player

	Visibles map[int32]bool
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		conn:      conn,
		Connected: true,
		incoming:  network.NewNetworkQueue(),

		Visibles: make(map[int32]bool),
	}
}

func (c *Connection) Start() {
	for c.Connected {
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
		if m == nil {
			c.Close(fmt.Sprintf("Unknown MessageId: %d", messageId))
			break
		}
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

	if !c.Connected {
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
	if !c.Connected {
		return
	}
	c.Connected = false

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
	case *network.MoveMessage:
		c.HandleMoveMessage(msg)
	case *network.UpdateAckMessage:
		c.HandleUpdateAckMessage(msg)
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
	c.updateSurroundings()
}

func (c *Connection) HandleCreateMessage(m *network.CreateMessage) {

	c.Player = c.World.CreatePlayer(c, 32.5, 32.5)

	c.SendMessage(network.CreateSuccessMessage(c.Player.Id, 0))
	c.updateSurroundings()
}

func (c *Connection) HandleMoveMessage(m *network.MoveMessage) {

	if !c.World.InBoundsFloat(m.NewX, m.NewY) {
		// c.Close("Out of bounds")
		return
	}

	c.Player.SetPosition(m.NewX, m.NewY)
}

func (c *Connection) HandleUpdateAckMessage(m *network.UpdateAckMessage) {

}

func (c *Connection) SendMessage(data []byte) {
	_, err := c.conn.Write(data)
	if err != nil {
		c.Close("Send failed")
	}
}

// Server State

func (c *Connection) NewTick(dt float64) {
	c.updateSurroundings()

	c.tickId++
	tickTime := int32(dt * 1000.0)
	c.SendMessage(network.NewTickMessage(c.tickId, tickTime))
}

func (c *Connection) updateSurroundings() {

	var tiles = c.updateTiles()
	var newObjs = c.updateObjects()
	var drops = c.updateDrops()

	if len(tiles) > 0 || len(newObjs) > 0 || len(drops) > 0 {
		c.SendMessage(network.UpdateMessage(tiles, newObjs, drops))
	}
}

func (c *Connection) updateTiles() []network.UpdateTileData {

	var tiles []network.UpdateTileData

	playerX := c.Player.X
	playerY := c.Player.Y

	for dx := -15; dx <= 15; dx++ {
		for dy := -15; dy <= 15; dy++ {

			if dx*dx+dy*dy > 15*15 {
				continue
			}

			tileX := int32(playerX + float32(dx))
			tileY := int32(playerY + float32(dy))

			if !c.World.InBoundsInt32(tileX, tileY) {
				continue
			}

			tile := c.World.tiles[tileX][tileY]

			tiles = append(tiles, network.UpdateTileData{
				X:    tileX,
				Y:    tileY,
				Type: tile.Type,
			})
		}
	}
	return tiles
}

func (c *Connection) updateObjects() []network.NewObjectData {
	var newObjs []network.NewObjectData

	// hacky way to add player
	_, ok := c.Visibles[c.Player.Id]
	if !ok {
		newObjs = append(newObjs, network.NewObjectData{
			ObjectType: 0x030e,
			StatusData: network.StatusData{
				ObjectId: c.Player.Id,
				X:        c.Player.X,
				Y:        c.Player.Y,
				Stats:    []network.StatData{},
			},
		})

		c.Visibles[c.Player.Id] = true
	}

	return newObjs
}

func (c *Connection) updateDrops() []int32 {
	var drops []int32

	return drops
}
