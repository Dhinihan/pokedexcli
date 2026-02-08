package pokeapicache

import (
	"fmt"
	"sync"
	"time"
)

func NewCache(ttl time.Duration) *Cache {
	c := Cache{
		registros: map[string]cacheRegistro{},
	}
	c.varredura(ttl)
	return &c
}

type cacheRegistro struct {
	criacao time.Time
	valor   []byte
}

type Cache struct {
	registros map[string]cacheRegistro
	mutex     sync.Mutex
}

func (c *Cache) Add(chave string, valor []byte) {
	registro := cacheRegistro{
		criacao: time.Now(),
		valor:   valor,
	}
	c.mutex.Lock()
	c.registros[chave] = registro
	c.mutex.Unlock()
}

func (c *Cache) Get(chave string) ([]byte, bool) {
	c.mutex.Lock()
	registro, ok := c.registros[chave]
	c.mutex.Unlock()
	return registro.valor, ok
}

func (c *Cache) varredura(duracao time.Duration) {
	ticker := time.Tick(duracao)
	go func() {
		for now := range ticker {
			fmt.Printf("tick %v\n", now)
			for key, r := range c.registros {
				if r.criacao.Add(duracao).Compare(time.Now()) < 0 {
					delete(c.registros, key)
				}
			}
		}
	}()
}
