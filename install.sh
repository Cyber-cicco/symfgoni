#!/bin/bash

rm -rf ~/go/bin/go-cap/resources && cp -r resources/ ~/go/bin/go-cap/resources/ && cp go-cap.json ~/go/bin/go-cap/resources/go-cap.json && cd internals && go install && mv ~/go/bin/internals ~/go/bin/gocap
