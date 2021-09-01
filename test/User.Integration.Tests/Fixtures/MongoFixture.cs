using System;
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
            var mongoSetup = new MongoSetup();

            _context = new MongoContext(mongoSetup.MongoConfiguration);

            UsersSeed.RunSeed(_context);
        }

        public MongoContext MongoContext => _context;

        public void Dispose()
        {
            _context.Database.Client.DropDatabase("UserTest");
        }
    }
}
