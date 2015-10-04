from sklearn.externals import joblib
import argparse
import numpy as np
import matplotlib.pyplot as plt
from sklearn import svm

def learn(file, type):
  f = open(file, 'r')
  clf = svm.SVC()
  if type == "a":
    dataset = np.loadtxt(f, delimiter=",", usecols=range(1,9))
    print dataset
    X = dataset[:,1:7]
    y = dataset[:,7]
    clf.fit(X, y)
    joblib.dump(clf, 'aclassifier.pkl')
  elif type == "p":
    dataset = np.loadtxt(f, delimiter=",", usecols=range(1,8))
    X = dataset[:,1:6]
    y = dataset[:,6]
    clf.fit(X, y)
    joblib.dump(clf, 'pclassifier.pkl')
  elif type == "img":
    dataset = np.loadtxt(f, delimiter=",", usecols=range(1,7))
    X = dataset[:,1:5]
    y = dataset[:,5]
    clf.fit(X, y)
    joblib.dump(clf, 'imgclassifier.pkl')
  f.close()


if __name__ == '__main__':
  parser = argparse.ArgumentParser(description="some bs SVM")
  parser.add_argument("file", help="in file",
    type=str)
  parser.add_argument("type", help="str",
    type=str)
  args = parser.parse_args()
  learn(args.file, args.type)

