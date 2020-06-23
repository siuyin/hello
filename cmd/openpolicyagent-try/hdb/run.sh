#!/bin/sh
opa eval -i input.yaml -d new_flat/ -f pretty data.hdb.new_flat.age_eligible
