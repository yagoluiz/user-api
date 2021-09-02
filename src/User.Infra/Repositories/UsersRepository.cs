using System.Collections.Generic;
using System.Threading.Tasks;
using MongoDB.Driver;
using User.Domain.Entities;
using User.Domain.Interfaces.Repositories;
using User.Infra.Contexts;

namespace User.Infra.Repositories
{
    public class UsersRepository : IUsersRepository
    {
        private readonly MongoContext _context;

        public UsersRepository(MongoContext context)
        {
            _context = context;
        }

        public async Task<IEnumerable<Users>> GetAllPaginationByTermAsync(int page, int limit, string term)
        {
            var filter = Builders<Users>.Filter.Text(term);

            var users = await _context.Users.Find(filter)
                .Skip(limit * page)
                .Limit(limit)
                .SortBy(user => user.Priority)
                .ToListAsync();

            return users;
        }
    }
}
