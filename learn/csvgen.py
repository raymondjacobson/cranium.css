import argparse
import csv
import json


def main(de0):
  dump0() if de0 else dumpAll()


def dumpAll():
  with open("de_dump_all.json", "r") as de_dump_all:
    dump = json.load(de_dump_all)
    a_file = open("de_td_all_a.csv", "wb+")
    p_file = open("de_td_all_p.csv", "wb+")
    img_file = open("de_td_all_img.csv", "wb+")

    a_writer = csv.writer(a_file)
    p_writer = csv.writer(p_file)
    img_writer = csv.writer(img_file)

    for de in dump["data"]:
      writeDataEntry(de, a_writer, p_writer, img_writer)
    
    a_file.close()
    p_file.close()
    img_file.close()
  de_dump_all.close()


def dump0():
  with open("de_dump0.json", "r") as de_dump0:
    dump = json.load(de_dump0)
    a_file = open("de_td0_a.csv", "wb+")
    p_file = open("de_td0_p.csv", "wb+")
    img_file = open("de_td0_img.csv", "wb+")

    a_writer = csv.writer(a_file)
    p_writer = csv.writer(p_file)
    img_writer = csv.writer(img_file)

    writeDataEntry(dump["data"][0], a_writer, p_writer, img_writer)
    
    a_file.close()
    p_file.close()
    img_file.close()
  de_dump0.close()


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
      elem["id"],
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
      elem["id"],
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
      elem["id"],
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
