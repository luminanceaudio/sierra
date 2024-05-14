package sourcesample

import (
	"github.com/sirupsen/logrus"
	"sierra/services/sierra/client/models"
	"sierra/services/sierra/internal/sierradb/sierraent"
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
