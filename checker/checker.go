package checker

import (
	inti "github.com/hazcod/go-intigriti"
	"github.com/hazcod/intigriti-cicd-plugin/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

func RunChecker(conf config.Config) (exit bool, err error) {
	intigriti := inti.New(conf.IntigritiClientID, conf.IntigritiClientSecret)
	submissions, err := intigriti.GetSubmissions()
	if err != nil {
		return exit, errors.Wrap(err, "could not retrieve submissions")
	}

	tresholds := map[string]uint{
		"critical": conf.Tresholds.MaxCritical,
		"high"    : conf.Tresholds.MaxHigh,
		"medium"  : conf.Tresholds.MaxMedium,
		"low"     : conf.Tresholds.MaxLow,
	}

	for _, submission := range submissions {
		logger := log.WithField("status", strings.ToLower(submission.State))
		if ! submission.IsReady() {
			logger.Debug("skipping non-active issue")
			continue
		}

		severity := strings.ToLower(submission.Severity)
		treshold, found := tresholds[severity]

		if treshold == 0 { treshold = 1 }

		logger = logger.WithField("severity", severity)

		if !found {
			logger.Error("unknown severity")
			return true, nil
		}

		newTreshold := treshold -1

		if newTreshold < 0 {
			logger.
				WithField("program", submission.Program).
				Info("outstanding issue found!")
			return true, nil
		}

		logger.WithField("new_treshold", newTreshold).Debug("adding one outstanding issue")
		tresholds[severity] = newTreshold
	}

	return false, nil
}