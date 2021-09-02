using System;
using System.Diagnostics.CodeAnalysis;
using System.IO;
using System.IO.Compression;
using System.Reflection;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Diagnostics.HealthChecks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.ResponseCompression;
using Microsoft.AspNetCore.Server.Kestrel.Core;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Diagnostics.HealthChecks;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.OpenApi.Models;
using MongoDB.Bson;
using MongoDB.Driver;
using MongoDB.Driver.Core.Events;
using User.API.Extensions;
using User.API.Middlewares;
using User.API.Services;
using User.API.Services.Interfaces;
using User.Domain.Interfaces.Repositories;
using User.Infra.Contexts;
using User.Infra.Repositories;

namespace User.API
{
    [ExcludeFromCodeCoverage]
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers();
            services.AddResponseCompression();
            services.Configure<GzipCompressionProviderOptions>(options =>
            {
                options.Level = CompressionLevel.Fastest;
            });
            services.Configure<KestrelServerOptions>(options => { options.AllowSynchronousIO = true; });
            services.AddSwaggerGen(options =>
            {
                options.SwaggerDoc("v1", new OpenApiInfo { Title = "User API", Version = "v1" });

                var xmlFile = $"{Assembly.GetExecutingAssembly().GetName().Name}.xml";
                var xmlPath = Path.Combine(AppContext.BaseDirectory, xmlFile);
                options.IncludeXmlComments(xmlPath);
            });

            AddHealthChecks(services);
            AddDatabaseContext(services);
            AddServiceScopes(services);
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            app.UseMongoContextSeed();

            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
                app.UseSwagger();
                app.UseSwaggerUI(c => c.SwaggerEndpoint("/swagger/v1/swagger.json", "User API v1"));
            }

            app.UseRouting();
            app.UseLogMiddleware();
            app.UseExceptionHandler(new ExceptionHandlerOptions
            {
                ExceptionHandler = new ErrorHandlerMiddleware(env).Invoke
            });
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
                endpoints.MapHealthChecks("/health", new HealthCheckOptions
                {
                    ResultStatusCodes =
                    {
                        [HealthStatus.Healthy] = StatusCodes.Status200OK,
                        [HealthStatus.Degraded] = StatusCodes.Status200OK,
                        [HealthStatus.Unhealthy] = StatusCodes.Status503ServiceUnavailable
                    }
                });
            });
        }

        private void AddHealthChecks(IServiceCollection services)
        {
            services.AddHealthChecks().AddMongoDb(Configuration["MONGO_HOST"]);
        }

        private void AddDatabaseContext(IServiceCollection services)
        {
            services.AddSingleton<IMongoClient>(service =>
            {
                var settings = MongoClientSettings.FromUrl(new MongoUrl(Configuration["MONGO_HOST"]));
                settings.ClusterConfigurator = builder =>
                {
                    var logger = service.GetRequiredService<ILogger<Startup>>();

                    builder.Subscribe<CommandStartedEvent>(@event =>
                    {
                        logger.LogInformation(
                            $"Command Started: {@event.CommandName}, Json: {@event.Command.ToJson()}");
                    });
                };

                return new MongoClient(settings);
            });

            services.AddSingleton<MongoContext>();

            services.AddScoped<IUsersRepository, UsersRepository>();
        }

        private void AddServiceScopes(IServiceCollection services)
        {
            services.AddScoped<IUsersService, UsersService>();
        }
    }
}
