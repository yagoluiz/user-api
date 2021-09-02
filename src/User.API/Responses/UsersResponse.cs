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

        /// <summary>
        /// User id
        /// </summary>
        public string Id { get; }

        /// <summary>
        /// User name
        /// </summary>
        public string Name { get; }

        /// <summary>
        /// User identify
        /// </summary>
        public string Username { get; }
    }
}
