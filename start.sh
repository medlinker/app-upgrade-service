#!/bin/bash

if [ -e ./bin/appUpgradeService ]
then
    bin/appUpgradeService env=conf/env.conf
else
    echo 'can not found execute file!'
    exit 1
fi

exit 0
