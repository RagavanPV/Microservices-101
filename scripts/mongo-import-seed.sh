authDB="admin"
username="admin"
password="admin"
importDB="productsDB"
collectionName="products"
file_type="csv"
folder_of_dataset="/Users/apple/Downloads/archive"
for csvfile in $folder_of_dataset/*.csv
do
    echo "Importing $csvfile"
    mongoimport --authenticationDatabase $authDB --username $username --password $password --db $importDB --collection $collectionName --type $file_type --headerline --ignoreBlanks --file $csvfile
    echo "Import of $csvfile complete"
done