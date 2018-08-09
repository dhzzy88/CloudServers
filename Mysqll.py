# /bin/python3
# author:leozhao
# author@email: dhzzy88@163.com

import mysql.connector as sqlcon

db = sqlcon.connect(user="root", password="8865439ZZY", database="bbq",use_unicode=True)
cursor = db.cursor()
cursor.execute("DROP TABLE IF EXISTS PRICE")


sql = """  
CREATE TABLE PRICE (
         Id  CHAR(50) DEFAULT "",
         Name_Foo  CHAR(50)  DEFAULT "",
         Desc_Foo CHAR(50) DEFAULT "",  
         Price CHAR(50) DEFAULT "",
         Number_Foo CHAR(50) DEFAULT "",
         Users CHAR(50) DEFAULT "",
         Date_list CHAR(50) DEFAULT "",
         DeskNumber CHAR(50) DEFAULT "",
         Status  CHAR(50) DEFAULT "")
"""

cursor.execute(sql)
cursor.close()
