#!/bin/bash
echo "Running up migration"
(cd sql/schema && goose postgres postgres://postgres:postgres@localhost:5432/gator?sslmode=disable up)