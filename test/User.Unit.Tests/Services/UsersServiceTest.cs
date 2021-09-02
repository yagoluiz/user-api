using System.Threading.Tasks;
using Moq;
using User.API.Services;
using User.Domain.Interfaces.Repositories;
using User.Unit.Tests.Builders;
using Xunit;

namespace User.Unit.Tests.Services
{
    public class UsersServiceTest
    {
        private readonly Mock<IUsersRepository> _usersRepositoryMock;

        public UsersServiceTest()
        {
            _usersRepositoryMock = new Mock<IUsersRepository>();
        }

        [Fact(DisplayName = "Get all pagination by term when items is not empty")]
        public async Task GetAllPaginationByTermWhenItemsIsNotEmptyTest()
        {
            var request = UsersBuilder.UsersSearchRequest;

            _usersRepositoryMock.Setup(setup => setup.GetAllPaginationByTermAsync(
                    request.From,
                    request.Size,
                    request.Query))
                .ReturnsAsync(UsersEntityBuilder.Users);

            var service = new UsersService(_usersRepositoryMock.Object);
            var result = await service.GetAllPaginationByTermAsync(request);

            Assert.NotEmpty(result.Data);
        }

        [Fact(DisplayName = "Get all pagination by term when items is empty")]
        public async Task GetAllPaginationByTermWhenItemsIsEmptyTest()
        {
            var request = UsersBuilder.UsersSearchRequest;

            _usersRepositoryMock.Setup(setup => setup.GetAllPaginationByTermAsync(
                    request.From,
                    request.Size,
                    request.Query))
                .ReturnsAsync(UsersEntityBuilder.Users);

            var service = new UsersService(_usersRepositoryMock.Object);
            var result = await service.GetAllPaginationByTermAsync(request);

            Assert.NotEmpty(result.Data);
        }
    }
}
