using System.Diagnostics.CodeAnalysis;
using Microsoft.AspNetCore.Builder;
using User.API.Middlewares;

namespace User.API.Extensions
{
    [ExcludeFromCodeCoverage]
    public static class LogExtensions
    {
        public static void UseLogMiddleware(this IApplicationBuilder builder)
        {
            builder.UseMiddleware<LogMiddleware>();
        }
    }
}
