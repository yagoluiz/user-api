using System.Collections.Generic;
using System.Threading.Tasks;
using User.Domain.Entities;

namespace User.Domain.Interfaces.Repositories
{
    public interface IUsersRepository
    {
        Task<IEnumerable<Users>> GetAllPaginationByTermAsync(int page, int limit, string term);
    }
}
