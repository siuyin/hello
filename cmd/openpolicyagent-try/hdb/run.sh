#!/bin/sh
opa eval -i input.yaml \
	-d new_flat \
	-d ../siuyin/time \
	-f pretty data.hdb.new_flat
