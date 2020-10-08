#!/bin/sh
set -e

echo "Applying migrations..."
/venv/bin/python manage.py migrate --noinput

# Start uWSGI
exec "$@"
