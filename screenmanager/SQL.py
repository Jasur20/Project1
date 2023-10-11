import sqlite3

def connect():
    global db,cur
    db=sqlite3.connect("main.db")
    cur=db.cursor()

def inizialize():
    cur.execute("CREATE TABLE IS NOT EXISTS Password(Old text, New text)")
    db.commit()

def add_value_for_table_password(Old=None,New=None):
    cur.execute(f"INSERT INTO Password(Old,New) VALUES (?,?)",
    [Old,New])
    db.commit()
