package chicago

import (
	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 47.751(lined) + 15.264(scatter) = 63.014554%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 63.015(sym) + 0.071909*67.897(fg)*4.2 = 83.520645%
var Reels835 = slot.Reels5x{
	{7, 6, 11, 2, 8, 5, 12, 10, 6, 7, 9, 5, 8, 10, 9, 4, 12, 11, 1, 2, 7, 4, 11, 12, 3, 10, 12, 9, 3, 6, 4, 13, 8, 5},
	{9, 7, 5, 8, 6, 3, 4, 7, 11, 2, 12, 6, 5, 11, 9, 4, 3, 12, 5, 10, 6, 13, 8, 12, 2, 11, 7, 10, 4, 9, 8, 1, 12, 10},
	{3, 10, 8, 9, 13, 10, 9, 11, 4, 7, 6, 11, 10, 7, 4, 2, 12, 11, 6, 8, 5, 12, 9, 5, 7, 3, 5, 12, 1, 2, 12, 6, 4, 8},
	{3, 9, 4, 11, 7, 5, 2, 7, 3, 12, 7, 9, 11, 12, 4, 8, 9, 2, 10, 4, 6, 13, 10, 8, 5, 6, 12, 1, 8, 11, 6, 10, 5, 12},
	{6, 12, 1, 7, 12, 11, 9, 2, 11, 4, 6, 10, 4, 8, 5, 11, 9, 6, 8, 13, 4, 9, 12, 10, 7, 8, 5, 2, 3, 10, 12, 7, 3, 5},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 48.932(lined) + 15.264(scatter) = 64.195263%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 64.195(sym) + 0.071909*69.169(fg)*4.2 = 85.085578%
var Reels850 = slot.Reels5x{
	{7, 6, 11, 2, 8, 5, 12, 10, 6, 7, 9, 5, 8, 10, 9, 4, 12, 11, 1, 2, 7, 4, 11, 12, 3, 10, 12, 9, 3, 6, 4, 13, 8, 5},
	{9, 7, 5, 8, 6, 3, 4, 7, 11, 2, 12, 6, 5, 11, 9, 4, 3, 12, 5, 10, 6, 13, 8, 12, 2, 11, 7, 10, 4, 9, 8, 1, 12, 10},
	{3, 10, 8, 9, 13, 10, 9, 11, 4, 7, 6, 11, 10, 7, 4, 2, 12, 11, 6, 8, 5, 12, 9, 5, 7, 3, 5, 12, 1, 2, 12, 6, 4, 8},
	{10, 8, 7, 5, 4, 2, 12, 3, 8, 9, 3, 4, 10, 1, 12, 11, 5, 6, 8, 7, 13, 6, 11, 4, 7, 6, 10, 11, 5, 9, 2, 12, 9, 3},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 49.77(lined) + 15.264(scatter) = 65.033794%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 65.034(sym) + 0.071909*70.073(fg)*4.2 = 86.196982%
var Reels861 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{9, 7, 5, 8, 6, 3, 4, 7, 11, 2, 12, 6, 5, 11, 9, 4, 3, 12, 5, 10, 6, 13, 8, 12, 2, 11, 7, 10, 4, 9, 8, 1, 12, 10},
	{3, 10, 8, 9, 13, 10, 9, 11, 4, 7, 6, 11, 10, 7, 4, 2, 12, 11, 6, 8, 5, 12, 9, 5, 7, 3, 5, 12, 1, 2, 12, 6, 4, 8},
	{3, 9, 4, 11, 7, 5, 2, 7, 3, 12, 7, 9, 11, 12, 4, 8, 9, 2, 10, 4, 6, 13, 10, 8, 5, 6, 12, 1, 8, 11, 6, 10, 5, 12},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 51.263(lined) + 15.264(scatter) = 66.527144%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 66.527(sym) + 0.071909*71.682(fg)*4.2 = 88.176295%
var Reels881 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{9, 7, 5, 8, 6, 3, 4, 7, 11, 2, 12, 6, 5, 11, 9, 4, 3, 12, 5, 10, 6, 13, 8, 12, 2, 11, 7, 10, 4, 9, 8, 1, 12, 10},
	{3, 10, 8, 9, 13, 10, 9, 11, 4, 7, 6, 11, 10, 7, 4, 2, 12, 11, 6, 8, 5, 12, 9, 5, 7, 3, 5, 12, 1, 2, 12, 6, 4, 8},
	{10, 8, 7, 5, 4, 2, 12, 3, 8, 9, 3, 4, 10, 1, 12, 11, 5, 6, 8, 7, 13, 6, 11, 4, 7, 6, 10, 11, 5, 9, 2, 12, 9, 3},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 52.547(lined) + 15.264(scatter) = 67.810922%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 67.811(sym) + 0.071909*73.065(fg)*4.2 = 89.877838%
var Reels898 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{7, 12, 4, 9, 5, 2, 7, 6, 8, 1, 9, 6, 2, 12, 5, 8, 11, 13, 10, 3, 8, 4, 11, 5, 4, 7, 10, 9, 3, 12, 6, 3, 10, 11},
	{3, 10, 8, 9, 13, 10, 9, 11, 4, 7, 6, 11, 10, 7, 4, 2, 12, 11, 6, 8, 5, 12, 9, 5, 7, 3, 5, 12, 1, 2, 12, 6, 4, 8},
	{3, 9, 4, 11, 7, 5, 2, 7, 3, 12, 7, 9, 11, 12, 4, 8, 9, 2, 10, 4, 6, 13, 10, 8, 5, 6, 12, 1, 8, 11, 6, 10, 5, 12},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 54.065(lined) + 15.264(scatter) = 69.328439%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 69.328(sym) + 0.071909*74.7(fg)*4.2 = 91.889182%
var Reels918 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{7, 12, 4, 9, 5, 2, 7, 6, 8, 1, 9, 6, 2, 12, 5, 8, 11, 13, 10, 3, 8, 4, 11, 5, 4, 7, 10, 9, 3, 12, 6, 3, 10, 11},
	{3, 10, 8, 9, 13, 10, 9, 11, 4, 7, 6, 11, 10, 7, 4, 2, 12, 11, 6, 8, 5, 12, 9, 5, 7, 3, 5, 12, 1, 2, 12, 6, 4, 8},
	{10, 8, 7, 5, 4, 2, 12, 3, 8, 9, 3, 4, 10, 1, 12, 11, 5, 6, 8, 7, 13, 6, 11, 4, 7, 6, 10, 11, 5, 9, 2, 12, 9, 3},
	{6, 12, 1, 7, 12, 11, 9, 2, 11, 4, 6, 10, 4, 8, 5, 11, 9, 6, 8, 13, 4, 9, 12, 10, 7, 8, 5, 2, 3, 10, 12, 7, 3, 5},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 54.64(lined) + 15.264(scatter) = 69.903364%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 69.903(sym) + 0.071909*75.32(fg)*4.2 = 92.651199%
var Reels926 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{7, 12, 4, 9, 5, 2, 7, 6, 8, 1, 9, 6, 2, 12, 5, 8, 11, 13, 10, 3, 8, 4, 11, 5, 4, 7, 10, 9, 3, 12, 6, 3, 10, 11},
	{3, 10, 8, 9, 13, 10, 9, 11, 4, 7, 6, 11, 10, 7, 4, 2, 12, 11, 6, 8, 5, 12, 9, 5, 7, 3, 5, 12, 1, 2, 12, 6, 4, 8},
	{10, 8, 7, 5, 4, 2, 12, 3, 8, 9, 3, 4, 10, 1, 12, 11, 5, 6, 8, 7, 13, 6, 11, 4, 7, 6, 10, 11, 5, 9, 2, 12, 9, 3},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 56.472(lined) + 15.264(scatter) = 71.736031%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 71.736(sym) + 0.071909*77.294(fg)*4.2 = 95.080250%
var Reels950 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{7, 12, 4, 9, 5, 2, 7, 6, 8, 1, 9, 6, 2, 12, 5, 8, 11, 13, 10, 3, 8, 4, 11, 5, 4, 7, 10, 9, 3, 12, 6, 3, 10, 11},
	{9, 1, 8, 2, 9, 12, 4, 5, 7, 8, 6, 10, 12, 11, 10, 4, 2, 6, 8, 10, 3, 5, 11, 12, 5, 7, 13, 4, 9, 3, 11, 6, 3, 7},
	{3, 9, 4, 11, 7, 5, 2, 7, 3, 12, 7, 9, 11, 12, 4, 8, 9, 2, 10, 4, 6, 13, 10, 8, 5, 6, 12, 1, 8, 11, 6, 10, 5, 12},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 58.546(lined) + 15.264(scatter) = 73.809920%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 73.81(sym) + 0.071909*79.529(fg)*4.2 = 97.829019%
var Reels978 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{7, 12, 4, 9, 5, 2, 7, 6, 8, 1, 9, 6, 2, 12, 5, 8, 11, 13, 10, 3, 8, 4, 11, 5, 4, 7, 10, 9, 3, 12, 6, 3, 10, 11},
	{9, 1, 8, 2, 9, 12, 4, 5, 7, 8, 6, 10, 12, 11, 10, 4, 2, 6, 8, 10, 3, 5, 11, 12, 5, 7, 13, 4, 9, 3, 11, 6, 3, 7},
	{10, 8, 7, 5, 4, 2, 12, 3, 8, 9, 3, 4, 10, 1, 12, 11, 5, 6, 8, 7, 13, 6, 11, 4, 7, 6, 10, 11, 5, 9, 2, 12, 9, 3},
	{6, 12, 1, 7, 12, 11, 9, 2, 11, 4, 6, 10, 4, 8, 5, 11, 9, 6, 8, 13, 4, 9, 12, 10, 7, 8, 5, 2, 3, 10, 12, 7, 3, 5},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 59.344(lined) + 15.264(scatter) = 74.607403%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 74.607(sym) + 0.071909*80.388(fg)*4.2 = 98.886019%
var Reels988 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{7, 12, 4, 9, 5, 2, 7, 6, 8, 1, 9, 6, 2, 12, 5, 8, 11, 13, 10, 3, 8, 4, 11, 5, 4, 7, 10, 9, 3, 12, 6, 3, 10, 11},
	{9, 1, 8, 2, 9, 12, 4, 5, 7, 8, 6, 10, 12, 11, 10, 4, 2, 6, 8, 10, 3, 5, 11, 12, 5, 7, 13, 4, 9, 3, 11, 6, 3, 7},
	{10, 8, 7, 5, 4, 2, 12, 3, 8, 9, 3, 4, 10, 1, 12, 11, 5, 6, 8, 7, 13, 6, 11, 4, 7, 6, 10, 11, 5, 9, 2, 12, 9, 3},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 69.633(lined) + 15.264(scatter) = 84.896952%
// free spins 3267216, q = 0.071909, sq = 1/(1-q) = 1.077481
// free games frequency: 1/166.88
// RTP = 84.897(sym) + 0.071909*91.475(fg)*4.2 = 112.523975%
var Reels112 = slot.Reels5x{
	{10, 2, 13, 11, 5, 6, 7, 3, 9, 10, 4, 12, 7, 5, 10, 4, 9, 6, 5, 3, 6, 7, 12, 2, 9, 8, 11, 1, 3, 8, 11, 4, 12, 8},
	{7, 12, 4, 9, 5, 2, 7, 6, 8, 1, 9, 6, 2, 12, 5, 8, 11, 13, 10, 3, 8, 4, 11, 5, 4, 7, 10, 9, 3, 12, 6, 3, 10, 11},
	{7, 1, 9, 3, 10, 11, 12, 5, 3, 13, 2, 8, 6, 7, 8, 11, 10, 5, 4, 1, 12, 4, 9, 8, 6, 7, 12, 2, 5, 6, 9, 4, 10, 11},
	{10, 8, 7, 5, 4, 2, 12, 3, 8, 9, 3, 4, 10, 1, 12, 11, 5, 6, 8, 7, 13, 6, 11, 4, 7, 6, 10, 11, 5, 9, 2, 12, 9, 3},
	{8, 5, 11, 12, 7, 5, 12, 7, 9, 10, 4, 3, 2, 11, 3, 9, 8, 1, 6, 12, 9, 6, 11, 4, 13, 5, 10, 3, 2, 8, 10, 7, 6, 4},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	83.520645:  &Reels835,
	85.085578:  &Reels850,
	86.196982:  &Reels861,
	88.176295:  &Reels881,
	89.877838:  &Reels898,
	91.889182:  &Reels918,
	92.651199:  &Reels926,
	95.080250:  &Reels950,
	97.829019:  &Reels978,
	98.886019:  &Reels988,
	112.523975: &Reels112,
}
