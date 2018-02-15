.PHONY: install test clean benchmark production serve

serve: clean
	export FLASK_APP=server.py && \
	export FLASK_DEBUG=1 && \
	cd src && \
	flask run --debugger --port 8002

production: clean
	export LC_ALL=C.UTF-8 && \
	export LANG=C.UTF-8 && \
	export FLASK_APP=server.py && \
	export FLASK_DEBUG=0 && \
	cd src && \
	flask run --port 8002

production2: clean
	export LC_ALL=C.UTF-8 && \
	export LANG=C.UTF-8 && \
	cd src && \
	gunicorn --bind 0.0.0.0:8002 server:app -w 8
test: clean
	cd src && py.test --benchmark-skip --cov=learn test_learn.py
	cd src && py.test --benchmark-skip --cov=ttldict test_ttldict.py

benchmark: clean
	cd src && py.test test_learn.py

install:
	python -m pip install --no-cache-dir --upgrade -r requirements.txt

clean:
	rm -rf __pycache__*
	rm -rf src/__pycache__*
	rm -rf venv/*
