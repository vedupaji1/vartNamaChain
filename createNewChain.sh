#!/bin/bash

existingChainLocation="/home/codezeros/.temp"

if [ -d "$existingChainLocation" ]; then
    rm -rf "$existingChainLocation"
fi

ignite chain build && ignite chain init
