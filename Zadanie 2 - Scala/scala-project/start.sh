#!/bin/bash

ngrok http 9000 > /dev/null &
java -jar /opt/zadanie2-assembly/zadanie2-assembly-1.0.jar