package waveform

import (
	"fmt"
	"io"
	"math"
)

// DrawWaveformSVG draws a waveform and outputs it as SVG with optimized path.
func DrawWaveformSVG(samples []int16, w io.Writer) error {
	amplitudeData := calculateAmplitudeData(samples, 100)
	//amplitudeData = filterAmplitudeData(amplitudeData, 70)

	svgPath := generateSmoothSVGPath(amplitudeData)

	_, err := fmt.Fprintf(w, `<svg xmlns="http://www.w3.org/2000/svg" width="100%%" height="%d" viewBox="0 0 %d %d">
	<path d="%s" stroke="black" fill="none"/>
</svg>`, 100, len(amplitudeData), 100, svgPath)
	if err != nil {
		return err
	}
	return nil
}

func generateSVGPath(amplitudeData []float64) string {
	path := "M0," + fmt.Sprintf("%f", amplitudeData[0])
	previousValue := amplitudeData[0]
	changeIndex := 0

	for i := 1; i < len(amplitudeData); i++ {
		currentValue := amplitudeData[i]
		if currentValue != previousValue {
			if i-changeIndex > 1 {
				// If there are repeating values, add only the first and last points of the repeating segment
				path += fmt.Sprintf(" L%d,%f", changeIndex, amplitudeData[changeIndex])
				path += fmt.Sprintf(" L%d,%f", i-1, amplitudeData[i-1])
			} else {
				// Otherwise, add the current point
				path += fmt.Sprintf(" L%d,%f", i, currentValue)
			}
			previousValue = currentValue
			changeIndex = i
		}
	}

	return path
}

func generateSmoothSVGPath(amplitudeData []int64) string {
	path := "M0," + fmt.Sprintf("%d", amplitudeData[0])
	previousValue := amplitudeData[0]
	changeIndex := 0

	for i := 1; i < len(amplitudeData); i++ {
		currentValue := amplitudeData[i]
		if currentValue != previousValue {
			if i-changeIndex > 1 {
				// If there are repeating values, add only the first and last points of the repeating segment
				path += fmt.Sprintf(" C%d,%d %d,%d %d,%d", changeIndex, amplitudeData[changeIndex], i-1, amplitudeData[i-1], i, currentValue)
			} else {
				// Otherwise, add the current point
				path += fmt.Sprintf(" L%d,%d", i, currentValue)
			}
			previousValue = currentValue
			changeIndex = i
		}
	}

	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateAmplitudeData(samples []int16, efficiency int64) []int64 {
	amplitudeData := make([]int64, len(samples))
	maxAmplitude := int16(0)

	for _, sample := range samples {
		if abs(int(sample)) > abs(int(maxAmplitude)) {
			maxAmplitude = sample
		}
	}

	if maxAmplitude == 0 {
		return amplitudeData // Return zero amplitude data if all samples are zero
	}

	for i, sample := range samples {
		amplitudeData[i] = int64(math.Round(float64(sample)))
		amplitudeData[i] -= amplitudeData[i] % efficiency
	}

	return amplitudeData
}

func filterAmplitudeData(amplitudeData []int64, samplesCount int) []int64 {
	blockSize := int64(math.Floor(float64(len(amplitudeData) / samplesCount)))

	var filteredData []int64
	for i := 0; i < samplesCount-1; i++ {
		sum := int64(0)

		for _, sample := range amplitudeData[i*int(blockSize) : i*int(blockSize+1)] {
			sum += sample
		}

		filteredData = append(filteredData, sum/blockSize)
	}

	return filteredData
}
