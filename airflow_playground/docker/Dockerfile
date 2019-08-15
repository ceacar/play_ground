FROM alpine
# airflow needs a home, ~/airflow is the default,
# but you can lay foundation somewhere else if you prefer
# (optional)

RUN apk update
RUN apk upgrade
RUN apk add bash
RUN echo "**** install Python ****" && \
    apk add --no-cache python3 && \
    if [ ! -e /usr/bin/python ]; then ln -sf python3 /usr/bin/python ; fi && \
    \
    echo "**** install pip ****" && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --no-cache --upgrade pip setuptools wheel && \
    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi
#RUN apt-get update
#RUN apt-get install python python3
#RUN apt-get autoclean
#RUN apt-get python3-pip
#RUN export AIRFLOW_HOME=~/airflow

RUN apk add alpine-sdk
#RUN apt-get -y install libc-dev
#RUN apt-get -y install build-essential
RUN pip install -U pip

RUN pip3 install numpy pandas
# install from pypi using pip
RUN pip3 install apache-airflow

# initialize the database
RUN airflow initdb

# start the web server, default port is 8080
# visit localhost:8080 in the browser and enable the example dag in the home page
# COMMAND ['airflow', 'webserver', '-p', '8080']
RUN airflow webserver -p 8080&

# start the scheduler
CMD ['airflow', 'scheduler']

WORKDIR /tmp

