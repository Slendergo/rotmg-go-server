package network

const (
	Failure                = byte(0)
	Hello                  = byte(1)
	Unused2                = byte(2)
	ClaimLoginRewardMsg    = byte(3)
	DeletePet              = byte(4)
	Requesttrade           = byte(5)
	QuestFetchResponse     = byte(6)
	Joinguild              = byte(7)
	Ping                   = byte(8)
	Newtick                = byte(9)
	Playertext             = byte(10)
	Useitem                = byte(11)
	Serverplayershoot      = byte(12)
	Showeffect             = byte(13)
	Tradeaccepted          = byte(14)
	Guildremove            = byte(15)
	Petupgraderequest      = byte(16)
	EnterArena             = byte(17)
	Goto                   = byte(18)
	Invswap                = byte(19)
	Otherhit               = byte(20)
	Nameresult             = byte(21)
	Unused22               = byte(22)
	HatchPet               = byte(23)
	ActivePetUpdateRequest = byte(24)
	Enemyhit               = byte(25)
	Guildresult            = byte(26)
	Editaccountlist        = byte(27)
	Tradechanged           = byte(28)
	Unused29               = byte(29)
	Playershoot            = byte(30)
	Pong                   = byte(31)
	Unused32               = byte(32)
	PetChangeSkinMsg       = byte(33)
	Tradedone              = byte(34)
	Enemyshoot             = byte(35)
	Accepttrade            = byte(36)
	Changetrade            = byte(37)
	Playsound              = byte(38)
	VerifyEmail            = byte(39)
	Squarehit              = byte(40)
	NewAbility             = byte(41)
	Move                   = byte(42)
	Petyardupdate          = byte(43)
	Text                   = byte(44)
	Reconnect              = byte(45)
	Death                  = byte(46)
	Unused47               = byte(47)
	QuestRoomMsg           = byte(48)
	Allyshoot              = byte(49)
	ImminentArenaWave      = byte(50)
	Unused51               = byte(51)
	ResetDailyQuests       = byte(52)
	PetChangeFormMsg       = byte(53)
	Unused54               = byte(54)
	Unused55               = byte(55)
	Unused56               = byte(56)
	Load                   = byte(57)
	QuestRedeem            = byte(58)
	Createguild            = byte(59)
	Setcondition           = byte(60)
	Create                 = byte(61)
	Update                 = byte(62)
	KeyInfoResponse        = byte(63)
	Aoe                    = byte(64)
	Gotoack                = byte(65)
	GlobalNotification     = byte(66)
	Notification           = byte(67)
	ArenaDeath             = byte(68)
	Clientstat             = byte(69)
	Unused70               = byte(70)
	Unused71               = byte(71)
	Unused72               = byte(72)
	Unused73               = byte(73)
	Unused74               = byte(74)
	Unused75               = byte(75)
	Activepetupdate        = byte(76)
	Unused77               = byte(77)
	Unused78               = byte(78)
	PasswordPrompt         = byte(79)
	AcceptArenaDeath       = byte(80)
	UpdateAck              = byte(81)
	Unused82               = byte(82)
	Unused83               = byte(83)
	RealmHeroLeftMsg       = byte(84)
	Unused85               = byte(85)
	Tradestart             = byte(86)
	EvolvePet              = byte(87)
	Traderequested         = byte(88)
	Unused89               = byte(89)
	Unused90               = byte(90)
	Canceltrade            = byte(91)
	MapInfo                = byte(92)
	LoginRewardMsg         = byte(93)
	KeyInfoRequest         = byte(94)
	Unused95               = byte(95)
	Unused96               = byte(96)
	Unused97               = byte(97)
	QuestFetchAsk          = byte(98)
	Accountlist            = byte(99)
	Shootack               = byte(100)
	CreateSuccess          = byte(101)
	Checkcredits           = byte(102)
	Grounddamage           = byte(103)
	Guildinvite            = byte(104)
	Escape                 = byte(105)
	File                   = byte(106)
	ReskinUnlock           = byte(107)
	Unused108              = byte(108)
	Unused109              = byte(109)
	Unused110              = byte(110)
	Unused111              = byte(111)
	Unused112              = byte(112)
	Unused113              = byte(113)
	Unused114              = byte(114)
	Unused115              = byte(115)
	Unused116              = byte(116)
	Unused117              = byte(117)
	Unused118              = byte(118)
	Unused119              = byte(119)
	Unused120              = byte(120)
	Unused121              = byte(121)
	Unused122              = byte(122)
	Unused123              = byte(123)
	Unused124              = byte(124)
	Unused125              = byte(125)
	Unused126              = byte(126)
	Unused127              = byte(127)
)

// Incoming

type IncomingMessage interface {
	Read(rdr *NetworkReader)
}

func NewIncomingMessage(id byte) IncomingMessage {
	switch id {

	case Hello:
		return &HelloMessage{}

	case Load:
		return &LoadMessage{}

	case Create:
		return &CreateMessage{}

	case Move:
		return &MoveMessage{}

	case UpdateAck:
		return &UpdateAckMessage{}

	default:
		return nil
	}
}

type HelloMessage struct {
	BuildVersion string
	GameId       int32
	Email        string
	Password     string
	KeyTime      int32
	Key          string
	MapJson      string
}

func (m *HelloMessage) Read(rdr *NetworkReader) {
	m.BuildVersion = rdr.ReadString()
	m.GameId = rdr.ReadInt()
	m.Email = rdr.ReadString()
	m.Password = rdr.ReadString()
	m.KeyTime = rdr.ReadInt()
	m.Key = rdr.ReadString()
	m.MapJson = rdr.ReadString32()
}

type LoadMessage struct {
	CharacterId  int32
	IsFromArena  bool
	IsChallenger bool
}

func (m *LoadMessage) Read(rdr *NetworkReader) {
	m.CharacterId = rdr.ReadInt()
	m.IsFromArena = rdr.ReadBool()
	m.IsChallenger = rdr.ReadBool()
}

type CreateMessage struct {
	ClassType    int16
	SkinType     int16
	IsChallenger bool
}

func (m *CreateMessage) Read(rdr *NetworkReader) {
	m.ClassType = rdr.ReadShort()
	m.SkinType = rdr.ReadShort()
	m.IsChallenger = rdr.ReadBool()
}

type MoveRecord struct {
	Time int32
	X    float32
	Y    float32
}

type MoveMessage struct {
	TickId                        int32
	Time                          int32
	ServerRealTimeMSofLastNewTick int32
	NewX                          float32
	NewY                          float32
	MoveRecords                   []MoveRecord
}

func (m *MoveMessage) Read(rdr *NetworkReader) {
	m.TickId = rdr.ReadInt()
	m.Time = rdr.ReadInt()
	m.ServerRealTimeMSofLastNewTick = rdr.ReadInt()
	m.NewX = rdr.ReadFloat()
	m.NewY = rdr.ReadFloat()
	len := rdr.ReadShort()
	if len > 0 && len <= 10 {
		m.MoveRecords = make([]MoveRecord, len)
		for i := 0; i < int(len); i++ {
			m.MoveRecords[i] = MoveRecord{
				Time: rdr.ReadInt(),
				X:    rdr.ReadFloat(),
				Y:    rdr.ReadFloat(),
			}
		}
	}
}

type UpdateAckMessage struct{}

func (m *UpdateAckMessage) Read(rdr *NetworkReader) {}

// Outgoing

func FailureMessage(id int32, message string) []byte {
	wtr := NewNetworkWriter(Failure)
	wtr.WriteInt(id)
	wtr.WriteString(message)
	return wtr.Buffer()
}

func MapInfoMessage(width int32, height int32, name string, displayName string, fp int32, background int32, difficulty int32, allowPlayerTeleport bool, showDisplays bool) []byte {
	wtr := NewNetworkWriter(MapInfo)
	wtr.WriteInt(width)
	wtr.WriteInt(height)
	wtr.WriteString(name)
	wtr.WriteString(displayName)
	wtr.WriteInt(fp)
	wtr.WriteInt(background)
	wtr.WriteInt(difficulty)
	wtr.WriteBool(allowPlayerTeleport)
	wtr.WriteBool(showDisplays)
	return wtr.Buffer()
}

func CreateSuccessMessage(objectId int32, characterId int32) []byte {
	wtr := NewNetworkWriter(CreateSuccess)
	wtr.WriteInt(objectId)
	wtr.WriteInt(characterId)
	return wtr.Buffer()
}

type UpdateTileData struct {
	X    int32
	Y    int32
	Type int32
}

func NewUpdateTileData(x int32, y int32, typ int32) UpdateTileData {
	return UpdateTileData{
		X:    x,
		Y:    y,
		Type: typ,
	}
}

type NewObjectData struct {
	ObjectType int32
	StatusData StatusData
}

type StatusData struct {
	ObjectId int32
	X        float32
	Y        float32
	Stats    []StatData
}

// todo
type StatData struct {
	Type        byte
	IntValue    int32
	StringValue string
}

func UpdateMessage(tiles []UpdateTileData, newObjs []NewObjectData, drops []int32) []byte {
	wtr := NewNetworkWriter(Update)

	length := len(tiles)
	wtr.WriteCompressedInt(length)
	for i := 0; i < length; i++ {
		wtr.WriteShort(int16(tiles[i].X))
		wtr.WriteShort(int16(tiles[i].Y))
		wtr.WriteUnsignedShort(uint16(tiles[i].Type))
	}

	length = len(newObjs)
	wtr.WriteCompressedInt(length)
	for i := 0; i < length; i++ {
		wtr.WriteUnsignedShort(uint16(newObjs[i].ObjectType))
		wtr.WriteCompressedInt(int(newObjs[i].StatusData.ObjectId))
		wtr.WriteFloat(newObjs[i].StatusData.X)
		wtr.WriteFloat(newObjs[i].StatusData.Y)
		wtr.WriteCompressedInt(0) // StatData
	}

	length = len(drops)
	wtr.WriteCompressedInt(length)
	for i := 0; i < length; i++ {
		wtr.WriteCompressedInt(int(drops[i]))
	}
	return wtr.Buffer()
}

func NewTickMessage(tickId int32, tickTime int32) []byte {
	wtr := NewNetworkWriter(Newtick)
	wtr.WriteInt(tickId)
	wtr.WriteInt(tickTime)
	wtr.WriteInt(0)   // serverRealTimeMS_
	wtr.WriteShort(0) // serverLastRTTMS_
	wtr.WriteShort(0)
	// todo status data
	return wtr.Buffer()
}
