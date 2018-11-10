/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"github.com/knative/eventing-sources/pkg/adapter/awssqs"
	"log"
	"os"

	"go.uber.org/zap"

	"golang.org/x/net/context"
)

const (
	// Sink for messages.
	envSinkURI = "SINK_URI"

	// envUrl is the URL of the SQS queue to consume messages from.
	envQueueUrl = "AWS_SQS_URL"
)

func getRequiredEnv(envKey string) string {
	val, defined := os.LookupEnv(envKey)
	if !defined {
		log.Fatalf("required environment variable not defined '%s'", envKey)
	}
	return val
}

func main() {
	flag.Parse()

	ctx := context.Background()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Unable to create logger: %v", err)
	}

	adapter := &awssqs.Adapter{
		QueueUrl: getRequiredEnv(envQueueUrl),
		SinkURI:  getRequiredEnv(envSinkURI),
	}

	logger.Info("Starting AWS SQS Receive Adapter. %v", zap.Reflect("adapter", adapter))
	if err := adapter.Start(ctx); err != nil {
		logger.Fatal("failed to start adapter: ", zap.Error(err))
	}
}
