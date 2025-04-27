package util

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
)

var node *snowflake.Node

func InitId() {
	n, _ := snowflake.NewNode(1)
	node = n
	return
}

// 生成ID
func GenID() string {
	return strconv.FormatInt(int64(node.Generate()), 10)
}
