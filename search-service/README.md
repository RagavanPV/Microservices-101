# Search Service

Phase 1:
Using Elastic Search as a search engine, because of its indexing capabilities
First we start index on the dataset which are indexed in elastic search from mongodb

Data from MongoDB are indexed to Elasticsearch using monstache

monstache -f ./docker-compose/databases/products-service/config.toml -verbose

Index files
./sist2 scan ../sist2-admin/data/ --output ../sist2-admin/documents.idx --very-verbose

./sist2 index --es-url https://elastic:elastic@es01:9200 ./documents.idx --es-insecure-ssl --very-verbose