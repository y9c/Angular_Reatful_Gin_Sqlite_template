#! /usr/bin/env python
"""
Read and Write Sqlite3 database
"""

import sqlite3
import pandas as pd

TABLE_NAME = "people"

if __name__ == '__main__':
    db = sqlite3.connect('../data/db.sqlite3')
    df = pd.read_sql_query("SELECT * from " + TABLE_NAME, db)
    print(df.head())
    # This block of code will delete table in database!
    # cursor = db.cursor()
    # cursor.execute('''DROP TABLE todo_models''')
    # db.commit()
    db.close()
