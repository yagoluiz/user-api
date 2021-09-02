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

        /// <summary>
        /// Pagination from
        /// </summary>
        public int From { get; }

        /// <summary>
        /// Size items
        /// </summary>
        public int Size { get; }

        /// <summary>
        /// Users
        /// </summary>
        public IEnumerable<UsersResponse> Data { get; }
    }
}
