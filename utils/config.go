package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/sirupsen/logrus"
)

func ReadConfig(source string) (*models.Configuration, error) {
	bytes, err := ioutil.ReadFile(source)
	if err != nil {
		eMsg := "Couldn't read configuration file"
		logrus.WithError(err).Error(eMsg)
		return nil, err
	}

	conf := &models.Configuration{}
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		eMsg := "Couldn't parse config file from json"
		logrus.WithError(err).Error(eMsg)
		return nil, err
	}

	return conf, nil
}
