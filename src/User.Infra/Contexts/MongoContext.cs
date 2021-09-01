using Microsoft.Extensions.Options;
using MongoDB.Driver;
using User.Domain.Entities;
using User.Domain.Settings;

namespace User.Infra.Contexts
{
    public class MongoContext
    {
        private readonly IOptions<UserDatabaseSettings> _options;
        private IMongoDatabase Database { get; }

        public MongoContext(IOptions<UserDatabaseSettings> options)
        {
            _options = options;

            var client = new MongoClient(options.Value.ConnectionString);
            Database = client.GetDatabase(options.Value.DatabaseName);
        }

        public IMongoCollection<Users> Users => Database.GetCollection<Users>(_options.Value.UsersCollectionName);
    }
}
