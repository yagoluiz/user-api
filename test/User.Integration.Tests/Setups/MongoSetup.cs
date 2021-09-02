using System.Collections.Generic;
using Microsoft.Extensions.Configuration;

namespace User.Integration.Tests.Setups
{
    public class MongoSetup
    {
        public MongoSetup()
        {
            MongoConfiguration = GetConfigurations();
        }

        public IConfigurationRoot MongoConfiguration { get; }

        private static IConfigurationRoot GetConfigurations()
        {
            var settings = new Dictionary<string, string>
            {
                { "MONGO_USER_DATABASE", "UserTest" },
                { "MONGO_USERS_COLLECTION", "UsersTest" }
            };

            return new ConfigurationBuilder()
                .AddInMemoryCollection(settings)
                .Build();
        }
    }
}
