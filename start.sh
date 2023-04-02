#!/bin/bash

echo "My goal is to:"
read goal
echo

# TODO start automatically
echo "In a separate window, start clippy with:"
echo "clippy --goal \"$goal\""
echo

# wait for /tmp/clippy-pipe to exist
while [ ! -p /tmp/clippy-pipe ]; do
    sleep 1
done

script -F /tmp/clippy-pipe
