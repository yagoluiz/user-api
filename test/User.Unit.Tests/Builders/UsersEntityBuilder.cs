using System.Collections.Generic;
using Bogus;
using User.Domain.Entities;

namespace User.Unit.Tests.Builders
{
    public static class UsersEntityBuilder
    {
        public static IEnumerable<Users> Users =>
            new Faker<Users>()
                .CustomInstantiator(faker => new Users(
                    faker.Random.Guid().ToString(),
                    faker.Person.FullName,
                    faker.Person.UserName,
                    faker.Random.Number(1, 10)))
                .Generate(10);
    }
}
