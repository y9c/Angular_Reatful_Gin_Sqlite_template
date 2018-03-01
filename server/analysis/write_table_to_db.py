#! /usr/bin/env python
"""
Read and Write Sqlite3 database
"""

import sqlite3
import pandas as pd

# TISSUES = [
# "Bladder",
# "Bone-Marrow",
# "Brain",
# "Bone-Marrow-c-kit",
# "Embryonic-Mesenchyme",
# "Embryonic-Stem-Cell",
# "Fetal-Liver",
# "Kidney",
# "Liver",
# "Lung",
# "Mammary-Gland-Involution",
# "Mammary-Gland-Lactation",
# "Mammary-Gland-Pregnancy",
# "Mammary-Gland-Virgin",
# "Muscle",
# "Neonatal-Calvaria",
# "Neonatal-Muscle",
# "Neonatal-Heart",
# "Neonatal-Rib",
# "Neonatal-Skin",
# "E18-Brain",
# "Ovary",
# "Pancreas",
# "Peripheral-Blood",
# "Placenta",
# "Preimplantation-Embryo",
# "Prostate",
# "Retina",
# "Small-Intestine",
# "Spleen",
# "Stomach",
# "Testis",
# "Thymus",
# "Trophoblast-Stem-Cell",
# "Uterus",
# "Lung-Mesenchyme",
# "Fetal-Brain",
# "Female-Fetal-Gonad",
# "Fetal-Intestine",
# "Fetal-Lung",
# "Fetal-Kidney",
# "Male-Fetal-Gonad",
# "Fetal-Stomache",
# "Bone-Marrow-Mesenchyme",
# "Neonatal-Brain",
# "Mesenchymal-Stem-Cell-Cultured",
# "E8.25-embryo",
# "Figure2-98Clusters",
# "Arcuate-hypothalamus-and-median-eminence",
# ]

TISSUES = ["Brain"]

if __name__ == '__main__':
    for tissue in TISSUES:
        # read csv file into python dataframe
        input_table = "./temp/MCA/tsne_{}.csv".format(tissue)
        df = pd.read_table(input_table, sep=",")
        # write dataframe into sqlite table
        db = sqlite3.connect('../data/db.sqlite3')
        db_table = "tsne_{}".format(tissue)

        # This block of code will delete table in database!
        cursor = db.cursor()
        cursor.execute('''DROP TABLE tsne_Brain''')
        db.commit()

        df.to_sql(db_table, db, if_exists="replace", index_label="id")
        db.close()
