rs.initiate(
    {
        _id: 'rs0',
        members: [
            {_id: 0, host: "mongo:27017"}
        ]
    }
)

while (!db.runCommand("ismaster").ismaster) {
    sleep(100)
}

db.createUser({
    user: "mongo_user",
    pwd: "mongo_pass",
    roles: [
        {
            role: "readWrite",
            db: "Backend"
        }
    ]
})
