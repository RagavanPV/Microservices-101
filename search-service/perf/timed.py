from functools import wraps
from time import time

def current_milli_time():
    return round(time() * 1000)

def timed(f):
  @wraps(f)
  def wrapper(*args, **kwds):
    start = current_milli_time()
    result = f(*args, **kwds)
    elapsed = current_milli_time() - start
    print("%s took %d millseconds to finish" % (f.__name__, elapsed))
    return result
  return wrapper