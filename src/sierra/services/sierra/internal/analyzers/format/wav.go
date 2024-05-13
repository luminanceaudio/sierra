package format

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sierra/services/sierra/internal/format"
)

var (
	riffSignature = []byte("RIFF")
	wavSignature  = []byte("WAVE")
	formatChunkId = []byte("fmt ")
	dataChunkId   = []byte("data")
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

	wavFormat, samples, err := handleChunks(reader)
	if err != nil {
		return nil, false, fmt.Errorf("failed handling chunks: %w", err)
	}

	return newFormat(format.TypeWav, wavFormat.Channels, wavFormat.SampleRate, wavFormat.BitsPerSample, samples), true, nil
}

func handleChunks(reader io.ReadSeeker) (wavFormat *WavFormat, samples []int16, err error) {
	var cursor int64
	var foundChunkId [4]byte

	for {
		err := binary.Read(reader, binary.LittleEndian, &foundChunkId)
		if err != nil {
			// No more chunks
			if errors.Is(err, io.EOF) {
				return wavFormat, samples, nil
			}

			if errors.Is(err, io.ErrUnexpectedEOF) {
				// Check for EOF, if it exists that means it's likely just padding
				var padding [1]byte
				_, paddingErr := reader.Read(padding[:])
				if paddingErr != io.EOF {
					return nil, nil, fmt.Errorf("expected padding after unexpected EOF reading chunk: %w", paddingErr)
				}

				return wavFormat, samples, nil
			}

			return nil, nil, fmt.Errorf("failed reading chunk id: %w", err)
		}

		var chunkSize int32
		err = binary.Read(reader, binary.LittleEndian, &chunkSize)
		if err != nil {
			return nil, nil, fmt.Errorf("failed reading chunk id: %w", err)
		}

		// Get cursor without actually seeking
		cursor, err = reader.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, nil, fmt.Errorf("failed seeking sample file to get cursor: %w", err)
		}

		switch {
		case bytes.Equal(foundChunkId[:], formatChunkId):
			wavFormat, err = parseFormat(reader, chunkSize)
			if err != nil {
				return nil, nil, fmt.Errorf("failed parsing format: %w", err)
			}
		case bytes.Equal(foundChunkId[:], dataChunkId):
			samples, err = readAudioSamples(reader, chunkSize, wavFormat.BitsPerSample)
			if err != nil {
				return nil, nil, fmt.Errorf("failed parsing format: %w", err)
			}
		}

		// Seek to end of chunk
		_, err = reader.Seek(cursor+int64(chunkSize), io.SeekStart)
		if err != nil {
			return nil, nil, fmt.Errorf("failed seeking sample file to get cursor: %w", err)
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

func readAudioSamples(reader io.ReadSeeker, dataSize int32, bitsPerSample uint16) ([]int16, error) {
	// Calculate the number of samples based on the data size and bits per sample
	numSamples := dataSize / int32(bitsPerSample/8)

	// Initialize a slice to hold the audio samples
	samples := make([]int16, numSamples)

	// Read audio samples based on the bit depth
	switch bitsPerSample {
	case 8:
		// 8-bit samples are unsigned, so we need to convert them to signed int16
		rawSamples := make([]uint8, numSamples)
		if err := binary.Read(reader, binary.LittleEndian, &rawSamples); err != nil {
			return nil, fmt.Errorf("failed reading 8-bit samples: %w", err)
		}
		for i, sample := range rawSamples {
			samples[i] = int16(sample) - 128 // Convert unsigned to signed
		}
	case 16:
		if err := binary.Read(reader, binary.LittleEndian, &samples); err != nil {
			return nil, fmt.Errorf("failed reading 16-bit samples: %w", err)
		}
	case 24:
		// 24-bit samples are stored in little-endian format across three bytes
		rawSamples := make([]byte, numSamples*3)
		if err := binary.Read(reader, binary.LittleEndian, &rawSamples); err != nil {
			return nil, fmt.Errorf("failed reading 24-bit samples: %w", err)
		}
		for i := 0; i < len(rawSamples); i += 3 {
			// Combine three bytes into a 24-bit integer
			sample := int32(rawSamples[i]) | (int32(rawSamples[i+1]) << 8) | (int32(rawSamples[i+2]) << 16)
			// Convert to signed int16 and adjust for the right alignment
			samples[i/3] = int16(sample >> 8)
		}
	default:
		return nil, fmt.Errorf("unsupported bit depth: %d", bitsPerSample)
	}

	return samples, nil
}
