print("Started creating DB User");
db = db.getSiblingDB("admin");
db.createUser({
    user: "band-admin",
    pwd: "band-admin",
    roles: [
        {
            role: "readWrite",
            db: "admin"
        }
    ]
});
print("Finished creating DB User");

print("Started creating DB");
db = db.getSiblingDB("band-app");
db.createCollection("bands");
db.createCollection("albums");
db.createCollection("tracks");
db.createCollection("users");
print("Finished creating DB");