using System.Linq;
using System.Threading.Tasks;
using User.API.Requests;
using User.API.Responses;
using User.API.Services.Interfaces;
using User.Domain.Interfaces.Repositories;

namespace User.API.Services
{
    public class UsersService : IUsersService
    {
        private readonly IUsersRepository _usersRepository;

        public UsersService(IUsersRepository usersRepository)
        {
            _usersRepository = usersRepository;
        }

        public async Task<UsersPaginationResponse> GetAllPaginationByTermAsync(UsersSearchRequest request)
        {
            var users = await _usersRepository.GetAllPaginationByTermAsync(
                request.From,
                request.Size,
                request.Query
            );

            return new UsersPaginationResponse(
                request.From,
                users.Count(),
                users.Select(user => new UsersResponse(
                    user.Id,
                    user.Name,
                    user.Username
                ))
            );
        }
    }
}
