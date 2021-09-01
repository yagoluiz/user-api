using CsvHelper.Configuration.Attributes;

namespace User.Domain.Models
{
    public class UsersPriorityCsvModel
    {
        [Index(0)] public string Id { get; set; }
    }
}
