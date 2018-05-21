#! /bin/bash

tar czf - mbus.reader | \
ssh smart 'tar xzf - -C .'
