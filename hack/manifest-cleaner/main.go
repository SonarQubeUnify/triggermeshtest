/*
Copyright 2022 TriggerMesh Inc.

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
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

/* A command which cleans up the YAML manifests generated by 'ko resolve' by:
    - Removing redundant copyright headers
    - Filtering out null YAML documents
    - Removing '+rbac-check' tags

   Usage:
     manifest-cleaner <manifest.yaml
     cat manifest.yaml | manifest-cleaner
*/

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error running command: %s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, stdin io.ReadCloser, stdout io.Writer) error {
	data, err := readInput(ctx, stdin)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	docs, err := yamlNodes(data)
	if err != nil {
		return fmt.Errorf("decoding YAML documents: %w", err)
	}

	copyrightHeader := findCopyrightHeader(docs)

	docs = filterOutCopyrightHeaders(docs)
	docs = filterOutNullDocs(docs)
	docs = filterOutRBACCheckTags(docs)

	if err := writeNodes(stdout, docs, copyrightHeader); err != nil {
		return fmt.Errorf("writing YAML nodes: %w", err)
	}

	return nil
}