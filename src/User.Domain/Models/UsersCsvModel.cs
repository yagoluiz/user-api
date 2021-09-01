using CsvHelper.Configuration.Attributes;

namespace User.Domain.Models
{
    public class UsersCsvModel
    {
        [Index(0)] public string Id { get; set; }
        [Index(1)] public string Name { get; set; }
        [Index(2)] public string Username { get; set; }
    }
}
