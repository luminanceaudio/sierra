package format

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/format"
	"io"
)

var (
	riffSignature = []byte("RIFF")
	wavSignature  = []byte("WAVE")
	formatChunkId = []byte("fmt ")
)

type WavHeader struct {
	RiffSignature [4]byte // = "RIFF"
	FileSize      int32
	WavSignature  [4]byte // = "WAVE"
}

type WavFormat struct {
	FormatTag          int16
	Channels           uint16
	SampleRate         uint32
	AverageBytesPerSec uint32
	BlockAlign         uint16
	BitsPerSample      uint16
}

func TryParseWav(reader io.ReadSeeker) (*Format, bool, error) {
	_, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, false, fmt.Errorf("failed seeking sample file: %w", err)
	}

	var wavHeader WavHeader
	err = binary.Read(reader, binary.LittleEndian, &wavHeader)
	if err != nil {
		return nil, false, fmt.Errorf("failed parsing header: %w", err)
	}

	if !bytes.Equal(wavHeader.RiffSignature[:], riffSignature) {
		return nil, false, nil
	}

	if !bytes.Equal(wavHeader.WavSignature[:], wavSignature) {
		return nil, false, nil
	}

	wavFormat, err := handleChunks(reader)
	if err != nil {
		return nil, false, fmt.Errorf("failed handling chunks: %w", err)
	}

	return NewFormat(format.TypeWav, wavFormat.Channels, wavFormat.SampleRate, wavFormat.BitsPerSample), true, nil
}

func handleChunks(reader io.ReadSeeker) (wavFormat *WavFormat, err error) {
	var cursor int64
	var foundChunkId [4]byte
	for {
		err := binary.Read(reader, binary.LittleEndian, &foundChunkId)
		if err != nil {
			// No more chunks
			if errors.Is(err, io.EOF) {
				return wavFormat, nil
			}

			if errors.Is(err, io.ErrUnexpectedEOF) {
				// Check for EOF, if it exists that means it's likely just padding
				var padding [1]byte
				_, paddingErr := reader.Read(padding[:])
				if paddingErr != io.EOF {
					return nil, fmt.Errorf("expected padding after unexpected EOF reading chunk: %w", paddingErr)
				}

				return wavFormat, nil
			}

			return nil, fmt.Errorf("failed reading chunk id: %w", err)
		}

		var chunkSize int32
		err = binary.Read(reader, binary.LittleEndian, &chunkSize)
		if err != nil {
			return nil, fmt.Errorf("failed reading chunk id: %w", err)
		}

		// Get cursor without actually seeking
		cursor, err = reader.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, fmt.Errorf("failed seeking sample file to get cursor: %w", err)
		}

		switch {
		case bytes.Equal(foundChunkId[:], formatChunkId):
			wavFormat, err = parseFormat(reader, chunkSize)
			if err != nil {
				return nil, fmt.Errorf("failed parsing format: %w", err)
			}
		}

		// Seek to end of chunk
		_, err = reader.Seek(cursor+int64(chunkSize), io.SeekStart)
		if err != nil {
			return nil, fmt.Errorf("failed seeking sample file to get cursor: %w", err)
		}
	}
}

func parseFormat(reader io.ReadSeeker, size int32) (*WavFormat, error) {
	var wavFormat WavFormat
	err := binary.Read(reader, binary.LittleEndian, &wavFormat)
	if err != nil {
		return nil, fmt.Errorf("failed parsing format: %w", err)
	}

	return &wavFormat, nil
}
