package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey        = "rank"     // 每日排名
	ElectricalRank = "elecRank" // 家电排名
	AccessoryRank  = "acceRank" // 配件排名
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
