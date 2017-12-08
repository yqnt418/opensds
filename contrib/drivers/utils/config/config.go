// Copyright (c) 2017 OpenSDS Authors.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package config

import (
	"io/ioutil"

	log "github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

type PoolProperties struct {
	DiskType string `yaml:"diskType,omitempty"`
	AZ       string `yaml:"AZ,omitempty"`
	Thin     bool   `yaml:"thin,omitempty"`
	Compress bool   `yaml:"compress,omitempty"`
	Dedupe   bool   `yaml:"dedupe,omitempty"`
}

func Parse(conf interface{}, p string) (interface{}, error) {
	confYaml, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalf("Read config yaml file (%s) failed, reason:(%v)", p, err)
		return nil, err
	}
	if err = yaml.Unmarshal(confYaml, conf); err != nil {
		log.Fatal("Parse error: %v", err)
		return nil, err
	}
	return conf, nil
}
