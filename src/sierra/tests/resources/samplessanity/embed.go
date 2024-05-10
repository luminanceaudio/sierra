package samplessanity

import "embed"

const (
	SierraDrum1_16BitWav = "sierra_drum1_16bit.wav"
	SierraDrum2_24BitWav = "sierra_drum2_24bit.wav"
	SierraPianoNoJunkWav = "sierra_piano_nojunk.wav"
)

//go:embed *
var Files embed.FS
