#!/usr/local/env bash

echo "This is just a test."
echo "Input 1: $1"
echo "Input 2: $2"
echo
echo $OSTYPE
echo $HOME
echo $PWD
echo $PATH

# success: exit code 0
exit 0

# failure: exit code 1
# exit 1