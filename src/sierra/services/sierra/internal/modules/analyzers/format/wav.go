package format

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type WavHeader struct {
	RiffSignature     [4]byte // = "RIFF"
	FileSize          int32
	WavSignature      [4]byte // = "WAVE"
	FormatChunkMarker [4]byte
	FormatDataLength  int32
	FormatType        int16 // 1 = PCM
	Channels          int16
	SampleRate        int32
	_                 int16
	BitsPerSample     int16
	DataChunkHeader   [4]byte // = "data"
	DataSize          int32
}

func TryParseWav(file *os.File) (*Format, bool, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, false, fmt.Errorf("failed seeking sample file: %w", err)
	}

	const wavHeaderLen = 44

	header := make([]byte, wavHeaderLen)
	n, err := file.Read(header)
	if err != nil {
		return nil, false, fmt.Errorf("failed reading header: %w", err)
	}
	if n != wavHeaderLen {
		return nil, false, nil
	}

	var wavHeader WavHeader
	err = binary.Read(file, binary.LittleEndian, &wavHeader)
	if err != nil {
		return nil, false, fmt.Errorf("failed parsing header: %w", err)
	}

	if !bytes.Equal(wavHeader.RiffSignature[:], []byte("RIFF")) {
		return nil, false, nil
	}

	if !bytes.Equal(wavHeader.WavSignature[:], []byte("WAVE")) {
		return nil, false, nil
	}

	return NewFormat(FormatWav), true, nil
}
