#!/bin/bash
docker pull $1
kind load docker-image $1 --name kne
