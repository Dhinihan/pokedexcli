package pokeapicache

import (
	"fmt"
	"testing"
	"time"
)

type testCase struct {
	chave string
	valor []byte
}

func TestAddGet(t *testing.T) {
	cases := []testCase{
		{"chave1", []byte("valor1")},
		{"chave2", []byte{}},
	}
	for ix, tc := range cases {
		t.Run(fmt.Sprintf("Test %d", ix), func(tt *testing.T) {
			tt.Parallel()
			cache := NewCache(1 * time.Second)
			_, ok := cache.Get(tc.chave)
			if ok {
				tt.Errorf("Não deveria ter achado %s", tc.chave)
			}
			cache.Add(tc.chave, tc.valor)
			encontrado, ok := cache.Get(tc.chave)
			if !ok {
				tt.Errorf("Deveria ter achado %s", tc.chave)
			}
			if string(encontrado) != string(tc.valor) {
				tt.Errorf("'%s' deveria ser igual a '%s'", encontrado, tc.valor)
			}
		})
	}
}
func TestVarredura(t *testing.T) {
	cases := []testCase{
		{"chave1", []byte("valor1")},
		{"chave2", []byte{}},
	}
	for ix, tc := range cases {
		t.Run(fmt.Sprintf("Test %d", ix), func(tt *testing.T) {
			tt.Parallel()
			cache := NewCache(2 * time.Millisecond)
			cache.Add(tc.chave, tc.valor)
			encontrado, ok := cache.Get(tc.chave)
			if !ok {
				tt.Errorf("Deveria ter achado %s", tc.chave)
			}
			if string(encontrado) != string(tc.valor) {
				tt.Errorf("'%s' deveria ser igual a '%s'", encontrado, tc.valor)
			}
			time.Sleep(3 * time.Millisecond)
			_, ok = cache.Get(tc.chave)
			if ok {
				tt.Errorf("Deveria ter apagado %s", tc.chave)
			}
		})
	}
}
