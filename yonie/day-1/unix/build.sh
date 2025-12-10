#!/bin/bash

OUTPUT_DIR="binary"
FILE_DIR="test-file"

mkdir -p $OUTPUT_DIR
mkdir -p $FILE_DIR

echo "Create $FILE_DIR/file.txt"
echo "Go, often referred to as Golang, is an open-source programming language developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It was designed to address challenges in modern software development, particularly for large-scale systems and cloud-native applications." > $FILE_DIR/file.txt

echo "Create $FILE_DIR/file2.txt"
echo "Common applications of Golang include: Web development (especially for backend services and APIs), Cloud and network programming, Building command-line tools, Developing microservices, and System programming and infrastructure tools." > $FILE_DIR/file2.txt

echo "Building binary..."
go build -o $OUTPUT_DIR ./cmd/...

ls -l $OUTPUT_DIR