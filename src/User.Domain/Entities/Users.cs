using MongoDB.Bson.Serialization.Attributes;

namespace User.Domain.Entities
{
    public class Users
    {
        public Users(string id, string name, string username, int? priority)
        {
            Id = id;
            Name = name;
            Username = username;
            Priority = priority;
        }

        public string Id { get; }
        [BsonElement("name")] public string Name { get; }
        [BsonElement("username")] public string Username { get; }

        [BsonIgnoreIfNull]
        [BsonElement("priority")]
        public int? Priority { get; }
    }
}
