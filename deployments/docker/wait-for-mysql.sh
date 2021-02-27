#!/bin/sh
# wait until MySQL is really available
maxcounter=30
if [ ! -z "$1" ]
then
      maxcounter=$1
fi

counter=1
while ! mysql -u"$DB_USER" -p"$DB_PASSWORD" -h"0.0.0.0" -P"3306" -e "show databases;" > /dev/null 2>&1; do
    sleep 1
    counter=`expr $counter + 1`
    if [ $counter -gt $maxcounter ]; then
        >&2 echo "We have been waiting for MySQL too long already; failing."
        exit 1
    fi;
done