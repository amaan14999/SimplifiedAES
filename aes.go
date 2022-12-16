package main

import "fmt"

type simplifiedAES struct {
	//16bit unsigned integer
	sBox        [16]uint16
	sBoxI       [16]uint16
	preRoundKey [4]uint16
	round1Key   [4]uint16
	round2Key   [4]uint16
}

func newSimplifiedAES(key uint16) *simplifiedAES {
	//initialised struct
	s := &simplifiedAES{
		sBox: [16]uint16{
			0x9, 0x4, 0xA, 0xB, 0xD, 0x1, 0x8, 0x5,
			0x6, 0x2, 0x0, 0x3, 0xC, 0xE, 0xF, 0x7,
		},
		sBoxI: [16]uint16{
			0xA, 0x5, 0x9, 0xB, 0x1, 0x7, 0x8, 0xF,
			0x6, 0x0, 0x2, 0x3, 0xC, 0x4, 0xD, 0xE,
		},
	}
	//set the keys for the rounds of encryption using the key expansion function
	s.preRoundKey, s.round1Key, s.round2Key = s.keyExpansion(key)

	return s
}

func (s *simplifiedAES) keyExpansion(key uint16) ([4]uint16, [4]uint16, [4]uint16) {
	rCon1, rCon2 := uint16(0x80), uint16(0x30)

	w := [6]uint16{}
	w[0] = (key & 0xFF00) >> 8
	w[1] = key & 0x00FF
	w[2] = w[0] ^ (s.subWord(s.rotWord(w[1])) ^ rCon1)
	w[3] = w[2] ^ w[1]
	w[4] = w[2] ^ (s.subWord(s.rotWord(w[3])) ^ rCon2)
	w[5] = w[4] ^ w[3]

	return s.intToState((w[0] << 8) + w[1]),
		s.intToState((w[2] << 8) + w[3]),
		s.intToState((w[4] << 8) + w[5])
}

func (s *simplifiedAES) rotWord(word uint16) uint16 {
	return ((word & 0xF) << 4) + ((word & 0xF0) >> 4)
}

//sub word subtitutes the nibbles of the word using the sbox
func (s *simplifiedAES) subWord(word uint16) uint16 {
	return (s.sBox[word>>4] << 4) + s.sBox[word&0xF]
}

func (s *simplifiedAES) intToState(n uint16) [4]uint16 {
	return [4]uint16{
		n >> 12 & 0xF,
		(n >> 4) & 0xF,
		(n >> 8) & 0xF,
		n & 0xF,
	}
}

func (s *simplifiedAES) addRoundKey(key1, key2 [4]uint16) [4]uint16 {
	return [4]uint16{
		key1[0] ^ key2[0],
		key1[1] ^ key2[1],
		key1[2] ^ key2[2],
		key1[3] ^ key2[3],
	}
}

func (s *simplifiedAES) gfMult(a, b uint16) uint16 {
	var product uint16

	a = a & 0xF
	b = b & 0xF

	for i := 0; i < 4; i++ {
		if b&1 == 1 {
			product ^= a
		}

		a <<= 1
		if a&(1<<4) != 0 {
			a ^= 0b10011
		}

		b >>= 1
	}

	return product
}

func (s *simplifiedAES) mixColumns(state [4]uint16) [4]uint16 {
	return [4]uint16{
		state[0] ^ s.gfMult(4, state[2]),
		state[1] ^ s.gfMult(4, state[3]),
		state[2] ^ s.gfMult(4, state[0]),
		state[3] ^ s.gfMult(4, state[1]),
	}
}

func (s *simplifiedAES) shiftRows(state [4]uint16) [4]uint16 {
	return [4]uint16{state[0], state[1], state[3], state[2]}
}

func (s *simplifiedAES) subNibbles(sBox [16]uint16, state [4]uint16) [4]uint16 {
	return [4]uint16{
		sBox[state[0]],
		sBox[state[1]],
		sBox[state[2]],
		sBox[state[3]],
	}
}

func (s *simplifiedAES) stateToInt(state [4]uint16) uint16 {
	return (state[0] << 12) +
		(state[2] << 8) +
		(state[1] << 4) +
		state[3]
}

func (s *simplifiedAES) inverseMixColumns(state [4]uint16) [4]uint16 {
	return [4]uint16{
		s.gfMult(9, state[0]) ^ s.gfMult(2, state[2]),
		s.gfMult(9, state[1]) ^ s.gfMult(2, state[3]),
		s.gfMult(9, state[2]) ^ s.gfMult(2, state[0]),
		s.gfMult(9, state[3]) ^ s.gfMult(2, state[1]),
	}
}

func (s *simplifiedAES) encrypt(plaintext uint16) uint16 {
	state := s.addRoundKey(s.preRoundKey, s.intToState(plaintext))
	state = s.mixColumns(s.shiftRows(s.subNibbles(s.sBox, state)))
	state = s.addRoundKey(s.round1Key, state)
	state = s.shiftRows(s.subNibbles(s.sBox, state))
	state = s.addRoundKey(s.round2Key, state)
	return s.stateToInt(state)
}

func (s *simplifiedAES) decrypt(ciphertext uint16) uint16 {
	state := s.addRoundKey(s.round2Key, s.intToState(ciphertext))
	state = s.subNibbles(s.sBoxI, s.shiftRows(state))
	state = s.inverseMixColumns(s.addRoundKey(s.round1Key, state))
	state = s.subNibbles(s.sBoxI, s.shiftRows(state))
	state = s.addRoundKey(s.preRoundKey, state)
	return s.stateToInt(state)
}

func main() {
	key := uint16(0b0011001001000010)
	plaintext := uint16(0b1100110011010000)

	s := newSimplifiedAES(key)
	ciphertext := s.encrypt(plaintext)

	fmt.Printf("Encrypted: %b\n", ciphertext)
	fmt.Printf("Decrypted: %b\n", s.decrypt(ciphertext))
}
