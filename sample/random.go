package sample

import (
	"github.com/google/uuid"
	"math/rand"
	"mbook/pb"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyBoardLayout() pb.KeyBoard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.KeyBoard_AZERTY
	case 2:
		return pb.KeyBoard_QWERTY
	case 3:
		return pb.KeyBoard_QWERTZ
	default:
		return pb.KeyBoard_UNKNOWN
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon 1",
			"Core 2",
			"Core 3",
			"Core 4",
		)
	}

	return randomStringFromSet(
		"Ryzen 1",
		"Ryzen 2",
		"Ryzen 3",
	)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return ((max - min) * rand.Float64()) + min
}

func randomFloat32(min float32, max float32) float32 {
	return ((max - min) * rand.Float32()) + min
}

func randomGPUBrand() (brand string) {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) (name string) {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 1",
			"RTX 2",
			"RTX 3",
		)
	}

	return randomStringFromSet(
		"RX 1",
		"RX 2",
		"RX 3",
		"RX 4",
	)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}

	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomFloat32(1080, 4320)
	width := height / 16 * 9

	return &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Voro", "XPS", "Alienware")
	default:
		return randomStringFromSet("ThinkPad 1", "ThinkPad 2", "ThinkPad 3")
	}
}
