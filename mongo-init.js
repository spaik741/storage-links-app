
db.createUser({
    user: "dbuser",
    pwd: "dbpassword",
    roles: [{
        role: "readWrite",
        db: "storage"
    }]
});

db.createCollection("storage-links")
