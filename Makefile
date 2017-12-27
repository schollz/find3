.PHONY: install test

test:
	py.test --cov=learn test_learn.py

install:
	python -m pip install --no-cache-dir --upgrade -r requirements.txt

clean:
	rm -rf __pycache__*
	rm -rf venv/*