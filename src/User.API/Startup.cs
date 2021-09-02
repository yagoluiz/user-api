using System;
using System.IO;
using System.Reflection;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.OpenApi.Models;
using User.API.Extensions;
using User.API.Services;
using User.API.Services.Interfaces;
using User.Domain.Interfaces.Repositories;
using User.Infra.Contexts;
using User.Infra.Repositories;

namespace User.API
{
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
            services.AddSwaggerGen(options =>
            {
                options.SwaggerDoc("v1", new OpenApiInfo { Title = "User API", Version = "v1" });

                var xmlFile = $"{Assembly.GetExecutingAssembly().GetName().Name}.xml";
                var xmlPath = Path.Combine(AppContext.BaseDirectory, xmlFile);
                options.IncludeXmlComments(xmlPath);
            });

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

            app.UseAuthorization();

            app.UseEndpoints(endpoints => { endpoints.MapControllers(); });
        }

        private void AddDatabaseContext(IServiceCollection services)
        {
            services.AddSingleton<MongoContext>();

            services.AddScoped<IUsersRepository, UsersRepository>();
        }

        private void AddServiceScopes(IServiceCollection services)
        {
            services.AddScoped<IUsersService, UsersService>();
        }
    }
}
