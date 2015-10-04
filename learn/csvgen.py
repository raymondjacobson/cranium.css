import argparse
import csv
import json
import pymongo

from pymongo import MongoClient

client = MongoClient('localhost', 27017)
db = client.cranium

def main(de0, vid):
  dump0(vid) if de0 else dumpAll(vid)


def dumpAll(vid):
  query = db.visitors.find_one({"vid": vid})
  print query['data']
  with open("learn/de_td_all.csv", "wb+") as de_td_all:
    writer = csv.writer(de_td_all)
    for de in query['data']:
      readDataEntry(de, writer)


def dump0(vid):
  query = db.visitors.find_one({"vid": vid})
  print query['data']
  with open("learn/de_td0.csv", "wb+") as de_td0:
    writer = csv.writer(de_td0)
    readDataEntry(query['data'][0], writer)



def readDataEntry(de, writer):
  for tagType, elems in de.iteritems():
    if tagType == "atags":
      writeAtags(elems, writer)
    elif tagType == "ptags":
      writePtags(elems, writer)
    else:
      writeImgtags(elems, writer)


def writeAtags(elems, writer):
  for elem in elems:
    writer.writerow([
      elem["fontsize"],
      elem["fontstyle"],
      1 if elem["color"] else 0,
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"],
      elem["id"]
    ])


def writePtags(elems, writer):
  for elem in elems:
    writer.writerow([
      elem["fontsize"],
      elem["fontstyle"],
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"],
      elem["id"]
    ])


def writeImgtags(elems, writer):
  for elem in elems:
    writer.writerow([
      elem["width"],
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"],
      elem["id"]
    ])


if __name__ == '__main__':
  parser = argparse.ArgumentParser(description="Dump JSON data to a CSV.")
  parser.add_argument("vid", help="str",
                    type=str)
  parser.add_argument("-de0", action="store_true", default=False)
  args = parser.parse_args()

  main(args.de0, args.vid)
