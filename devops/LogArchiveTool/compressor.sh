#!/bin/bash

current_date=$(date +"%Y%m%d")
current_time=$(date +"%H%M%S")

echo "Current Date: $current_date"
echo "Current Time: $current_time"

tar -czvf "logs_archive_$(echo "$current_date")_$(echo "$current_time").tar.gz" test.log
echo "Archive created successfully."

delete_old_logs() {
  find . -name "*.log" -type f -mtime +30 -exec rm {} \;
  echo "Old logs deleted."
}

delete_old_logs
echo "Script execution completed."
# End of LogArchiveTool/compressor.sh
# This script compresses log files and deletes old logs older than 30 days.
# Ensure to run this script in the directory containing the log files.
# Usage: ./compressor.sh
# Note: The tar command will create a compressed archive of the specified log file.
# The archive will be named with the current date and time for uniqueness.
# The script also includes a function to delete log files older than 30 days.
# Make sure to have the necessary permissions to delete files in the directory.
# This script is part of the LogArchiveTool project.
# It is designed to help manage log files efficiently by archiving and cleaning up old logs.