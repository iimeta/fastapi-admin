package util

import (
	"math/rand"
	"slices"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
)

var node *snowflake.Node

func init() {

	var err error
	if node, err = snowflake.NewNode(rand.Int63n(1023)); err != nil {
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

			l := (length - len(prefix)) / len(ids)
			i := len(id)
			n := (l - i) / i
			k := ""

			for j := 0; j < i; j++ {
				k += gstr.Join(slices.Insert(strings.Split(grand.Letters(n), ""), grand.Intn(n), id[j:j+1]), "")
			}

			if len(k) < l {
				k += grand.Letters(l - len(k))
			}

			key += k
		}

		if len(key) < length {
			key += grand.Letters(length - len(key))
		}

	} else {
		key += grand.Letters(length - len(prefix))
	}

	return key
}
