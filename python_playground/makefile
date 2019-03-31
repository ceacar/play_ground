SHELL := /bin/bash
tester_version = 1.0.0
python = $(shell pipenv --py) # evaluate this once to improve performance in test
integration_test = $(npm_bin)/take-home-integration-test
log_file = integration-test.log

pip_version = $(word 2,$(shell pipenv run pip --version))
pipenv_version = $(word 3,$(shell pipenv --version))
source_env = $(source ./venv_py3/bin/activate)

start:
	source ./venv_py3/bin/activate
	export PS1='loan_processor $PS1'
	export PYTHONPATH=$PWD:$PYTHONPATH
	python ./run.py


install:
	source ./venv_py3/bin/activate
	export PS1='loan_processor $PS1'
	export PYTHONPATH=$PWD:$PYTHONPATH

utest:
	source ./venv_py3/bin/activate
	py.test -vv ./loan_processor/tests/utests/*

