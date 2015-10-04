import argparse
from sklearn.externals import joblib
import numpy as np
import matplotlib.pyplot as plt
from sklearn import svm

import pymongo

from pymongo import MongoClient

client = MongoClient('localhost', 27017)
db = client.cranium

def classify(file, type, vid):
	pid = ""
	if type == "a":
		f = open(file, 'r')
		dataset = np.loadtxt(f, delimiter=",", usecols=range(1,8))
		f.close()
		f = open(file, 'r')
		pid = np.loadtxt(f, dtype=np.str, delimiter=",", usecols=range(0,1))
		f.close()
		print pid
		clf = joblib.load('aclassifier.pkl') 
		X = dataset[:,1:7]
		Y = []
		for i in xrange(len(X)):
			y = clf.predict(X[i])
			y = list(y)[0]
			print vid, pid
			print y
			Y.append(y)
			# Insert classification into mongo
			db.visitors.update_one({"vid":vid, "data.0.atags.id": pid[i]},
								   {"$set": {"data.0.atags.$.classification": y}})
		print Y
	elif type == "p":
		f = open(file, 'r')
		dataset = np.loadtxt(f, delimiter=",", usecols=range(1,7))
		f.close()
		f = open(file, 'r')
		pid = np.loadtxt(f, dtype=np.str, delimiter=",", usecols=range(0,1))
		f.close()
		print pid
		clf = joblib.load('pclassifier.pkl') 
		X = dataset[:,1:6]
		Y = []
		for i in xrange(len(X)):
			y = clf.predict(X[i])
			y = list(y)[0]
			print vid, pid
			print y
			Y.append(y)
			# Insert classification into mongo
			db.visitors.update_one({"vid":vid, "data.0.ptags.id": pid[i]},
								   {"$set": {"data.0.ptags.$.classification": y}})
		print Y
	elif type == "img":
		f = open(file, 'r')
		dataset = np.loadtxt(f, delimiter=",", usecols=range(1,6))
		f.close()
		f = open(file, 'r')
		pid = np.loadtxt(f, dtype=np.str, delimiter=",", usecols=range(0,1))
		f.close()
		print pid
		clf = joblib.load('imgclassifier.pkl') 
		print dataset
		X = dataset[:,1:5]
		Y = []
		for i in xrange(len(X)):
			y = clf.predict(X[i])
			y = list(y)[0]
			print vid, pid
			print y
			Y.append(y)
			# Insert classification into mongo
			db.visitors.update_one({"vid":vid, "data.0.imgtags.id": pid[i]},
								   {"$set": {"data.0.imgtags.$.classification": y}})
		print Y

if __name__ == '__main__':
	parser = argparse.ArgumentParser(description="some bs SVM")
	parser.add_argument("file", help="in file",
		type=str)
	parser.add_argument("type", help="str",
		type=str)
	parser.add_argument("vid", help="str",
                    type=str)
	args = parser.parse_args()
	classify(args.file, args.type, args.vid)
