import time 
from ttldict import TTLDict

def test_ttldict():
  d=TTLDict(ttl=1)
  d['foo']='bar'
  assert 'foo' in d
  assert d['foo'] == 'bar'
  time.sleep(1)
  assert 'foo' not in d