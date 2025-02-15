import pandas as pd
df = pd.read_csv('dac.csv')
df.columns = [c.lower() for c in df.columns] # PostgreSQL doesn't like capitals or spaces

from sqlalchemy import create_engine
engine = create_engine('postgresql://admin:admin@localhost:5432/dac')

df.to_sql("dac-store", engine)