/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
package basic

import (
	"testing"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

func TestRandom(t *testing.T) {
	key := randomKey()
	t.Log(key)
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	swap(&a, &b)
	t.Log(a, b)
}
