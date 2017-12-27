.PHONY: install test

test: clean
	cd src && py.test --cov=learn test_learn.py

install:
	python -m pip install --no-cache-dir --upgrade -r requirements.txt

clean:
	rm -rf __pycache__*
	rm -rf src/__pycache__*
	rm -rf venv/*