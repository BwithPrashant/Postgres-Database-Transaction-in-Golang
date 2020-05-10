echo "Creating tables"
HOST="localhost"
PORT="5432"
DB_NAME="postgres"
USER="postgres"
PASSWORD="password"

psql "host=$HOST port=$PORT dbname=$DB_NAME user=$USER password=$PASSWORD" <<EOF

CREATE SCHEMA IF NOT EXISTS sales;
CREATE TABLE IF NOT EXISTS sales.employee (id int PRIMARY KEY, name TEXT, role TEXT, salary int);

EOF