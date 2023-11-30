#!/bin/bash
set -e

DATABASE_NAME="users"

DB_EXISTS=$(mysql -u"$MYSQL_ROOT_USER" -p"$MYSQL_PASSWORD" -e "SHOW DATABASES LIKE '$DATABASE_NAME';" | grep "$DATABASE_NAME" > /dev/null; echo "$?")

# Check if the database already exists
if [ "$DB_EXISTS" -eq 0 ]; then
    echo "Database '$DATABASE_NAME' already exists. Skipping creation."
else
    echo "Database '$DATABASE_NAME' does not exist. Creating..."
    mysql -u"$MYSQL_ROOT_USER" -p"$MYSQL_PASSWORD" -e "CREATE DATABASE $DATABASE_NAME;"
    echo "Database '$DATABASE_NAME' created successfully."
fi
