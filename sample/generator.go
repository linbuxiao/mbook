package sample

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"mbook/pb"
)

func NewKeyBoard() (keyBoard *pb.KeyBoard) {
	keyBoard = &pb.KeyBoard{
		Layout:  randomKeyBoardLayout(),
		Backlit: randomBool(),
	}

	return
}

func NewCPU() (CPU *pb.CPU) {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	CPU = &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return
}

func NewGPU() (GPU *pb.GPU) {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	GPU = &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return
}

func NewRAM() (memory *pb.Memory) {
	memory = &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return
}

func NewSSD() (ssd *pb.Storage) {
	ssd = &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return
}

func NewHDD() (hdd *pb.Storage) {
	hdd = &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return
}

func NewScreen() (screen *pb.Screen) {
	screen = &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return
}

func NewLaptop() (laptop *pb.Laptop) {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop = &pb.Laptop{
		Id:    randomID(),
		Brand: brand,
		Name:  name,
		Ram:   NewRAM(),
		Gpus: []*pb.GPU{
			NewGPU(),
		},
		Cpu: NewCPU(),
		Storages: []*pb.Storage{
			NewSSD(),
			NewHDD(),
		},
		Screen:   NewScreen(),
		Keyboard: NewKeyBoard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}

	return
}
