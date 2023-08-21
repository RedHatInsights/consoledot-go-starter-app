#!/bin/bash

# This script is executed after the template is created.
rm -rf README.md
mv DERIVED.md README.md
git init
git add -A
git commit -am "Initial commit"
