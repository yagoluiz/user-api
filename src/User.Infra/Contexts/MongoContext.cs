using Microsoft.Extensions.Configuration;
using MongoDB.Driver;
using User.Domain.Entities;

namespace User.Infra.Contexts
{
    public class MongoContext
    {
        private readonly IConfiguration _configuration;
        public IMongoDatabase Database { get; }

        public MongoContext(IMongoClient mongoClient, IConfiguration configuration)
        {
            _configuration = configuration;

            Database = mongoClient.GetDatabase(configuration["MONGO_USER_DATABASE"]);
        }

        public IMongoCollection<Users> Users => Database.GetCollection<Users>(_configuration["MONGO_USERS_COLLECTION"]);
    }
}
