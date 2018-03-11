// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: dynamicglobalproperties.go

package objects

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *DynamicGlobalProperties) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *DynamicGlobalProperties) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"current_witness":`)

	{

		obj, err = j.CurrentWitness.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"last_budget_time":`)

	{

		obj, err = j.LastBudgetTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"time":`)

	{

		obj, err = j.Time.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"next_maintenance_time":`)

	{

		obj, err = j.NextMaintenanceTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"accounts_registered_this_interval":`)
	fflib.FormatBits2(buf, uint64(j.AccountsRegisteredThisInterval), 10, j.AccountsRegisteredThisInterval < 0)
	buf.WriteString(`,"dynamic_flags":`)
	fflib.FormatBits2(buf, uint64(j.DynamicFlags), 10, j.DynamicFlags < 0)
	buf.WriteString(`,"head_block_id":`)
	fflib.WriteJsonString(buf, string(j.HeadBlockID))
	buf.WriteString(`,"recent_slots_filled":`)
	fflib.WriteJsonString(buf, string(j.RecentSlotsFilled))
	buf.WriteString(`,"head_block_number":`)
	fflib.FormatBits2(buf, uint64(j.HeadBlockNumber), 10, false)
	buf.WriteString(`,"last_irreversible_block_num":`)
	fflib.FormatBits2(buf, uint64(j.LastIrreversibleBlockNum), 10, false)
	buf.WriteString(`,"current_aslot":`)
	fflib.FormatBits2(buf, uint64(j.CurrentAslot), 10, j.CurrentAslot < 0)
	buf.WriteString(`,"witness_budget":`)
	fflib.FormatBits2(buf, uint64(j.WitnessBudget), 10, j.WitnessBudget < 0)
	buf.WriteString(`,"recently_missed_count":`)
	fflib.FormatBits2(buf, uint64(j.RecentlyMissedCount), 10, j.RecentlyMissedCount < 0)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtDynamicGlobalPropertiesbase = iota
	ffjtDynamicGlobalPropertiesnosuchkey

	ffjtDynamicGlobalPropertiesID

	ffjtDynamicGlobalPropertiesCurrentWitness

	ffjtDynamicGlobalPropertiesLastBudgetTime

	ffjtDynamicGlobalPropertiesTime

	ffjtDynamicGlobalPropertiesNextMaintenanceTime

	ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval

	ffjtDynamicGlobalPropertiesDynamicFlags

	ffjtDynamicGlobalPropertiesHeadBlockID

	ffjtDynamicGlobalPropertiesRecentSlotsFilled

	ffjtDynamicGlobalPropertiesHeadBlockNumber

	ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum

	ffjtDynamicGlobalPropertiesCurrentAslot

	ffjtDynamicGlobalPropertiesWitnessBudget

	ffjtDynamicGlobalPropertiesRecentlyMissedCount
)

var ffjKeyDynamicGlobalPropertiesID = []byte("id")

var ffjKeyDynamicGlobalPropertiesCurrentWitness = []byte("current_witness")

var ffjKeyDynamicGlobalPropertiesLastBudgetTime = []byte("last_budget_time")

var ffjKeyDynamicGlobalPropertiesTime = []byte("time")

var ffjKeyDynamicGlobalPropertiesNextMaintenanceTime = []byte("next_maintenance_time")

var ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval = []byte("accounts_registered_this_interval")

var ffjKeyDynamicGlobalPropertiesDynamicFlags = []byte("dynamic_flags")

var ffjKeyDynamicGlobalPropertiesHeadBlockID = []byte("head_block_id")

var ffjKeyDynamicGlobalPropertiesRecentSlotsFilled = []byte("recent_slots_filled")

var ffjKeyDynamicGlobalPropertiesHeadBlockNumber = []byte("head_block_number")

var ffjKeyDynamicGlobalPropertiesLastIrreversibleBlockNum = []byte("last_irreversible_block_num")

var ffjKeyDynamicGlobalPropertiesCurrentAslot = []byte("current_aslot")

var ffjKeyDynamicGlobalPropertiesWitnessBudget = []byte("witness_budget")

var ffjKeyDynamicGlobalPropertiesRecentlyMissedCount = []byte("recently_missed_count")

// UnmarshalJSON umarshall json - template of ffjson
func (j *DynamicGlobalProperties) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *DynamicGlobalProperties) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtDynamicGlobalPropertiesbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtDynamicGlobalPropertiesnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval, kn) {
						currentKey = ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesCurrentWitness, kn) {
						currentKey = ffjtDynamicGlobalPropertiesCurrentWitness
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesCurrentAslot, kn) {
						currentKey = ffjtDynamicGlobalPropertiesCurrentAslot
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesDynamicFlags, kn) {
						currentKey = ffjtDynamicGlobalPropertiesDynamicFlags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesHeadBlockID, kn) {
						currentKey = ffjtDynamicGlobalPropertiesHeadBlockID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesHeadBlockNumber, kn) {
						currentKey = ffjtDynamicGlobalPropertiesHeadBlockNumber
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesID, kn) {
						currentKey = ffjtDynamicGlobalPropertiesID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesLastBudgetTime, kn) {
						currentKey = ffjtDynamicGlobalPropertiesLastBudgetTime
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesLastIrreversibleBlockNum, kn) {
						currentKey = ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesNextMaintenanceTime, kn) {
						currentKey = ffjtDynamicGlobalPropertiesNextMaintenanceTime
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesRecentSlotsFilled, kn) {
						currentKey = ffjtDynamicGlobalPropertiesRecentSlotsFilled
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesRecentlyMissedCount, kn) {
						currentKey = ffjtDynamicGlobalPropertiesRecentlyMissedCount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesTime, kn) {
						currentKey = ffjtDynamicGlobalPropertiesTime
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesWitnessBudget, kn) {
						currentKey = ffjtDynamicGlobalPropertiesWitnessBudget
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesRecentlyMissedCount, kn) {
					currentKey = ffjtDynamicGlobalPropertiesRecentlyMissedCount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesWitnessBudget, kn) {
					currentKey = ffjtDynamicGlobalPropertiesWitnessBudget
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesCurrentAslot, kn) {
					currentKey = ffjtDynamicGlobalPropertiesCurrentAslot
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesLastIrreversibleBlockNum, kn) {
					currentKey = ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesHeadBlockNumber, kn) {
					currentKey = ffjtDynamicGlobalPropertiesHeadBlockNumber
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesRecentSlotsFilled, kn) {
					currentKey = ffjtDynamicGlobalPropertiesRecentSlotsFilled
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesHeadBlockID, kn) {
					currentKey = ffjtDynamicGlobalPropertiesHeadBlockID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesDynamicFlags, kn) {
					currentKey = ffjtDynamicGlobalPropertiesDynamicFlags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval, kn) {
					currentKey = ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyDynamicGlobalPropertiesNextMaintenanceTime, kn) {
					currentKey = ffjtDynamicGlobalPropertiesNextMaintenanceTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyDynamicGlobalPropertiesTime, kn) {
					currentKey = ffjtDynamicGlobalPropertiesTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesLastBudgetTime, kn) {
					currentKey = ffjtDynamicGlobalPropertiesLastBudgetTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesCurrentWitness, kn) {
					currentKey = ffjtDynamicGlobalPropertiesCurrentWitness
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyDynamicGlobalPropertiesID, kn) {
					currentKey = ffjtDynamicGlobalPropertiesID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtDynamicGlobalPropertiesnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtDynamicGlobalPropertiesID:
					goto handle_ID

				case ffjtDynamicGlobalPropertiesCurrentWitness:
					goto handle_CurrentWitness

				case ffjtDynamicGlobalPropertiesLastBudgetTime:
					goto handle_LastBudgetTime

				case ffjtDynamicGlobalPropertiesTime:
					goto handle_Time

				case ffjtDynamicGlobalPropertiesNextMaintenanceTime:
					goto handle_NextMaintenanceTime

				case ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval:
					goto handle_AccountsRegisteredThisInterval

				case ffjtDynamicGlobalPropertiesDynamicFlags:
					goto handle_DynamicFlags

				case ffjtDynamicGlobalPropertiesHeadBlockID:
					goto handle_HeadBlockID

				case ffjtDynamicGlobalPropertiesRecentSlotsFilled:
					goto handle_RecentSlotsFilled

				case ffjtDynamicGlobalPropertiesHeadBlockNumber:
					goto handle_HeadBlockNumber

				case ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum:
					goto handle_LastIrreversibleBlockNum

				case ffjtDynamicGlobalPropertiesCurrentAslot:
					goto handle_CurrentAslot

				case ffjtDynamicGlobalPropertiesWitnessBudget:
					goto handle_WitnessBudget

				case ffjtDynamicGlobalPropertiesRecentlyMissedCount:
					goto handle_RecentlyMissedCount

				case ffjtDynamicGlobalPropertiesnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ID:

	/* handler: j.ID type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CurrentWitness:

	/* handler: j.CurrentWitness type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.CurrentWitness.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastBudgetTime:

	/* handler: j.LastBudgetTime type=objects.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LastBudgetTime.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Time:

	/* handler: j.Time type=objects.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Time.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NextMaintenanceTime:

	/* handler: j.NextMaintenanceTime type=objects.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.NextMaintenanceTime.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AccountsRegisteredThisInterval:

	/* handler: j.AccountsRegisteredThisInterval type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.AccountsRegisteredThisInterval = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DynamicFlags:

	/* handler: j.DynamicFlags type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.DynamicFlags = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HeadBlockID:

	/* handler: j.HeadBlockID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.HeadBlockID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RecentSlotsFilled:

	/* handler: j.RecentSlotsFilled type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.RecentSlotsFilled = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HeadBlockNumber:

	/* handler: j.HeadBlockNumber type=objects.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.HeadBlockNumber.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastIrreversibleBlockNum:

	/* handler: j.LastIrreversibleBlockNum type=objects.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LastIrreversibleBlockNum.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CurrentAslot:

	/* handler: j.CurrentAslot type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.CurrentAslot = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WitnessBudget:

	/* handler: j.WitnessBudget type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.WitnessBudget = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RecentlyMissedCount:

	/* handler: j.RecentlyMissedCount type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.RecentlyMissedCount = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
