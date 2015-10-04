import argparse
import csv
import json

def main(de0):
  dump0() if de0 else dumpAll()


def dumpAll():
  with open("de_dump_all.json", "r") as de_dump_all:
    dump = json.load(de_dump_all)

    with open("de_td_all.csv", "wb+") as de_td_all:
      writer = csv.writer(de_td_all)
      for de in dump["data"]:
        readDataEntry(de, writer)
    de_td_all.close()
  de_dump_all.close()


def dump0():
  with open("de_dump0.json", "r") as de_dump0:
    dump = json.load(de_dump0)

    with open("de_td0.csv", "wb+") as de_td0:
      writer = csv.writer(de_td0)
      readDataEntry(dump["data"][0], writer)
    de_td0.close()
  de_dump0.close()


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
      "atag",
      elem["fontsize"],
      elem["fontstyle"],
      1 if elem["color"] else 0,
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"]
    ])


def writePtags(elems, writer):
  for elem in elems:
    writer.writerow([
      "ptag",
      elem["fontsize"],
      elem["fontstyle"],
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"]
    ])


def writeImgtags(elems, writer):
  for elem in elems:
    writer.writerow([
      "imgtag",
      elem["width"],
      elem["padding"],
      elem["hover"],
      elem["click"],
      elem["frametime"]
    ])


if __name__ == '__main__':
  parser = argparse.ArgumentParser(description="Dump JSON data to a CSV.")
  parser.add_argument("-de0", action="store_true", default=False)
  args = parser.parse_args()

  main(args.de0)
