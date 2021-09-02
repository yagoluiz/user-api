using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Moq;
using User.API.Controllers;
using User.API.Services.Interfaces;
using User.Unit.Tests.Builders;
using Xunit;

namespace User.Unit.Tests.Controllers
{
    public class UsersControllerTest
    {
        private readonly Mock<IUsersService> _usersServiceMock;

        public UsersControllerTest()
        {
            _usersServiceMock = new Mock<IUsersService>();
        }

        [Fact(DisplayName = "Get all pagination by term when is success")]
        public async Task GetAllPaginationByTermWhenIsSuccessTest()
        {
            var request = UsersBuilder.UsersSearchRequest;

            _usersServiceMock.Setup(setup => setup.GetAllPaginationByTermAsync(request))
                .ReturnsAsync(UsersBuilder.UsersPaginationResponse);

            var controller = new UsersController(_usersServiceMock.Object);
            var result = await controller.GetAllPaginationByTermAsync(request);

            Assert.IsType<OkObjectResult>(result.Result);
        }
    }
}
