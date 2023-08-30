// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package orders

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *OrderInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{170}); err != nil {
		return err
	}

	// t.Msg (string) (string)
	if len("Msg") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Msg\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Msg"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Msg")); err != nil {
		return err
	}

	if len(t.Msg) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Msg was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Msg))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Msg)); err != nil {
		return err
	}

	// t.User (string) (string)
	if len("User") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"User\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("User"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("User")); err != nil {
		return err
	}

	if len(t.User) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.User was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.User))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.User)); err != nil {
		return err
	}

	// t.State (orders.OrderState) (int64)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("State")); err != nil {
		return err
	}

	if t.State >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.State)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.State-1)); err != nil {
			return err
		}
	}

	// t.Value (string) (string)
	if len("Value") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Value\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Value"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Value")); err != nil {
		return err
	}

	if len(t.Value) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Value was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Value))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Value)); err != nil {
		return err
	}

	// t.VpsID (int64) (int64)
	if len("VpsID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"VpsID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("VpsID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("VpsID")); err != nil {
		return err
	}

	if t.VpsID >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.VpsID)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.VpsID-1)); err != nil {
			return err
		}
	}

	// t.OrderID (orders.OrderHash) (string)
	if len("OrderID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"OrderID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("OrderID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("OrderID")); err != nil {
		return err
	}

	if len(t.OrderID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.OrderID was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.OrderID))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.OrderID)); err != nil {
		return err
	}

	// t.CycleTime (string) (string)
	if len("CycleTime") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CycleTime\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CycleTime"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CycleTime")); err != nil {
		return err
	}

	if len(t.CycleTime) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.CycleTime was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.CycleTime))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.CycleTime)); err != nil {
		return err
	}

	// t.DoneState (orders.OrderDoneState) (int64)
	if len("DoneState") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DoneState\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("DoneState"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DoneState")); err != nil {
		return err
	}

	if t.DoneState >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.DoneState)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.DoneState-1)); err != nil {
			return err
		}
	}

	// t.GoodsInfo (orders.GoodsInfo) (struct)
	if len("GoodsInfo") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"GoodsInfo\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("GoodsInfo"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("GoodsInfo")); err != nil {
		return err
	}

	if err := t.GoodsInfo.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.OrderType (int64) (int64)
	if len("OrderType") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"OrderType\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("OrderType"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("OrderType")); err != nil {
		return err
	}

	if t.OrderType >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.OrderType)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.OrderType-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *OrderInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = OrderInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("OrderInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Msg (string) (string)
		case "Msg":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Msg = string(sval)
			}
			// t.User (string) (string)
		case "User":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.User = string(sval)
			}
			// t.State (orders.OrderState) (int64)
		case "State":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.State = OrderState(extraI)
			}
			// t.Value (string) (string)
		case "Value":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Value = string(sval)
			}
			// t.VpsID (int64) (int64)
		case "VpsID":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.VpsID = int64(extraI)
			}
			// t.OrderID (orders.OrderHash) (string)
		case "OrderID":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.OrderID = OrderHash(sval)
			}
			// t.CycleTime (string) (string)
		case "CycleTime":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.CycleTime = string(sval)
			}
			// t.DoneState (orders.OrderDoneState) (int64)
		case "DoneState":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.DoneState = OrderDoneState(extraI)
			}
			// t.GoodsInfo (orders.GoodsInfo) (struct)
		case "GoodsInfo":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.GoodsInfo = new(GoodsInfo)
					if err := t.GoodsInfo.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.GoodsInfo pointer: %w", err)
					}
				}

			}
			// t.OrderType (int64) (int64)
		case "OrderType":
			{
				maj, extra, err := cr.ReadHeader()
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.OrderType = int64(extraI)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *GoodsInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.ID (string) (string)
	if len("ID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ID")); err != nil {
		return err
	}

	if len(t.ID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ID was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.ID))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.ID)); err != nil {
		return err
	}

	// t.Password (string) (string)
	if len("Password") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Password\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Password"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Password")); err != nil {
		return err
	}

	if len(t.Password) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Password was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Password))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Password)); err != nil {
		return err
	}
	return nil
}

func (t *GoodsInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = GoodsInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("GoodsInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.ID (string) (string)
		case "ID":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.ID = string(sval)
			}
			// t.Password (string) (string)
		case "Password":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Password = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
