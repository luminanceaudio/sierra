package waveform

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
)

// DrawWaveformSVG draws a waveform and outputs it as SVG with optimized path.
func DrawWaveformSVG(samples []int16, w io.Writer) error {
	amplitudeData := calculateAmplitudeData(samples)
	fmt.Println(amplitudeData[:100])
	svgPath := generateSVGPath(amplitudeData)

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

//func generateSVGPath(amplitudeData []float64) string {
//	path := "M0," + fmt.Sprintf("%f", amplitudeData[0])
//
//	for i := 1; i < len(amplitudeData); i++ {
//		path += fmt.Sprintf(" L%d,%f", i, amplitudeData[i])
//	}
//
//	return path
//}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateAmplitudeData(samples []int16) []float64 {
	amplitudeData := make([]float64, len(samples))
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
		amplitudeData[i] = float64(sample)
	}

	return amplitudeData
}

//func DrawWaveformPNG(samples []int16, w io.Writer) error {
//	amplitudeData := calculateAmplitudeData(samples, 10)
//
//	// Image dimensions
//	width := len(amplitudeData)
//	height := 400
//
//	// Create a new grayscale image
//	img := image.NewGray(image.Rect(0, 0, width, height))
//	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
//
//	// Draw the waveform
//	for x := 0; x < width-1; x++ {
//		y1 := int((1 - amplitudeData[x]) * float64(height) / 2)
//		y2 := int((1 - amplitudeData[x+1]) * float64(height) / 2)
//		drawLine(img, x, y1, x+1, y2)
//	}
//
//	err := png.Encode(w, img)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//
//// calculateAmplitudeData converts samples to amplitude data with reduced points
//func calculateAmplitudeData(samples []int16, reductionFactor int) []float64 {
//	numSamples := len(samples)
//	numAmplitudeData := numSamples / reductionFactor
//	amplitudeData := make([]float64, numAmplitudeData)
//	maxAmplitude := math.MaxInt16 // Maximum positive amplitude for int16
//
//	for i := 0; i < numAmplitudeData; i++ {
//		// Calculate the starting index for the current amplitude data point
//		startIndex := i * reductionFactor
//
//		// Calculate the sum of samples within the reduction factor
//		sum := 0
//		for j := 0; j < reductionFactor; j++ {
//			index := startIndex + j
//			if index < numSamples {
//				sum += int(samples[index])
//			}
//		}
//
//		// Calculate the average amplitude within the reduction factor
//		averageSample := float64(sum) / float64(reductionFactor)
//
//		// Normalize sample value to range [-1, 1]
//		amplitudeData[i] = averageSample / float64(maxAmplitude)
//	}
//
//	return amplitudeData
//}

func drawLine(img *image.Gray, x1, y1, x2, y2 int) {
	dx := math.Abs(float64(x2 - x1))
	dy := math.Abs(float64(y2 - y1))
	sx := 1
	if x1 > x2 {
		sx = -1
	}
	sy := 1
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy

	for {
		img.SetGray(x1, y1, color.Gray{Y: 0}) // Draw pixel

		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}
