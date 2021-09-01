using System;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using User.Infra.Contexts;
using User.Infra.Seeds;

namespace User.API.Extensions
{
    public static class MongoContextSeedExtension
    {
        public static void UseMongoContextSeed(this IApplicationBuilder builder)
        {
            using var scope = builder.ApplicationServices.CreateScope();

            var services = scope.ServiceProvider;
            var logger = services.GetRequiredService<ILogger<Startup>>();

            try
            {
                var context = services.GetRequiredService<MongoContext>();

                UsersSeed.RunSeed(context);

                logger.LogInformation("Seed Users collection successfully");
            }
            catch (Exception exception)
            {
                logger.LogError(exception, "An error occurred while seeding the database");
            }
        }
    }
}
