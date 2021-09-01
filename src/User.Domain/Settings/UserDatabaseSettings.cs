namespace User.Domain.Settings
{
    public class UserDatabaseSettings
    {
        public string ConnectionString { get; set; }
        public string UsersCollectionName { get; set; }
        public string DatabaseName { get; set; }
    }
}
