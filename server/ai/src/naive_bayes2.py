import sqlite3
import mmap
import os
import sys
import copy
import math
import tempfile

from tqdm import tqdm
from scipy.stats import norm
from expiringdict import ExpiringDict

cache = ExpiringDict(max_len=100000,max_age_seconds=600)

def get_num_lines(file_path):
    fp = open(file_path, "r+")
    buf = mmap.mmap(fp.fileno(), 0)
    lines = 0
    while buf.readline():
        lines += 1
    return lines

class ExtendedNaiveBayes2:
    
    def __init__(self,family,path_to_data="."):
        self.family = family 
        self.db_name = os.path.join(path_to_data,family+".nb2.db")

    def fit(self,csv_file):
        db = sqlite3.connect(":memory:")
        c = db.cursor()
        try:
            c.execute('''CREATE TABLE data (loc TEXT, mac TEXT, val INTEGER, count INTEGER)''')
            db.commit()
        except sqlite3.OperationalError:
            pass

        headers = []
        with open(csv_file,"r") as f:
            for i,line in enumerate(tqdm(f, total=get_num_lines(csv_file))):
                line = line.strip()
                if i == 0:
                    headers = line.split(",")
                    continue
                loc = ""
                for j,signal in enumerate(line.split(",")):
                    if j == 0:
                        loc = signal 
                        continue
                    if signal.strip() == "":
                        continue
                    mac = headers[j]
                    val = int(round(float(signal.strip())))
                    c.execute('''SELECT count FROM data WHERE loc = ? AND mac = ? AND val = ?''',(loc, mac, val
        ))
                    count = c.fetchone()
                    if count == None:
                        c.execute('''INSERT INTO data(loc,mac,val,count)
                          VALUES(?,?,?,?)''', (loc,mac,val,1))
                    else:
                        c.execute('''UPDATE data SET count = ? WHERE loc = ? AND mac = ? AND val = ?''',(count[0]+1,loc,mac,val))
                    db.commit()

        # with open("dump.sql","w") as f:
        #     for line in db.iterdump():
        #         f.write('%s\n' % line)
        f = tempfile.TemporaryFile()
        for line in db.iterdump():
            f.write('{}\n'.format(line).encode('utf-8'))

        db.close()

        # Write disk to file
        try:
            os.remove(self.db_name)
        except:
            pass
        db = sqlite3.connect(self.db_name)
        c = db.cursor()
        f.seek(0)
        c.executescript(f.read().decode('utf-8'))
        f.close()
        db.commit()
        db.close()
        # os.remove("dump.sql")


    def get_locations(self):
        db = sqlite3.connect(self.db_name)
        c = db.cursor()
        c.execute('''SELECT loc FROM data GROUP BY loc''')
        locs = c.fetchall()
        db.close()
        locations = []
        for l in locs:
            locations.append(l[0])
        return locations

    def prob_mac_given_loc(self,mac,val,loc,positive):
        """
        Determine the P(mac | loc) (positive)
        Determine the P(mac | ~loc) (not positive)
        """
        name = "{}{}{}".format(mac,loc,positive)
        cached = cache.get(name)
        if cached != None:
            return cached

        # First find all the values for mac at loc
        db = sqlite3.connect(self.db_name)
        c = db.cursor()
        if positive:
            c.execute('''SELECT 10000*(SELECT SUM(count) FROM data WHERE loc = ? AND mac = ?)/(SELECT SUM(count) FROM data WHERE mac = ?)''',(loc,mac,mac))
        else:
            c.execute('''SELECT 10000*(SELECT SUM(count) FROM data WHERE loc != ? AND mac = ?)/(SELECT SUM(count) FROM data WHERE mac = ?)''',(loc,mac,mac))
        Pvalue_times_10000 = c.fetchone()[0]
        if Pvalue_times_10000 == None:
            Pvalue_times_10000 = 0
        db.close()
        P = float(Pvalue_times_10000) / float(10000)
        if P == 0:
            P = 0.005
        cache[name] = P 
        return P

    def predict_proba(self,header_unfiltered,csv_data_unfiltered):
        header = []
        csv_data = []
        for i,dat in enumerate(csv_data_unfiltered):
            if dat == 0:
                continue
            csv_data.append(dat)
            header.append(header_unfiltered[i])

        locations = self.get_locations()
        num_locations = len(locations)
        NA = 1/num_locations
        NnotA = 1-NA
        Ps = {}
        for i,mac in enumerate(header):
            val = int(round(float(csv_data[i])))
            for location in locations:
                if location not in Ps:
                    Ps[location] = []
                PA = self.prob_mac_given_loc(mac,val,location,True)
                PnotA = self.prob_mac_given_loc(mac,val,location,False)
                P = PA*NA / (PA*NA + PnotA*NnotA)
                Ps[location].append(math.log(P))
        P_sum = 0
        for location in Ps:
            P_sum += math.exp(sum(Ps[location]))
        d = {}
        for location in Ps:
            d[location] = math.exp(sum(Ps[location]))/P_sum
        return [(k, d[k]) for k in sorted(d, key=d.get, reverse=True)]



def testit():
    a =ExtendedNaiveBayes2("testing1")
    print("fitting data")
    file_to_test = "reverse.csv"
    a.fit(file_to_test)
    print("done")
    with open(file_to_test,"r") as f:
        for i,line in enumerate(f):
            line = line.strip()
            if i == 0:
                headers = line.split(",")
                continue
            headers_submit = []
            csv_data_submit = []
            loc = ""
            for j,signal in enumerate(line.split(",")):
                if j == 0:
                    loc = signal 
                    continue
                if signal.strip() == "":
                    continue
                headers_submit.append(headers[j])
                csv_data_submit.append(int(round(float(signal.strip()))))
            print(loc)
            a.predict_proba(headers_submit,csv_datat5rrrrrrrr1_submit)
