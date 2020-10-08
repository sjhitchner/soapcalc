FROM python:3.7-alpine3.8
ARG VERSION
ARG PORT=8000
RUN test -n "${VERSION}"
ARG NAME=tools

# Local directory with project source
# DJANGO_SETTINGS_MODULE=${NAME}.settings.production \
ENV SRVHOME=/srv \
    DJANGO_ALLOWED_HOSTS=.internal,.us-east-1.elb.amazonaws.com,*,localhost,127.0.0.1,[::1] \
    PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    VERSION=${VERSION} \
    DJANGO_SETTINGS_MODULE=${NAME}.settings \
    SOAP_DB_USER=soap \
    SOAP_DB_PW=soap \
    SOAP_DB_HOST=localhost \
    SOAP_DB_PORT=15432 \

COPY requirements.txt ${SRVHOME}/requirements.txt

# Install build deps, then run `pip install`, then remove unneeded build deps all in a single step. Correct the path to your production requirements file, if needed.
RUN set -ex \
    && apk add --no-cache --virtual .build-deps \
            gcc \
            make \
            libc-dev \
            musl-dev \
            linux-headers \
            pcre-dev \
            postgresql-dev \
            uwsgi-python3 \
            mailcap \
            libffi-dev \
    && pyvenv /venv \
    && /venv/bin/pip install -U pip==18.1 \
	&& /venv/bin/pip install uwsgi \
    && LIBRARY_PATH=/lib:/usr/lib /bin/sh -c "/venv/bin/pip install --no-cache-dir -r ${SRVHOME}/requirements.txt" \
    && runDeps="$( \
            scanelf --needed --nobanner --recursive /venv \
                    | awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
                    | sort -u \
                    | xargs -r apk info --installed \
                    | sort -u \
    )" \
    && apk add --virtual .python-rundeps $runDeps \
    && apk del .build-deps \
	&& mkdir ${SRVHOME}/static

COPY mimetypes.txt /etc/mime.types
COPY ${SRC} ${SRVHOME}
WORKDIR ${SRVHOME}

# uWSGI will listen on this port
EXPOSE ${PORT}

# uWSGI configuration (customize as needed):
ENV UWSGI_VIRTUALENV=/venv UWSGI_WSGI_FILE=${SRVHOME}/${NAME}/wsgi.py UWSGI_HTTP=:${PORT} UWSGI_MASTER=1 UWSGI_WORKERS=4 UWSGI_THREADS=8 UWSGI_UID=1000 UWSGI_GID=2000 UWSGI_LAZY_APPS=1 UWSGI_WSGI_ENV_BEHAVIOR=holy

# Call collectstatic (customize the following line with the minimal environment variables needed for manage.py to run):
RUN /venv/bin/python manage.py collectstatic --noinput
RUN echo VERSION=${VERSION}

ENTRYPOINT ["/srv/docker-entrypoint.sh"]
CMD ["/venv/bin/uwsgi", "--http-auto-chunked", "--http-keepalive", "--static-map", "/static=/srv/static", "--mimefile", "/etc/mime.types"]
