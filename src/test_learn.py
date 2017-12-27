from learn import *

def test_hello(benchmark):
  assert hello() == "hello"
  result = benchmark(hello)