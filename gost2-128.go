// go run gost2-128.go

// go build gost2-128.go

package main

import (
	"fmt"
)

const n1 = 512 // 4096-bit key -> 64 * 64-bit subkeys

type word64 = uint64

var (
	x1, x2 int
	h2     [n1]byte
	h1     [n1 * 3]byte
)

// init sets everything to zero
func initState() {
	x1, x2 = 0, 0
	for i := 0; i < n1; i++ {
		h2[i] = 0
		h1[i] = 0
	}
}

func hashing(t1 []byte, b6 int) {
	s4 := [256]byte{
		13, 199, 11, 67, 237, 193, 164, 77, 115, 184, 141, 222, 73,
		38, 147, 36, 150, 87, 21, 104, 12, 61, 156, 101, 111, 145,
		119, 22, 207, 35, 198, 37, 171, 167, 80, 30, 219, 28, 213,
		121, 86, 29, 214, 242, 6, 4, 89, 162, 110, 175, 19, 157,
		3, 88, 234, 94, 144, 118, 159, 239, 100, 17, 182, 173, 238,
		68, 16, 79, 132, 54, 163, 52, 9, 58, 57, 55, 229, 192,
		170, 226, 56, 231, 187, 158, 70, 224, 233, 245, 26, 47, 32,
		44, 247, 8, 251, 20, 197, 185, 109, 153, 204, 218, 93, 178,
		212, 137, 84, 174, 24, 120, 130, 149, 72, 180, 181, 208, 255,
		189, 152, 18, 143, 176, 60, 249, 27, 227, 128, 139, 243, 253,
		59, 123, 172, 108, 211, 96, 138, 10, 215, 42, 225, 40, 81,
		65, 90, 25, 98, 126, 154, 64, 124, 116, 122, 5, 1, 168,
		83, 190, 131, 191, 244, 240, 235, 177, 155, 228, 125, 66, 43,
		201, 248, 220, 129, 188, 230, 62, 75, 71, 78, 34, 31, 216,
		254, 136, 91, 114, 106, 46, 217, 196, 92, 151, 209, 133, 51,
		236, 33, 252, 127, 179, 69, 7, 183, 105, 146, 97, 39, 15,
		205, 112, 200, 166, 223, 45, 48, 246, 186, 41, 148, 140, 107,
		76, 85, 95, 194, 142, 50, 49, 134, 23, 135, 169, 221, 210,
		203, 63, 165, 82, 161, 202, 53, 14, 206, 232, 103, 102, 195,
		117, 250, 99, 0, 74, 160, 241, 2, 113,
	}
	b4 := 0
	for b6 > 0 {
		for b6 > 0 && x2 < n1 {
			b5 := t1[b4]
			b4++
			h1[x2+n1] = b5
			h1[x2+(n1*2)] = b5 ^ h1[x2]
			x1 = int(h2[x2] ^ s4[byte(b5^byte(x1))])
			h2[x2] = byte(x1)
			x2++
			b6--
		}
		if x2 == n1 {
			b2 := 0
			x2 = 0
			for b3 := 0; b3 < n1+2; b3++ {
				for b1 := 0; b1 < n1*3; b1++ {
					b2 = int(h1[b1] ^ s4[byte(b2)])
					h1[b1] = byte(b2)
				}
				b2 = (b2 + b3) % 256
			}
		}
	}
}

func endHash(h4 *[n1]byte) {
	var h3 [n1]byte
	n4 := n1 - x2
	for i := 0; i < n4; i++ {
		h3[i] = byte(n4)
	}
	hashing(h3[:], n4)
	hashing(h2[:], len(h2))
	for i := 0; i < n1; i++ {
		h4[i] = h1[i]
	}
}

func createKeys(h4 *[n1]byte, key *[64]word64) {
	k := 0
	for i := 0; i < 64; i++ {
		key[i] = 0
		for z := 0; z < 8; z++ {
			key[i] = (key[i] << 8) + word64(h4[k]&0xff)
			k++
		}
	}
}

// Substitution tables
var (
	k1 = [16]byte{0x4, 0xA, 0x9, 0x2, 0xD, 0x8, 0x0, 0xE, 0x6, 0xB, 0x1, 0xC, 0x7, 0xF, 0x5, 0x3}
	k2 = [16]byte{0xE, 0xB, 0x4, 0xC, 0x6, 0xD, 0xF, 0xA, 0x2, 0x3, 0x8, 0x1, 0x0, 0x7, 0x5, 0x9}
	k3 = [16]byte{0x5, 0x8, 0x1, 0xD, 0xA, 0x3, 0x4, 0x2, 0xE, 0xF, 0xC, 0x7, 0x6, 0x0, 0x9, 0xB}
	k4 = [16]byte{0x7, 0xD, 0xA, 0x1, 0x0, 0x8, 0x9, 0xF, 0xE, 0x4, 0x6, 0xC, 0xB, 0x2, 0x5, 0x3}
	k5 = [16]byte{0x6, 0xC, 0x7, 0x1, 0x5, 0xF, 0xD, 0x8, 0x4, 0xA, 0x9, 0xE, 0x0, 0x3, 0xB, 0x2}
	k6 = [16]byte{0x4, 0xB, 0xA, 0x0, 0x7, 0x2, 0x1, 0xD, 0x3, 0x6, 0x8, 0x5, 0x9, 0xC, 0xF, 0xE}
	k7 = [16]byte{0xD, 0xB, 0x4, 0x1, 0x3, 0xF, 0x5, 0x9, 0x0, 0xA, 0xE, 0x7, 0x6, 0x8, 0x2, 0xC}
	k8 = [16]byte{0x1, 0xF, 0xD, 0x0, 0x5, 0x7, 0xA, 0x4, 0x9, 0x2, 0x3, 0xE, 0x6, 0xB, 0x8, 0xC}
	k9 = [16]byte{0xC, 0x4, 0x6, 0x2, 0xA, 0x5, 0xB, 0x9, 0xE, 0x8, 0xD, 0x7, 0x0, 0x3, 0xF, 0x1}
	k10 = [16]byte{0x6, 0x8, 0x2, 0x3, 0x9, 0xA, 0x5, 0xC, 0x1, 0xE, 0x4, 0x7, 0xB, 0xD, 0x0, 0xF}
	k11 = [16]byte{0xB, 0x3, 0x5, 0x8, 0x2, 0xF, 0xA, 0xD, 0xE, 0x1, 0x7, 0x4, 0xC, 0x9, 0x6, 0x0}
	k12 = [16]byte{0xC, 0x8, 0x2, 0x1, 0xD, 0x4, 0xF, 0x6, 0x7, 0x0, 0xA, 0x5, 0x3, 0xE, 0x9, 0xB}
	k13 = [16]byte{0x7, 0xF, 0x5, 0xA, 0x8, 0x1, 0x6, 0xD, 0x0, 0x9, 0x3, 0xE, 0xB, 0x4, 0x2, 0xC}
	k14 = [16]byte{0x5, 0xD, 0xF, 0x6, 0x9, 0x2, 0xC, 0xA, 0xB, 0x7, 0x8, 0x1, 0x4, 0x3, 0xE, 0x0}
	k15 = [16]byte{0x8, 0xE, 0x2, 0x5, 0x6, 0x9, 0x1, 0xC, 0xF, 0x4, 0xB, 0x0, 0xD, 0xA, 0x3, 0x7}
	k16 = [16]byte{0x1, 0x7, 0xE, 0xD, 0x0, 0x5, 0x8, 0x3, 0x4, 0xF, 0xA, 0x6, 0x9, 0xC, 0xB, 0x2}

	k175, k153, k131, k109, k87, k65, k43, k21 [256]byte
)

func kboxinit() {
	for i := 0; i < 256; i++ {
		k175[i] = k16[i>>4]<<4 | k15[i&15]
		k153[i] = k14[i>>4]<<4 | k13[i&15]
		k131[i] = k12[i>>4]<<4 | k11[i&15]
		k109[i] = k10[i>>4]<<4 | k9[i&15]
		k87[i] = k8[i>>4]<<4 | k7[i&15]
		k65[i] = k6[i>>4]<<4 | k5[i&15]
		k43[i] = k4[i>>4]<<4 | k3[i&15]
		k21[i] = k2[i>>4]<<4 | k1[i&15]
	}
}

func f(x word64) word64 {
	y := x >> 32
	z := x & 0xffffffff
	y = word64(k87[(y>>24)&255])<<24 | word64(k65[(y>>16)&255])<<16 |
		word64(k43[(y>>8)&255])<<8 | word64(k21[y&255])
	z = word64(k175[(z>>24)&255])<<24 | word64(k153[(z>>16)&255])<<16 |
		word64(k131[(z>>8)&255])<<8 | word64(k109[z&255])
	x = (y << 32) | (z & 0xffffffff)
	return (x << 11) | (x >> (64 - 11))
}

func gostcrypt(in [2]word64, key [64]word64) (out [2]word64) {
	ng1, ng2 := in[0], in[1]
	k := 0
	for i := 0; i < 32; i++ {
		ng2 ^= f(ng1 + key[k])
		k++
		ng1 ^= f(ng2 + key[k])
		k++
	}
	out[0] = ng2
	out[1] = ng1
	return
}

func gostdecrypt(in [2]word64, key [64]word64) (out [2]word64) {
	ng1, ng2 := in[0], in[1]
	k := 63
	for i := 0; i < 32; i++ {
		ng2 ^= f(ng1 + key[k])
		k--
		ng1 ^= f(ng2 + key[k])
		k--
	}
	out[0] = ng2
	out[1] = ng1
	return
}

func printHex(label string, v [2]word64) {
	fmt.Printf("%s %016X%016X\n", label, v[0], v[1])
}

func main() {
	kboxinit()
	fmt.Println("GOST2-128 by Alexander PUKALL 2016")
	fmt.Println("128-bit block, 4096-bit subkeys, 64 rounds")
	fmt.Println("Public domain implementation\n")

	var key [64]word64
	var plain, cipher, decrypted [2]word64
	var h4 [n1]byte

	// Example 1
	initState()
	pass := []byte("My secret password!0123456789abc")
	hashing(pass, len(pass))
	endHash(&h4)
	createKeys(&h4, &key)
	plain[0], plain[1] = 0xFEFEFEFEFEFEFEFE, 0xFEFEFEFEFEFEFEFE
	fmt.Println("Key 1:", string(pass))
	printHex("Plaintext  1:", plain)
	cipher = gostcrypt(plain, key)
	printHex("Encryption 1:", cipher)
	decrypted = gostdecrypt(cipher, key)
	printHex("Decryption 1:", decrypted)
	fmt.Println()

	// Example 2
	initState()
	pass = []byte("My secret password!0123456789ABC")
	hashing(pass, len(pass))
	endHash(&h4)
	createKeys(&h4, &key)
	plain[0], plain[1] = 0x0, 0x0
	fmt.Println("Key 2:", string(pass))
	printHex("Plaintext  2:", plain)
	cipher = gostcrypt(plain, key)
	printHex("Encryption 2:", cipher)
	decrypted = gostdecrypt(cipher, key)
	printHex("Decryption 2:", decrypted)
	fmt.Println()

	// Example 3
	initState()
	pass = []byte("My secret password!0123456789abZ")
	hashing(pass, len(pass))
	endHash(&h4)
	createKeys(&h4, &key)
	plain[0], plain[1] = 0x0, 0x1
	fmt.Println("Key 3:", string(pass))
	printHex("Plaintext  3:", plain)
	cipher = gostcrypt(plain, key)
	printHex("Encryption 3:", cipher)
	decrypted = gostdecrypt(cipher, key)
	printHex("Decryption 3:", decrypted)
}
