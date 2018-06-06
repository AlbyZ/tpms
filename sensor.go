package tpms

import (
	"encoding/binary"
	"github.com/go-ble/ble"
)

type Sensor struct {
	Id         int
	Address    ble.Addr
	Kilopascal int
	Psi int
	Celsius    int
}

func (this *Sensor) ParseData(b []byte) {
	// Bytes 8 to 11 are pressure in kPa.
	this.Kilopascal = int(binary.LittleEndian.Uint32(b[8:]) / 1000)
	// Bytes 8 to 11 are pressure in kPa.
	this.Psi = int(binary.LittleEndian.Uint32(b[8:]) / 6894.757280343134)
	// Bytes 12 to 15 are temperature in Celsius.
	this.Celsius = int(binary.LittleEndian.Uint32(b[12:]) / 100)
}
