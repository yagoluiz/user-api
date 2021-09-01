using System.Threading.Tasks;
using User.Infra.Repositories;
using User.Integration.Tests.Fixtures;
using Xunit;

namespace User.Integration.Tests.Repositories
{
    public class UsersRepositoryTest : IClassFixture<MongoFixture>
    {
        private readonly MongoFixture _mongoFixture;

        public UsersRepositoryTest(MongoFixture mongoFixture)
        {
            _mongoFixture = mongoFixture;
        }

        [Fact(DisplayName = "Get all pagination by term records in database")]
        public async Task GetAllPaginationByTermRecordsInDatabaseTest()
        {
            var context = _mongoFixture.MongoContext;

            var repository = new UsersRepository(context);

            var users = await repository.GetAllPaginationByTermAsync(0, 15, "yago");

            Assert.NotEmpty(users);
        }
    }
}
