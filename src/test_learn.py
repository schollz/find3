from learn import *
from server import *

def test_to_base64():
  assert to_base64("testdb") == "dGVzdGRi"

def basic_learning():
    ai = AI()
    ai.load_data('../testing/testdb.csv')
    ai.learn()
    ai.save('dGVzdGRi.de0gee.ai')

def test_basic_learning(benchmark):
    result = benchmark(basic_learning)

def basic_classifying():
    ai = AI()
    ai.load('dGVzdGRi.de0gee.ai')
    return ai.classify()

def test_basic_classifying(benchmark):
    result = benchmark(basic_classifying)