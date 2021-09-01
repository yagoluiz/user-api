using System.Collections.Generic;

namespace User.API.Responses
{
    public class UsersPaginationResponse
    {
        public UsersPaginationResponse(int @from, int size, IEnumerable<UsersResponse> data)
        {
            From = @from;
            Size = size;
            Data = data;
        }

        public int From { get; }
        public int Size { get; }
        public IEnumerable<UsersResponse> Data { get; }
    }
}
