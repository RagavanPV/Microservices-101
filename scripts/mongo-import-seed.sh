authDB="admin"
username="admin"
password="admin"
importDB="RoadAccidentDB"
collectionName="RoadAccidentList"
file_type="csv"
folder_of_dataset="/Users/apple/Ragavan/M.Tech/BDS/Assignment-1/"
port="27021"
for csvfile in $folder_of_dataset/*.csv
do
    echo "Importing $csvfile"
    mongoimport --authenticationDatabase admin --username admin --password admin --port 27017 --db 'RoadAccidentDB' --collection 'RoadAccidentList' --type csv --headerline --ignoreBlanks --file 'India_Injury_Road_Accident_Fatality_2017-2020.csv'
    echo "Import of $csvfile complete"
done