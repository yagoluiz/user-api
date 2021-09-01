using Microsoft.Extensions.Configuration;
using MongoDB.Driver;
using User.Domain.Entities;

namespace User.Infra.Contexts
{
    public class MongoContext
    {
        private readonly IConfiguration _configuration;
        private IMongoDatabase Database { get; }

        public MongoContext(IConfiguration configuration)
        {
            _configuration = configuration;

            var client = new MongoClient(configuration["MONGO_HOST"]);
            Database = client.GetDatabase(configuration["MONGO_DATABASE"]);
        }

        public IMongoCollection<Users> Users => Database.GetCollection<Users>(_configuration["MONGO_USERS_COLLECTION"]);
    }
}
