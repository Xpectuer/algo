package basic

import (
	"math/rand"
	"strings"
	"time"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
func randomKey() string {
	var key strings.Builder
	for i := 0; i < 32; i++ {

		rand.Seed(time.Now().UnixNano())
		b := byte(rand.Intn(26) + 65)
		key.WriteByte(b)

	}
	return key.String()
}
