from pymongo import MongoClient
from elasticsearch import Elasticsearch
from perf.timed import timed

es_client = Elasticsearch(
    "https://localhost:9200",
    ca_certs="/Users/apple/Ragavan/Learning/E-Commerce/docker-compose/elastic-stack/es-certs/ca/ca.crt",
    basic_auth=("elastic", "elastic")
)
mongo_client = MongoClient("mongodb://admin:admin@localhost:27021/?replicaSet=dbrs&directConnection=true")

@timed
def elasticsearch():
    result = es_client.search(index="productsdb.products", body={"query": {"query_string": {"query": "accessories","fields"  : ["main_category.keyword"]}}})
    return result

@timed
def mongosearch():
    result = mongo_client.productsDB.products.find({"main_category": "accessories"})
    return result

result = mongosearch()
print(len(list(result.clone())))

result = elasticsearch()
print(result)

