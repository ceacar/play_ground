FROM ubuntu
# airflow needs a home, ~/airflow is the default,
# but you can lay foundation somewhere else if you prefer
# (optional)

RUN apt-get update
RUN apt-get install python python3 python3-pip libc-dev build-essential
RUN pip install -U pip
RUN pip3 install numpy pandas apache_airflow

RUN export AIRFLOW_HOME=~/airflow
# initialize the database
RUN airflow initdb

# start the web server, default port is 8080
# visit localhost:8080 in the browser and enable the example dag in the home page
# COMMAND ['airflow', 'webserver', '-p', '8080']
RUN airflow webserver -p 8080&

# start the scheduler
CMD ['airflow', 'scheduler']

WORKDIR /tmp

