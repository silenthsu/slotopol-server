package spi

import "errors"

const (
	SECnull = iota

	// authorization
	SEC_unauthorized
	SEC_nourl

	// POST /game/join
	SEC_game_join_nobind
	SEC_game_join_norid
	SEC_game_join_nouid
	SEC_game_join_nodata
	SEC_game_join_noalias
	SEC_game_join_noroom
	SEC_game_join_nouser
	SEC_prop_join_noaccess
	SEC_game_join_noreels
	SEC_game_join_open
	SEC_game_join_props

	// POST /game/part
	SEC_game_part_nobind
	SEC_game_part_nogid
	SEC_game_part_notopened
	SEC_game_part_nouser
	SEC_prop_part_noaccess

	// POST /game/state
	SEC_game_state_nobind
	SEC_game_state_nogid
	SEC_game_state_notopened
	SEC_game_state_nouser
	SEC_prop_state_noaccess
	SEC_game_state_noprops

	// POST /game/bet
	SEC_game_betget_nobind
	SEC_game_betget_nogid
	SEC_game_betget_notopened
	SEC_prop_betget_noaccess

	// POST /game/bet
	SEC_game_betset_nobind
	SEC_game_betset_nogid
	SEC_game_betset_nodata
	SEC_game_betset_notopened
	SEC_prop_betset_noaccess
	SEC_game_betset_badbet

	// POST /game/sbl
	SEC_game_sblget_nobind
	SEC_game_sblget_nogid
	SEC_game_sblget_notopened
	SEC_prop_sblget_noaccess

	// POST /game/sbl
	SEC_game_sblset_nobind
	SEC_game_sblset_nogid
	SEC_game_sblset_nodata
	SEC_game_sblset_notopened
	SEC_prop_sblset_noaccess
	SEC_game_sblset_badlines

	// POST /game/spin
	SEC_game_spin_nobind
	SEC_game_spin_nogid
	SEC_game_spin_notopened
	SEC_game_spin_noroom
	SEC_game_spin_nouser
	SEC_prop_spin_noaccess
	SEC_game_spin_noprops
	SEC_game_spin_nomoney
	SEC_game_spin_badbank
	SEC_game_spin_sqlbank
	SEC_game_spin_sqlupdate

	// POST /prop/wallet
	SEC_prop_walletget_nobind
	SEC_prop_walletget_norid
	SEC_prop_walletget_nouid
	SEC_prop_walletget_noroom
	SEC_prop_walletget_nouser
	SEC_prop_walletget_noaccess

	// POST /prop/wallet
	SEC_prop_walletadd_nobind
	SEC_prop_walletadd_norid
	SEC_prop_walletadd_nouid
	SEC_prop_walletadd_noadd
	SEC_prop_walletadd_limit
	SEC_prop_walletadd_noroom
	SEC_prop_walletadd_nouser
	SEC_prop_walletadd_nomoney
	SEC_prop_walletadd_noaccess
	SEC_prop_walletadd_sqlupdate
	SEC_prop_walletadd_sqlinsert
	SEC_prop_walletadd_sqllog

	// POST /game/doubleup
	SEC_game_doubleup_nobind
	SEC_game_doubleup_nogid
	SEC_game_doubleup_nomult
	SEC_game_doubleup_bigmult
	SEC_game_doubleup_notopened
	SEC_game_doubleup_noroom
	SEC_game_doubleup_nouser
	SEC_prop_doubleup_noaccess
	SEC_game_doubleup_noprops
	SEC_game_doubleup_nomoney

	// POST /user/rename
	SEC_user_rename_nobind
	SEC_user_rename_nouid
	SEC_user_rename_noname
	SEC_user_rename_nouser
	SEC_prop_rename_noaccess
	SEC_user_rename_update
)

var (
	Err404       = errors.New("page not found")
	ErrNoUID     = errors.New("user ID does not provided")
	ErrNoRID     = errors.New("room ID does not provided")
	ErrNoGID     = errors.New("game ID does not provided")
	ErrNoData    = errors.New("data does not provided or empty")
	ErrNoRoom    = errors.New("room with given ID does not found")
	ErrNoUser    = errors.New("user with given ID does not found")
	ErrNoWallet  = errors.New("wallet for given user and room does not found")
	ErrNoMoney   = errors.New("not enough money")
	ErrNoAccess  = errors.New("no access rights for this feature")
	ErrNoMult    = errors.New("gamble multiplier not given")
	ErrBigMult   = errors.New("gamble multiplier too big")
	ErrZero      = errors.New("given value is zero")
	ErrTooBig    = errors.New("given value exceeds the limit")
	ErrNoAliase  = errors.New("no game alias")
	ErrNoReels   = errors.New("no reels for given game with selected percentage")
	ErrNotOpened = errors.New("game with given ID is not opened")
	ErrBadBank   = errors.New("can not generate spin with current bank balance")
)
