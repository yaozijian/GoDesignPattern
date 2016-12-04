#!/bin/bash

find . -exec /bin/bash -c '[ "{}" != "./build.sh" ] && [ "{}" != "./clear.sh" ] && [ ! -d {} ] && [ -x {} ] && rm -f {}' \; 

