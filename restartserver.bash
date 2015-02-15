#!/bin/bash
./stopserver.bash
nohup ./runserver.bash >> log.txt &
