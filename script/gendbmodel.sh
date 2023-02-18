#!/bin/bash
goctl model mysql ddl --src="internal/db/*.sql" --dir "internal/db/operations/model" -c