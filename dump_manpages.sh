#!/bin/bash

COMMANDS=(
  cat cp mv grep awk sed tar rsync ssh find du df chmod chown ls ps top kill mkdir rm
)

mkdir -p ./manpages

for cmd in "${COMMANDS[@]}"; do
  man "$cmd" | col -bx > "./manpages/${cmd}.txt"
done

