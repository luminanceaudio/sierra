package format

import (
	"github.com/stretchr/testify/require"
	"io"
	"sierra/services/sierra/internal/format"
	"sierra/tests/resources/samplessanity"
	"testing"
)

func TestWavParse16Bit(t *testing.T) {
	f, err := samplessanity.Files.Open(samplessanity.SierraDrum1_16BitWav)
	require.NoError(t, err)

	readSeeker, isReadSeeker := f.(io.ReadSeeker)
	require.True(t, isReadSeeker)

	formatStruct, valid, err := TryParseWav(readSeeker)
	require.NoError(t, err)

	require.True(t, valid)
	require.Equal(t, formatStruct.Type, format.TypeWav)
	require.Equal(t, uint16(2), formatStruct.Channels)
	require.Equal(t, uint32(48000), formatStruct.SampleRate)
	require.Equal(t, uint16(16), formatStruct.BitsPerSample)
}

func TestWavParse24Bit(t *testing.T) {
	f, err := samplessanity.Files.Open(samplessanity.SierraDrum2_24BitWav)
	require.NoError(t, err)

	readSeeker, isReadSeeker := f.(io.ReadSeeker)
	require.True(t, isReadSeeker)

	formatStruct, valid, err := TryParseWav(readSeeker)
	require.NoError(t, err)

	require.True(t, valid)
	require.Equal(t, formatStruct.Type, format.TypeWav)
	require.Equal(t, uint16(2), formatStruct.Channels)
	require.Equal(t, uint32(48000), formatStruct.SampleRate)
	require.Equal(t, uint16(24), formatStruct.BitsPerSample)
}

func TestWavParseNoJunk(t *testing.T) {
	f, err := samplessanity.Files.Open(samplessanity.SierraPianoNoJunkWav)
	require.NoError(t, err)

	readSeeker, isReadSeeker := f.(io.ReadSeeker)
	require.True(t, isReadSeeker)

	formatStruct, valid, err := TryParseWav(readSeeker)
	require.NoError(t, err)

	require.True(t, valid)
	require.Equal(t, formatStruct.Type, format.TypeWav)
	require.Equal(t, uint16(2), formatStruct.Channels)
	require.Equal(t, uint32(48000), formatStruct.SampleRate)
	require.Equal(t, uint16(16), formatStruct.BitsPerSample)
}
