import threading
import time


class TTLDict(dict):
    """ A dictionary with keys having specific time to live """

    def __init__(self, ttl=5):
        self._ttl = ttl
        # Internal key mapping keys of _d
        # to times of access
        self.__map = {}
        self._flag = True
        self._t = threading.Thread(target=self._collect)
        self._t.setDaemon(True)
        self._t.start()

    def _collect(self):
        """ Collect expired keys in a loop """

        while self._flag:
            now = time.time()
            keys = list(self.__map.keys())

            # Pop the top first 'sz' keys
            for key in keys:
                val = self.__map[key]
                diff = now - val
                if diff > self._ttl:
                    # print 'Dropping key',key
                    del self[key]
                    # Drop it from this also
                    del self.__map[key]

    def _set(self, key):
        self.__map[key] = time.time()

    def __setitem__(self, key, value):
        # Set internal map
        self._set(key)
        # Set it
        return dict.__setitem__(self, key, value)

    def __getitem__(self, key):
        # Return
        val = dict.__getitem__(self, key)
        return val

    def __del__(self):
        self._flag = False
        self._t.join()
