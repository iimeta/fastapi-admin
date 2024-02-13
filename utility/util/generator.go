package util

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"slices"
	"strings"
)

var node *snowflake.Node

func init() {

	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

}

func GenerateId() string {
	return node.Generate().String()
}

func NewKey(prefix string, length int, ids ...string) string {

	key := prefix

	if len(ids) > 0 {

		for _, id := range ids {

			n := (length - len(prefix)) / len(ids)
			l := len(id)

			n = (n - l) / l

			for i := 0; i < l; i++ {
				key += gstr.Join(slices.Insert(strings.Split(grand.Letters(n), ""), grand.Intn(n), id[i:i+1]), "")
			}
		}

		if len(key) < length {
			key += grand.Letters(length - len(key))
		}

	} else {
		key += grand.Letters(length - len(prefix))
	}

	return key
}
