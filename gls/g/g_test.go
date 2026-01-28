package g

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestG1(t *testing.T) {
	gp1 := G()
	assert.True(t, gp1 != nil)

	t.Run("G in another goroutine", func(t *testing.T) {
		gp2 := G()
		assert.True(t, gp2 != nil)
		assert.True(t, gp1 != gp2)
	})

	gType := reflect.TypeOf(G0())
	sf, ss := gType.FieldByName("labels")
	assert.True(t, ss && sf.Offset > 0)
}

func TestG2(t *testing.T) {
	gp1 := G()

	if gp1 == nil {
		t.Fatalf("fail to get G.")
	}

	t.Run("G in another goroutine", func(t *testing.T) {
		gp2 := G()

		if gp2 == nil {
			t.Fatalf("fail to get G.")
		}

		if gp2 == gp1 {
			t.Fatalf("every living G must be different. [gp1:%p] [gp2:%p]", gp1, gp2)
		}
	})
}

func TestUnsafe(t *testing.T) {
	type Header struct {
		Id    int64
		Magic uint32
		Len   uint32
	}

	hdr := Header{Id: 123, Magic: 9999, Len: 8}
	p := unsafe.Pointer(&hdr)
	buf := (*[8]byte)(unsafe.Pointer(&hdr))[:]
	fmt.Printf("sizeOf hdr.Magic: %v\n", unsafe.Sizeof(hdr.Magic))
	fmt.Printf("buf: %v\n", buf)

	idPtr := unsafe.Pointer(uintptr(p))
	fmt.Printf("point addr: %v\n", uintptr(p))
	fmt.Printf("idPtr: %v\n", idPtr)
	fmt.Printf("Id: %v\n", *(*int64)(idPtr))

	magicPtr := unsafe.Pointer(uintptr(p) + unsafe.Sizeof(hdr.Id))
	fmt.Printf("magicPtr: %v\n", magicPtr)
	fmt.Printf("Magic: %v\n", *(*uint32)(magicPtr))

	lenPtr := unsafe.Pointer(uintptr(p) + unsafe.Sizeof(hdr.Id) + unsafe.Sizeof(hdr.Magic))
	fmt.Printf("lenPtr: %v\n", lenPtr)
	fmt.Printf("Len: %v\n", *(*uint32)(lenPtr))
}

func TestGoroutineG(t *testing.T) {
	gp1 := G()
	assert.True(t, gp1 != nil)
	fmt.Printf("gp1: %v\n", gp1)

	t.Run("G in another goroutine", func(t *testing.T) {
		gp2 := G()
		fmt.Printf("gp2: %v\n", gp2)
		assert.True(t, gp2 != nil)
		assert.True(t, gp1 != gp2)
	})
}
