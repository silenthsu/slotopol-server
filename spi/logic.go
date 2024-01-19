package spi

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slotopol/server/game"
	"github.com/slotopol/server/util"
)

// Room means independent bank into which gambles some users.
type Room struct {
	RID  uint64  `xorm:"pk autoincr" json:"rid" yaml:"rid" xml:"rid,attr"`
	Name string  `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Bank float64 `xorm:"notnull" json:"bank" yaml:"bank" xml:"bank"` // users win/lost balance, in coins
	Fund float64 `xorm:"notnull" json:"fund" yaml:"fund" xml:"fund"` // jackpot fund, in coins
	Lock float64 `xorm:"notnull" json:"lock" yaml:"lock" xml:"lock"` // not changed deposit within games

	JptRate float64 `xorm:"'jptrate' notnull default 0.015" json:"jptrate" yaml:"jptrate" xml:"jptrate"`
	GainRTP float64 `xorm:"'gainrtp' notnull default 95.00" json:"gainrtp" yaml:"gainrtp" xml:"gainrtp"`

	mux sync.RWMutex
}

// User means registration of somebody. Each user can have splitted
// wallet with some coins balance in each Room. User can opens several
// games without any limitation.
type User struct {
	UID    uint64 `xorm:"pk autoincr" json:"uid" yaml:"uid" xml:"uid,attr"`
	Email  string `xorm:"notnull unique index" json:"email" yaml:"email" xml:"email"`
	Secret string `xorm:"notnull" json:"secret" yaml:"secret" xml:"secret"` // auth password
	Name   string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	games  util.RWMap[uint64, OpenGame]
	props  util.RWMap[uint64, *Props]
}

// OpenGame is opened game for user with UID at room with RID.
// Each instance of game have own GID. Alias - is game type identifier.
type OpenGame struct {
	GID   uint64 `xorm:"pk autoincr" json:"gid" yaml:"gid" xml:"gid,attr"`
	RID   uint64 `xorm:"notnull" json:"rid" yaml:"rid" xml:"rid,attr"`
	UID   uint64 `xorm:"notnull" json:"uid" yaml:"uid" xml:"uid,attr"`
	Alias string `xorm:"notnull" json:"alias" yaml:"alias" xml:"alias"`
	game  game.SlotGame
}

func (OpenGame) TableName() string {
	return "game"
}

// Access level.
type AL uint

const (
	ALgame  AL = 1 << iota // can change room game settings
	ALuser                 // can change user balance and move user money to/from room deposit
	ALbank                 // can change room bank, fund, deposit
	ALadmin                // can change same access levels to other users
	ALall   AL = 0xffff    // all available rights.
)

// Props contains properties for user at some room.
// Any property can be zero by default, or if object does not created at DB.
type Props struct {
	RID    uint64 `xorm:"notnull index(bid)" json:"rid" yaml:"rid" xml:"rid,attr"`
	UID    uint64 `xorm:"notnull index(bid)" json:"uid" yaml:"uid" xml:"uid,attr"`
	Wallet int    `xorm:"notnull" json:"wallet" yaml:"wallet" xml:"wallet"` // in coins
	Access AL     `xorm:"notnull" json:"access" yaml:"access" xml:"access"`
}

type Spinlog struct {
	SID    uint64    `xorm:"pk autoincr" json:"sid" yaml:"sid" xml:"sid,attr"`
	GID    uint64    `xorm:"notnull" json:"gid" yaml:"gid" xml:"gid,attr"`
	Game   string    `xorm:"notnull" json:"game" yaml:"game" xml:"game"`
	Screen string    `xorm:"notnull" json:"screen,omitempty" yaml:"screen,omitempty" xml:"screen,omitempty"`
	Wins   string    `json:"wins,omitempty" yaml:"wins,omitempty" xml:"wins,omitempty"`
	Gain   int       `xorm:"notnull" json:"gain" yaml:"gain" xml:"gain"`
	Wallet int       `xorm:"notnull" json:"wallet" yaml:"wallet" xml:"wallet"`
	CTime  time.Time `xorm:"created" json:"ctime" yaml:"ctime" xml:"ctime"`
}

type Walletlog struct {
	RID    uint64    `xorm:"notnull index(bid)" json:"rid" yaml:"rid" xml:"rid,attr"`
	UID    uint64    `xorm:"notnull index(bid)" json:"uid" yaml:"uid" xml:"uid,attr"`
	AdmID  uint64    `xorm:"notnull" json:"admid" yaml:"admid" xml:"admid"`
	Wallet int       `xorm:"notnull" json:"wallet" yaml:"wallet" xml:"wallet"` // in coins
	Addend int       `xorm:"notnull" json:"addend" yaml:"addend" xml:"addend"`
	CTime  time.Time `xorm:"created" json:"ctime" yaml:"ctime" xml:"ctime"`
}

// All created rooms, by RID.
var Rooms util.RWMap[uint64, *Room]

// All registered users, by UID.
var Users util.RWMap[uint64, *User]

// All opened games, by GID.
var OpenGames util.RWMap[uint64, OpenGame]

func (user *User) Init() {
	user.games.Init(0)
	user.props.Init(0)
}

func (user *User) GetWallet(rid uint64) int {
	if props, ok := user.props.Get(rid); ok {
		return props.Wallet
	}
	return 0
}

func (user *User) GetAL(rid uint64) AL {
	if props, ok := user.props.Get(rid); ok {
		return props.Access
	}
	return 0
}

func (user *User) InsertProps(props *Props) {
	user.props.Set(props.RID, props)
}

// GetAdmin always returns User pointer for authorized requests.
// And access level for it. Or returns nil pointer for not authorized request.
func GetAdmin(c *gin.Context, rid uint64) (*User, AL) {
	if v, ok := c.Get(identityKey); ok {
		var admin = v.(*User)
		var alg = admin.GetAL(0)
		var alr = admin.GetAL(rid)
		return admin, alg | alr
	}
	return nil, 0
}

func init() {
	Rooms.Init(0)
	Users.Init(0)
	OpenGames.Init(0)
}
