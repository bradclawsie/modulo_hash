package modulo_hash

import (
	"testing"
	"strconv"
	"fmt"
	"math/rand"
	"encoding/base64"
)

func TestMHash(t *testing.T) {
	mh := NewModuloHash()
	items := []string{"127.0.0.1","17.0.1.1","1.1.0.1","27.99.0.111","64.0.8.8","8.8.8.8","10.100.0.100",
		"128.4.4.4","28.28.1.1","28.10.0.10","12.9.0.10","11.11.8.1","13.10.0.19","128.19.19.1"}
	mh.SetTargets(items)
	// test distribution
	dist := make(map [string] int)
	total := 10000
	for j := 0; j < total; j++ {
		b64 := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(rand.Int31()))))
		tgt,tgt_err := mh.Find(b64)
		if tgt_err != nil || tgt == "" {
			e := fmt.Sprintf("err returned on finding target for %s",b64)
			t.Errorf(e)
		}
		dist[tgt]++
	}
	for k,v := range dist {
		fmt.Printf("%s %d (%f pct) \n",k,v,(float64(v)/float64(total))*100)
	}
}

func TestEmpty(t *testing.T) {
	mh := NewModuloHash()
	tgt,tgt_err := mh.Find("hello")
	if tgt_err == nil || tgt != "" {
		t.Errorf("no error on empty target list")
	}
	mh.SetTargets([]string{"hi","there","how","are","you"})
	tgt,tgt_err = mh.Find("")
	if tgt_err == nil || tgt != "" {
		t.Errorf("no error on empty target list")
	}	
}