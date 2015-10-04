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

  a_file = open("learn/de_td_all_a.csv", "wb+")
  p_file = open("learn/de_td_all_p.csv", "wb+")
  img_file = open("learn/de_td_all_img.csv", "wb+")

  a_writer = csv.writer(a_file)
  p_writer = csv.writer(p_file)
  img_writer = csv.writer(img_file)

  for de in query["data"]:
    writeDataEntry(de, a_writer, p_writer, img_writer)
  
  a_file.close()
  p_file.close()
  img_file.close()

def dump0(vid):
  query = db.visitors.find_one({"vid": vid})
  print query['data']

  a_file = open("learn/de_td0_a.csv", "wb+")
  p_file = open("learn/de_td0_p.csv", "wb+")
  img_file = open("learn/de_td0_img.csv", "wb+")

  a_writer = csv.writer(a_file)
  p_writer = csv.writer(p_file)
  img_writer = csv.writer(img_file)

  writeDataEntry(query["data"][0], a_writer, p_writer, img_writer)
  
  a_file.close()
  p_file.close()
  img_file.close()


def writeDataEntry(de, a_writer, p_writer, img_writer):
  for tagType, elems in de.iteritems():
    if tagType == "atags":
      writeAtags(elems, a_writer)
    elif tagType == "ptags":
      writePtags(elems, p_writer)
    else:
      writeImgtags(elems, img_writer)


def writeAtags(elems, writer):
  for elem in elems:
    writer.writerow([
      elem["id"]
      elem["fontsize"],
      elem["fontstyle"],
      1 if elem["color"] else 0,
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"],
    ])


def writePtags(elems, writer):
  for elem in elems:
    writer.writerow([
      elem["id"]
      elem["fontsize"],
      elem["fontstyle"],
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"],
    ])


def writeImgtags(elems, writer):
  for elem in elems:
    writer.writerow([
      elem["id"]
      elem["width"],
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"],
    ])


if __name__ == '__main__':
  parser = argparse.ArgumentParser(description="Dump JSON data to a CSV.")
  parser.add_argument("vid", help="str",
                    type=str)
  parser.add_argument("-de0", action="store_true", default=False)
  args = parser.parse_args()

  main(args.de0, args.vid)
