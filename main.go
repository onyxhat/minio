/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/kardianos/service"
	minio "github.com/minio/minio/cmd"
)

type program struct{}

//Runtime
func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	log.Info("minio Started")

	minio.Main()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "minio",
		DisplayName: "minio Service",
		Description: "minio S3 Compatible Service",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
