import argparse
from sklearn.externals import joblib
import numpy as np
import matplotlib.pyplot as plt
from sklearn import svm

def classify(file, type):
	dataset = np.loadtxt(file, delimiter=",")
	if type == "a":
		clf = joblib.load('learn/aclassifier.pkl') 
		X = dataset[:,0:7]
	elif type == "p":
		clf = joblib.load('learn/pclassifier.pkl') 
		X = dataset[:,0:7]
	elif type == "img":
		clf = joblib.load('learn/iclassifier.pkl') 
		X = dataset[:,0:7]
	Y = []
	for x in X:
		Y.append(clf.predict(x))
	print Y

if __name__ == '__main__':
	parser = argparse.ArgumentParser(description="some bs SVM")
	parser.add_argument("file", help="in file",
		type=str)
	parser.add_argument("type", help="str",
		type=str)
	args = parser.parse_args()
	classify(file, type)
