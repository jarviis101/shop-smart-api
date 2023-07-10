db.createUser({
    user: "user",
    pwd: "password",
    roles: [{
        role: "readWrite",
        db: "shop-smart"
    }]
});
db.createCollection('users');
db.users.createIndex({ phone: 1 }, { unique: true });