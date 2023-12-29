package smux

import (
	"runtime"
	"math/rand"
	"time"
	"crypto/sha256"
)

type Keyring struct {
	p    string
	// keys map[string][]byte
}

func NewKeyring(s string) *Keyring {
	Keyring := &Keyring{
		p:    s,
	}
	return Keyring
}

func (k *Keyring) Extract(iv []byte, s string) []byte {
	tail := []byte(k.p + s)
	b := make([]byte, len(iv)+len(tail))
	copy(b[:len(iv)], iv)
	copy(b[len(iv):], tail)
	return SHA256(b)
}

func SHA256(b []byte) []byte {
	s := sha256.Sum256(b)
	return s[:]
}


func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func ChunkBytes(b []byte, chunkSize int) [][]byte {
	if chunkSize <= 0 {
		return nil
	}
	var chunks [][]byte
	for i := 0; i < len(b); i += chunkSize {
		end := i + chunkSize
		if end > len(b) {
			end = len(b)
		}
		chunks = append(chunks, b[i:end])
	}
	return chunks
}

func randomASCIIChar() byte {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(95) + 32
	return byte(randomInt)
}

func XORBytes(a, b []byte) []byte {
	buf := make([]byte, len(a))
	for i, _ := range a {
		buf[i] = a[i] ^ b[i]
	}
	return buf
}

func increment(b *[24]byte) {
	for i := range b {
		b[i]++
		if b[i] != 0 {
			return
		}
	}
}
