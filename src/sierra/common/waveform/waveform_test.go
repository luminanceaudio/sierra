package waveform

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"sierra/common/format"
	"sierra/tests/resources/samplessanity"
	"testing"
)

func TestDrawWaveform(t *testing.T) {
	f, err := samplessanity.Files.Open(samplessanity.SierraDrum1_16BitWav)
	require.NoError(t, err)
	defer f.Close()

	readSeeker, isReadSeeker := f.(io.ReadSeeker)
	require.True(t, isReadSeeker)

	formatStruct, valid, err := format.TryParseWav(readSeeker)
	require.NoError(t, err)
	require.True(t, valid)

	svgFile, err := os.CreateTemp("", "waveform_*.svg")
	require.NoError(t, err)
	defer svgFile.Close()

	err = DrawWaveformSVG(formatStruct.Samples, svgFile)
	require.NoError(t, err)

	fmt.Println(svgFile.Name())
}
