package sourcesample

import (
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent"
	"github.com/sirupsen/logrus"
)

func translateSample(sample *sierraent.SourceSample) *models.Sample {
	if sample.Edges.Source == nil {
		logrus.Error("sourcesample's source is nil")
		return nil
	}

	if sample.Edges.Sample == nil {
		logrus.Error("sourcesample's sample is nil")
		return nil
	}

	return models.NewSample(
		sample.ID,
		sample.Edges.Sample.ID,
		sample.Edges.Sample.Format,
		sample.Edges.Source.ID,
		sample.Edges.Sample.Duration,
	)
}

func translateSamples(sourceSamples []*sierraent.SourceSample) []*models.Sample {
	var results []*models.Sample

	for _, sample := range sourceSamples {
		translatedSample := translateSample(sample)
		if translatedSample == nil {
			continue
		}

		results = append(results, translatedSample)
	}

	return results
}
