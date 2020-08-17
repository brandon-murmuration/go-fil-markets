// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package piecestore

import (
	"fmt"
	"io"

	abi "github.com/filecoin-project/specs-actors/actors/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *PieceInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PieceCID (cid.Cid) (struct)
	if len("PieceCID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceCID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PieceCID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceCID")); err != nil {
		return err
	}

	if err := cbg.WriteCidBuf(scratch, w, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	// t.Deals ([]piecestore.DealInfo) (slice)
	if len("Deals") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Deals\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Deals"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Deals")); err != nil {
		return err
	}

	if len(t.Deals) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Deals was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Deals))); err != nil {
		return err
	}
	for _, v := range t.Deals {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *PieceInfo) UnmarshalCBOR(r io.Reader) error {
	*t = PieceInfo{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("PieceInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.PieceCID (cid.Cid) (struct)
		case "PieceCID":

			{

				c, err := cbg.ReadCid(br)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
				}

				t.PieceCID = c

			}
			// t.Deals ([]piecestore.DealInfo) (slice)
		case "Deals":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Deals: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Deals = make([]DealInfo, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v DealInfo
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Deals[i] = v
			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
func (t *DealInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{164}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealID (abi.DealID) (uint64)
	if len("DealID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.DealID)); err != nil {
		return err
	}

	// t.SectorID (abi.SectorNumber) (uint64)
	if len("SectorID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SectorID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SectorID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SectorID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SectorID)); err != nil {
		return err
	}

	// t.Offset (abi.PaddedPieceSize) (uint64)
	if len("Offset") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Offset\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Offset"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Offset")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Offset)); err != nil {
		return err
	}

	// t.Length (abi.PaddedPieceSize) (uint64)
	if len("Length") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Length\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Length"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Length")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Length)); err != nil {
		return err
	}

	return nil
}

func (t *DealInfo) UnmarshalCBOR(r io.Reader) error {
	*t = DealInfo{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("DealInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.DealID (abi.DealID) (uint64)
		case "DealID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.DealID = abi.DealID(extra)

			}
			// t.SectorID (abi.SectorNumber) (uint64)
		case "SectorID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.SectorID = abi.SectorNumber(extra)

			}
			// t.Offset (abi.PaddedPieceSize) (uint64)
		case "Offset":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Offset = abi.PaddedPieceSize(extra)

			}
			// t.Length (abi.PaddedPieceSize) (uint64)
		case "Length":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Length = abi.PaddedPieceSize(extra)

			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
func (t *BlockLocation) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.RelOffset (uint64) (uint64)
	if len("RelOffset") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"RelOffset\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("RelOffset"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("RelOffset")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.RelOffset)); err != nil {
		return err
	}

	// t.BlockSize (uint64) (uint64)
	if len("BlockSize") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"BlockSize\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("BlockSize"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("BlockSize")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.BlockSize)); err != nil {
		return err
	}

	return nil
}

func (t *BlockLocation) UnmarshalCBOR(r io.Reader) error {
	*t = BlockLocation{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("BlockLocation: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.RelOffset (uint64) (uint64)
		case "RelOffset":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.RelOffset = uint64(extra)

			}
			// t.BlockSize (uint64) (uint64)
		case "BlockSize":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.BlockSize = uint64(extra)

			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
func (t *PieceBlockLocation) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.BlockLocation (piecestore.BlockLocation) (struct)
	if len("BlockLocation") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"BlockLocation\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("BlockLocation"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("BlockLocation")); err != nil {
		return err
	}

	if err := t.BlockLocation.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)
	if len("PieceCID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceCID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PieceCID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceCID")); err != nil {
		return err
	}

	if err := cbg.WriteCidBuf(scratch, w, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	return nil
}

func (t *PieceBlockLocation) UnmarshalCBOR(r io.Reader) error {
	*t = PieceBlockLocation{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("PieceBlockLocation: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.BlockLocation (piecestore.BlockLocation) (struct)
		case "BlockLocation":

			{

				if err := t.BlockLocation.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.BlockLocation: %w", err)
				}

			}
			// t.PieceCID (cid.Cid) (struct)
		case "PieceCID":

			{

				c, err := cbg.ReadCid(br)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
				}

				t.PieceCID = c

			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
func (t *CIDInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.CID (cid.Cid) (struct)
	if len("CID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("CID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CID")); err != nil {
		return err
	}

	if err := cbg.WriteCidBuf(scratch, w, t.CID); err != nil {
		return xerrors.Errorf("failed to write cid field t.CID: %w", err)
	}

	// t.PieceBlockLocations ([]piecestore.PieceBlockLocation) (slice)
	if len("PieceBlockLocations") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceBlockLocations\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PieceBlockLocations"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceBlockLocations")); err != nil {
		return err
	}

	if len(t.PieceBlockLocations) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.PieceBlockLocations was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.PieceBlockLocations))); err != nil {
		return err
	}
	for _, v := range t.PieceBlockLocations {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *CIDInfo) UnmarshalCBOR(r io.Reader) error {
	*t = CIDInfo{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("CIDInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.CID (cid.Cid) (struct)
		case "CID":

			{

				c, err := cbg.ReadCid(br)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.CID: %w", err)
				}

				t.CID = c

			}
			// t.PieceBlockLocations ([]piecestore.PieceBlockLocation) (slice)
		case "PieceBlockLocations":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.PieceBlockLocations: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.PieceBlockLocations = make([]PieceBlockLocation, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v PieceBlockLocation
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.PieceBlockLocations[i] = v
			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
