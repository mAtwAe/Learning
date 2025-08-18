echo 'Top 5 IP addresses with the most requests:'

awk '{print $1}' ngnx.log | sort | uniq -c | sort -nr | head -5 | awk '{print $2, "-", $1, "requests"}'

echo 'Top 5 most requested paths:'

awk '{print $7}' ngnx.log | sort | uniq -c | sort -nr | head -5 | awk '{print $2, "-", $1, "requests"}'

echo 'Top 5 HTTP status codes:'

awk '{print $9}' ngnx.log | sort | uniq -c | sort -nr | head -5 | awk '{print $2, "-", $1, "responses"}'

echo 'Top 5 user agents:'

awk -F\" '{print $6}' ngnx.log | sort | uniq -c | sort -nr | head -5 | awk '{print $2, "-", $1, "requests"}'
