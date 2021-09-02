using System;
using MongoDB.Driver;
using User.Infra.Contexts;
using User.Infra.Seeds;
using User.Integration.Tests.Setups;

namespace User.Integration.Tests.Fixtures
{
    public class MongoFixture : IDisposable
    {
        private readonly MongoContext _context;

        public MongoFixture()
        {
            var mongoClient = new MongoClient("mongodb://localhost:27017");
            var mongoSetup = new MongoSetup();

            _context = new MongoContext(mongoClient, mongoSetup.MongoConfiguration);

            UsersSeed.RunSeed(_context);
        }

        public MongoContext MongoContext => _context;

        public void Dispose()
        {
            _context.Database.Client.DropDatabase("UserTest");
        }
    }
}
