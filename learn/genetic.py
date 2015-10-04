import scipy.spacial.distance
import csv
import numpy as np
import argparse

import pymongo

from pymongo import MongoClient

client = MongoClient('localhost', 27017)
db = client.cranium

def readCorrectPoints(filename):
	data = {'imp':[], 'nimp':[]}
	with open(filename, 'r') as f:
		r = csv.reader(f)
		for row in r:
			if row[0].split('-')[0] == 'nimp':
				imp = False
			elif row[0].split('-')[0] == 'imp':
				imp = True
			else:
				print "ERROR: First element in data row is neither imp nor nimp"
				return

			if (row[-1] == '-1' and not imp) or (row[-1] == '1' and imp):
				correct = True
			elif (row[-1] == '-1' and imp) or (row[-1] == '1' and not imp):
				correct = False
			else:
				print "ERROR: Last element in data row is neither -1 nor 1"
				return

			if correct:
				datarow = []
				for i in xrange(1, len(row)-1):
					datarow.append(int(row[i]))
				if imp:
					data['imp'].append(datarow)
				else:
					data['nimp'].append(datarow)
	return data

def findNearestNeighbors(point, train_points):
	with_dists = []

	curr_min = -1
	for t_pt in train_points:
		dist = scipy.spacial.distance.euclidean(point, t_pt)
		extended = dist + t_pt
		with_dists.append(extended)

	key = lambda x: x[0]
	with_dists.sort(key=key)

	gen_pts = with_dists[:2]
	gen_pts += point

	return gen_pts

def averagePoints(points):
	arr = np.array(points)
	avg = np.mean(arr, axis=0)

	return avg

def genNewPoint(csvfile):
	


def main():
	filename = 'learn/training data/ptag.csv'
	print readCorrectPoints(filename)

if __name__ == "__main__":
	parser = argparse.ArgumentParser(description='Genetic algorithm to fix mismatched points.')
	parser.add_argument('csvfile', type=str,
	                    help='csv file with training points')

	args = parser.parse_args()
	genNewPoint(args.csvfile)