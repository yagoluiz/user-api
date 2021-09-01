namespace User.API.Responses
{
    public class UsersResponse
    {
        public UsersResponse(string id, string name, string username)
        {
            Id = id;
            Name = name;
            Username = username;
        }

        public string Id { get; }
        public string Name { get; }
        public string Username { get; }
    }
}
