package config

const (
	MovePattern = `^([1-9]{1}|1[0-5]{1})-[A-O]{1}$`
	Ipv4Pattern = `^(1[0-1]\d|1[3-9]\d|12[0-6]|12[8-9]|2[0-1]\d|22[0-3]|[1-9]\d|[1-9])\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$`
)

var Map [15][15]uint8
var Move string
var NewGame string
var YourTurn bool
var MapChan chan [15][15]uint8 = make(chan [15][15]uint8)

// HTTPRsp is the type of http common response payload
type HTTPRsp struct {
	Code    int           `yaml:"code" json:"code"`
	Message string        `yaml:"message" json:"message"`
	Data    [15][15]uint8 `yaml:"data" json:"data"`
}
