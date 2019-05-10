/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package arithmetic

import (
	"github.com/unidoc/unidoc/common"
	"github.com/unidoc/unidoc/pdf/internal/jbig2/reader"
	"io"
	"math"
)

var (
	qe = [][4]uint32{{0x5601, 1, 1, 1}, {0x3401, 2, 6, 0},
		{0x1801, 3, 9, 0}, {0x0AC1, 4, 12, 0}, {0x0521, 5, 29, 0}, {0x0221, 38, 33, 0},
		{0x5601, 7, 6, 1}, {0x5401, 8, 14, 0}, {0x4801, 9, 14, 0}, {0x3801, 10, 14, 0},
		{0x3001, 11, 17, 0}, {0x2401, 12, 18, 0}, {0x1C01, 13, 20, 0},
		{0x1601, 29, 21, 0}, {0x5601, 15, 14, 1}, {0x5401, 16, 14, 0},
		{0x5101, 17, 15, 0}, {0x4801, 18, 16, 0}, {0x3801, 19, 17, 0},
		{0x3401, 20, 18, 0}, {0x3001, 21, 19, 0}, {0x2801, 22, 19, 0},
		{0x2401, 23, 20, 0}, {0x2201, 24, 21, 0}, {0x1C01, 25, 22, 0},
		{0x1801, 26, 23, 0}, {0x1601, 27, 24, 0}, {0x1401, 28, 25, 0},
		{0x1201, 29, 26, 0}, {0x1101, 30, 27, 0}, {0x0AC1, 31, 28, 0},
		{0x09C1, 32, 29, 0}, {0x08A1, 33, 30, 0}, {0x0521, 34, 31, 0},
		{0x0441, 35, 32, 0}, {0x02A1, 36, 33, 0}, {0x0221, 37, 34, 0},
		{0x0141, 38, 35, 0}, {0x0111, 39, 36, 0}, {0x0085, 40, 37, 0},
		{0x0049, 41, 38, 0}, {0x0025, 42, 39, 0}, {0x0015, 43, 40, 0},
		{0x0009, 44, 41, 0}, {0x0005, 45, 42, 0}, {0x0001, 45, 43, 0},
		{0x5601, 46, 46, 0}}

	qeTable = []int{0x56010000, 0x34010000, 0x18010000, 0x0AC10000, 0x05210000, 0x02210000, 0x56010000, 0x54010000, 0x48010000, 0x38010000, 0x30010000, 0x24010000, 0x1C010000, 0x16010000, 0x56010000, 0x54010000, 0x51010000, 0x48010000, 0x38010000, 0x34010000, 0x30010000, 0x28010000, 0x24010000, 0x22010000, 0x1C010000, 0x18010000, 0x16010000, 0x14010000, 0x12010000, 0x11010000, 0x0AC10000, 0x09C10000, 0x08A10000, 0x05210000, 0x04410000, 0x02A10000, 0x02210000, 0x01410000, 0x01110000, 0x00850000, 0x00490000, 0x00250000, 0x00150000, 0x00090000, 0x00050000, 0x00010000,
		0x56010000}

	nmpsTable = []int{
		1, 2, 3, 4, 5, 38, 7, 8, 9, 10, 11, 12, 13, 29, 15,
		16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
		44, 45, 45, 46,
	}

	nlpsTable = []int{
		1, 6, 9, 12, 29, 33, 6, 14, 14, 14, 17, 18, 20, 21, 14, 14, 15, 16, 17, 18, 19,
		19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
		39, 40, 41, 42, 43, 46,
	}

	switchTable = []int{
		1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
)

// Decoder is the arithmetic Decoder structure that is used to decode the
// segments in an arithmetic method.
type Decoder struct {
	ContextSize          []int
	ReferedToContextSize []int

	r reader.StreamReader
	b int

	c        uint64
	a        uint32
	previous int64

	ct int

	prvCtr int

	streamPosition int64
}

// New creates new arithmetic Decoder
func New(r reader.StreamReader) (*Decoder, error) {
	d := &Decoder{
		r:                    r,
		ContextSize:          []int{16, 13, 10, 10},
		ReferedToContextSize: []int{13, 10},
	}

	if err := d.init(); err != nil {
		return nil, err
	}

	return d, nil
}

// DecodeBit decodes the arithmetic Bit
func (d *Decoder) DecodeBit(stats *DecoderStats) (int, error) {
	var (
		bit     int
		qeValue = qe[stats.cx()][0]
		icx     = int(stats.cx())
	)

	defer func() {

		d.prvCtr++
		// common.Log.Trace("Decoder 'a' value: %b", d.a)
		// common.Log.Trace("%d, D: %01b C: %08X A: %04X, CTR: %d, B: %02X  QE: %04X", d.prvCtr, bit, d.c, d.a, d.ct, d.b, qeValue)
	}()

	d.a -= qeValue

	// common.Log.Trace("Icx: %d, qeValue: %d", icx, qeValue)

	if (d.c >> 16) < uint64(qeValue) {
		// common.Log.Trace("d.c >> 16 < qeValue")
		bit = d.lpsExchange(stats, icx, qeValue)

		if err := d.renormalize(); err != nil {
			return 0, err
		}
	} else {
		// common.Log.Trace("d.c >> 16 >= qeValue")
		d.c -= (uint64(qeValue) << 16)

		if (d.a & 0x8000) == 0 {
			// common.Log.Trace("mpsExchange, renormalize")
			bit = d.mpsExchange(stats, icx)
			if err := d.renormalize(); err != nil {
				return 0, err
			}
		} else {
			// common.Log.Trace("stats.getMPS")
			bit = int(stats.getMps())
		}
	}

	// if bit > 0 {
	// 	bit = 1
	// }

	return bit, nil
}

// DecodeInt decodes the Integer from the arithmetic Decoder
func (d *Decoder) DecodeInt(stats *DecoderStats) (int, error) {
	var (
		value, bit, s, bitsToRead, offset int
		err                               error
	)
	if stats == nil {
		common.Log.Trace("Nil stats")
		stats = NewStats(512, 1)
	}

	d.previous = 1

	// first bit defines the sign of the integer
	s, err = d.decodeIntBit(stats)
	if err != nil {
		return 0, err
	}

	// common.Log.Trace("Sign int bit: '%01b'", s)

	bit, err = d.decodeIntBit(stats)
	if err != nil {
		return 0, err
	}

	// common.Log.Trace("First bit value: %b", bit)
	// First read
	if bit == 1 {
		bit, err = d.decodeIntBit(stats)
		if err != nil {
			return 0, err
		}

		// Second Read
		if bit == 1 {

			bit, err = d.decodeIntBit(stats)
			if err != nil {
				return 0, err
			}

			// Third Read
			if bit == 1 {

				bit, err = d.decodeIntBit(stats)
				if err != nil {
					return 0, err
				}

				// Fourth Read
				if bit == 1 {

					bit, err = d.decodeIntBit(stats)
					if err != nil {
						return 0, err
					}
					// Fifth Read
					if bit == 1 {

						bitsToRead = 32
						offset = 4436

					} else {
						// Fifth Read

						bitsToRead = 12
						offset = 340
					}
				} else {
					// Fourth Read
					bitsToRead = 8
					offset = 84

				}
			} else {
				// Third Read
				bitsToRead = 6
				offset = 20

			}
		} else {
			// SecondRead
			bitsToRead = 4
			offset = 4

		}
	} else {
		// First read
		bitsToRead = 2
		offset = 0

		// common.Log.Trace("Read First value: %v ", value)
	}

	for i := 0; i < bitsToRead; i++ {
		bit, err = d.decodeIntBit(stats)
		if err != nil {
			return 0, err
		}
		value = (value << 1) | bit
	}
	value += offset

	// common.Log.Trace("Value decoded: %v with sign: %b", value, s)

	if s == 0 {

		return int(value), nil
	} else if s == 1 && value > 0 {

		return int(-value), nil
	}

	return math.MaxInt64, nil
}

// DecodeIAID decodes the IAID procedure, Annex A.3
func (d *Decoder) DecodeIAID(codeLen uint64, stats *DecoderStats) (int64, error) {

	// common.Log.Trace("DecodeIAID with codeLen: %d", codeLen)
	// A.3 1)
	d.previous = 1

	var i uint64

	// A.3 2)
	for i = 0; i < codeLen; i++ {

		stats.SetIndex(int(d.previous))
		bit, err := d.DecodeBit(stats)
		if err != nil {
			return 0, err
		}

		d.previous = (d.previous << 1) | int64(bit)
	}

	// A.3 3) & 5)
	result := d.previous - (1 << codeLen)
	// common.Log.Trace("decodeIAID result: %d", result)
	return result, nil
}

func (d *Decoder) init() error {
	d.streamPosition = d.r.StreamPosition()
	b, err := d.r.ReadByte()
	if err != nil {
		common.Log.Trace("Buffer0 readByte failed. %v", err)
		return err
	}
	d.b = int(b)

	d.c = (uint64(b) << 16)
	if err = d.readByte(); err != nil {
		return err
	}

	// Shift c value again
	d.c <<= 7

	// set
	d.ct -= 7

	d.a = 0x8000

	d.prvCtr++
	// common.Log.Trace("Decoder 'a' value: %b", d.a)
	// common.Log.Trace("%d, C: %08x A: %04x, CTR: %d, B0: %02x", d.prvCtr, d.c, d.a, d.ct, d.b)

	return nil
}

func (d *Decoder) readByte() error {

	if d.r.StreamPosition() > d.streamPosition {
		if _, err := d.r.Seek(-1, io.SeekCurrent); err != nil {
			return err
		}
	}

	b, err := d.r.ReadByte()
	if err != nil {
		return err
	}

	d.b = int(b)

	if d.b == 0xFF {
		b1, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		if b1 > 0x8F {
			d.c += 0xFF00
			d.ct = 8
			if _, err := d.r.Seek(-2, io.SeekCurrent); err != nil {
				return err
			}
		} else {
			d.c += uint64(b1) << 9
			d.ct = 7
		}
	} else {
		b, err = d.r.ReadByte()
		if err != nil {
			return err
		}
		d.b = int(b)

		d.c += uint64(d.b) << 8
		d.ct = 8
	}

	d.c &= 0xFFFFFFFFFF

	return nil
}

func (d *Decoder) renormalize() error {
fl:
	for {
		if d.ct == 0 {
			if err := d.readByte(); err != nil {
				return err
			}
		}

		d.a <<= 1
		d.c <<= 1
		d.ct--

		if (d.a & 0x8000) != 0 {
			break fl
		}
	}

	d.c &= 0xffffffff
	return nil
}

func (d *Decoder) decodeIntBit(stats *DecoderStats) (int, error) {
	stats.SetIndex(int(d.previous))
	bit, err := d.DecodeBit(stats)
	if err != nil {
		common.Log.Trace("ArithmeticDecoder 'decodeIntBit'-> DecodeBit failed. %v", err)
		return bit, err
	}

	// common.Log.Trace("decodedIntBit: %d", bit)

	// setPrev(bit)
	if d.previous < 256 {
		// common.Log.Trace("Lower than 256 - %d", d.previous)
		d.previous = ((d.previous << uint64(1)) | int64(bit)) & 0x1ff
	} else {
		// common.Log.Trace("Greater than 256 - %d", d.previous)
		d.previous = (((d.previous<<uint64(1) | int64(bit)) & 511) | 256) & 0x1ff
	}

	// common.Log.Trace("Previous: %d", d.previous)

	return bit, nil
}

func (d *Decoder) mpsExchange(stats *DecoderStats, icx int) int {
	mps := stats.mps[stats.index]

	if d.a < qe[icx][0] {
		if qe[icx][3] == 1 {
			stats.toggleMps()
		}

		stats.SetEntry(int(qe[icx][2]))
		return int(1 - mps)
	}
	stats.SetEntry(int(qe[icx][1]))
	return int(mps)

}

func (d *Decoder) lpsExchange(stats *DecoderStats, icx int, qeValue uint32) int {
	// common.Log.Trace("LPS Exchange for icx: %d...", icx)
	mps := stats.getMps()
	if d.a < qeValue {
		// common.Log.Trace("Smaller than qeValue")
		stats.SetEntry(int(qe[icx][1]))
		d.a = qeValue
		return int(mps)
	}

	if qe[icx][3] == 1 {
		// common.Log.Trace("toggleMps")
		stats.toggleMps()
	}

	stats.SetEntry(int(qe[icx][2]))
	d.a = qeValue
	return int(1 - mps)

}