// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmd

import (
	"github.com/pulsar-sigs/pulsar-client/pkg/pulsar-client/consumer"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pulsar-client",
		Short: "A cli for apache pulsar client",
		Long:  `Hi here is pulsar client.`,
	}
)

// Execute executes the root command.
func Execute() error {
	rootCmd.AddCommand(consumer.NewConsumerCommand())
	return rootCmd.Execute()
}

func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func init() {
}
