#!/bin/bash


# Stopping the go server if already running
# Check if the program is running
if pgrep -f "existserver" > /dev/null; then
    # Get the process ID
    pid=$(pgrep -f "existserver")
    # Send the kill signal
    echo "Killing process with ID: $pid"
    kill $pid
else
    echo "The program is not currently running"
fi

