import csv
import argparse
from sklearn.externals import joblib
import numpy as np
import matplotlib.pyplot as plt
from sklearn import svm

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
		print "point", point
		print "tp", t_pt
		dist = np.linalg.norm(np.array(point) - np.array(t_pt))
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
	db.visitors.find({})

def classify(file, type, vid, acsvfile, pcsvfile, imgcsvfile):
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


	elif type == "solve":

		visitor = db.visitors.find_one({"vid": vid})
		print visitor['data']
		tags = visitor['data'][0]['atags'] + visitor['data'][0]['ptags'] + visitor['data'][0]['imgtags']
		a_tags = []
		p_tags = []
		img_tags = []
		for tag in tags:
			new_point_dict = {}
			new_point_dict['important'] = tag['important']
			new_point_dict['id'] = tag['id']
			if tag['classification'] > 0 and not tag['important']:
				tag_type = tag['id'].split('-')[1]
				print tag_type
				if tag_type == 'a':
					read_cor_pts = readCorrectPoints(acsvfile)
					old_point = [tag['fontsize'], tag['fontstyle'], int(tag['color']), tag['padding'], tag['hover'], tag['click'], tag['frametime']]
					target_points = read_cor_pts['imp']
					new_point = averagePoints(findNearestNeighbors(old_point, target_points))
					new_point_dict['fontsize'] = new_point[0]
					new_point_dict['fontstyle'] = new_point[1]
					new_point_dict['color'] = int(new_point[2])
					new_point_dict['padding'] = new_point[3]
					a_tags.append(new_point_dict)
				elif tag_type == 'p':
					read_cor_pts = readCorrectPoints(pcsvfile)
					old_point = [tag['fontsize'], tag['fontstyle'], tag['padding'], tag['hover'], tag['click'], tag['frametime']]
					target_points = read_cor_pts['imp']
					new_point = averagePoints(findNearestNeighbors(old_point, target_points))
					new_point_dict['fontsize'] = new_point[0]
					new_point_dict['fontstyle'] = new_point[1]
					new_point_dict['padding'] = new_point[2]
					p_tags.append(new_point_dict)
				elif tag_type == 'img':
					read_cor_pts = readCorrectPoints(imgcsvfile)
					old_point = [tag['width'], tag['padding'], tag['hover'], tag['click'], tag['frametime']]
					target_points = read_cor_pts['imp']
					new_point = averagePoints(findNearestNeighbors(old_point, target_points))
					new_point_dict['width'] = new_point[0],
					new_point_dict['padding'] = new_point[1]
					img_tags.append(new_point_dict)
			elif tag['classification'] < 0 and tag['important']:
				tag_type = tag['id'].split('-')[1]
				print tag_type
				if tag_type == 'a':
					read_cor_pts = readCorrectPoints(acsvfile)
					old_point = [tag['fontsize'], tag['fontstyle'], int(tag['color']), tag['padding'], tag['hover'], tag['click'], tag['frametime']]
					target_points = read_cor_pts['nimp']
					new_point = averagePoints(findNearestNeighbors(old_point, target_points))
					new_point_dict['fontsize'] = new_point[0]
					new_point_dict['fontstyle'] = new_point[1]
					new_point_dict['color'] = int(new_point[2])
					new_point_dict['padding'] = new_point[3]
					a_tags.append(new_point_dict)
				elif tag_type == 'p':
					read_cor_pts = readCorrectPoints(pcsvfile)
					old_point = [tag['fontsize'], tag['fontstyle'], tag['padding'], tag['hover'], tag['click'], tag['frametime']]
					target_points = read_cor_pts['nimp']
					new_point = averagePoints(findNearestNeighbors(old_point, target_points))
					new_point_dict['fontsize'] = new_point[0]
					new_point_dict['fontstyle'] = new_point[1]
					new_point_dict['padding'] = new_point[2]
					p_tags.append(new_point_dict)
				elif tag_type == 'img':
					read_cor_pts = readCorrectPoints(imgcsvfile)
					old_point = [tag['width'], tag['padding'], tag['hover'], tag['click'], tag['frametime']]
					target_points = read_cor_pts['nimp']
					new_point = averagePoints(findNearestNeighbors(old_point, target_points))
					new_point_dict['width'] = new_point[0],
					new_point_dict['padding'] = new_point[1]
					img_tags.append(new_point_dict)
			else:
				tag_type = tag['id'].split('-')[1]
				if tag_type == 'a':
					a_tags.append(tag)
				elif tag_type == 'p':
					p_tags.append(tag)
				elif tag_type == 'img':
					img_tags.append(tag)
		data_entry = {
			'atags': a_tags,
			'ptags': p_tags,
			'imgtags': img_tags
		}

		db.visitors.update_one({"vid":vid}, {"$addToSet": {"data": data_entry}})

if __name__ == '__main__':
	parser = argparse.ArgumentParser(description="some bs SVM")
	parser.add_argument("file", help="in file",
		type=str)
	parser.add_argument("type", help="str",
		type=str)
	parser.add_argument("vid", help="str",
                    type=str)
	parser.add_argument('acsvfile', type=str,
	                    help='csv file with training points')
	parser.add_argument('pcsvfile', type=str,
	                    help='csv file with training points')
	parser.add_argument('imgcsvfile', type=str,
	                    help='csv file with training points')
	args = parser.parse_args()
	classify(args.file, args.type, args.vid, args.acsvfile, args.pcsvfile, args.imgcsvfile)
