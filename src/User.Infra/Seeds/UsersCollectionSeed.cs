using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using CsvHelper;
using CsvHelper.Configuration;
using MongoDB.Bson;
using MongoDB.Driver;
using User.Domain.Entities;
using User.Domain.Models;
using User.Infra.Contexts;

namespace User.Infra.Seeds
{
    public static class UsersSeed
    {
        public static void RunSeed(MongoContext context)
        {
            var users = context.Users.Find(new BsonDocument()).FirstOrDefault();

            if (users != null) return;

            AddUsersIndexes(context);
            AddUsers(context);
        }

        private static void AddUsersIndexes(MongoContext context)
        {
            var indexes = Builders<Users>.IndexKeys
                .Text(users => users.Name)
                .Text(users => users.Username);

            context.Users.Indexes.CreateOne(new CreateIndexModel<Users>(indexes, new CreateIndexOptions
            {
                DefaultLanguage = "pt"
            }));
        }

        private static void AddUsers(MongoContext context)
        {
            var usersCsvRecords = GetUsersCsvRecords();
            var usersPriorityFirstCsvRecords = GetUsersPriorityFirstCsvRecords();
            var usersPrioritySecondCsvRecords = GetUsersPrioritySecondCsvRecords();

            var users = usersCsvRecords.Select(userCsv => new Users(
                userCsv.Id,
                userCsv.Name,
                userCsv.Username,
                usersPriorityFirstCsvRecords.Any(userPriority => userPriority.Id == userCsv.Id)
                    ? 1
                    : usersPrioritySecondCsvRecords.Any(userPriority => userPriority.Id == userCsv.Id)
                        ? 2
                        : null
            ));

            context.Users.InsertMany(users);
        }

        private static IEnumerable<UsersCsvModel> GetUsersCsvRecords()
        {
            using var csv = GetCsvReader("Imports/users.csv");

            var records = csv.GetRecords<UsersCsvModel>();

            return records.ToList();
        }

        private static IEnumerable<UsersPriorityCsvModel> GetUsersPriorityFirstCsvRecords()
        {
            using var csv = GetCsvReader("Imports/users_priority1.csv");

            var records = csv.GetRecords<UsersPriorityCsvModel>();

            return records.ToList();
        }

        private static IEnumerable<UsersPriorityCsvModel> GetUsersPrioritySecondCsvRecords()
        {
            using var csv = GetCsvReader("Imports/users_priority2.csv");

            var records = csv.GetRecords<UsersPriorityCsvModel>();

            return records.ToList();
        }

        private static CsvReader GetCsvReader(string filePath)
        {
            var path = Path.Combine(
                Path.GetDirectoryName(AppDomain.CurrentDomain.BaseDirectory) ?? string.Empty,
                filePath
            );

            var csvConfiguration = new CsvConfiguration(CultureInfo.InvariantCulture)
            {
                MissingFieldFound = null
            };

            var reader = new StreamReader(path);
            var csv = new CsvReader(reader, csvConfiguration);

            return csv;
        }
    }
}
