using System.Threading.Tasks;
using User.API.Requests;
using User.API.Responses;

namespace User.API.Services.Interfaces
{
    public interface IUsersService
    {
        Task<UsersPaginationResponse> GetAllPaginationByTermAsync(UsersSearchRequest request);
    }
}
