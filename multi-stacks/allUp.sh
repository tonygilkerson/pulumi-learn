#! /bin/bash

#
# Note, you should not need to use '--refresh' but I keep deleteding my kind cluster making my state invalid
#       --refresh = Refresh the state of the stack's resources before this update
#
pulumi up --cwd ./infr --refresh
pulumi up --cwd ./app1 --refresh
pulumi up --cwd ./app2 --refresh
