#!/usr/bin/env python3

import csv
from fuzzywuzzy import process

def csv_reader(file, key):
    with open(file) as csvfile:
        reader = csv.DictReader(csvfile)
        return [row[key] for row in reader]

const_data = csv_reader("constituents.csv", "Name")
target_data = csv_reader("sp500.csv", "company")

with open('results_python.csv', 'w', newline='') as csvfile:
    fieldnames = ["name", "match_name", "match_score"]
    writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
    writer.writeheader()
    for row in const_data:
        match = process.extractOne(row, target_data)
        writer.writerow({'name': row, 'match_name': match[0], 'match_score': match[1]})
   
