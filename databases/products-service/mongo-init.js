db.createUser({
    user: 'root',
    pwd: 'root',
    roles: [
        {
            role: 'readWrite',
            db: 'productsDB',
        },
    ],
});

db = new Mongo().getDB("productsDB");

db.createCollection('products', { capped: false });