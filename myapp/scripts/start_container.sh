#!/bin/bash

# Define the container ID or name
CONTAINER_ID_OR_NAME="your_container_id_or_name"

# Start the container if it is stopped and attach a Bash shell
docker start -i $CONTAINER_ID_OR_NAME && docker exec -it $CONTAINER_ID_OR_NAME /bin/bash

