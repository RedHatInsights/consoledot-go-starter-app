#!/bin/bash

# This script is executed after the template is created.
echo "Creating README.md"
rm -rf README.md
mv DERIVED.md README.md
echo "Creating Git Repository"
git init
git add -A
git commit -am "Initial commit"
