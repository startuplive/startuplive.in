#!/bin/bash
until startuplive.in; do
    echo "Server 'startuplive.in' crashed with exit code $?.  Respawning..." >&2
    sleep 1
done
