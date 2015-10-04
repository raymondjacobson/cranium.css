import csv
import argparse

def mergecsv(csv1, csv2, out):
	with open(out, 'wb') as out_f:
		w = csv.writer(out_f)
		with open(csv1, 'rb') as f1:
			r = csv.reader(f1)
			for row in r:
				w.writerow(row)
		with open(csv2, 'rb') as f2:
			r = csv.reader(f2)
			for row in r:
				w.writerow(row)

if __name__ == "__main__":
	main()