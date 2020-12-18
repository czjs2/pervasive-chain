#!/usr/bin/env bash
echo "start bak ..."
rm -rf bak
mkdir bak
cp pervasive-chain bak

echo "start compile ..."
cd pervasive-chain
git pull origin master
go build

\cp timertask ../pervasive-chain
echo "stop application ..."
PIDS=`ps -ef | grep pervasive-chain | awk '{print $2}'`
for pid in $PIDS
do      
  kill -9 $pid
done

echo "start application ..."
cd  ../
nohup ./pervasive-chain >> pervasive-chain.log 2>&1 &
sleep 1
echo "done"
