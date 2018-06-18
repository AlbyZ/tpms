package tpms

import (
	"encoding/binary"
	"github.com/go-ble/ble"
)

type Sensor struct {
	Id         int
	Address    ble.Addr
	Kilopascal int
	Psi 	   float32
	Celsius    float32
	Percentage int
}

func (this *Sensor) ParseData(b []byte) {
	// Bytes 8 to 11 are pressure in kPa.
	this.Kilopascal = int(binary.LittleEndian.Uint32(b[8:]) / 1000)
	// Bytes 8 to 11 are pressure in kPa.
	this.Psi = float32(binary.LittleEndian.Uint32(b[8:]) / 6894)
	// Bytes 12 to 15 are temperature in Celsius.
	this.Celsius = float32(binary.LittleEndian.Uint32(b[12:]) / 100)
	// Bytes 16 to 17 are Battery Charge in Percentage.
	this.Percentage = int(binary.LittleEndian.Uint16(b[16:]))
}
