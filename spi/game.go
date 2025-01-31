package spi

import (
	"encoding/xml"
	"sync/atomic"

	cfg "github.com/slotopol/server/config"
	game "github.com/slotopol/server/game"
	keno "github.com/slotopol/server/game/keno"
	slot "github.com/slotopol/server/game/slot"
	"github.com/slotopol/server/util"

	"github.com/gin-gonic/gin"
)

var (
	SpinBuf util.SqlBuf[Spinlog]
	MultBuf util.SqlBuf[Multlog]
	BankBat = map[uint64]*SqlBank{}
	JoinBuf = SqlStory{}
)

// Joins to game and creates new instance of game.
func SpiGameJoin(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		CID     uint64   `json:"cid" yaml:"cid" xml:"cid,attr" form:"cid"`
		UID     uint64   `json:"uid" yaml:"uid" xml:"uid,attr" form:"uid"`
		Alias   string   `json:"alias" yaml:"alias" xml:"alias" form:"alias" binding:"required"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr"`
		Game    any      `json:"game" yaml:"game" xml:"game"`
		Scrn    any      `json:"scrn" yaml:"scrn" xml:"scrn"`
		Wallet  float64  `json:"wallet" yaml:"wallet" xml:"wallet"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_game_join_nobind, err)
		return
	}
	if arg.CID == 0 {
		Ret400(c, SEC_game_join_norid, ErrNoCID)
		return
	}
	if arg.UID == 0 {
		Ret400(c, SEC_game_join_nouid, ErrNoUID)
		return
	}

	var club *Club
	if club, ok = Clubs.Get(arg.CID); !ok {
		Ret404(c, SEC_game_join_noclub, ErrNoClub)
		return
	}

	var user *User
	if user, ok = Users.Get(arg.UID); !ok {
		Ret404(c, SEC_game_join_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, arg.CID)
	if (al&ALmem == 0) || (admin != user && al&ALgame == 0) {
		Ret403(c, SEC_game_join_noaccess, ErrNoAccess)
		return
	}

	var alias = util.ToID(arg.Alias)
	var maker, has = game.GameFactory[alias]
	if !has {
		Ret400(c, SEC_game_join_noalias, ErrNoAliase)
		return
	}

	var anygame = maker()
	var gid = atomic.AddUint64(&StoryCounter, 1)
	var scene = &Scene{
		Story: Story{
			GID:   gid,
			Alias: alias,
			CID:   arg.CID,
			UID:   arg.UID,
			Flow:  true,
		},
		Game: anygame,
	}

	// insert new story entry
	if Cfg.ClubInsertBuffer > 1 {
		go JoinBuf.Join(cfg.XormStorage, &scene.Story)
	} else if err = JoinBuf.Join(cfg.XormStorage, &scene.Story); err != nil {
		Ret500(c, SEC_game_join_sql, err)
		return
	}

	Scenes.Set(scene.GID, scene)
	user.games.Set(scene.GID, scene)

	ret.GID = gid
	ret.Game = anygame
	// make game screen object
	club.mux.RLock()
	var rtp = GetRTP(user, club)
	club.mux.RUnlock()
	switch game := anygame.(type) {
	case slot.SlotGame:
		var scrn = game.NewScreen()
		game.Spin(scrn, rtp)
		ret.Scrn = scrn
	case keno.KenoGame:
		var scrn keno.Screen
		game.Spin(&scrn, rtp)
		ret.Scrn = &scrn
	}
	ret.Wallet = user.GetWallet(arg.CID)

	RetOk(c, ret)
}

// Removes instance of opened game.
func SpiGamePart(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_game_part_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_part_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_game_part_notopened, ErrNotOpened)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_game_part_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin != user && al&ALgame == 0 {
		Ret403(c, SEC_game_part_noaccess, ErrNoAccess)
		return
	}

	// update story entry
	if Cfg.ClubUpdateBuffer > 1 {
		go JoinBuf.Flow(cfg.XormStorage, arg.GID, false)
	} else if err = JoinBuf.Flow(cfg.XormStorage, arg.GID, false); err != nil {
		Ret500(c, SEC_game_part_sql, err)
		return
	}

	scene.Flow = false
	Scenes.Delete(arg.GID)
	user.games.Delete(arg.GID)

	Ret204(c)
}

// Returns full info of game scene with given GID, and balance on wallet.
func SpiGameInfo(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr"`
		Alias   string   `json:"alias" yaml:"alias" xml:"alias"`
		CID     uint64   `json:"cid" yaml:"cid" xml:"cid,attr"`
		UID     uint64   `json:"uid" yaml:"uid" xml:"uid,attr"`
		SID     uint64   `json:"sid" yaml:"sid" xml:"sid,attr"`
		Game    any      `json:"game" yaml:"game" xml:"game"`
		Wallet  float64  `json:"wallet" yaml:"wallet" xml:"wallet"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_game_info_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_info_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_game_info_notopened, ErrNotOpened)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_game_info_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin != user && al&ALgame == 0 {
		Ret403(c, SEC_game_info_noaccess, ErrNoAccess)
		return
	}

	var props *Props
	if props, ok = user.props.Get(scene.CID); !ok {
		Ret500(c, SEC_game_info_noprops, ErrNoProps)
		return
	}

	ret.GID = arg.GID
	ret.Alias = scene.Alias
	ret.CID = scene.CID
	ret.UID = scene.UID
	ret.SID = scene.SID
	ret.Game = scene.Game
	ret.Wallet = props.Wallet

	RetOk(c, ret)
}

// Returns master RTP for given GID.
func SpiGameRtpGet(c *gin.Context) {
	var err error
	var ok bool
	var arg struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"arg"`
		GID     uint64   `json:"gid" yaml:"gid" xml:"gid,attr" form:"gid"`
	}
	var ret struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"ret"`
		RTP     float64  `json:"rtp" yaml:"rtp" xml:"rtp"`
	}

	if err = c.ShouldBind(&arg); err != nil {
		Ret400(c, SEC_game_rdget_nobind, err)
		return
	}
	if arg.GID == 0 {
		Ret400(c, SEC_game_rdget_nogid, ErrNoGID)
		return
	}

	var scene *Scene
	if scene, ok = Scenes.Get(arg.GID); !ok {
		Ret404(c, SEC_game_rdget_notopened, ErrNotOpened)
		return
	}

	var club *Club
	if club, ok = Clubs.Get(scene.CID); !ok {
		Ret500(c, SEC_game_rdget_noclub, ErrNoClub)
		return
	}

	var user *User
	if user, ok = Users.Get(scene.UID); !ok {
		Ret500(c, SEC_game_rdget_nouser, ErrNoUser)
		return
	}

	var admin, al = MustAdmin(c, scene.CID)
	if admin.UID != scene.UID && al&ALgame == 0 {
		Ret403(c, SEC_game_rdget_noaccess, ErrNoAccess)
		return
	}

	club.mux.RLock()
	ret.RTP = GetRTP(user, club)
	club.mux.RUnlock()

	RetOk(c, ret)
}
